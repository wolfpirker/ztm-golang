//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type SecurityTag bool

//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
type Item struct {
	name  string
	tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activateSecurityTag(t *SecurityTag){
	*t = true
}

func deactivateSecurityTag(t *SecurityTag){
	*t = false
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items[] Item){
	for i := range(items){
		items[i].tag = false
	}
}

func printSlice(items []Item){
	for i := range items {
		fmt.Println("Item", i, items[i].name, "Status", items[i].tag)
	}
}


func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	//  - Deactivate any one security tag in the array/slice
	//  - Call the checkout() function to deactivate all tags
	//  - Print out the array/slice after each change
	
	//items := make([]Item, 4) // must'nt be preallocated
	var items = []Item{}
	itemNames := []string{"Shoes", "Pants", "T-Shirt", "Jacket"}

	for i := range(itemNames){
		items = append(items, Item{
			name: itemNames[i],
			tag: true,
		})
	}
	printSlice(items)

	deactivateSecurityTag(&items[0].tag)
	printSlice(items)

	checkout(items)
	printSlice(items)
}
