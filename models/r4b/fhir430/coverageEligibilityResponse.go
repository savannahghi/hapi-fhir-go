
package fhir430

import "encoding/json"
// CoverageEligibilityResponse is documented here http://hl7.org/fhir/StructureDefinition/CoverageEligibilityResponse
// This resource provides eligibility and plan details from the processing of an CoverageEligibilityRequest resource.
type CoverageEligibilityResponse struct {
	ID                *string                                `json:"ID,omitempty"`
	Meta              *Meta                                  `json:"meta,omitempty"`
	ImplicitRules     *string                                `json:"implicitRules,omitempty"`
	Language          *string                                `json:"language,omitempty"`
	Text              *Narrative                             `json:"text,omitempty"`
	Contained         []json.RawMessage                      `json:"contained,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                           `json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes           `json:"status"`
	Purpose           []EligibilityResponsePurpose           `json:"purpose"`
	Patient           Reference                              `json:"patient"`
	ServicedDate      *string                                `json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                `json:"servicedPeriod,omitempty"`
	Created           string                                 `json:"created"`
	Requestor         *Reference                             `json:"requestor,omitempty"`
	Request           Reference                              `json:"request"`
	Outcome           ClaimProcessingCodes                   `json:"outcome"`
	Disposition       *string                                `json:"disposition,omitempty"`
	Insurer           Reference                              `json:"insurer"`
	Insurance         []CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
	PreAuthRef        *string                                `json:"preAuthRef,omitempty"`
	Form              *CodeableConcept                       `json:"form,omitempty"`
	Error             []CoverageEligibilityResponseError     `json:"error,omitempty"`
}

// Financial instruments for reimbursement for the health care products and services.
// All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
type CoverageEligibilityResponseInsurance struct {
	ID                *string                                    `json:"ID,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Coverage          Reference                                  `json:"coverage"`
	Inforce           *bool                                      `json:"inforce,omitempty"`
	BenefitPeriod     *Period                                    `json:"benefitPeriod,omitempty"`
	Item              []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

// Benefits and optionally current balances, and authorization details by category or service.
type CoverageEligibilityResponseInsuranceItem struct {
	ID                      *string                                           `json:"ID,omitempty"`
	Extension               []Extension                                       `json:"extension,omitempty"`
	ModifierExtension       []Extension                                       `json:"modifierExtension,omitempty"`
	Category                *CodeableConcept                                  `json:"category,omitempty"`
	ProductOrService        *CodeableConcept                                  `json:"productOrService,omitempty"`
	Modifier                []CodeableConcept                                 `json:"modifier,omitempty"`
	Provider                *Reference                                        `json:"provider,omitempty"`
	Excluded                *bool                                             `json:"excluded,omitempty"`
	Name                    *string                                           `json:"name,omitempty"`
	Description             *string                                           `json:"description,omitempty"`
	Network                 *CodeableConcept                                  `json:"network,omitempty"`
	Unit                    *CodeableConcept                                  `json:"unit,omitempty"`
	Term                    *CodeableConcept                                  `json:"term,omitempty"`
	Benefit                 []CoverageEligibilityResponseInsuranceItemBenefit `json:"benefit,omitempty"`
	AuthorizationRequired   *bool                                             `json:"authorizationRequired,omitempty"`
	AuthorizationSupporting []CodeableConcept                                 `json:"authorizationSupporting,omitempty"`
	AuthorizationUrl        *string                                           `json:"authorizationUrl,omitempty"`
}

// Benefits used to date.
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	ID                 *string         `json:"ID,omitempty"`
	Extension          []Extension     `json:"extension,omitempty"`
	ModifierExtension  []Extension     `json:"modifierExtension,omitempty"`
	Type               CodeableConcept `json:"type"`
	AllowedUnsignedInt *int            `json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string         `json:"allowedString,omitempty"`
	AllowedMoney       *Money          `json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *int            `json:"usedUnsignedInt,omitempty"`
	UsedString         *string         `json:"usedString,omitempty"`
	UsedMoney          *Money          `json:"usedMoney,omitempty"`
}

// Errors encountered during the processing of the request.
type CoverageEligibilityResponseError struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
}

// This function returns resource reference information
func (r CoverageEligibilityResponse) ResourceRef() (string, *string) {
	return "CoverageEligibilityResponse", r.ID
}

// This function returns resource reference information
func (r CoverageEligibilityResponse) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCoverageEligibilityResponse CoverageEligibilityResponse

// MarshalJSON marshals the given CoverageEligibilityResponse as JSON into a byte slice
func (r CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityResponse: OtherCoverageEligibilityResponse(r),
		ResourceType:                     "CoverageEligibilityResponse",
	})
}

// UnmarshalCoverageEligibilityResponse unmarshals a CoverageEligibilityResponse.
func UnmarshalCoverageEligibilityResponse(b []byte) (CoverageEligibilityResponse, error) {
	var coverageEligibilityResponse CoverageEligibilityResponse
	if err := json.Unmarshal(b, &coverageEligibilityResponse); err != nil {
		return coverageEligibilityResponse, err
	}
	return coverageEligibilityResponse, nil
}
