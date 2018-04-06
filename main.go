package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/shmurakami/blog_header_generator/engine"
)

func generate() error {
	// TODO get output path from option

	jekyll := engine.New()

	// actually want to ask/set dynamically to support another blog template engine in future
	// seems golang does not support dynamic value get/set without reflection
	// for now support only jekyll

	var err error
	if jekyll.Title, err = ask("Title", jekyll.Title); err != nil {
		return err
	}
	if jekyll.Description, err = ask("Description", jekyll.Description); err != nil {
		return err
	}
	if jekyll.Date, err = ask("Date", jekyll.Date); err != nil {
		return err
	}
	if jekyll.Filename, err = ask("Filename", jekyll.Filename); err != nil {
		return err
	}

	err = jekyll.Output("./_posts")
	if err != nil {
		return err
	}

	fmt.Println("File generated")
	return nil
}

func main() {
	err := generate()
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func ask(t, initial string) (string, error) {
	fmt.Printf("%s? (%s): ", t, initial)
	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return "", err
	}
	line := strings.TrimSpace(string(b))

	if line == "" {
		return initial, nil
	}
	return line, nil
}
