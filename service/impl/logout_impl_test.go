package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogoutService(t *testing.T) {
	assertion := assert.New(t)

	createdLogoutService := NewLogoutService()

	assertion.NotNil(createdLogoutService)
	assertion.IsType(new(logoutService), createdLogoutService)
}