package models

import (
	"errors"
	"strings"

	"github.com/savannahghi/scalarutils"
)

// Patient definition: demographics and other administrative information
// about an individual or animal receiving care or other health-related services.
type Patient struct {
	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID *string `json:"id,omitempty"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human.
	// The narrative need not encode all the structured data, but is required to
	// contain sufficient detail to make it "clinically safe" for a human to just read the narrative.
	// Resource definitions may define what content should be represented in the
	// narrative to ensure clinical safety.
	Text *Narrative `json:"text,omitempty"`

	// An identifier for this patient.
	Identifier []*Identifier `json:"identifier,omitempty"`

	// Whether this patient record is in active use.
	// Many systems use this property to mark as non-current patients, such as
	// those that have not been seen for a period of time based on an organization's business rules.
	//
	// It is often used to filter patient lists to exclude inactive patients
	//
	// Deceased patients may also be marked as inactive for the same reasons, but
	// may be active for some time after death.
	Active *bool `json:"active,omitempty"`

	// A name associated with the individual.
	Name []*HumanName `json:"name,omitempty"`

	// A contact detail (e.g. a telephone number or an email address) by which the
	// individual may be contacted.
	Telecom []*ContactPoint `json:"telecom,omitempty"`

	// Administrative Gender - the gender that the patient is considered to have
	// for administration and record keeping purposes.
	Gender *PatientGenderEnum `json:"gender,omitempty"`

	// The date of birth for the individual.
	BirthDate *scalarutils.Date `json:"birthDate,omitempty"`

	// Indicates if the individual is deceased or not.
	DeceasedBoolean *bool `json:"deceasedBoolean,omitempty"`

	// Indicates if the individual is deceased or not.
	DeceasedDateTime *scalarutils.Date `json:"deceasedDateTime,omitempty"`

	// An address for the individual.
	Address []*Address `json:"address,omitempty"`

	// This field contains a patient's most recent marital (civil) status.
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`

	// Indicates whether the patient is part of a multiple (boolean) or indicates
	// the actual birth order (integer).
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`

	// Indicates whether the patient is part of a multiple (boolean) or indicates
	// the actual birth order (integer).
	MultipleBirthInteger *string `json:"multipleBirthInteger,omitempty"`

	// Image of the patient.
	Photo []*Attachment `json:"photo,omitempty"`

	// A contact party (e.g. guardian, partner, friend) for the patient.
	Contact []*PatientContact `json:"contact,omitempty"`

	// A language which may be used to communicate with the patient about his or
	// her health.
	Communication []*PatientCommunication `json:"communication,omitempty"`

	// Patient's nominated care provider.
	GeneralPractitioner []*Reference `json:"generalPractitioner,omitempty"`

	// Organization that is the custodian of the patient record.
	ManagingOrganization *Reference `json:"managingOrganization,omitempty"`

	// Link to another patient resource that concerns the same actual patient.
	Link []*PatientLink `json:"link,omitempty"`

	// Meta stores more information about the resource
	Meta *Meta `json:"meta,omitempty"`

	// Extension is an optional element that provides additional information not
	// captured in the basic resource definition
	Extension []*Extension `json:"extension,omitempty"`
}

// GetReceivingFacilityDetails is used to fetch the details of the receiving
// facility.
func (p Patient) GetPatientHealthIDIdentifier() string {
	var healthID string

	for _, identifier := range p.Identifier {
		healthIDDocumentType := "HEALTH_ID"
		for _, id := range identifier.Type.Coding {
			if *id.System == healthIDDocumentType {
				healthID = id.Display

				break
			}
		}
	}

	return healthID
}

// PatientCommunication definition: demographics and other administrative
// information about an individual or animal receiving care or other health-related services.
type PatientCommunication struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The ISO-639-1 alpha 2 code in lower case for the language, optionally
	// followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper case;
	// e.g. "en" for English, or "en-US" for American English versus "en-EN" for England English.
	Language *CodeableConcept `json:"language,omitempty"`

	// Indicates whether or not the patient prefers this language (over other
	// languages he masters up a certain level).
	Preferred *bool `json:"preferred,omitempty"`
}

// PatientContact definition: demographics and other administrative
// information about an individual or animal receiving care or other health-related services.
type PatientContact struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The nature of the relationship between the patient and the contact person.
	Relationship []*CodeableConcept `json:"relationship,omitempty"`

	// A name associated with the contact person.
	Name *HumanName `json:"name,omitempty"`

	// A contact detail for the person, e.g. a telephone number or an email
	// address.
	Telecom []*ContactPoint `json:"telecom,omitempty"`

	// Address for the contact person.
	Address *Address `json:"address,omitempty"`

	// Administrative Gender - the gender that the contact person is considered to
	// have for administration and record keeping purposes.
	Gender *PatientContactGenderEnum `json:"gender,omitempty"`

	// Organization on behalf of which the contact is acting or for which the
	// contact is working.
	Organization *Reference `json:"organization,omitempty"`

	// The period during which this contact person or organization is valid to be
	// contacted relating to this patient.
	Period *Period `json:"period,omitempty"`
}

// PatientLink definition: demographics and other administrative
// information about an individual or animal receiving care or other health-related services.
type PatientLink struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The other patient resource that the link refers to.
	Other *Reference `json:"other,omitempty"`

	// The type of link between this patient resource and another patient resource.
	Type *PatientLinkTypeEnum `json:"type,omitempty"`
}

// PatientRelayConnection is a Relay connection for Patient.
type PatientRelayConnection struct {
	Edges           []*PatientRelayEdge `json:"edges,omitempty"`
	HasOpenEpisodes bool                `json:"hasOpenEpisodes,omitempty"`
	PageInfo        *PageInfo           `json:"pageInfo,omitempty"`
}

// PatientRelayEdge is a Relay edge for Patient.
type PatientRelayEdge struct {
	Cursor          *string  `json:"cursor,omitempty"`
	HasOpenEpisodes bool     `json:"hasOpenEpisodes,omitempty"`
	Node            *Patient `json:"node,omitempty"`
}

// PatientRelayPayload is used to return single instances of Patient.
type PatientRelayPayload struct {
	Resource *Patient `json:"resource,omitempty"`
}

// PatientEdge is a Relay style edge for listings of  patient records.
type PatientEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Patient `json:"node"`
}

// PatientConnection is a Relay style connection for use in listings of
// patient records.
type PatientConnection struct {
	Edges    []*PatientEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

// Names renders the patient's names as a string.
func (p Patient) Names() string {
	name := ""
	if p.Name == nil {
		return name
	}

	names := []string{}

	for _, hn := range p.Name {
		if hn == nil {
			continue
		}

		if hn.Text == "" {
			continue
		}

		names = append(names, hn.Text)
	}

	name = strings.Join(names, " | ")

	return name
}

// IsEntity ...
func (p Patient) IsEntity() {}

// GetPhoneNumbers extract the patients phonenumbers.
func (p Patient) GetPhoneNumbers() ([]string, error) {
	if p.Telecom == nil {
		return nil, errors.New("telecom information is nil")
	}

	var phoneNumbers []string

	for _, telecom := range p.Telecom {
		if telecom != nil && telecom.Value != nil && *telecom.System == ContactPointSystemEnumPhone {
			phoneNumbers = append(phoneNumbers, *telecom.Value)
		}
	}

	if len(phoneNumbers) == 0 {
		return nil, errors.New("no patient phone numbers found")
	}

	return phoneNumbers, nil
}

// Identifiers models the IDs that can be assigned to a patient.
type Identifiers struct {
	HealthID   string
	NationalID string
}

// GetIDs safely extracts the Health ID and National ID from the patient.
func (p Patient) GetIDs() (*Identifiers, error) {
	if p.Identifier == nil {
		return nil, errors.New("identifier information is nil")
	}

	var healthID, nationalID string

	for _, identifier := range p.Identifier {
		if identifier != nil && len(identifier.Type.Coding) > 0 {
			for _, coding := range identifier.Type.Coding {
				if coding.System != nil && coding.Code != nil {
					switch *coding.System {
					case string("HEALTH_ID"):
						healthID = *coding.Code
					case string("NATIONAL_ID"):
						nationalID = *coding.Code
					}
				}
			}
		}
	}

	return &Identifiers{
		HealthID:   healthID,
		NationalID: nationalID,
	}, nil
}

// PageInfo is used to add pagination information to Relay edges.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

// PatientPayload is used to return patient records and ancillary data after
// mutations.
type PatientPayload struct {
	PatientRecord *Patient `json:"patientRecord,omitempty"`
}
