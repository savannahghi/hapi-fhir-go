
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ResourceVersionPolicy is documented here http://hl7.org/fhir/ValueSet/versioning-policy
type ResourceVersionPolicy int

const (
	ResourceVersionPolicyNoVersion ResourceVersionPolicy = iota
	ResourceVersionPolicyVersioned
	ResourceVersionPolicyVersionedUpdate
)

func (code ResourceVersionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResourceVersionPolicy) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ResourceVersionPolicy code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "no-version":
		*code = ResourceVersionPolicyNoVersion
	case "versioned":
		*code = ResourceVersionPolicyVersioned
	case "versioned-update":
		*code = ResourceVersionPolicyVersionedUpdate
	default:
		return fmt.Errorf("unknown ResourceVersionPolicy code `%s`", s)
	}
	return nil
}
func (code ResourceVersionPolicy) String() string {
	return code.Code()
}
func (code ResourceVersionPolicy) Code() string {
	switch code {
	case ResourceVersionPolicyNoVersion:
		return "no-version"
	case ResourceVersionPolicyVersioned:
		return "versioned"
	case ResourceVersionPolicyVersionedUpdate:
		return "versioned-update"
	}
	return "<unknown>"
}
func (code ResourceVersionPolicy) Display() string {
	switch code {
	case ResourceVersionPolicyNoVersion:
		return "No VersionId Support"
	case ResourceVersionPolicyVersioned:
		return "Versioned"
	case ResourceVersionPolicyVersionedUpdate:
		return "VersionId tracked fully"
	}
	return "<unknown>"
}
func (code ResourceVersionPolicy) Definition() string {
	switch code {
	case ResourceVersionPolicyNoVersion:
		return "VersionId meta-property is not supported (server) or used (client)."
	case ResourceVersionPolicyVersioned:
		return "VersionId meta-property is supported (server) or used (client)."
	case ResourceVersionPolicyVersionedUpdate:
		return "VersionId must be correct for updates (server) or will be specified (If-match header) for updates (client)."
	}
	return "<unknown>"
}
