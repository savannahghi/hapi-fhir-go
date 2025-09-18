
package fhir430

import "encoding/json"
// ClaimResponse is documented here http://hl7.org/fhir/StructureDefinition/ClaimResponse
// This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse struct {
	ID                   *string                         `json:"ID,omitempty"`
	Meta                 *Meta                           `json:"meta,omitempty"`
	ImplicitRules        *string                         `json:"implicitRules,omitempty"`
	Language             *string                         `json:"language,omitempty"`
	Text                 *Narrative                      `json:"text,omitempty"`
	Contained            []json.RawMessage               `json:"contained,omitempty"`
	Extension            []Extension                     `json:"extension,omitempty"`
	ModifierExtension    []Extension                     `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                    `json:"identifier,omitempty"`
	Status               FinancialResourceStatusCodes    `json:"status"`
	Type                 CodeableConcept                 `json:"type"`
	SubType              *CodeableConcept                `json:"subType,omitempty"`
	Use                  Use                             `json:"use"`
	Patient              Reference                       `json:"patient"`
	Created              string                          `json:"created"`
	Insurer              Reference                       `json:"insurer"`
	Requestor            *Reference                      `json:"requestor,omitempty"`
	Request              *Reference                      `json:"request,omitempty"`
	Outcome              ClaimProcessingCodes            `json:"outcome"`
	Disposition          *string                         `json:"disposition,omitempty"`
	PreAuthRef           *string                         `json:"preAuthRef,omitempty"`
	PreAuthPeriod        *Period                         `json:"preAuthPeriod,omitempty"`
	PayeeType            *CodeableConcept                `json:"payeeType,omitempty"`
	Item                 []ClaimResponseItem             `json:"item,omitempty"`
	AddItem              []ClaimResponseAddItem          `json:"addItem,omitempty"`
	Adjudication         []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Total                []ClaimResponseTotal            `json:"total,omitempty"`
	Payment              *ClaimResponsePayment           `json:"payment,omitempty"`
	FundsReserve         *CodeableConcept                `json:"fundsReserve,omitempty"`
	FormCode             *CodeableConcept                `json:"formCode,omitempty"`
	Form                 *Attachment                     `json:"form,omitempty"`
	ProcessNote          []ClaimResponseProcessNote      `json:"processNote,omitempty"`
	CommunicationRequest []Reference                     `json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsurance        `json:"insurance,omitempty"`
	Error                []ClaimResponseError            `json:"error,omitempty"`
}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ClaimResponseItem struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence      int                             `json:"itemSequence"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication"`
	Detail            []ClaimResponseItemDetail       `json:"detail,omitempty"`
}

// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
type ClaimResponseItemAdjudication struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Value             *json.Number     `json:"value,omitempty"`
}

// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimResponseItemDetail struct {
	ID                *string                            `json:"ID,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	DetailSequence    int                                `json:"detailSequence"`
	NoteNumber        []int                              `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication    `json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

// A sub-detail adjudication of a simple product or service.
type ClaimResponseItemDetailSubDetail struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	SubDetailSequence int                             `json:"subDetailSequence"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
}

// The first-tier service adjudications for payor added product or service lines.
type ClaimResponseAddItem struct {
	ID                      *string                         `json:"ID,omitempty"`
	Extension               []Extension                     `json:"extension,omitempty"`
	ModifierExtension       []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence            []int                           `json:"itemSequence,omitempty"`
	DetailSequence          []int                           `json:"detailSequence,omitempty"`
	SubdetailSequence       []int                           `json:"subdetailSequence,omitempty"`
	Provider                []Reference                     `json:"provider,omitempty"`
	ProductOrService        CodeableConcept                 `json:"productOrService"`
	Modifier                []CodeableConcept               `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept               `json:"programCode,omitempty"`
	ServicedDate            *string                         `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                         `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                        `json:"locationAddress,omitempty"`
	LocationReference       *Reference                      `json:"locationReference,omitempty"`
	Quantity                *Quantity                       `json:"quantity,omitempty"`
	UnitPrice               *Money                          `json:"unitPrice,omitempty"`
	Factor                  *json.Number                    `json:"factor,omitempty"`
	Net                     *Money                          `json:"net,omitempty"`
	BodySite                *CodeableConcept                `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept               `json:"subSite,omitempty"`
	NoteNumber              []int                           `json:"noteNumber,omitempty"`
	Adjudication            []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Detail                  []ClaimResponseAddItemDetail    `json:"detail,omitempty"`
}

// The second-tier service adjudications for payor added services.
type ClaimResponseAddItemDetail struct {
	ID                *string                               `json:"ID,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                       `json:"productOrService"`
	Modifier          []CodeableConcept                     `json:"modifier,omitempty"`
	Quantity          *Quantity                             `json:"quantity,omitempty"`
	UnitPrice         *Money                                `json:"unitPrice,omitempty"`
	Factor            *json.Number                          `json:"factor,omitempty"`
	Net               *Money                                `json:"net,omitempty"`
	NoteNumber        []int                                 `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication       `json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

// The third-tier service adjudications for payor added services.
type ClaimResponseAddItemDetailSubDetail struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                 `json:"productOrService"`
	Modifier          []CodeableConcept               `json:"modifier,omitempty"`
	Quantity          *Quantity                       `json:"quantity,omitempty"`
	UnitPrice         *Money                          `json:"unitPrice,omitempty"`
	Factor            *json.Number                    `json:"factor,omitempty"`
	Net               *Money                          `json:"net,omitempty"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
}

// Categorized monetary totals for the adjudication.
// Totals for amounts submitted, co-pays, benefits payable etc.
type ClaimResponseTotal struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// Payment details for the adjudication of the claim.
type ClaimResponsePayment struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              CodeableConcept  `json:"type"`
	Adjustment        *Money           `json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `json:"adjustmentReason,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Amount            Money            `json:"amount"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
}

// A note that describes or explains adjudication results in a human readable form.
type ClaimResponseProcessNote struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *NoteType        `json:"type,omitempty"`
	Text              string           `json:"text"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
// All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
type ClaimResponseInsurance struct {
	ID                  *string     `json:"ID,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Sequence            int         `json:"sequence"`
	Focal               bool        `json:"focal"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
	ClaimResponse       *Reference  `json:"claimResponse,omitempty"`
}

// Errors encountered during the processing of the adjudication.
// If the request contains errors then an error element should be provided and no adjudication related sections (item, addItem, or payment) should be present.
type ClaimResponseError struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	ItemSequence      *int            `json:"itemSequence,omitempty"`
	DetailSequence    *int            `json:"detailSequence,omitempty"`
	SubDetailSequence *int            `json:"subDetailSequence,omitempty"`
	Code              CodeableConcept `json:"code"`
}

// This function returns resource reference information
func (r ClaimResponse) ResourceRef() (string, *string) {
	return "ClaimResponse", r.ID
}

// This function returns resource reference information
func (r ClaimResponse) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherClaimResponse ClaimResponse

// MarshalJSON marshals the given ClaimResponse as JSON into a byte slice
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
}

// UnmarshalClaimResponse unmarshals a ClaimResponse.
func UnmarshalClaimResponse(b []byte) (ClaimResponse, error) {
	var claimResponse ClaimResponse
	if err := json.Unmarshal(b, &claimResponse); err != nil {
		return claimResponse, err
	}
	return claimResponse, nil
}
