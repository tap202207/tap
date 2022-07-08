package membershipcard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	membershipcard "github.com/tap202207/tap/member_ship_card"
)

func TestNewCompanyCodeNG(t *testing.T) {
	_, err := membershipcard.NewCompanyCode(
		"12345",
	)
	assert.NotEqual(t, nil, err)
}

func TestNewCompanyCodeOK(t *testing.T) {
	code, err := membershipcard.NewCompanyCode(
		"1234",
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1234", code.Value())
}
