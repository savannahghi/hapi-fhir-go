
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// QuestionnaireItemOperator is documented here http://hl7.org/fhir/ValueSet/questionnaire-enable-operator
type QuestionnaireItemOperator int

const (
	QuestionnaireItemOperatorExists QuestionnaireItemOperator = iota
	QuestionnaireItemOperatorEquals
	QuestionnaireItemOperatorNotEquals
	QuestionnaireItemOperatorGreaterThan
	QuestionnaireItemOperatorLessThan
	QuestionnaireItemOperatorGreaterOrEquals
	QuestionnaireItemOperatorLessOrEquals
)

func (code QuestionnaireItemOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuestionnaireItemOperator) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal QuestionnaireItemOperator code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "exists":
		*code = QuestionnaireItemOperatorExists
	case "=":
		*code = QuestionnaireItemOperatorEquals
	case "!=":
		*code = QuestionnaireItemOperatorNotEquals
	case ">":
		*code = QuestionnaireItemOperatorGreaterThan
	case "<":
		*code = QuestionnaireItemOperatorLessThan
	case ">=":
		*code = QuestionnaireItemOperatorGreaterOrEquals
	case "<=":
		*code = QuestionnaireItemOperatorLessOrEquals
	default:
		return fmt.Errorf("unknown QuestionnaireItemOperator code `%s`", s)
	}
	return nil
}
func (code QuestionnaireItemOperator) String() string {
	return code.Code()
}
func (code QuestionnaireItemOperator) Code() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "exists"
	case QuestionnaireItemOperatorEquals:
		return "="
	case QuestionnaireItemOperatorNotEquals:
		return "!="
	case QuestionnaireItemOperatorGreaterThan:
		return ">"
	case QuestionnaireItemOperatorLessThan:
		return "<"
	case QuestionnaireItemOperatorGreaterOrEquals:
		return ">="
	case QuestionnaireItemOperatorLessOrEquals:
		return "<="
	}
	return "<unknown>"
}
func (code QuestionnaireItemOperator) Display() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "Exists"
	case QuestionnaireItemOperatorEquals:
		return "Equals"
	case QuestionnaireItemOperatorNotEquals:
		return "Not Equals"
	case QuestionnaireItemOperatorGreaterThan:
		return "Greater Than"
	case QuestionnaireItemOperatorLessThan:
		return "Less Than"
	case QuestionnaireItemOperatorGreaterOrEquals:
		return "Greater or Equals"
	case QuestionnaireItemOperatorLessOrEquals:
		return "Less or Equals"
	}
	return "<unknown>"
}
func (code QuestionnaireItemOperator) Definition() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "True if whether an answer exists is equal to the enableWhen answer (which must be a boolean)."
	case QuestionnaireItemOperatorEquals:
		return "True if whether at least one answer has a value that is equal to the enableWhen answer."
	case QuestionnaireItemOperatorNotEquals:
		return "True if whether at least no answer has a value that is equal to the enableWhen answer."
	case QuestionnaireItemOperatorGreaterThan:
		return "True if whether at least no answer has a value that is greater than the enableWhen answer."
	case QuestionnaireItemOperatorLessThan:
		return "True if whether at least no answer has a value that is less than the enableWhen answer."
	case QuestionnaireItemOperatorGreaterOrEquals:
		return "True if whether at least no answer has a value that is greater or equal to the enableWhen answer."
	case QuestionnaireItemOperatorLessOrEquals:
		return "True if whether at least no answer has a value that is less or equal to the enableWhen answer."
	}
	return "<unknown>"
}
