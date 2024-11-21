package models

import (
	"errors"
	"strings"

	"github.com/savannahghi/scalarutils"
)

type FHIRServiceRequest struct {
	ID                      *string                    `json:"id,omitempty"`
	Text                    *FHIRNarrative             `json:"text,omitempty"`
	Identifier              []*FHIRIdentifier          `json:"identifier,omitempty"`
	InstantiatesCanonical   *scalarutils.Canonical     `json:"instantiatesCanonical,omitempty"`
	InstantiatesURI         *scalarutils.Instant       `json:"instantiatesURI,omitempty"`
	BasedOn                 []*FHIRReference           `json:"basedOn,omitempty"`
	Replaces                []*FHIRReference           `json:"replaces,omitempty"`
	Requisition             *FHIRIdentifier            `json:"requisition,omitempty"`
	Status                  ServiceRequestStatusEnum   `json:"status,omitempty"`
	Intent                  ServiceRequestIntentEnum   `json:"intent,omitempty"`
	Category                []*FHIRCodeableConcept     `json:"category,omitempty"`
	Priority                ServiceRequestPriorityEnum `json:"priority,omitempty"`
	DoNotPerform            *bool                      `json:"doNotPerform,omitempty"`
	Code                    *FHIRCodeableConcept       `json:"code,omitempty"`
	OrderDetail             []*FHIRCodeableConcept     `json:"orderDetail,omitempty"`
	QuantityQuantity        *FHIRQuantity              `json:"quantityQuantity,omitempty"`
	QuantityRatio           *FHIRRatio                 `json:"quantityRatio,omitempty"`
	QuantityRange           *FHIRRange                 `json:"quantityRange,omitempty"`
	Subject                 *FHIRReference             `json:"subject,omitempty"`
	Encounter               *FHIRReference             `json:"encounter,omitempty"`
	OccurrenceDateTime      *scalarutils.Date          `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *FHIRPeriod                `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *FHIRTiming                `json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool                      `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *scalarutils.Code          `json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *scalarutils.DateTime      `json:"authoredOn,omitempty"`
	Requester               *FHIRReference             `json:"requester,omitempty"`
	PerformerType           *FHIRCodeableConcept       `json:"performerType,omitempty"`
	Performer               []*FHIRReference           `json:"performer,omitempty"`
	LocationCode            *scalarutils.Code          `json:"locationCode,omitempty"`
	LocationReference       []*FHIRReference           `json:"locationReference,omitempty"`
	ReasonCode              *scalarutils.Code          `json:"reasonCode,omitempty"`
	ReasonReference         []*FHIRReference           `json:"reasonReference,omitempty"`
	Insurance               []*FHIRReference           `json:"insurance,omitempty"`
	SupportingInfo          []*FHIRReference           `json:"supportingInfo,omitempty"`
	Specimen                []*FHIRReference           `json:"specimen,omitempty"`
	BodySite                []*FHIRCodeableConcept     `json:"bodySite,omitempty"`
	Note                    []*FHIRAnnotation          `json:"note,omitempty"`
	PatientInstruction      *string                    `json:"patientInstruction,omitempty"`
	RelevantHistory         []*FHIRReference           `json:"relevantHistory,omitempty"`
	Meta                    *FHIRMeta                  `json:"meta,omitempty"`
	Extension               []*FHIRExtension           `json:"extension,omitempty"`
}

// FHIRServiceRequestRelayPayload is used to return single instances of ServiceRequest.
type FHIRServiceRequestRelayPayload struct {
	Resource *FHIRServiceRequest `json:"resource,omitempty"`
}

// ReceivingFacility is the facility that is accepting the patient into its care.
type ReceivingFacility struct {
	FacilityName    string `json:"facilityName,omitempty"`
	FacilityCounty  string `json:"facilityCounty,omitempty"`
	FacilityContact string `json:"facilityContact,omitempty"`
	FacilityEmail   string `json:"facilityEmail,omitempty"`
}

// GetReceivingFacilityDetails is used to fetch the details of the receiving facility.
func (f *FHIRServiceRequest) GetReceivingFacilityDetails() (*ReceivingFacility, error) {
	if f.Extension == nil {
		return nil, errors.New("extension details is nil")
	}

	receivingFacility := &ReceivingFacility{}

	for _, extension := range f.Extension {
		if extension.URL == "http://savannahghi.org/fhir/StructureDefinition/referred-facility" {
			for _, ext := range extension.Extension {
				if ext.URL == "facilityName" {
					receivingFacility.FacilityName = ext.ValueString
				}

				if ext.URL == "facilityCounty" {
					receivingFacility.FacilityCounty = ext.ValueString
				}

				if ext.URL == "facilityContact" {
					receivingFacility.FacilityContact = ext.ValueString
				}

				if ext.URL == "facilityEmail" {
					receivingFacility.FacilityEmail = ext.ValueString
				}
			}
		}
	}

	return receivingFacility, nil
}

// GetSubject returns the patient details.
func (f *FHIRServiceRequest) GetSubject() *FHIRReference {
	return f.Subject
}

func (f *FHIRServiceRequest) GetFacilityFromMeta() *FHIRMeta {
	return f.Meta
}

// GetFacilityName is used to get the name of the facility.
func (f *FHIRServiceRequest) GetFacilityName() string {
	facilitySystem := scalarutils.URI("http://mycarehub/tenant-identification/facility")

	if f.Meta.Tag != nil {
		for _, meta := range f.Meta.Tag {
			if string(*meta.System) == string(facilitySystem) {
				return meta.Display
			}
		}
	}

	return ""
}

// GetPatientReferralReason fetches the reason why a patient has been referred from a FHIR service request.
// It specifically looks for the code matching the common.ReferralReasonCIELCode in the service request's coding.
// If a matching code is found, the display name of that code is returned as the name of the referred test.
// It defaults to "test" if a test name is not found.
func (f *FHIRServiceRequest) GetPatientReferralReason() string {
	referralReasonCIELCode := "159623"

	if f.Code != nil && len(f.Code.Coding) > 0 {
		for _, coding := range f.Code.Coding {
			if coding.Code != nil && *coding.Code == scalarutils.Code(referralReasonCIELCode) && coding.Display != "" {
				return coding.Display
			}
		}
	}

	return "test"
}

// GetPatientReferralTest returns the test that the patient has been referred for.
func (f *FHIRServiceRequest) GetPatientReferralTest() string {
	if f.Code != nil && len(f.Code.Coding) > 0 {
		for _, coding := range f.Code.Coding {
			if coding.Code != nil && *coding.Code == scalarutils.Code("TEST") && coding.Display != "" {
				return coding.Display
			}
		}
	}

	return ""
}

// GetRequestedServices retrieves What is being requested/ordered, in this case, it could be the tests that the patient
// has been referred for.
func (f *FHIRServiceRequest) GetRequestedServices(serviceCIELCode string) []string {
	var services []string

	for _, coding := range f.Code.Coding {
		if coding.Code != nil && *coding.Code == scalarutils.Code(serviceCIELCode) {
			services = append(services, coding.Display)
		}
	}

	return services
}

// GetPractitionersNotes returns all practitioner's notes from the FHIR service request concatenated with a newline separator.
func (f *FHIRServiceRequest) GetPractitionersNotes() string {
	if len(f.Note) == 0 {
		return ""
	}

	var notes []string

	for _, note := range f.Note {
		if note.Text != nil {
			notes = append(notes, string(*note.Text))
		}
	}

	return strings.Join(notes, "\n")
}
