package fhir430

import "encoding/json"

// Extension is documented here http://hl7.org/fhir/StructureDefinition/Extension
// Base StructureDefinition for Extension Type: Optional Extension Element - found in all resources.
type Extension struct {
	ID                       *string              `json:"id,omitempty"`
	Extension                []Extension          `json:"extension,omitempty"`
	Url                      string               `json:"url"`
	ValueBase64Binary        *string              `json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                `json:"valueBoolean,omitempty"`
	ValueCanonical           *string              `json:"valueCanonical,omitempty"`
	ValueCode                *string              `json:"valueCode,omitempty"`
	ValueDate                *string              `json:"valueDate,omitempty"`
	ValueDateTime            *string              `json:"valueDateTime,omitempty"`
	ValueDecimal             *json.Number         `json:"valueDecimal,omitempty"`
	ValueId                  *string              `json:"valueId,omitempty"`
	ValueInstant             *string              `json:"valueInstant,omitempty"`
	ValueInteger             *int                 `json:"valueInteger,omitempty"`
	ValueMarkdown            *string              `json:"valueMarkdown,omitempty"`
	ValueOid                 *string              `json:"valueOid,omitempty"`
	ValuePositiveInt         *int                 `json:"valuePositiveInt,omitempty"`
	ValueString              *string              `json:"valueString,omitempty"`
	ValueTime                *string              `json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                 `json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string              `json:"valueUri,omitempty"`
	ValueUrl                 *string              `json:"valueUrl,omitempty"`
	ValueUuid                *string              `json:"valueUuid,omitempty"`
	ValueAddress             *Address             `json:"valueAddress,omitempty"`
	ValueAge                 *Age                 `json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation          `json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment          `json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *CodeableConcept     `json:"valueCodeableConcept,omitempty"`
	ValueCoding              *Coding              `json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint        `json:"valueContactPoint,omitempty"`
	ValueCount               *Count               `json:"valueCount,omitempty"`
	ValueDistance            *Distance            `json:"valueDistance,omitempty"`
	ValueDuration            *Duration            `json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName           `json:"valueHumanName,omitempty"`
	ValueIdentifier          *Identifier          `json:"valueIdentifier,omitempty"`
	ValueMoney               *Money               `json:"valueMoney,omitempty"`
	ValuePeriod              *Period              `json:"valuePeriod,omitempty"`
	ValueQuantity            *Quantity            `json:"valueQuantity,omitempty"`
	ValueRange               *Range               `json:"valueRange,omitempty"`
	ValueRatio               *Ratio               `json:"valueRatio,omitempty"`
	ValueReference           *Reference           `json:"valueReference,omitempty"`
	ValueSampledData         *SampledData         `json:"valueSampledData,omitempty"`
	ValueSignature           *Signature           `json:"valueSignature,omitempty"`
	ValueTiming              *Timing              `json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail       `json:"valueContactDetail,omitempty"`
	ValueContributor         *Contributor         `json:"valueContributor,omitempty"`
	ValueDataRequirement     *DataRequirement     `json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression          `json:"valueExpression,omitempty"`
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact     `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition   `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext        `json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage              `json:"valueDosage,omitempty"`
	ValueMeta                *Meta                `json:"valueMeta,omitempty"`
}
