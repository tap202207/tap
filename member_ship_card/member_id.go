package membershipcard

import "errors"

type MemberId struct {
	value string
}

func NewMemberId(value string) (*MemberId, error) {
	if len(value) == 64 {
		return &MemberId{
			value: value,
		}, nil
	}
	return nil, errors.New(
		"MemberId length must be 64",
	)
}
