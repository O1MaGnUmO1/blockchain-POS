package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertNotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotNil(t, object, msgAndArgs...)
}

func AssertNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return assert.Nil(t, object, msgAndArgs...)
}

func AsserEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, actual, msgAndArgs...)
}

func AssertNotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEqual(t, expected, actual, msgAndArgs...)
}

func AssertTrue(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	return assert.True(t, value, msgAndArgs)
}

func AssertFalse(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	return assert.False(t, value, msgAndArgs)
}
