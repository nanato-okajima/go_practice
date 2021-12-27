package calendar

import (
	"errors"
)

type Date struct {
	// change to lowercase
	year  int
	month int
	day   int
}

// Set Getter Method using a pointer receiver type
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

// Change Value Receiver to Pointer Receiver
func (d *Date) SetYear(year int) error {
	if year < 1 || year > 2100 {
		return errors.New("Invalid year")
	}
	// change to lowercase
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	// change to lowercase
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}
	// change to lowercase
	d.day = day
	return nil
}
