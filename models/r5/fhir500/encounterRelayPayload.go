package fhir500

// EncounterRelayPayload is used to return single instances of Encounter.
type EncounterRelayPayload struct {
	Resource *Encounter `json:"resource,omitempty"`
}
