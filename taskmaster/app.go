package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"log"
)

type Task struct {
	task, status string
}

func newTask( task string, status string) *Task {
	p := Task{task : task}
	p.status = "done"
	return &p
}
	
func main(){
	
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	
	var file := pwd + "/data/db.json"

	dat, err := os.Readfile(file)
	if err != nil{
		if os.IsNotExist(err){
			f, err2 := os.Create(file)
			if err2 != nil {
				log.Fatal(err2)
			}
		} else {
			log.Fatal(err)
	}



	fmt.Printf("Initializing database....\n...\n..\n")
	

	fmt.Printf("Heyoo! Welcome to Taskmaster 2000\n...")
	fmt.Printf("What would you like to do?\n")

	for {

		fmt.Printf("Add Tasks 		   : Press 1 and enter\n")
		fmt.Printf("List Tasks 		   : Press 2 and enter\n")
		fmt.Printf("Mark tasks as done : Press 3 and enter\n")
		fmt.Printf("Delete tasks	   : Press 4 and enter\n")
		fmt.Printf("To quit			   : Press q and enter\n")

	}

}
