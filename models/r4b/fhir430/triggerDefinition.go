
package fhir430
// TriggerDefinition is documented here http://hl7.org/fhir/StructureDefinition/TriggerDefinition
// Base StructureDefinition for TriggerDefinition Type: A description of a triggering event. Triggering events can be named events, data events, or periodic, as determined by the type element.
type TriggerDefinition struct {
	ID              *string           `json:"ID,omitempty"`
	Extension       []Extension       `json:"extension,omitempty"`
	Type            TriggerType       `json:"type"`
	Name            *string           `json:"name,omitempty"`
	TimingTiming    *Timing           `json:"timingTiming,omitempty"`
	TimingReference *Reference        `json:"timingReference,omitempty"`
	TimingDate      *string           `json:"timingDate,omitempty"`
	TimingDateTime  *string           `json:"timingDateTime,omitempty"`
	Data            []DataRequirement `json:"data,omitempty"`
	Condition       *Expression       `json:"condition,omitempty"`
}
