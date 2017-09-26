package jsonpath

import "errors"

// Expression represents a compiled jsonpath expression ready to be evaluated
type Expression struct {
}

// FromString parse the jsonpath expression passed as argument
func FromString(jsonpath string) *Expression {
	return &Expression{}
}

// Evaluate against the json passed as string
//
// Return a json array containing the results
func (e *Expression) Evaluate(json string) ([]interface{}, error) {
	return nil, errors.New("not implemented")
}

// MustEvaluate has the same behaviour as Evaluate but panics in case of error
func (e *Expression) MustEvaluate(json string) []interface{} {
	return nil
}

// EvaluateJSON against a the json passed as a parsed structure e.g. obtained by json.Unmarshal
//
// Return a json array containing the results
func (e *Expression) EvaluateJSON(json interface{}) ([]interface{}, error) {
	return nil, errors.New("not implemented")
}

// MustEvaluateJSON against a the json passed as a parsed structure e.g. obtained by json.Unmarshal
// Return a json array containing the results
func (e *Expression) MustEvaluateJSON(json interface{}) []interface{} {
	return nil
}
