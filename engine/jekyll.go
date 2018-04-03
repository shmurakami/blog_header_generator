package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Jekyll struct {
	Title       string
	Description string
	Date        string
	Filename    string
}

func New() *Jekyll {
	today := today()
	return &Jekyll{
		Title:       "",
		Description: "",
		Date:        today,
		Filename:    "",
	}
}

func (j *Jekyll) Output(directory string) error {
	filename := j.filename()
	output := fmt.Sprintf("%s/%s", directory, filename)

	// check if file already exsists
	_, error := ioutil.ReadFile(output)
	if error == nil {
		return errors.New("File exists")
	}

	pattern := `---
layout: post
posted: %s
title: %s
description: %s
---

`

	h := fmt.Sprintf(pattern, j.Date, j.Title, j.Description)
	vec := []byte(h)

	err := ioutil.WriteFile(output, vec, 0644)
	if err != nil {
		return err
	}

	return nil
}

func today() string {
	year, month, day := time.Now().Date()
	y := fmt.Sprintf("%04d", year)
	m := fmt.Sprintf("%02d", month)
	d := fmt.Sprintf("%02d", day)
	today := fmt.Sprintf("%s-%s-%s", y, m, d)

	return today
}

func (j *Jekyll) filename() string {
	// jekyll post file must has date prefix
	has := strings.HasPrefix(j.Filename, j.Date)
	if has {
		return j.Filename
	}
	return fmt.Sprintf("%s-%s", j.Date, j.Filename)
}
