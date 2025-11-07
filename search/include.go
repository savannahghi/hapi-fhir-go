package search

// Include adds an _include search parameter to the SearchBuilder.
//
// In FHIR, _include instructs the server to return related resources that are
// referenced by the primary resource. For example, including a performer on a
// ServiceRequest would look like:
//
//   _include=ServiceRequest:performer
//
// This method appends a new SearchParams entry to the builder's Params slice,
// marking it as an Include type, and then returns the builder to allow for
// fluent/chainable calls.
func (s *SearchBuilder) Include(targetResource, searchFiled string) *SearchBuilder {
	s.Input.Params = append(s.Input.Params, SearchParams{
		Type: Include,
		TargetResource: targetResource,
		SearchField: searchFiled,
	})
	return s
}