package fhir500

// ObservationRelayPayload is used to return single instances of Observation.
type ObservationRelayPayload struct {
	Resource *Observation `json:"resource,omitempty"`
}
