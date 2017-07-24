package arangolite

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func assertTrue(t *testing.T, a interface{}, message string) {
	assertEqual(t, true, a, message)
}

func TestIsErrNotFound(t *testing.T) {
	err := errors.Wrap(errors.New("Error message"), "the database execution returned an error")
	err = withErrorNum(err, 1202)
	err = withStatusCode(err, 403)
	assertTrue(t, IsErrNotFound(err), "Error code does not match!")
}

func TestIsErrUnique(t *testing.T) {
	err := errors.Wrap(errors.New("Error message"), "the database execution returned an error")
	err = withErrorNum(err, 1210)
	err = withStatusCode(err, 403)
	assertTrue(t, IsErrUnique(err), "Error code does not match!")
}

func TestGetErrorNum(t *testing.T) {
	err := errors.Wrap(errors.New("Error message"), "the database execution returned an error")
	err = withErrorNum(err, 1207)
	err = withStatusCode(err, 409)
	errNum, ok := GetErrorNum(err)
	assertEqual(t, ok, true, "Err in not of type numberedError!")
	assertEqual(t, errNum, 1207, "Error code does not match!")
}

func TestHowToGetErrorNumber(t *testing.T) {
	err := errors.Wrap(errors.New("Error message"), "the database execution returned an error")
	err = withErrorNum(err, 1207)
	err = withStatusCode(err, 409)
	assertTrue(t, HasErrorNum(err, 1207), "Error code does not match!")
	// How to access error number here....
}
