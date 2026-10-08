package fhir500

// EpisodeOfCareRelayEdge is a Relay edge for EpisodeOfCare.
type EpisodeOfCareRelayEdge struct {
	Cursor *string        `json:"cursor,omitempty"`
	Node   *EpisodeOfCare `json:"node,omitempty"`
}

// EpisodeOfCareRelayPayload is used to return single instances of EpisodeOfCare.
type EpisodeOfCareRelayPayload struct {
	Resource *EpisodeOfCare `json:"resource,omitempty"`
}

// EpisodeOfCarePayload is used to return the results after creation of
// episodes of care.
type EpisodeOfCarePayload struct {
	EpisodeOfCare *EpisodeOfCare `json:"episodeOfCare"`
}
