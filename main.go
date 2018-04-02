package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	//h := header{
	//	title:       "foo",
	//	description: "brah brah",
	//	date:        "2018-04-02",
	//	filename:    "test",
	//}

	h := map[string]string{
		"title":       "foo",
		"description": "brah brah",
		"date":        "2018-04-02",
		"filename":    "test",
	}

	year, month, day := time.Now().Date()
	m := fmt.Sprintf("%02d", month)
	d := fmt.Sprintf("%02d", day)

	fmt.Printf("Today is %d-%s-%s\n", year, m, d)

	scans := map[string]string{}

	var stdin string
	for k, v := range h {
		stdin = ""
		fmt.Printf("%s? (%s): ", k, v)
		fmt.Scan(&stdin)
		if stdin == "" {
			stdin = v
		}
		fmt.Println(stdin)
		scans[k] = stdin
	}

	fmt.Println(scans)

	makeFile(scans)
}

func makeFile(headers map[string]string) error {
	directory := "./_posts"
	filename := headers["filename"]
	output := fmt.Sprintf("%s/%s", directory, filename)

	// check if file exsists or not
	_, error := ioutil.ReadFile(output)
	if error == nil {
		// how to return new error?
		return nil
	}

	h := fmt.Sprintf("---\nlayout: post\nposted: %s\ntitle: %s\ndescription: %s\n---\n\n",
		headers["date"],
		headers["title"],
		headers["description"],
	)
	vec := []byte(h)

	err := ioutil.WriteFile(output, vec, 0644)
	if err != nil {
		return err
	}

	return nil
}
