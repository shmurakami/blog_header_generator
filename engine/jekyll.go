package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var now = time.Now()

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
	if _, err := time.Parse("2006-01-02", j.Date); err != nil {
		return err
	}

	filename := j.filename()
	output := filepath.Join(directory, filename)

	// check if file already exsists
	_, err := ioutil.ReadFile(output)
	if err == nil {
		return errors.New("File exists")
	}

	format := `---
layout: post
posted: %s
title: %s
description: %s
---

`

	h := fmt.Sprintf(format, j.Date, j.Title, j.Description)
	vec := []byte(h)

	err = ioutil.WriteFile(output, vec, 0644)
	if err != nil {
		return err
	}

	return nil
}

func today() string {
	//year, month, day := time.Now().Date()
	year, month, day := now.Date()
	y := fmt.Sprintf("%04d", year)
	m := fmt.Sprintf("%02d", month)
	d := fmt.Sprintf("%02d", day)
	today := fmt.Sprintf("%s-%s-%s", y, m, d)

	return today
}

func normalize(s string) string {
	return regexp.MustCompile("[^\\p{L}\\d_]+").ReplaceAllString(s, "_")
}

func (j *Jekyll) filename() string {
	if j.Filename == "" {
		j.Filename = normalize(j.Title)
	}

	// jekyll post file must has date prefix
	hasPrefix := strings.HasPrefix(j.Filename, j.Date)
	if hasPrefix == false {
		j.Filename = fmt.Sprintf("%s-%s", j.Date, j.Filename)
	}

	// filename must has markdown extension
	hasExtension := strings.HasSuffix(j.Filename, ".md")
	if hasExtension == false {
		j.Filename = fmt.Sprintf("%s.md", j.Filename)
	}
	return j.Filename
}
