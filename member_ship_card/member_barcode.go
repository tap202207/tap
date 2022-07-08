package membershipcard

import "errors"

type MemberBarcode struct {
	value string
}

func NewMemberBarcode(value string) (*MemberBarcode, error) {
	if len(value) == 16 {
		return &MemberBarcode{
			value: value,
		}, nil
	}
	return nil, errors.New(
		"MemberBarcode length must be 16",
	)
}
