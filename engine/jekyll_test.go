package engine

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	now = time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC)

	result := today()
	if result != "2018-01-02" {
		t.Fatal()
	}
}

func TestFilename(t *testing.T) {
	now = time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC)

	j := New()
	j.Filename = "test"
	result := j.filename()
	if result != "2018-01-02-test.md" {
		t.Fatal("failed auto prefix failename")
	}

	j.Filename = "2018-01-02-manual"
	result = j.filename()
	if result != "2018-01-02-manual.md" {
		t.Fatal("failed manual prefix filename")
	}
}
