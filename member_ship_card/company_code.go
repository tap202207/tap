package membershipcard

import "errors"

type CompanyCode struct {
	value string
}

func NewCompanyCode(value string) (*CompanyCode, error) {
	for _, v := range []bool{
		len(value) == 4,
	} {
		if !v {
			return nil, errors.New(
				"company code assertion error",
			)
		}
	}
	return &CompanyCode{
		value: value,
	}, nil
}

func (x CompanyCode) Value() string {
	return x.value
}
