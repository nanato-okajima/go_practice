package calendar

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid Title")
	}
	e.title = title
	return nil
}

func (e *Event) Title() {
	fmt.Println(e.title)
}
