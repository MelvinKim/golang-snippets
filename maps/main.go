package main

import (
	"fmt"
	"sort"
)

type user struct {
	name    string
	surname string
}

type car struct {
	name string
	year int
}

func main() {
	users := make(map[string]user)
	users["Melvin"] = user{name: "Melvin", surname: "Kimathi"}
	users["Bill"] = user{name: "Bill", surname: "Kennedy"}

	// keys and values
	for key, val := range users {
		fmt.Printf("Key: %s, value: %s\n", key, val) // pass by value --> copies
	}

	// keys
	for key := range users {
		fmt.Println("key: ", key) // pass by value --> copies
	}

	// declare and intialize the map with values
	cars := map[string]car{
		"Toyota":        {"Toyota", 2020},
		"VW":            {"VW", 2018},
		"Mercedes Benz": {"Benz", 2018},
	}
	for key, val := range cars {
		fmt.Println(key, val)
	}

	// delete a car
	fmt.Println("-------------------------")
	delete(cars, "Toyota")
	fmt.Println("cars: ", cars)

	// find a key
	/*
		- when you try to find a key in a map and it doesn't exist, the zero value of the key is returned
		- not everything can be a key in maps, but everything can be a value
	*/
	u, found := cars["VW"]
	fmt.Println("VW", found, u)

	// sort keys
	var keys []string
	for car := range cars {
		keys = append(keys, car)
	}
	sort.Strings(keys)
	// walk through the sorted keys and pull each value from the map
	for _, key := range keys {
		fmt.Println(key, cars[key])
	}

}

/*
- iterating over a map is randomized
*/
