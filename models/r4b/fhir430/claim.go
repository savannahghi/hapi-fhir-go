
package fhir430

import "encoding/json"
// Claim is documented here http://hl7.org/fhir/StructureDefinition/Claim
// A provider issued list of professional services and products which have been provided, or are to be provided, to a patient which is sent to an insurer for reimbursement.
type Claim struct {
	ID                   *string                      `json:"ID,omitempty"`
	Meta                 *Meta                        `json:"meta,omitempty"`
	ImplicitRules        *string                      `json:"implicitRules,omitempty"`
	Language             *string                      `json:"language,omitempty"`
	Text                 *Narrative                   `json:"text,omitempty"`
	Contained            []json.RawMessage            `json:"contained,omitempty"`
	Extension            []Extension                  `json:"extension,omitempty"`
	ModifierExtension    []Extension                  `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                 `json:"identifier,omitempty"`
	Status               FinancialResourceStatusCodes `json:"status"`
	Type                 CodeableConcept              `json:"type"`
	SubType              *CodeableConcept             `json:"subType,omitempty"`
	Use                  Use                          `json:"use"`
	Patient              Reference                    `json:"patient"`
	BillablePeriod       *Period                      `json:"billablePeriod,omitempty"`
	Created              string                       `json:"created"`
	Enterer              *Reference                   `json:"enterer,omitempty"`
	Insurer              *Reference                   `json:"insurer,omitempty"`
	Provider             Reference                    `json:"provider"`
	Priority             CodeableConcept              `json:"priority"`
	FundsReserve         *CodeableConcept             `json:"fundsReserve,omitempty"`
	Related              []ClaimRelated               `json:"related,omitempty"`
	Prescription         *Reference                   `json:"prescription,omitempty"`
	OriginalPrescription *Reference                   `json:"originalPrescription,omitempty"`
	Payee                *ClaimPayee                  `json:"payee,omitempty"`
	Referral             *Reference                   `json:"referral,omitempty"`
	Facility             *Reference                   `json:"facility,omitempty"`
	CareTeam             []ClaimCareTeam              `json:"careTeam,omitempty"`
	SupportingInfo       []ClaimSupportingInfo        `json:"supportingInfo,omitempty"`
	Diagnosis            []ClaimDiagnosis             `json:"diagnosis,omitempty"`
	Procedure            []ClaimProcedure             `json:"procedure,omitempty"`
	Insurance            []ClaimInsurance             `json:"insurance"`
	Accident             *ClaimAccident               `json:"accident,omitempty"`
	Item                 []ClaimItem                  `json:"item,omitempty"`
	Total                *Money                       `json:"total,omitempty"`
}

// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
// For example,  for the original treatment and follow-up exams.
type ClaimRelated struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

// The party to be reimbursed for cost of the products and services according to the terms of the policy.
// Often providers agree to receive the benefits payable to reduce the near-term costs to the patient. The insurer may decline to pay the provider and choose to pay the subscriber instead.
type ClaimPayee struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Party             *Reference      `json:"party,omitempty"`
}

// The members of the team who provided the products and services.
type ClaimCareTeam struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Provider          Reference        `json:"provider"`
	Responsible       *bool            `json:"responsible,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Qualification     *CodeableConcept `json:"qualification,omitempty"`
}

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
// Often there are multiple jurisdiction specific valuesets which are required.
type ClaimSupportingInfo struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Category          CodeableConcept  `json:"category"`
	Code              *CodeableConcept `json:"code,omitempty"`
	TimingDate        *string          `json:"timingDate,omitempty"`
	TimingPeriod      *Period          `json:"timingPeriod,omitempty"`
	ValueBoolean      *bool            `json:"valueBoolean,omitempty"`
	ValueString       *string          `json:"valueString,omitempty"`
	ValueQuantity     *Quantity        `json:"valueQuantity,omitempty"`
	ValueAttachment   *Attachment      `json:"valueAttachment,omitempty"`
	ValueReference    *Reference       `json:"valueReference,omitempty"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
}

// Information about diagnoses relevant to the claim items.
type ClaimDiagnosis struct {
	ID                       *string           `json:"ID,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	DiagnosisCodeableConcept CodeableConcept   `json:"diagnosisCodeableConcept"`
	DiagnosisReference       Reference         `json:"diagnosisReference"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `json:"onAdmission,omitempty"`
	PackageCode              *CodeableConcept  `json:"packageCode,omitempty"`
}

// Procedures performed on the patient relevant to the billing items with the claim.
type ClaimProcedure struct {
	ID                       *string           `json:"ID,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	Date                     *string           `json:"date,omitempty"`
	ProcedureCodeableConcept CodeableConcept   `json:"procedureCodeableConcept"`
	ProcedureReference       Reference         `json:"procedureReference"`
	Udi                      []Reference       `json:"udi,omitempty"`
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
// All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'Coverage.subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
type ClaimInsurance struct {
	ID                  *string     `json:"ID,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Sequence            int         `json:"sequence"`
	Focal               bool        `json:"focal"`
	Identifier          *Identifier `json:"identifier,omitempty"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
	PreAuthRef          []string    `json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference  `json:"claimResponse,omitempty"`
}

// Details of an accident which resulted in injuries which required the products and services listed in the claim.
type ClaimAccident struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Date              string           `json:"date"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress,omitempty"`
	LocationReference *Reference       `json:"locationReference,omitempty"`
}

