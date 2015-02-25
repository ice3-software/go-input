package goinput

import (
	"errors"
	"testing"
)

func TestNewValidationError(t *testing.T) {

	errors := make([]error, 0)
	children := make(map[string]*ValidationError)
	valErr := NewValidationError(errors, children)

	if valErr == nil {
		t.Errorf("Validation error is nil")
	}
	if valErr.Children == nil {
		t.Errorf("Error not initialised with children")
	}
	if valErr.Errors == nil {
		t.Errorf("Error not initialised with strings")
	}

}

func TestNewValidationErrorCreatesEmptyCollections(t *testing.T) {

	valErr := NewValidationError(nil, nil)

	if valErr.Children == nil {
		t.Errorf("Children not initialised")
	}
	if valErr.Errors == nil {
		t.Errorf("Errors not initialised")
	}

}

func TestValidaionErrorString(t *testing.T) {

	valErr := NewValidationError([]error{
		errors.New("Error1"),
		errors.New("Error2"),
		errors.New("Error3"),
	}, map[string]*ValidationError{
		"Child1": NewValidationError([]error{errors.New("Error1")}, nil),
		"Child2": NewValidationError([]error{errors.New("Error1")}, nil),
		"Child3": NewValidationError([]error{errors.New("Error1")}, nil),
	})
	expErrStr := `["Error1" "Error2" "Error3"], map["Child1":"[\"Error1\"], map[]" "Child2":"[\"Error1\"], map[]" "Child3":"[\"Error1\"], map[]"]`

	if errStr := valErr.Error(); errStr != expErrStr{
		t.Errorf("Expected `%q`, for `%q`", expErrStr, errStr)
	}
}
