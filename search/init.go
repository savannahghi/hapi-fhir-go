// Package search provides utilities for constructing FHIR-compliant RESTful
// search queries, including support for the _include and _revinclude
// parameters defined in the HL7® FHIR specification.
//
// The package exposes two main components:
//
//   1. InlineSearchInput
//      Represents the data needed to construct a query path. This includes
//      the base resource being searched, the resource ID, and a list of
//      include/revinclude parameters.
//
//   2. SearchBuilder
//      A fluent builder used to incrementally construct an InlineSearchInput
//      without manually assembling the underlying structure.
//
// # Overview
//
// Many FHIR servers support inline expansion of related resources using the
// _include and _revinclude parameters. These allow clients to retrieve a
// primary resource along with related referenced resources in a single request.
//
//   _include     follows a reference from the primary resource outward.
//   _revinclude  returns resources that reference the primary resource.
//
// Example:
//   // Return a ServiceRequest and its performer
//   GET /ServiceRequest?_id=123&_include=ServiceRequest:performer
//
//   // Return a ServiceRequest and all Observations that reference it
//   GET /ServiceRequest?_id=123&_revinclude=Observation:based-on
//
// # Using SearchBuilder
//
// The SearchBuilder provides a simple, chainable API for constructing these queries.
//
// Example:
//
//   sb := search.NewSearchBuilder("ServiceRequest", "123")
//
//   sb.Include("ServiceRequest", "performer").
//      RevInclude("Observation", "subject")
//
//   input := sb.Build()
//   query := input.BuildQueryPath()
//
//   // Result:
//   // ServiceRequest/?_id=123&_include=ServiceRequest%3Aperformer&_revinclude=Observation%3Asubject
//
// # URL Escaping
//
// FHIR include expressions (e.g., "Resource:field") contain reserved characters.
// BuildQueryPath automatically URL-escapes these values to ensure compliance
// with the HTTP specification.
//
// # Intended Usage
//
// The package is designed for:
//
//   • FHIR client libraries that need to construct complex search queries  
//   • Services performing inline expansion of FHIR resources  
//   • REST API wrappers that encapsulate FHIR search logic  
//
// This package focuses only on query construction; the execution of HTTP
// requests is expected to be performed by higher-level components.
package search


// NewSearchBuilder creates and initializes a SearchBuilder instance that is used
// to construct FHIR inline search queries (e.g., using _include and _revinclude).
//
// It requires a baseResource (such as "ServiceRequest" or "Patient") and the
// resourceID of the primary resource being searched. These values form the
// foundation of the query path.
func NewSearchBuilder(baseResource, resourceID string) *SearchBuilder {
	return &SearchBuilder{
		Input: InlineSearchInput{
			BaseResource: baseResource,
			ResourceID:   resourceID,
			Params:       make([]SearchParams, 0),
		},
	}
}
