//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:

package main

import "fmt"

type status int

const (
	Online      status = 0
	Offline     status = 1
	Maintenance status = 2
	Retired     status = 3
)

func (s status)String() string {
	switch s {
	case Online:
		return "Online"
	case Offline:
		return "Offline"
	case Maintenance:
		return "Maintenance"
	case Retired:
		return "Retired"
	default:
		return "error"
	}
}

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
func printServerStatus(serversStats map[string]status) {
	var totalServers int
	serverStatusCount := make(map[string]int)
	for _, stat := range serversStats {
		totalServers += 1
		serverStatusCount[stat.String()] += 1
	}
	fmt.Println("There are", totalServers, "servers")
	for stat, count := range serverStatusCount {
		fmt.Println(count, "servers are", stat)
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	
	//* Create a map using the server names as the key and the server status
	//*  as the value
	//* Set all of the server statuses to `Online` when creating the map
	serverStatus := make(map[string]status)
	for _, server := range servers {
		serverStatus[server] = Online
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	printServerStatus(serverStatus)
	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverStatus["aiur"] = Offline
	//  - call display server info function
	printServerStatus(serverStatus)
	//  - change server status of all servers to `Maintenance`
	for server := range serverStatus {
		serverStatus[server] = Maintenance
	}
	//  - call display server info function
	printServerStatus(serverStatus)
}
