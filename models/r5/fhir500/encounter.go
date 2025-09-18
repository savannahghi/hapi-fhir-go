package fhir500

import (
	"github.com/savannahghi/scalarutils"
)

// Encounter definition: an interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s)
// or assessing the health status of a patient.
type Encounter struct {
	ID              *string                   `json:"id,omitempty"`
	Text            *Narrative                `json:"text,omitempty"`
	Identifier      []*Identifier             `json:"identifier,omitempty"`
	Status          EncounterStatus           `json:"status,omitempty"`
	StatusHistory   []*EncounterStatushistory `json:"statusHistory,omitempty"`
	Class           Coding                    `json:"class,omitempty"`
	ClassHistory    []*EncounterClasshistory  `json:"classHistory,omitempty"`
	Type            []*CodeableConcept        `json:"type,omitempty"`
	ServiceType     *CodeableConcept          `json:"serviceType,omitempty"`
	Priority        *CodeableConcept          `json:"priority,omitempty"`
	Subject         *Reference                `json:"subject,omitempty"`
	EpisodeOfCare   []*Reference              `json:"episodeOfCare,omitempty"`
	BasedOn         []*Reference              `json:"basedOn,omitempty"`
	Participant     []*EncounterParticipant   `json:"participant,omitempty"`
	Appointment     []*Reference              `json:"appointment,omitempty"`
	Period          *Period                   `json:"period,omitempty"`
	Length          *Duration                 `json:"length,omitempty"`
	ReasonCode      *string                   `json:"reasonCode,omitempty"`
	ReasonReference []*Reference              `json:"reasonReference,omitempty"`
	Diagnosis       []*EncounterDiagnosis     `json:"diagnosis,omitempty"`
	Account         []*Reference              `json:"account,omitempty"`
	Hospitalization *EncounterHospitalization `json:"hospitalization,omitempty"`
	Location        []*EncounterLocation      `json:"location,omitempty"`
	ServiceProvider *Reference                `json:"serviceProvider,omitempty"`
	PartOf          *Reference                `json:"partOf,omitempty"`
	Meta            *Meta                     `json:"meta,omitempty"`
	Extension       []*Extension              `json:"extension,omitempty"`
}

type EncounterClasshistory struct {
	ID     *string `json:"id,omitempty"`
	Class  *Coding `json:"class,omitempty"`
	Period *Period `json:"period,omitempty"`
}

type EncounterDiagnosis struct {
	ID        *string          `json:"id,omitempty"`
	Condition *Reference       `json:"condition,omitempty"`
	Use       *CodeableConcept `json:"use,omitempty"`
	Rank      *string          `json:"rank,omitempty"`
}

type EncounterHospitalization struct {
	ID                     *string            `json:"id,omitempty"`
	PreAdmissionIdentifier *Identifier        `json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference         `json:"origin,omitempty"`
	AdmitSource            *CodeableConcept   `json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept   `json:"reAdmission,omitempty"`
	DietPreference         []*CodeableConcept `json:"dietPreference,omitempty"`
	SpecialCourtesy        []*CodeableConcept `json:"specialCourtesy,omitempty"`
	SpecialArrangement     []*CodeableConcept `json:"specialArrangement,omitempty"`
	Destination            *Reference         `json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept   `json:"dischargeDisposition,omitempty"`
}

type EncounterLocation struct {
	ID           *string                  `json:"id,omitempty"`
	Location     *Reference               `json:"location,omitempty"`
	Status       *EncounterLocationStatus `json:"status,omitempty"`
	PhysicalType *CodeableConcept         `json:"physicalType,omitempty"`
	Period       *Period                  `json:"period,omitempty"`
}

type EncounterParticipant struct {
	ID         *string            `json:"id,omitempty"`
	Type       []*CodeableConcept `json:"type,omitempty"`
	Period     *Period            `json:"period,omitempty"`
	Individual *Reference         `json:"individual,omitempty"`
}

type EncounterStatushistory struct {
	ID     *string          `json:"id,omitempty"`
	Status *EncounterStatus `json:"status,omitempty"`
	Period *Period          `json:"period,omitempty"`
}

// EncounterRelayPayload is used to return single instances of Encounter.
type EncounterRelayPayload struct {
	Resource *Encounter `json:"resource,omitempty"`
}

// Duration definition: a length of time.
type Duration struct {
	ID         *string                 `json:"id,omitempty"`
	Value      *scalarutils.Decimal    `json:"value,omitempty"`
	Comparator *DurationComparatorEnum `json:"comparator,omitempty"`
	Unit       *string                 `json:"unit,omitempty"`
	System     *string                 `json:"system,omitempty"`
	Code       *string                 `json:"code,omitempty"`
}
