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

func TestDate_Parse(t *testing.T) {
	date := NewDate(2001, 6, 2)
	assert.Equal(t, 2001, date.Year())
	assert.Equal(t, 6, date.Month())
	assert.Equal(t, 2, date.Day())
	err := date.Parse("1999-02-01")
	assert.NoError(t, err)
	assert.Equal(t, 1999, date.Year())
	assert.Equal(t, 2, date.Month())
	assert.Equal(t, 1, date.Day())
}
