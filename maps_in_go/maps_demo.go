package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var m map[string]Vertex

func main() {
	//maps are like slices,
	//but with keys instead of indexes
	//keys can be of any type that is comparable

	//by default maps are nil
	//a nil map has no keys and cannot be assigned to

	//make function can be used to creat a map

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{40, -74}

	fmt.Println(m["Bell Labs"])

	//maps can be literals too
	var m1 = map[string]Vertex{
		"Bell Labs": Vertex{40, -74},
		"Google":    Vertex{37, -122},
	}

	fmt.Println(m1)

	//the value data type 
	// can be omitted from the elements of a
	// map literal 

	var m2 = map[string]Vertex{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}	
	fmt.Println(m2)

	//insert and item in a map
	m2["Amazon"] = Vertex{47, -122}
	fmt.Println(m2)

	//delete an item from a map 
	delete(m2, "Google")
	fmt.Println(m2)

	//update an item in a map
	m2["Bell Labs"] = Vertex{40, -75}
	fmt.Println(m2)

	//check if a key is present in a map
	_, ok := m2["Google"]
	fmt.Println("Is Google present in m2 ? ", ok)

	//iterate over a map using range
	for key, value := range m2 {
		fmt.Printf("%s: %v\n", key, value)
	}
}