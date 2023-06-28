package main

import "fmt"

func main() {
	//declare and initialise map
	myMap := make(map[string]int)
	myMap["apple"] = 1
	myMap["banana"] = 2
	myMap["orange"] = 3
	//Accessig values using keys
	appleValue := myMap["apple"]
	bananaValue := myMap["banana"]
	fmt.Println("Value of apple", appleValue)
	fmt.Println("Value of banana", bananaValue)
	//updating values
	myMap["apple"] = 5
	fmt.Println("Updated value of apple", myMap["apple"])
	//Deleting a keyvalue pair
	delete(myMap, "orange")
	fmt.Println("After deleting orange:", myMap)
	//checking existence of a key
	value, exists := myMap["banana"]
	if exists {
		fmt.Println("Value of banana:", value)
	} else {
		fmt.Println("Banana not found in the map")
	}
	//iterating over the map
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
	//Length of the map
	fmt.Println("Length of the map:",len(myMap))

}