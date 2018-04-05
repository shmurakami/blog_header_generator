package main

import (
	"fmt"
	"log"

	"github.com/shmurakami/blog_header_generator/engine"
)

func main() {
	// TODO get output path from option

	jekyll := engine.New()

	// actually want to ask/set dynamically to support another blog template engine in future
	// seems golang does not support dynamic value get/set without reflection
	// for now support only jekyll

	jekyll.Title = ask("Title", jekyll.Title)
	jekyll.Description = ask("Description", jekyll.Description)
	jekyll.Date = ask("Date", jekyll.Date)
	jekyll.Filename = ask("Filename", jekyll.Filename)

	err := jekyll.Output("./_posts")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File generated")
}

func ask(t, initial string) string {
	stdin := ""
	fmt.Printf("%s? (%s): ", t, initial)
	fmt.Scanln(&stdin)

	if stdin == "" {
		return initial
	}
	return stdin
}
