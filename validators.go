package goinput

import (
	"errors"
	"fmt"
	"regexp"
)


//
// @note Requires string
//
type NotEmptyValidator struct{}

func (v NotEmptyValidator) Validate(value interface{}) (err error) {

	if len(value.(string)) == 0 {
		err = errors.New("String is empty")
	}

	return
}


//
// @note Requires string
//
type RegexValidator struct {
	Reg string
}

func (v RegexValidator) Validate(value interface{}) (err error) {

	if matched, _ := regexp.Match(value.(string), []byte(v.Reg)); !matched {
		err = errors.New(fmt.Sprintf("'%q' does not match the valid regex format %q", value, v.Reg))
	}

	return
}

func NewUUIDValidator() RegexValidator {
	return RegexValidator{
		Reg: "^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$",
	}
}


//
// @note Requires string
//
type LengthValidator struct {
	Min int
	Max int
}

func (v LengthValidator) Validate(value interface{}) (err error) {

	if leng := len(value.(string)); leng < v.Min || leng > v.Max {
		err = errors.New("String is not within bounds")
	}
	return
}


//
// @note Requires string
//
type InArrayValidator struct {
	Array []string
}

func (v InArrayValidator) Validate(value interface{}) (err error) {

	if !InSlice(v.Array, value.(string)) {
		err = errors.New("Value is not in the given array")
	}
	return
}


//
// Helper function for detecting whether a string is in another string
// of slices
//
func InSlice(slice []string, value string) (in bool) {
	for _, arrVal := range slice {
		if arrVal == value {
			in = true
			break
		}
	}
	return
}
