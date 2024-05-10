package model

import "errors"

type Category struct {
	Name string
	Time int
}

type CategoryStorage map[string]int

func (cs CategoryStorage) Add(category string, timeToAdd int) {
	cs[category] = cs[category] + timeToAdd
}

func (cs CategoryStorage) Subtract(category string, timeToSubstrct int) error {

	val, ok := cs[category]
	if !ok {
		return errors.New("no category in storage")
	}
	if timeToSubstrct > val {
		return errors.New("trying to subtract more time than stored")
	}
	cs[category] -= timeToSubstrct
	return nil
}

func (cs CategoryStorage) Load(category string) int {
	return cs[category]
}
