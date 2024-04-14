package model

import "errors"

type Category struct {
	Name string
	Time int
}

type CategoryStorage map[string]int

func (cs CategoryStorage) Add(category string, time int) {
	_, ok := cs[category]
	if !ok {
		cs[category] += time

	} else {
		cs[category] = time
	}
}

func (cs CategoryStorage) Subtract(categoty string, time int) error {

	val, ok := cs[categoty]
	if !ok {
		return errors.New("no category in storage")
	}
	if time > val {
		return errors.New("trying to subtract more time than stored")
	}
	cs[categoty] -= time
	return nil
}
