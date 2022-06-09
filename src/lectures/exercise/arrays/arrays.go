//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name 	string   
	Price 	int32	 
}

func printStats(list [4]Product){
	var totalCost, totalItemsCount int32
	for i := 0; i < len(list); i++ {
		totalCost += list[i].Price
		if list[i].Name != ""{
			totalItemsCount++
		}		
	}

	fmt.Println("Last item on the list:", list[totalItemsCount-1])
	fmt.Println("Total items:", totalItemsCount)
	fmt.Println("Total cost:", totalCost)
}

func main() {
	var myShoppingList [4]Product
	myShoppingList[0] = Product{Name: "Eggs", Price: 55}
	myShoppingList[1] = Product{Name: "Carrots", Price: 35}
	myShoppingList[2] = Product{Name: "Bread", Price: 65}	
	printStats(myShoppingList)

	myShoppingList[3] = Product{Name: "Salt", Price: 25}
	printStats(myShoppingList)
}
