package models

import (
	"github.com/savannahghi/scalarutils"
)

// FHIREncounter definition: an interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s)
// or assessing the health status of a patient.
type FHIREncounter struct {
	ID              *string                       `json:"id,omitempty"`
	Text            *FHIRNarrative                `json:"text,omitempty"`
	Identifier      []*FHIRIdentifier             `json:"identifier,omitempty"`
	Status          EncounterStatus               `json:"status,omitempty"`
	StatusHistory   []*FHIREncounterStatushistory `json:"statusHistory,omitempty"`
	Class           FHIRCoding                    `json:"class,omitempty"`
	ClassHistory    []*FHIREncounterClasshistory  `json:"classHistory,omitempty"`
	Type            []*FHIRCodeableConcept        `json:"type,omitempty"`
	ServiceType     *FHIRCodeableConcept          `json:"serviceType,omitempty"`
	Priority        *FHIRCodeableConcept          `json:"priority,omitempty"`
	Subject         *FHIRReference                `json:"subject,omitempty"`
	EpisodeOfCare   []*FHIRReference              `json:"episodeOfCare,omitempty"`
	BasedOn         []*FHIRReference              `json:"basedOn,omitempty"`
	Participant     []*FHIREncounterParticipant   `json:"participant,omitempty"`
	Appointment     []*FHIRReference              `json:"appointment,omitempty"`
	Period          *FHIRPeriod                   `json:"period,omitempty"`
	Length          *FHIRDuration                 `json:"length,omitempty"`
	ReasonCode      *scalarutils.Code             `json:"reasonCode,omitempty"`
	ReasonReference []*FHIRReference              `json:"reasonReference,omitempty"`
	Diagnosis       []*FHIREncounterDiagnosis     `json:"diagnosis,omitempty"`
	Account         []*FHIRReference              `json:"account,omitempty"`
	Hospitalization *FHIREncounterHospitalization `json:"hospitalization,omitempty"`
	Location        []*FHIREncounterLocation      `json:"location,omitempty"`
	ServiceProvider *FHIRReference                `json:"serviceProvider,omitempty"`
	PartOf          *FHIRReference                `json:"partOf,omitempty"`
	Meta            *FHIRMeta                     `json:"meta,omitempty"`
	Extension       []*FHIRExtension              `json:"extension,omitempty"`
}

type FHIREncounterClasshistory struct {
	ID     *string     `json:"id,omitempty"`
	Class  *FHIRCoding `json:"class,omitempty"`
	Period *FHIRPeriod `json:"period,omitempty"`
}

type FHIREncounterDiagnosis struct {
	ID        *string              `json:"id,omitempty"`
	Condition *FHIRReference       `json:"condition,omitempty"`
	Use       *FHIRCodeableConcept `json:"use,omitempty"`
	Rank      *string              `json:"rank,omitempty"`
}

type FHIREncounterHospitalization struct {
	ID                     *string                `json:"id,omitempty"`
	PreAdmissionIdentifier *FHIRIdentifier        `json:"preAdmissionIdentifier,omitempty"`
	Origin                 *FHIRReference         `json:"origin,omitempty"`
	AdmitSource            *FHIRCodeableConcept   `json:"admitSource,omitempty"`
	ReAdmission            *FHIRCodeableConcept   `json:"reAdmission,omitempty"`
	DietPreference         []*FHIRCodeableConcept `json:"dietPreference,omitempty"`
	SpecialCourtesy        []*FHIRCodeableConcept `json:"specialCourtesy,omitempty"`
	SpecialArrangement     []*FHIRCodeableConcept `json:"specialArrangement,omitempty"`
	Destination            *FHIRReference         `json:"destination,omitempty"`
	DischargeDisposition   *FHIRCodeableConcept   `json:"dischargeDisposition,omitempty"`
}

type FHIREncounterLocation struct {
	ID           *string                  `json:"id,omitempty"`
	Location     *FHIRReference           `json:"location,omitempty"`
	Status       *EncounterLocationStatus `json:"status,omitempty"`
	PhysicalType *FHIRCodeableConcept     `json:"physicalType,omitempty"`
	Period       *FHIRPeriod              `json:"period,omitempty"`
}

type FHIREncounterParticipant struct {
	ID         *string                `json:"id,omitempty"`
	Type       []*FHIRCodeableConcept `json:"type,omitempty"`
	Period     *FHIRPeriod            `json:"period,omitempty"`
	Individual *FHIRReference         `json:"individual,omitempty"`
}

type FHIREncounterStatushistory struct {
	ID     *string          `json:"id,omitempty"`
	Status *EncounterStatus `json:"status,omitempty"`
	Period *FHIRPeriod      `json:"period,omitempty"`
}

// FHIREncounterRelayPayload is used to return single instances of Encounter.
type FHIREncounterRelayPayload struct {
	Resource *FHIREncounter `json:"resource,omitempty"`
}

// FHIRDuration definition: a length of time.
type FHIRDuration struct {
	ID         *string                 `json:"id,omitempty"`
	Value      *scalarutils.Decimal    `json:"value,omitempty"`
	Comparator *DurationComparatorEnum `json:"comparator,omitempty"`
	Unit       *string                 `json:"unit,omitempty"`
	System     *scalarutils.URI        `json:"system,omitempty"`
	Code       *scalarutils.Code       `json:"code,omitempty"`
}
