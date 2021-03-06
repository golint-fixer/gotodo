package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/marthjod/gotodo/model/todotxt"
	"github.com/marthjod/gotodo/read"
)

func main() {
	var (
		todo         todotxt.TodoTxt
		js           []byte
		err          error
		todoFilename = flag.String("t", "todo.txt", "todo.txt file to use")
	)

	flag.Parse()

	if todo, err = read.Read(*todoFilename); err != nil {
		log.Fatal(err)
	}

	js, err = json.MarshalIndent(todo.Entries, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", js)

}
