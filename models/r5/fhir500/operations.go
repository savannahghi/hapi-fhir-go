package fhir500

import "encoding/json"

type Parameters struct {
	ID            *string               `json:"id,omitempty"`
	ResourceType  string                `json:"resourceType,omitempty"`
	Meta          *Meta                 `json:"meta,omitempty"`
	ImplicitRules *string               `json:"implicitRules,omitempty"`
	Language      *string               `json:"language,omitempty"`
	Parameter     []ParametersParameter `json:"parameter,omitempty"`
}

type ParametersParameter struct {
	ID                         *string                `json:"id,omitempty"`
	Extension                  []Extension            `json:"extension,omitempty"`
	ModifierExtension          []Extension            `json:"modifierExtension,omitempty"`
	Name                       string                 `json:"name"`
	ValueBase64Binary          *string                `json:"valueBase64Binary,omitempty"`
	ValueBoolean               *bool                  `json:"valueBoolean,omitempty"`
	ValueCanonical             *string                `json:"valueCanonical,omitempty"`
	ValueCode                  *string                `json:"valueCode,omitempty"`
	ValueDate                  *string                `json:"valueDate,omitempty"`
	ValueDateTime              *string                `json:"valueDateTime,omitempty"`
	ValueDecimal               *json.Number           `json:"valueDecimal,omitempty"`
	ValueID                    *string                `json:"valueId,omitempty"`
	ValueInstant               *string                `json:"valueInstant,omitempty"`
	ValueInteger               *int                   `json:"valueInteger,omitempty"`
	ValueInteger64             *string                `json:"valueInteger64,omitempty"`
	ValueMarkdown              *string                `json:"valueMarkdown,omitempty"`
	ValueOid                   *string                `json:"valueOid,omitempty"`
	ValuePositiveInt           *int                   `json:"valuePositiveInt,omitempty"`
	ValueString                *string                `json:"valueString,omitempty"`
	ValueTime                  *string                `json:"valueTime,omitempty"`
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt,omitempty"`
	ValueUrI                   *string                `json:"valueUri,omitempty"`
	ValueUrL                   *string                `json:"valueUrl,omitempty"`
	ValueUUID                  *string                `json:"valueUuid,omitempty"`
	ValueAddress               *Address               `json:"valueAddress,omitempty"`
	ValueAge                   *Age                   `json:"valueAge,omitempty"`
	ValueAnnotation            *Annotation            `json:"valueAnnotation,omitempty"`
	ValueAttachment            *Attachment            `json:"valueAttachment,omitempty"`
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept,omitempty"`
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference,omitempty"`
	ValueCoding                *Coding                `json:"valueCoding,omitempty"`
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint,omitempty"`
	ValueCount                 *Count                 `json:"valueCount,omitempty"`
	ValueDistance              *Distance              `json:"valueDistance,omitempty"`
	ValueDuration              *Duration              `json:"valueDuration,omitempty"`
	ValueHumanName             *HumanName             `json:"valueHumanName,omitempty"`
	ValueIdentifier            *Identifier            `json:"valueIdentifier,omitempty"`
	ValueMoney                 *Money                 `json:"valueMoney,omitempty"`
	ValuePeriod                *Period                `json:"valuePeriod,omitempty"`
	ValueQuantity              *Quantity              `json:"valueQuantity,omitempty"`
	ValueRange                 *Range                 `json:"valueRange,omitempty"`
	ValueRatio                 *Ratio                 `json:"valueRatio,omitempty"`
	ValueRatioRange            *RatioRange            `json:"valueRatioRange,omitempty"`
	ValueReference             *Reference             `json:"valueReference,omitempty"`
	ValueSampledData           *SampledData           `json:"valueSampledData,omitempty"`
	ValueSignature             *Signature             `json:"valueSignature,omitempty"`
	ValueTiming                *Timing                `json:"valueTiming,omitempty"`
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail,omitempty"`
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement,omitempty"`
	ValueExpression            *Expression            `json:"valueExpression,omitempty"`
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext          *UsageContext          `json:"valueUsageContext,omitempty"`
	ValueAvailability          *Availability          `json:"valueAvailability,omitempty"`
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	ValueDosage                *Dosage                `json:"valueDosage,omitempty"`
	ValueMeta                  *Meta                  `json:"valueMeta,omitempty"`
	Resource                   json.RawMessage        `json:"resource,omitempty"`
	Part                       []ParametersParameter  `json:"part,omitempty"`
}
