
package fhir430

import "encoding/json"
// ServiceRequest is documented here http://hl7.org/fhir/StructureDefinition/ServiceRequest
// A record of a request for service such as diagnostic investigations, treatments, or operations to be performed.
type ServiceRequest struct {
	ID                      *string           `json:"ID,omitempty"`
	Meta                    *Meta             `json:"meta,omitempty"`
	ImplicitRules           *string           `json:"implicitRules,omitempty"`
	Language                *string           `json:"language,omitempty"`
	Text                    *Narrative        `json:"text,omitempty"`
	Contained               []json.RawMessage `json:"contained,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Identifier              []Identifier      `json:"identifier,omitempty"`
	InstantiatesCanonical   []string          `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri         []string          `json:"instantiatesUri,omitempty"`
	BasedOn                 []Reference       `json:"basedOn,omitempty"`
	Replaces                []Reference       `json:"replaces,omitempty"`
	Requisition             *Identifier       `json:"requisition,omitempty"`
	Status                  RequestStatus     `json:"status"`
	Intent                  RequestIntent     `json:"intent"`
	Category                []CodeableConcept `json:"category,omitempty"`
	Priority                *RequestPriority  `json:"priority,omitempty"`
	DoNotPerform            *bool             `json:"doNotPerform,omitempty"`
	Code                    *CodeableConcept  `json:"code,omitempty"`
	OrderDetail             []CodeableConcept `json:"orderDetail,omitempty"`
	QuantityQuantity        *Quantity         `json:"quantityQuantity,omitempty"`
	QuantityRatio           *Ratio            `json:"quantityRatio,omitempty"`
	QuantityRange           *Range            `json:"quantityRange,omitempty"`
	Subject                 Reference         `json:"subject"`
	Encounter               *Reference        `json:"encounter,omitempty"`
	OccurrenceDateTime      *string           `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period           `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing           `json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool             `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept  `json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *string           `json:"authoredOn,omitempty"`
	Requester               *Reference        `json:"requester,omitempty"`
	PerformerType           *CodeableConcept  `json:"performerType,omitempty"`
	Performer               []Reference       `json:"performer,omitempty"`
	LocationCode            []CodeableConcept `json:"locationCode,omitempty"`
	LocationReference       []Reference       `json:"locationReference,omitempty"`
	ReasonCode              []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference         []Reference       `json:"reasonReference,omitempty"`
	Insurance               []Reference       `json:"insurance,omitempty"`
	SupportingInfo          []Reference       `json:"supportingInfo,omitempty"`
	Specimen                []Reference       `json:"specimen,omitempty"`
	BodySite                []CodeableConcept `json:"bodySite,omitempty"`
	Note                    []Annotation      `json:"note,omitempty"`
	PatientInstruction      *string           `json:"patientInstruction,omitempty"`
	RelevantHistory         []Reference       `json:"relevantHistory,omitempty"`
}

// This function returns resource reference information
func (r ServiceRequest) ResourceRef() (string, *string) {
	return "ServiceRequest", r.ID
}

// This function returns resource reference information
func (r ServiceRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherServiceRequest ServiceRequest

// MarshalJSON marshals the given ServiceRequest as JSON into a byte slice
func (r ServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherServiceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherServiceRequest: OtherServiceRequest(r),
		ResourceType:        "ServiceRequest",
	})
}

// UnmarshalServiceRequest unmarshals a ServiceRequest.
func UnmarshalServiceRequest(b []byte) (ServiceRequest, error) {
	var serviceRequest ServiceRequest
	if err := json.Unmarshal(b, &serviceRequest); err != nil {
		return serviceRequest, err
	}
	return serviceRequest, nil
}
