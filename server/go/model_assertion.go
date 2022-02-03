/*
 * Project X
 *
 * OpenAPI definition for project X endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Assertion struct {

	// ID
	Id string `json:"id,omitempty"`

	OperationName string `json:"operationName,omitempty"`

	Duration string `json:"duration,omitempty"`

	NumOfSPans int32 `json:"numOfSPans,omitempty"`

	Attributes []string `json:"attributes,omitempty"`
}

// AssertAssertionRequired checks if the required fields are not zero-ed
func AssertAssertionRequired(obj Assertion) error {
	return nil
}

// AssertRecurseAssertionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Assertion (e.g. [][]Assertion), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssertionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssertion, ok := obj.(Assertion)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssertionRequired(aAssertion)
	})
}
