//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
func printStatus(servers map[string]int){
	var serversCount int
	statusCount := make(map[int]int)
	//serversCount := len(servers)

	for _, value := range servers {			
		statusCount[value]++
		serversCount++	
	}
	fmt.Println("Number of servers:", serversCount)
	fmt.Println(statusCount[Online], "servers are online")
	fmt.Println(statusCount[Offline], "servers are offline")
	fmt.Println(statusCount[Maintenance], "servers are undergoing maintenance")
	fmt.Println(statusCount[Retired], "servers are retired")

}

func setStatus(servers map[string]int, status int){
	for key, _ := range servers {		
		servers[key] = status
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	//* Set all of the server statuses to `Online` when creating the map
	serversMap := make(map[string]int)
	for name := range servers{
		serversMap[servers[name]] = Online
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	printStatus(serversMap)
	//  - change server status of `darkstar` to `Retired`
	serversMap["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serversMap["aiur"] = Offline
	//  - call display server info function
	printStatus(serversMap)
	//  - change server status of all servers to `Maintenance`
	//  - call display server info function
	setStatus( serversMap, Maintenance)
	printStatus(serversMap)
}
