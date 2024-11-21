package models

// FHIREpisodeOfCare definition: an association between a patient and an organization / healthcare provider(s) during which time encounters may occur.
// the managing organization assumes a level of responsibility for the patient during this time.
type FHIREpisodeOfCare struct {
	ID                   *string                           `json:"id,omitempty"`
	Text                 *FHIRNarrative                    `json:"text,omitempty"`
	Identifier           []*FHIRIdentifier                 `json:"identifier,omitempty"`
	Status               *EpisodeOfCareStatus              `json:"status,omitempty"`
	StatusHistory        []*FHIREpisodeofcareStatushistory `json:"statusHistory,omitempty"`
	Type                 []*FHIRCodeableConcept            `json:"type,omitempty"`
	Diagnosis            []*FHIREpisodeofcareDiagnosis     `json:"diagnosis,omitempty"`
	Patient              *FHIRReference                    `json:"patient,omitempty"`
	ManagingOrganization *FHIRReference                    `json:"managingOrganization,omitempty"`
	Period               *FHIRPeriod                       `json:"period,omitempty"`
	ReferralRequest      []*FHIRReference                  `json:"referralRequest,omitempty"`
	CareManager          *FHIRReference                    `json:"careManager,omitempty"`
	Team                 []*FHIRReference                  `json:"team,omitempty"`
	Account              []*FHIRReference                  `json:"account,omitempty"`
	Meta                 *FHIRMeta                         `json:"meta,omitempty"`
	Extension            []*FHIRExtension                  `json:"extension,omitempty"`
}

// FHIREpisodeOfCareRelayEdge is a Relay edge for EpisodeOfCare.
type FHIREpisodeOfCareRelayEdge struct {
	Cursor *string            `json:"cursor,omitempty"`
	Node   *FHIREpisodeOfCare `json:"node,omitempty"`
}

// FHIREpisodeOfCareRelayPayload is used to return single instances of EpisodeOfCare.
type FHIREpisodeOfCareRelayPayload struct {
	Resource *FHIREpisodeOfCare `json:"resource,omitempty"`
}

// FHIREpisodeofcareDiagnosis definition: an association between a patient and an organization / healthcare provider(s) during which time encounters may occur.
type FHIREpisodeofcareDiagnosis struct {
	ID        *string              `json:"id,omitempty"`
	Condition *FHIRReference       `json:"condition,omitempty"`
	Role      *FHIRCodeableConcept `json:"role,omitempty"`
	Rank      *string              `json:"rank,omitempty"`
}

// FHIREpisodeofcareStatushistory definition: an association between a patient and an organization / healthcare provider during which time encounters may occur.
type FHIREpisodeofcareStatushistory struct {
	ID     *string              `json:"id,omitempty"`
	Status *EpisodeOfCareStatus `json:"status,omitempty"`
	Period *FHIRPeriod          `json:"period,omitempty"`
}

// EpisodeOfCarePayload is used to return the results after creation of
// episodes of care.
type EpisodeOfCarePayload struct {
	EpisodeOfCare *FHIREpisodeOfCare `json:"episodeOfCare"`
}
