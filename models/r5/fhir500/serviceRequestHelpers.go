package fhir500

import (
	"errors"
	"strings"
)

// ServiceRequestRelayPayload is used to return single instances of ServiceRequest.
type ServiceRequestRelayPayload struct {
	Resource *ServiceRequest `json:"resource,omitempty"`
}

// ReceivingFacility is the facility that is accepting the patient into its care.
type ReceivingFacility struct {
	FacilityName    string `json:"facilityName,omitempty"`
	FacilityCounty  string `json:"facilityCounty,omitempty"`
	FacilityContact string `json:"facilityContact,omitempty"`
	FacilityEmail   string `json:"facilityEmail,omitempty"`
}

// GetReceivingFacilityDetails is used to fetch the details of the receiving facility.
func (f *ServiceRequest) GetReceivingFacilityDetails() (*ReceivingFacility, error) {
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
func (f *ServiceRequest) GetSubject() *Reference {
	return &f.Subject
}

func (f *ServiceRequest) GetFacilityFromMeta() *Meta {
	return f.Meta
}

// GetFacilityName is used to get the name of the facility.
func (f *ServiceRequest) GetFacilityName() string {
	facilitySystem := "http://mycarehub/tenant-identification/facility"

	if f.Meta.Tag != nil {
		for _, meta := range f.Meta.Tag {
			if *meta.System == facilitySystem {
				return meta.Display
			}
		}
	}

	return ""
}

// GetPatientReferralReason fetches the reason why a patient has been referred from a  service request.
// It specifically looks for the code matching the common.ReferralReasonCIELCode in the service request's coding.
// If a matching code is found, the display name of that code is returned as the name of the referred test.
// It defaults to "test" if a test name is not found.
func (f *ServiceRequest) GetPatientReferralReason() string {
	referralReasonCIELCode := "159623"

	if f.Code != nil && f.Code.Concept != nil {
		for _, coding := range f.Code.Concept.Coding {
			if coding.Code != nil && *coding.Code == referralReasonCIELCode && coding.Display != "" {
				return coding.Display
			}
		}
	}

	return "test"
}

// GetPatientReferralTest returns the test that the patient has been referred for.
func (f *ServiceRequest) GetPatientReferralTest() string {
	if f.Code != nil && f.Code.Concept != nil {
		for _, coding := range f.Code.Concept.Coding {
			if coding.Code != nil && *coding.Code == "TEST" && coding.Display != "" {
				return coding.Display
			}
		}
	}

	return ""
}

// GetRequestedServices retrieves What is being requested/ordered, in this case, it could be the tests that the patient
// has been referred for.
func (f *ServiceRequest) GetRequestedServices(serviceCIELCode string) []string {
	var services []string

	if f.Code == nil || f.Code.Concept == nil {
		return services
	}

	for _, coding := range f.Code.Concept.Coding {
		if coding.Code != nil && *coding.Code == serviceCIELCode {
			services = append(services, coding.Display)
		}
	}

	return services
}

// GetPractitionersNotes returns all practitioner's notes from the  service request concatenated with a newline separator.
func (f *ServiceRequest) GetPractitionersNotes() string {
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
