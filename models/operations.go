package models

import "encoding/json"

type Parameters struct {
	ID            *string               `json:"id,omitempty"`
	ResourceType  string                `json:"resourceType,omitempty"`
	Meta          *FHIRMeta             `json:"meta,omitempty"`
	ImplicitRules *string               `json:"implicitRules,omitempty"`
	Language      *string               `json:"language,omitempty"`
	Parameter     []ParametersParameter `json:"parameter,omitempty"`
}

type ParametersParameter struct {
	ID                   *string               `json:"id,omitempty"`
	Extension            []Extension           `json:"extension,omitempty"`
	ModifierExtension    []Extension           `json:"modifierExtension,omitempty"`
	Name                 string                `json:"name"`
	ValueBase64Binary    *string               `json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool                 `json:"valueBoolean,omitempty"`
	ValueCanonical       *string               `json:"valueCanonical,omitempty"`
	ValueCode            *string               `json:"valueCode,omitempty"`
	ValueDate            *string               `json:"valueDate,omitempty"`
	ValueDateTime        *string               `json:"valueDateTime,omitempty"`
	ValueDecimal         *json.Number          `json:"valueDecimal,omitempty"`
	ValueID              *string               `json:"valueId,omitempty"`
	ValueInstant         *string               `json:"valueInstant,omitempty"`
	ValueInteger         *int                  `json:"valueInteger,omitempty"`
	ValueMarkdown        *string               `json:"valueMarkdown,omitempty"`
	ValueOid             *string               `json:"valueOid,omitempty"`
	ValuePositiveInt     *int                  `json:"valuePositiveInt,omitempty"`
	ValueString          *string               `json:"valueString,omitempty"`
	ValueTime            *string               `json:"valueTime,omitempty"`
	ValueUnsignedInt     *int                  `json:"valueUnsignedInt,omitempty"`
	ValueUrI             *string               `json:"valueUrI,omitempty"`
	ValueUrL             *string               `json:"valueUrL,omitempty"`
	ValueUUID            *string               `json:"valueUUID,omitempty"`
	ValueAddress         *FHIRAddress          `json:"valueAddress,omitempty"`
	ValueAge             *FHIRAge              `json:"valueAge,omitempty"`
	ValueAnnotation      *FHIRAnnotation       `json:"valueAnnotation,omitempty"`
	ValueAttachment      *FHIRAttachment       `json:"valueAttachment,omitempty"`
	ValueCodeableConcept *FHIRCodeableConcept  `json:"valueCodeableConcept,omitempty"`
	ValueCoding          *FHIRCoding           `json:"valueCoding,omitempty"`
	ValueContactPoint    *FHIRContactPoint     `json:"valueContactPoint,omitempty"`
	ValueDuration        *FHIRDuration         `json:"valueDuration,omitempty"`
	ValueHumanName       *FHIRHumanName        `json:"valueHumanName,omitempty"`
	ValueIdentifier      *FHIRIdentifier       `json:"valueIdentifier,omitempty"`
	ValuePeriod          *FHIRPeriod           `json:"valuePeriod,omitempty"`
	ValueQuantity        *FHIRQuantity         `json:"valueQuantity,omitempty"`
	ValueRange           *FHIRRange            `json:"valueRange,omitempty"`
	ValueRatio           *FHIRRatio            `json:"valueRatio,omitempty"`
	ValueReference       *FHIRReference        `json:"valueReference,omitempty"`
	ValueSampledData     *FHIRSampledData      `json:"valueSampledData,omitempty"`
	ValueTiming          *FHIRTiming           `json:"valueTiming,omitempty"`
	ValueExpression      *FHIRExpression       `json:"valueExpression,omitempty"`
	ValueMeta            *FHIRMeta             `json:"valueMeta,omitempty"`
	Resource             json.RawMessage       `json:"resource,omitempty"`
	Part                 []ParametersParameter `json:"part,omitempty"`
}
