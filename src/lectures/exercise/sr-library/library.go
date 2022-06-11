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

func checkOutBook(lib *Library, title Title, mem *Member) bool{
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
	mem.books[title] = LendAudit{
		lastCheckedOut: time.Now()}
	return true
}

func checkInBook(lib *Library, title Title, mem *Member) bool{
	// only allow check in when it is a known library book
	bookEntry, bookInLib := lib.books[title]
	if !bookInLib {
		fmt.Println("book does not belong to library!")
		return false
	}
	
	// check if book was borrowed by this member; if not don't accept book#
	audit, bookLendedByMember := mem.books[title]
	if !bookLendedByMember{
		fmt.Println("book was not borrowed by this member! No check-in possible!")
		return false
	}

	// decrease the lended bookEntry counter and put the LendAudit timestamp
	bookEntry.lended -= 1;
	lib.books[title] = bookEntry

	// mem.books[title].lastReturned = time.Now() // Not like this!
	audit.lastReturned = time.Now();
	mem.books[title] = audit

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
	if len(mem.books) > 0 {
		fmt.Println("Name of Member:", mem.name)
		fmt.Println("borrowed books:")
	}

	for title, entry := range mem.books {
		// check whether a book was not returned yet
		var returnTime string
		if entry.lastReturned.IsZero() {
			returnTime = "[not returned yet]"
		} else{
			returnTime = entry.lastReturned.String()
		}
		fmt.Println("borrowed", title, "on", entry.lastCheckedOut, "last returned:", returnTime)
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
		// Note: books Map requires initialization with make!
		// otherwise "panic: assignment to entry in nil map!" possible
		var lendHistory = make(map[Title]LendAudit, 4) 
		memberMap[members[i]] = Member{
			name: members[i],
			books: lendHistory}	
	}

	//  - Check out a book
	member := myLibrary.members[members[1]]
	
	checkedOut := checkOutBook(&myLibrary, titles[0], &member) // Note: use of references!

	if checkedOut {
		fmt.Println("\nCheck out a book. New Status:")
		printLibraryBooks(&myLibrary)
		printMemberAudits(&myLibrary)
	}

	//  - Check in a book
	returned := checkInBook(&myLibrary, titles[1], &member)

	if returned {
		fmt.Println("\nCheck in a book. New Status:")
		printLibraryBooks(&myLibrary)
		printMemberAudits(&myLibrary)
	}

	//  - Check in a book again, with success
	returned = checkInBook(&myLibrary, titles[0], &member)
	if returned {
		fmt.Println("\nCheck in a book. New Status:")
		printLibraryBooks(&myLibrary)
		printMemberAudits(&myLibrary)
	}

}
