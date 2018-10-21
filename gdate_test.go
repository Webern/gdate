// gdate is Copyright (c) 2018 by Matthew James Briggs, MIT License

package gdate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGdate(t *testing.T) {
	date := Date{}
	assert.NotNil(t, &date)
}
