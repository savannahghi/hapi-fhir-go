
package fhir430

import "encoding/json"
// ExplanationOfBenefit is documented here http://hl7.org/fhir/StructureDefinition/ExplanationOfBenefit
// This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit struct {
	ID                    *string                                `json:"ID,omitempty"`
	Meta                  *Meta                                  `json:"meta,omitempty"`
	ImplicitRules         *string                                `json:"implicitRules,omitempty"`
	Language              *string                                `json:"language,omitempty"`
	Text                  *Narrative                             `json:"text,omitempty"`
	Contained             []json.RawMessage                      `json:"contained,omitempty"`
	Extension             []Extension                            `json:"extension,omitempty"`
	ModifierExtension     []Extension                            `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                           `json:"identifier,omitempty"`
	Status                ExplanationOfBenefitStatus             `json:"status"`
	Type                  CodeableConcept                        `json:"type"`
	SubType               *CodeableConcept                       `json:"subType,omitempty"`
	Use                   Use                                    `json:"use"`
	Patient               Reference                              `json:"patient"`
	BillablePeriod        *Period                                `json:"billablePeriod,omitempty"`
	Created               string                                 `json:"created"`
	Enterer               *Reference                             `json:"enterer,omitempty"`
	Insurer               Reference                              `json:"insurer"`
	Provider              Reference                              `json:"provider"`
	Priority              *CodeableConcept                       `json:"priority,omitempty"`
	FundsReserveRequested *CodeableConcept                       `json:"fundsReserveRequested,omitempty"`
	FundsReserve          *CodeableConcept                       `json:"fundsReserve,omitempty"`
	Related               []ExplanationOfBenefitRelated          `json:"related,omitempty"`
	Prescription          *Reference                             `json:"prescription,omitempty"`
	OriginalPrescription  *Reference                             `json:"originalPrescription,omitempty"`
	Payee                 *ExplanationOfBenefitPayee             `json:"payee,omitempty"`
	Referral              *Reference                             `json:"referral,omitempty"`
	Facility              *Reference                             `json:"facility,omitempty"`
	Claim                 *Reference                             `json:"claim,omitempty"`
	ClaimResponse         *Reference                             `json:"claimResponse,omitempty"`
	Outcome               ClaimProcessingCodes                   `json:"outcome"`
	Disposition           *string                                `json:"disposition,omitempty"`
	PreAuthRef            []string                               `json:"preAuthRef,omitempty"`
	PreAuthRefPeriod      []Period                               `json:"preAuthRefPeriod,omitempty"`
	CareTeam              []ExplanationOfBenefitCareTeam         `json:"careTeam,omitempty"`
	SupportingInfo        []ExplanationOfBenefitSupportingInfo   `json:"supportingInfo,omitempty"`
	Diagnosis             []ExplanationOfBenefitDiagnosis        `json:"diagnosis,omitempty"`
	Procedure             []ExplanationOfBenefitProcedure        `json:"procedure,omitempty"`
	Precedence            *int                                   `json:"precedence,omitempty"`
	Insurance             []ExplanationOfBenefitInsurance        `json:"insurance"`
	Accident              *ExplanationOfBenefitAccident          `json:"accident,omitempty"`
	Item                  []ExplanationOfBenefitItem             `json:"item,omitempty"`
	AddItem               []ExplanationOfBenefitAddItem          `json:"addItem,omitempty"`
	Adjudication          []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Total                 []ExplanationOfBenefitTotal            `json:"total,omitempty"`
	Payment               *ExplanationOfBenefitPayment           `json:"payment,omitempty"`
	FormCode              *CodeableConcept                       `json:"formCode,omitempty"`
	Form                  *Attachment                            `json:"form,omitempty"`
	ProcessNote           []ExplanationOfBenefitProcessNote      `json:"processNote,omitempty"`
	BenefitPeriod         *Period                                `json:"benefitPeriod,omitempty"`
	BenefitBalance        []ExplanationOfBenefitBenefitBalance   `json:"benefitBalance,omitempty"`
}

// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
// For example,  for the original treatment and follow-up exams.
type ExplanationOfBenefitRelated struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

// The party to be reimbursed for cost of the products and services according to the terms of the policy.
// Often providers agree to receive the benefits payable to reduce the near-term costs to the patient. The insurer may decline to pay the provider and may choose to pay the subscriber instead.
type ExplanationOfBenefitPayee struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Party             *Reference       `json:"party,omitempty"`
}

// The members of the team who provided the products and services.
type ExplanationOfBenefitCareTeam struct {
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
type ExplanationOfBenefitSupportingInfo struct {
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
	Reason            *Coding          `json:"reason,omitempty"`
}

// Information about diagnoses relevant to the claim items.
type ExplanationOfBenefitDiagnosis struct {
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
type ExplanationOfBenefitProcedure struct {
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
type ExplanationOfBenefitInsurance struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Focal             bool        `json:"focal"`
	Coverage          Reference   `json:"coverage"`
	PreAuthRef        []string    `json:"preAuthRef,omitempty"`
}

// Details of a accident which resulted in injuries which required the products and services listed in the claim.
type ExplanationOfBenefitAccident struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress,omitempty"`
	LocationReference *Reference       `json:"locationReference,omitempty"`
}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ExplanationOfBenefitItem struct {
	ID                      *string                                `json:"ID,omitempty"`
	Extension               []Extension                            `json:"extension,omitempty"`
	ModifierExtension       []Extension                            `json:"modifierExtension,omitempty"`
	Sequence                int                                    `json:"sequence"`
	CareTeamSequence        []int                                  `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int                                  `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int                                  `json:"procedureSequence,omitempty"`
	InformationSequence     []int                                  `json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept                       `json:"revenue,omitempty"`
	Category                *CodeableConcept                       `json:"category,omitempty"`
	ProductOrService        CodeableConcept                        `json:"productOrService"`
	Modifier                []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                      `json:"programCode,omitempty"`
	ServicedDate            *string                                `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                               `json:"locationAddress,omitempty"`
	LocationReference       *Reference                             `json:"locationReference,omitempty"`
	Quantity                *Quantity                              `json:"quantity,omitempty"`
	UnitPrice               *Money                                 `json:"unitPrice,omitempty"`
	Factor                  *json.Number                           `json:"factor,omitempty"`
	Net                     *Money                                 `json:"net,omitempty"`
	Udi                     []Reference                            `json:"udi,omitempty"`
	BodySite                *CodeableConcept                       `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept                      `json:"subSite,omitempty"`
	Encounter               []Reference                            `json:"encounter,omitempty"`
	NoteNumber              []int                                  `json:"noteNumber,omitempty"`
	Adjudication            []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitItemDetail       `json:"detail,omitempty"`
}

// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
type ExplanationOfBenefitItemAdjudication struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Value             *json.Number     `json:"value,omitempty"`
}

// Second-tier of goods and services.
type ExplanationOfBenefitItemDetail struct {
	ID                *string                                   `json:"ID,omitempty"`
	Extension         []Extension                               `json:"extension,omitempty"`
	ModifierExtension []Extension                               `json:"modifierExtension,omitempty"`
	Sequence          int                                       `json:"sequence"`
	Revenue           *CodeableConcept                          `json:"revenue,omitempty"`
	Category          *CodeableConcept                          `json:"category,omitempty"`
	ProductOrService  CodeableConcept                           `json:"productOrService"`
	Modifier          []CodeableConcept                         `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                         `json:"programCode,omitempty"`
	Quantity          *Quantity                                 `json:"quantity,omitempty"`
	UnitPrice         *Money                                    `json:"unitPrice,omitempty"`
	Factor            *json.Number                              `json:"factor,omitempty"`
	Net               *Money                                    `json:"net,omitempty"`
	Udi               []Reference                               `json:"udi,omitempty"`
	NoteNumber        []int                                     `json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication    `json:"adjudication,omitempty"`
	SubDetail         []ExplanationOfBenefitItemDetailSubDetail `json:"subDetail,omitempty"`
}

// Third-tier of goods and services.
type ExplanationOfBenefitItemDetailSubDetail struct {
	ID                *string                                `json:"ID,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Sequence          int                                    `json:"sequence"`
	Revenue           *CodeableConcept                       `json:"revenue,omitempty"`
	Category          *CodeableConcept                       `json:"category,omitempty"`
	ProductOrService  CodeableConcept                        `json:"productOrService"`
	Modifier          []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                      `json:"programCode,omitempty"`
	Quantity          *Quantity                              `json:"quantity,omitempty"`
	UnitPrice         *Money                                 `json:"unitPrice,omitempty"`
	Factor            *json.Number                           `json:"factor,omitempty"`
	Net               *Money                                 `json:"net,omitempty"`
	Udi               []Reference                            `json:"udi,omitempty"`
	NoteNumber        []int                                  `json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
}

// The first-tier service adjudications for payor added product or service lines.
type ExplanationOfBenefitAddItem struct {
	ID                      *string                                `json:"ID,omitempty"`
	Extension               []Extension                            `json:"extension,omitempty"`
	ModifierExtension       []Extension                            `json:"modifierExtension,omitempty"`
	ItemSequence            []int                                  `json:"itemSequence,omitempty"`
	DetailSequence          []int                                  `json:"detailSequence,omitempty"`
	SubDetailSequence       []int                                  `json:"subDetailSequence,omitempty"`
	Provider                []Reference                            `json:"provider,omitempty"`
	ProductOrService        CodeableConcept                        `json:"productOrService"`
	Modifier                []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                      `json:"programCode,omitempty"`
	ServicedDate            *string                                `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                               `json:"locationAddress,omitempty"`
	LocationReference       *Reference                             `json:"locationReference,omitempty"`
	Quantity                *Quantity                              `json:"quantity,omitempty"`
	UnitPrice               *Money                                 `json:"unitPrice,omitempty"`
	Factor                  *json.Number                           `json:"factor,omitempty"`
	Net                     *Money                                 `json:"net,omitempty"`
	BodySite                *CodeableConcept                       `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept                      `json:"subSite,omitempty"`
	NoteNumber              []int                                  `json:"noteNumber,omitempty"`
	Adjudication            []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitAddItemDetail    `json:"detail,omitempty"`
}

// The second-tier service adjudications for payor added services.
type ExplanationOfBenefitAddItemDetail struct {
	ID                *string                                      `json:"ID,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                              `json:"productOrService"`
	Modifier          []CodeableConcept                            `json:"modifier,omitempty"`
	Quantity          *Quantity                                    `json:"quantity,omitempty"`
	UnitPrice         *Money                                       `json:"unitPrice,omitempty"`
	Factor            *json.Number                                 `json:"factor,omitempty"`
	Net               *Money                                       `json:"net,omitempty"`
	NoteNumber        []int                                        `json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication       `json:"adjudication,omitempty"`
	SubDetail         []ExplanationOfBenefitAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

// The third-tier service adjudications for payor added services.
type ExplanationOfBenefitAddItemDetailSubDetail struct {
	ID                *string                                `json:"ID,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                        `json:"productOrService"`
	Modifier          []CodeableConcept                      `json:"modifier,omitempty"`
	Quantity          *Quantity                              `json:"quantity,omitempty"`
	UnitPrice         *Money                                 `json:"unitPrice,omitempty"`
	Factor            *json.Number                           `json:"factor,omitempty"`
	Net               *Money                                 `json:"net,omitempty"`
	NoteNumber        []int                                  `json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
}

// Categorized monetary totals for the adjudication.
// Totals for amounts submitted, co-pays, benefits payable etc.
type ExplanationOfBenefitTotal struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// Payment details for the adjudication of the claim.
type ExplanationOfBenefitPayment struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Adjustment        *Money           `json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `json:"adjustmentReason,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
}

// A note that describes or explains adjudication results in a human readable form.
type ExplanationOfBenefitProcessNote struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *NoteType        `json:"type,omitempty"`
	Text              *string          `json:"text,omitempty"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// Balance by Benefit Category.
type ExplanationOfBenefitBenefitBalance struct {
	ID                *string                                       `json:"ID,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Category          CodeableConcept                               `json:"category"`
	Excluded          *bool                                         `json:"excluded,omitempty"`
	Name              *string                                       `json:"name,omitempty"`
	Description       *string                                       `json:"description,omitempty"`
	Network           *CodeableConcept                              `json:"network,omitempty"`
	Unit              *CodeableConcept                              `json:"unit,omitempty"`
	Term              *CodeableConcept                              `json:"term,omitempty"`
	Financial         []ExplanationOfBenefitBenefitBalanceFinancial `json:"financial,omitempty"`
}

// Benefits Used to date.
type ExplanationOfBenefitBenefitBalanceFinancial struct {
	ID                 *string         `json:"ID,omitempty"`
	Extension          []Extension     `json:"extension,omitempty"`
	ModifierExtension  []Extension     `json:"modifierExtension,omitempty"`
	Type               CodeableConcept `json:"type"`
	AllowedUnsignedInt *int            `json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string         `json:"allowedString,omitempty"`
	AllowedMoney       *Money          `json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *int            `json:"usedUnsignedInt,omitempty"`
	UsedMoney          *Money          `json:"usedMoney,omitempty"`
}

// This function returns resource reference information
func (r ExplanationOfBenefit) ResourceRef() (string, *string) {
	return "ExplanationOfBenefit", r.ID
}

// This function returns resource reference information
func (r ExplanationOfBenefit) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherExplanationOfBenefit ExplanationOfBenefit

// MarshalJSON marshals the given ExplanationOfBenefit as JSON into a byte slice
func (r ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExplanationOfBenefit
		ResourceType string `json:"resourceType"`
	}{
		OtherExplanationOfBenefit: OtherExplanationOfBenefit(r),
		ResourceType:              "ExplanationOfBenefit",
	})
}

// UnmarshalExplanationOfBenefit unmarshals a ExplanationOfBenefit.
func UnmarshalExplanationOfBenefit(b []byte) (ExplanationOfBenefit, error) {
	var explanationOfBenefit ExplanationOfBenefit
	if err := json.Unmarshal(b, &explanationOfBenefit); err != nil {
		return explanationOfBenefit, err
	}
	return explanationOfBenefit, nil
}
