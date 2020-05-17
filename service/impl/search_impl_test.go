package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSearchService(t *testing.T) {
	assertion := assert.New(t)

	assertion.NotNil(testSearchS)
	assertion.IsType(new(searchService), testSearchS)
}