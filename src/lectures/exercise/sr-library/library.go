//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Title string
type Name string

//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
type LendAudit struct {
	lastCheckedOut	time.Time
	lastReturned	time.Time
}
type Member struct{
	name	Name
	books	map[Title]LendAudit
}

type BookEntry struct{
	totalAmount int
	lended		int 
}

type Library struct{
	members map[Name] Member
	books   map[Title] BookEntry
}

func checkOut(lib *Library, title Title, mem *Member) bool{
	// also make sure the book is part of the library 
	bookEntry, bookInLib := lib.books[title]
	if !bookInLib {
		fmt.Println("book not in library!")
		return false
	}
	
	// and we have enough spare to lend
	if bookEntry.lended == bookEntry.totalAmount {
		fmt.Println("all books lended!")
		return false
	}
	
	// Update library
	bookEntry.lended += 1
	lib.books[title] = bookEntry // required to update bookEntry!

	// Update member info
	/* ATTENTION: leads to panic: assignment to entry in nil map!
	mem.books[title] = LendAudit{
		lastCheckedOut: time.Now()}
		FIXME!!!
	*/
	return true
}

func printLibraryBooks(lib *Library){
	for title, book := range lib.books{
		fmt.Println("Book:", title)
		fmt.Println(" total:", book.totalAmount, "lended:", book.lended)
	}
}

func printMemberAudits(lib *Library){
	for _, member := range lib.members{
		printSingleMemberAudit(&member)
	}
}

func printSingleMemberAudit(mem *Member){
	fmt.Println("Name of Member:", mem.name)
	fmt.Println("borrowed books...")
	for title, entry := range mem.books {
		fmt.Println("borrowed", title, "on", entry.lastCheckedOut)
	}
}


func main() {
	//* Perform the following:
	//  - Add at least 4 books and at least 3 members to the library	
	bookMap := make(map[Title]BookEntry, 4)
	memberMap := make(map[Name]Member, 3)
	myLibrary := Library{
		books: bookMap,
		members: memberMap,
	}
	rand.Seed(time.Now().UnixNano()) // addition, to change amount of total books, from 1-4

	titles := []Title{"Get Programming with Go", "Programming Typscript", "Neverending story", "Yet another Go book"}
	members := []Name{"Wolfgang", "Jacob", "Sunny"}

	for i := range(titles) {
		bookMap[titles[i]] = BookEntry{
			totalAmount: rand.Intn(3) + 1}	
	}
	for i := range(members) {
		memberMap[members[i]] = Member{
			name: members[i]}	
	}

	//  - Check out a book
	member := myLibrary.members[members[1]]
	
	checkedOut := checkOut(&myLibrary, titles[0], &member) // Note: use of references!

	if checkedOut {
		fmt.Println("\nCheck out a book. New Status:")
		printLibraryBooks(&myLibrary)
		printMemberAudits(&myLibrary)
	}

	//  - Check in a book

	//  - Print out initial library information, and after each change
	//* There must only ever be one copy of the library in memory at any time	


}