// A claim line. Either a simple  product or service or a 'group' of details which can each be a simple items or groups of sub-details.
type ClaimItem struct {
	ID                      *string           `json:"ID,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Sequence                int               `json:"sequence"`
	CareTeamSequence        []int             `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int             `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int             `json:"procedureSequence,omitempty"`
	InformationSequence     []int             `json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept  `json:"revenue,omitempty"`
	Category                *CodeableConcept  `json:"category,omitempty"`
	ProductOrService        CodeableConcept   `json:"productOrService"`
	Modifier                []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept `json:"programCode,omitempty"`
	ServicedDate            *string           `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period           `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept  `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address          `json:"locationAddress,omitempty"`
	LocationReference       *Reference        `json:"locationReference,omitempty"`
	Quantity                *Quantity         `json:"quantity,omitempty"`
	UnitPrice               *Money            `json:"unitPrice,omitempty"`
	Factor                  *json.Number      `json:"factor,omitempty"`
	Net                     *Money            `json:"net,omitempty"`
	Udi                     []Reference       `json:"udi,omitempty"`
	BodySite                *CodeableConcept  `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept `json:"subSite,omitempty"`
	Encounter               []Reference       `json:"encounter,omitempty"`
	Detail                  []ClaimItemDetail `json:"detail,omitempty"`
}

// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimItemDetail struct {
	ID                *string                    `json:"ID,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Sequence          int                        `json:"sequence"`
	Revenue           *CodeableConcept           `json:"revenue,omitempty"`
	Category          *CodeableConcept           `json:"category,omitempty"`
	ProductOrService  CodeableConcept            `json:"productOrService"`
	Modifier          []CodeableConcept          `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept          `json:"programCode,omitempty"`
	Quantity          *Quantity                  `json:"quantity,omitempty"`
	UnitPrice         *Money                     `json:"unitPrice,omitempty"`
	Factor            *json.Number               `json:"factor,omitempty"`
	Net               *Money                     `json:"net,omitempty"`
	Udi               []Reference                `json:"udi,omitempty"`
	SubDetail         []ClaimItemDetailSubDetail `json:"subDetail,omitempty"`
}

// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimItemDetailSubDetail struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Sequence          int               `json:"sequence"`
	Revenue           *CodeableConcept  `json:"revenue,omitempty"`
	Category          *CodeableConcept  `json:"category,omitempty"`
	ProductOrService  CodeableConcept   `json:"productOrService"`
	Modifier          []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept `json:"programCode,omitempty"`
	Quantity          *Quantity         `json:"quantity,omitempty"`
	UnitPrice         *Money            `json:"unitPrice,omitempty"`
	Factor            *json.Number      `json:"factor,omitempty"`
	Net               *Money            `json:"net,omitempty"`
	Udi               []Reference       `json:"udi,omitempty"`
}

// This function returns resource reference information
func (r Claim) ResourceRef() (string, *string) {
	return "Claim", r.ID
}

// This function returns resource reference information
func (r Claim) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherClaim Claim

// MarshalJSON marshals the given Claim as JSON into a byte slice
func (r Claim) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaim
		ResourceType string `json:"resourceType"`
	}{
		OtherClaim:   OtherClaim(r),
		ResourceType: "Claim",
	})
}

// UnmarshalClaim unmarshals a Claim.
func UnmarshalClaim(b []byte) (Claim, error) {
	var claim Claim
	if err := json.Unmarshal(b, &claim); err != nil {
		return claim, err
	}
	return claim, nil
}
