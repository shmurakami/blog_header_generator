package main

import (
	"fmt"
	"log"

	"./engine"
)

func main() {
	// TODO get output path from option

	jekyll := engine.New()

	// actually want to ask/set dynamically to support another blog template engine in future
	// seems golang does not support dynamic value get/set without reflection
	// for now support only jekyll

	title := ask("Title", jekyll.Title)
	description := ask("Description", jekyll.Description)
	date := ask("Date", jekyll.Date)
	filename := ask("Filename", jekyll.Filename)
	jekyll.Title = title
	jekyll.Description = description
	jekyll.Date = date
	jekyll.Filename = filename

	err := jekyll.Output("./_posts")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File generated")
}

func ask(t, initial string) string {
	stdin := ""
	fmt.Printf("%s?(%s) :", t, initial)
	fmt.Scanln(&stdin)
	fmt.Println("")

	if stdin == "" {
		return initial
	}
	return stdin
}
