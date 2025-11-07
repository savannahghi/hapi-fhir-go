package search

import (
	"fmt"
	"net/url"
	"strings"
)

// BuildQueryPath constructs the full FHIR search query string based on the
// InlineSearchInput fields.
//
// It generates a query following the FHIR RESTful search specification,
// including support for _include and _revinclude parameters. All query
// components that may contain reserved characters (such as ":" in include
// expressions) are URL-escaped to ensure safe transmission.
//
// The generated path has the format:
//   <BaseResource>/?_id=<ResourceID>&_<type>=<targetResource>:<searchField>...
//
// Example:
//   BaseResource: "ServiceRequest"
//   ResourceID:   "123"
//   Params: [
//       { Type: "include",    TargetResource: "ServiceRequest", SearchField: "performer" },
//       { Type: "revinclude", TargetResource: "Observation",    SearchField: "subject" },
//   ]
//
// Produces:
//   ServiceRequest/?_id=123&_include=ServiceRequest%3Aperformer&_revinclude=Observation%3Asubject
//
// Returns the fully assembled query string. The function does not perform
// validation; callers should ensure required fields (such as BaseResource
// and ResourceID) are populated.
func (i *InlineSearchInput) BuildQueryPath() string {
	var builder strings.Builder

	builder.WriteString(i.BaseResource)
	builder.WriteString("/?_id=")
	builder.WriteString(url.QueryEscape(i.ResourceID))

	for _, item := range i.Params {
		builder.WriteString("&_")
		builder.WriteString(string(item.Type))
		builder.WriteString("=")

		value := fmt.Sprintf("%s:%s", item.TargetResource, item.SearchField)
		builder.WriteString(url.QueryEscape(value))

	}

	return builder.String()
}

// BuildPath generates the final FHIR search query path by delegating to the
// InlineSearchInput's BuildQueryPath method.
//
// This function acts as a convenience wrapper, allowing callers to work
// directly with the SearchBuilder API without accessing the inner
// InlineSearchInput struct. It returns the fully constructed FHIR query string,
// including the base resource, resource ID, and any configured _include or
// _revinclude parameters.
//
// Example:
//   sb := NewSearchBuilder("ServiceRequest", "123")
//   sb.Include("ServiceRequest", "performer")
//   path := sb.BuildPath()
//   // path == "ServiceRequest/?_id=123&_include=ServiceRequest%3Aperformer"
//
// Returns the constructed search path as a string.
func (s *SearchBuilder) BuildPath() string {
	return s.Input.BuildQueryPath()
}

// Build finalizes the SearchBuilder and returns a pointer to its underlying
// InlineSearchInput structure.
//
// This method allows callers to retrieve the fully populated search input,
// which can then be passed to other components—such as HTTP clients or
// higher-level FHIR query executors. It does not generate the URL itself;
// callers can invoke BuildQueryPath() on the returned InlineSearchInput to
// produce the actual query string.
//
// Example:
//   sb := NewSearchBuilder("ServiceRequest", "123")
//   sb.Include("ServiceRequest", "performer")
//   input := sb.Build()
//   query := input.BuildQueryPath()
//
// Returns a pointer to the constructed InlineSearchInput
func (s *SearchBuilder) Build() *InlineSearchInput {
	return &s.Input
}