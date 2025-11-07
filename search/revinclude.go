package search

// RevInclude adds a _revinclude search parameter to the SearchBuilder.
//
// In FHIR, _revinclude instructs the server to return resources that reference
// the primary resource. Where _include follows a reference *outward*, 
// _revinclude follows references *back inward*. 
//
// Example FHIR usage:
//   _revinclude=Observation:subject
//
// This means: "return any Observation resource whose 'subject' field points to
// the resource we are searching for."
//
// The method appends a new SearchParams entry to the builder's Params slice,
// marking it as a RevInclude type, and returns the updated builder to allow
// fluent/chainable calls.
func (s *SearchBuilder) RevInclude(targetResource, searchField string) *SearchBuilder {
	s.Input.Params = append(s.Input.Params, SearchParams{
		Type:           RevInclude,
		TargetResource: targetResource,
		SearchField:    searchField,
	})
	return s
}