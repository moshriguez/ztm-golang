package timeparse

import (
	"errors"
	"fmt"
	"testing"
)

func TestTimeFromString(t *testing.T) {
	tests := []struct {
		name     string
		testStrs []string
		wantErr  error
		result   Time
	}{
		{
			name:     "successfully parses a formatted string",
			testStrs: []string{"03:30:45"},
			result: Time{
				hour:   3,
				minute: 30,
				second: 45,
			},
		},
		{
			name:     "throws error if string is not formatted correctly",
			testStrs: []string{"blah, blah, blah", "123:30:22", "20:c9:30", "30:293:088", "20:40:10:30", "20:39:20test", "test20:39:20", "22:22"},
			wantErr:  errors.New("bro, your string is not in the correct format. Try 'HH:MM:SS'. only use numbers and don't forget those starting zeros!"),
		},
		{
			name:     "throws error if hour is above 23",
			testStrs: []string{"55:20:20"},
			wantErr:  fmt.Errorf("invalid hour"),
		},
		{
			name:     "throws error if minute is above 59",
			testStrs: []string{"15:89:20"},
			wantErr:  fmt.Errorf("invalid minute"),
		},
		{
			name:     "throws error if second is above 59",
			testStrs: []string{"15:20:90"},
			wantErr:  fmt.Errorf("invalid second"),
		},
	}
	for _, test := range tests {
		for _, testStr := range test.testStrs {
			result, err := TimeFromString(testStr)
			if err != nil {
				if errors.Is(err, test.wantErr) {
					t.Errorf("Incorrect error thrown. want: %v, got: %v", test.wantErr, err)
				}
			}
			if test.wantErr == nil {
				if test.result != result {
					t.Errorf("correct result not found. wanted: %v, got: %v", test.result, result)
				}
			}
		}
	}
}
