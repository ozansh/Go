package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
    choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    selected map[int]struct{}   // which to-do items are selected
}

func main(){


	fmt.Printf("Initializing database....\n...\n..\n")
	// db function

	fmt.Printf("Heyoo! Welcome to Taskmaster 2000\n...")
	fmt.Printf("What would you like to do?\n")

}
