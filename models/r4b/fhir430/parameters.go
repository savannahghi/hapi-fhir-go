
package fhir430

import "encoding/json"
// Parameters is documented here http://hl7.org/fhir/StructureDefinition/Parameters
// This resource is a non-persisted resource used to pass information into and back from an [operation](operations.html). It has no other use, and there is no RESTful endpoint associated with it.
type Parameters struct {
	ID            *string               `json:"ID,omitempty"`
	Meta          *Meta                 `json:"meta,omitempty"`
	ImplicitRules *string               `json:"implicitRules,omitempty"`
	Language      *string               `json:"language,omitempty"`
	Parameter     []ParametersParameter `json:"parameter,omitempty"`
}

// A parameter passed to or received from the operation.
type ParametersParameter struct {
	ID                       *string               `json:"ID,omitempty"`
	Extension                []Extension           `json:"extension,omitempty"`
	ModifierExtension        []Extension           `json:"modifierExtension,omitempty"`
	Name                     string                `json:"name"`
	ValueBase64Binary        *string               `json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                 `json:"valueBoolean,omitempty"`
	ValueCanonical           *string               `json:"valueCanonical,omitempty"`
	ValueCode                *string               `json:"valueCode,omitempty"`
	ValueDate                *string               `json:"valueDate,omitempty"`
	ValueDateTime            *string               `json:"valueDateTime,omitempty"`
	ValueDecimal             *json.Number          `json:"valueDecimal,omitempty"`
	ValueId                  *string               `json:"valueId,omitempty"`
	ValueInstant             *string               `json:"valueInstant,omitempty"`
	ValueInteger             *int                  `json:"valueInteger,omitempty"`
	ValueMarkdown            *string               `json:"valueMarkdown,omitempty"`
	ValueOid                 *string               `json:"valueOid,omitempty"`
	ValuePositiveInt         *int                  `json:"valuePositiveInt,omitempty"`
	ValueString              *string               `json:"valueString,omitempty"`
	ValueTime                *string               `json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                  `json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string               `json:"valueUri,omitempty"`
	ValueUrl                 *string               `json:"valueUrl,omitempty"`
	ValueUuid                *string               `json:"valueUuid,omitempty"`
	ValueAddress             *Address              `json:"valueAddress,omitempty"`
	ValueAge                 *Age                  `json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation           `json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment           `json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *CodeableConcept      `json:"valueCodeableConcept,omitempty"`
	ValueCoding              *Coding               `json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint         `json:"valueContactPoint,omitempty"`
	ValueCount               *Count                `json:"valueCount,omitempty"`
	ValueDistance            *Distance             `json:"valueDistance,omitempty"`
	ValueDuration            *Duration             `json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName            `json:"valueHumanName,omitempty"`
	ValueIdentifier          *Identifier           `json:"valueIdentifier,omitempty"`
	ValueMoney               *Money                `json:"valueMoney,omitempty"`
	ValuePeriod              *Period               `json:"valuePeriod,omitempty"`
	ValueQuantity            *Quantity             `json:"valueQuantity,omitempty"`
	ValueRange               *Range                `json:"valueRange,omitempty"`
	ValueRatio               *Ratio                `json:"valueRatio,omitempty"`
	ValueReference           *Reference            `json:"valueReference,omitempty"`
	ValueSampledData         *SampledData          `json:"valueSampledData,omitempty"`
	ValueSignature           *Signature            `json:"valueSignature,omitempty"`
	ValueTiming              *Timing               `json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail        `json:"valueContactDetail,omitempty"`
	ValueContributor         *Contributor          `json:"valueContributor,omitempty"`
	ValueDataRequirement     *DataRequirement      `json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression           `json:"valueExpression,omitempty"`
	ValueParameterDefinition *ParameterDefinition  `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact      `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition    `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext         `json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage               `json:"valueDosage,omitempty"`
	ValueMeta                *Meta                 `json:"valueMeta,omitempty"`
	Resource                 json.RawMessage       `json:"resource,omitempty"`
	Part                     []ParametersParameter `json:"part,omitempty"`
}

// This function returns resource reference information
func (r Parameters) ResourceRef() (string, *string) {
	return "Parameters", r.ID
}

// This function returns resource reference information
func (r Parameters) ContainedResources() []json.RawMessage {
	return nil
}

type OtherParameters Parameters

// MarshalJSON marshals the given Parameters as JSON into a byte slice
func (r Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherParameters
		ResourceType string `json:"resourceType"`
	}{
		OtherParameters: OtherParameters(r),
		ResourceType:    "Parameters",
	})
}

// UnmarshalParameters unmarshals a Parameters.
func UnmarshalParameters(b []byte) (Parameters, error) {
	var parameters Parameters
	if err := json.Unmarshal(b, &parameters); err != nil {
		return parameters, err
	}
	return parameters, nil
}
