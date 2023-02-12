package main

import (
	"fmt"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"
)

func DiscoverDirs(wl *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("ReadDir error:", err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			nextpath := filepath.Join(path, entry.Name())
			DiscoverDirs(wl, nextpath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

var args struct {
	SearchTerm string `arg:"positional,required"`
	SearchDir  string `arg:"positional"`
}

func main() {
	arg.MustParse(&args)

	var workersWg sync.WaitGroup

	wl := worklist.New(100)

	results := make(chan worker.Result, 100)

	numWorkers := 10

	workersWg.Add(1)
	go func() {
		defer workersWg.Done()
		DiscoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workResult := worker.FindInFilePath(workEntry.Path, args.SearchTerm)
					if workResult != nil {
						for _, r := range workResult.Inner {
							results <- r
						}
					}
				} else {
					return
				}
			}
		}()
	}

	blockWorkersWg := make(chan struct{})
	go func() {
		workersWg.Wait()
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup
	displayWg.Add(1)
	go func() {
		for {
			select {
			case r := <-results:
				fmt.Printf("%v[%v]:%v\n", r.Path, r.LineNum, r.Line)
			case <-blockWorkersWg:
				if len(results) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()
	displayWg.Wait()
}
