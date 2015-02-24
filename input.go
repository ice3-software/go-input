package goinput

//
// A Filter. This is a function which transforms a value and returns the
// results.
//
type Filter func(value interface{}) interface{}

//
// A Validator. Should return an error describing a failure in validation in
// one occurs.
//
type Validator interface {
	Validate(value interface{}) error
}

//
// A type for a validation error. Validation errors can have child
// ValidationErrors. In this way, a tree of validation errors can be built
// up, with the leaf nodes having `Error` string values.
//
type ValidationError struct {
	Errors []error
	Children map[string]*ValidationError
}

//
// Creates a new validation error with either the given children or an empty
// map if nil.
//
func NewValidationError(errs []error, children map[string]*ValidationError) *ValidationError{

	if children == nil {
		children = make(map[string]*ValidationError)
	}
	if errs == nil {
		errs = make([]error, 0)
	}
	return &ValidationError{
		Errors: errs,
		Children: children,
	}
}

//
// Checks whether the map contains any validation errors - its a bit more
// semmantically pleasing than checking the map's length directly
//
func (errs ValidationError) HasErrors() bool {
	return len(errs.Errors) == 0 || len(errs.Children) == 0
}

//
// An interface for a group of input values
//
type Input interface {

	//
	// Validates and filters the input group. Validation errors are bubbled up
	// in the ValidationErrors return value and the filtered input group is
	// returned in the Input return value.
	//
	FilterAndValidate() (Input, *ValidationError)

}

//
// A generic Input struct that allows you to build a filterable / validatable
// input.
//
type BasicInput struct {
	Value                 interface{}
	BreaksValidationChain bool
	Validators            []Validator
	Filters               []Filter
}

//
// Runs through the Input's filters and then validators
//
func (input BasicInput) FilterAndValidate() (Input, *ValidationError) {

	errors := NewValidationError(nil, nil)

	for _, filter := range input.Filters {
		input.Value = filter(input.Value)
	}

	for _, validator := range input.Validators {
		if valErr := validator.Validate(input.Value); valErr != nil {
			errors.Errors = append(errors.Errors, valErr)
			if input.BreaksValidationChain {
				break
			}
		}
	}
	return input, errors
}

//
// A basic group of related input values
//
type BasicInputGroup map[string]BasicInput

//
// Sequentially filters and validates all the inputs in this group. Does not
// break the chain of filtering / validation if any of the inputs fail validation.
//
func (ig BasicInputGroup) FilterAndValidate() (filtered Input, errs *ValidationError) {

	errs = NewValidationError(nil, nil)
	filteredGroup := BasicInputGroup{}

	for fieldName, input := range ig {
		filteredInput, valErrs := input.FilterAndValidate()
		if valErrs.HasErrors() {
			errs.Children[fieldName] = valErrs
		}
		filteredGroup[fieldName] = filteredInput.(BasicInput)
	}

	filtered = filteredGroup
	return
}

//
// Gets a value from an input in the group
//
func (ig BasicInputGroup) Value(fieldName string) interface{} {

	if input, exists := ig[fieldName]; exists {
		return input.Value
	} else {
		panic("Key does not exist")
	}
}
