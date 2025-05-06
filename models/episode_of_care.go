package models

// EpisodeOfCare definition: an association between a patient and an organization / healthcare provider(s) during which time encounters may occur.
// the managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare struct {
	ID                   *string                       `json:"id,omitempty"`
	Text                 *Narrative                    `json:"text,omitempty"`
	Identifier           []*Identifier                 `json:"identifier,omitempty"`
	Status               *EpisodeOfCareStatus          `json:"status,omitempty"`
	StatusHistory        []*EpisodeofcareStatushistory `json:"statusHistory,omitempty"`
	Type                 []*CodeableConcept            `json:"type,omitempty"`
	Diagnosis            []*EpisodeofcareDiagnosis     `json:"diagnosis,omitempty"`
	Patient              *Reference                    `json:"patient,omitempty"`
	ManagingOrganization *Reference                    `json:"managingOrganization,omitempty"`
	Period               *Period                       `json:"period,omitempty"`
	ReferralRequest      []*Reference                  `json:"referralRequest,omitempty"`
	CareManager          *Reference                    `json:"careManager,omitempty"`
	Team                 []*Reference                  `json:"team,omitempty"`
	Account              []*Reference                  `json:"account,omitempty"`
	Meta                 *Meta                         `json:"meta,omitempty"`
	Extension            []*Extension                  `json:"extension,omitempty"`
}

// EpisodeOfCareRelayEdge is a Relay edge for EpisodeOfCare.
type EpisodeOfCareRelayEdge struct {
	Cursor *string        `json:"cursor,omitempty"`
	Node   *EpisodeOfCare `json:"node,omitempty"`
}

// EpisodeOfCareRelayPayload is used to return single instances of EpisodeOfCare.
type EpisodeOfCareRelayPayload struct {
	Resource *EpisodeOfCare `json:"resource,omitempty"`
}

// EpisodeofcareDiagnosis definition: an association between a patient and an organization / healthcare provider(s) during which time encounters may occur.
type EpisodeofcareDiagnosis struct {
	ID        *string          `json:"id,omitempty"`
	Condition *Reference       `json:"condition,omitempty"`
	Role      *CodeableConcept `json:"role,omitempty"`
	Rank      *string          `json:"rank,omitempty"`
}

// EpisodeofcareStatushistory definition: an association between a patient and an organization / healthcare provider during which time encounters may occur.
type EpisodeofcareStatushistory struct {
	ID     *string              `json:"id,omitempty"`
	Status *EpisodeOfCareStatus `json:"status,omitempty"`
	Period *Period              `json:"period,omitempty"`
}

// EpisodeOfCarePayload is used to return the results after creation of
// episodes of care.
type EpisodeOfCarePayload struct {
	EpisodeOfCare *EpisodeOfCare `json:"episodeOfCare"`
}
