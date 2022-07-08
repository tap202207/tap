package membershipcard

import "errors"

type Pin struct {
	value string
}

func NewPin(value string) (*Pin, error) {
	if len(value) == 4 {
		return &Pin{
			value: value,
		}, nil
	}
	return nil, errors.New(
		"PIN length must be four",
	)
}
