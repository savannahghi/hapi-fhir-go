
package fhir430

import "encoding/json"
// StructureMap is documented here http://hl7.org/fhir/StructureDefinition/StructureMap
// A Map of relationships between 2 structures that can be used to transform data.
type StructureMap struct {
	ID                *string                 `json:"ID,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []json.RawMessage       `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Url               string                  `json:"url"`
	Identifier        []Identifier            `json:"identifier,omitempty"`
	Version           *string                 `json:"version,omitempty"`
	Name              string                  `json:"name"`
	Title             *string                 `json:"title,omitempty"`
	Status            PublicationStatus       `json:"status"`
	Experimental      *bool                   `json:"experimental,omitempty"`
	Date              *string                 `json:"date,omitempty"`
	Publisher         *string                 `json:"publisher,omitempty"`
	Contact           []ContactDetail         `json:"contact,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	UseContext        []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose           *string                 `json:"purpose,omitempty"`
	Structure         []StructureMapStructure `json:"structure,omitempty"`
	Import            []string                `json:"import,omitempty"`
	Group             []StructureMapGroup     `json:"group"`
}

// A structure definition used by this map. The structure definition may describe instances that are converted, or the instances that are produced.
// It is not necessary for a structure map to identify any dependent structures, though not listing them may restrict its usefulness.
type StructureMapStructure struct {
	ID                *string               `json:"ID,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Url               string                `json:"url"`
	Mode              StructureMapModelMode `json:"mode"`
	Alias             *string               `json:"alias,omitempty"`
	Documentation     *string               `json:"documentation,omitempty"`
}

// Organizes the mapping into manageable chunks for human review/ease of maintenance.
type StructureMapGroup struct {
	ID                *string                   `json:"ID,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Name              string                    `json:"name"`
	Extends           *string                   `json:"extends,omitempty"`
	TypeMode          StructureMapGroupTypeMode `json:"typeMode"`
	Documentation     *string                   `json:"documentation,omitempty"`
	Input             []StructureMapGroupInput  `json:"input"`
	Rule              []StructureMapGroupRule   `json:"rule"`
}

// A name assigned to an instance of data. The instance must be provided when the mapping is invoked.
// If no inputs are named, then the entry mappings are type based.
type StructureMapGroupInput struct {
	ID                *string               `json:"ID,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Name              string                `json:"name"`
	Type              *string               `json:"type,omitempty"`
	Mode              StructureMapInputMode `json:"mode"`
	Documentation     *string               `json:"documentation,omitempty"`
}

// Transform Rule from source to target.
type StructureMapGroupRule struct {
	ID                *string                          `json:"ID,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Name              string                           `json:"name"`
	Source            []StructureMapGroupRuleSource    `json:"source"`
	Target            []StructureMapGroupRuleTarget    `json:"target,omitempty"`
	Rule              []StructureMapGroupRule          `json:"rule,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `json:"dependent,omitempty"`
	Documentation     *string                          `json:"documentation,omitempty"`
}

// Source inputs to the mapping.
type StructureMapGroupRuleSource struct {
	ID                              *string                     `json:"ID,omitempty"`
	Extension                       []Extension                 `json:"extension,omitempty"`
	ModifierExtension               []Extension                 `json:"modifierExtension,omitempty"`
	Context                         string                      `json:"context"`
	Min                             *int                        `json:"min,omitempty"`
	Max                             *string                     `json:"max,omitempty"`
	Type                            *string                     `json:"type,omitempty"`
	DefaultValueBase64Binary        *string                     `json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean             *bool                       `json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical           *string                     `json:"defaultValueCanonical,omitempty"`
	DefaultValueCode                *string                     `json:"defaultValueCode,omitempty"`
	DefaultValueDate                *string                     `json:"defaultValueDate,omitempty"`
	DefaultValueDateTime            *string                     `json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal             *json.Number                `json:"defaultValueDecimal,omitempty"`
	DefaultValueId                  *string                     `json:"defaultValueId,omitempty"`
	DefaultValueInstant             *string                     `json:"defaultValueInstant,omitempty"`
	DefaultValueInteger             *int                        `json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown            *string                     `json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                 *string                     `json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt         *int                        `json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString              *string                     `json:"defaultValueString,omitempty"`
	DefaultValueTime                *string                     `json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt         *int                        `json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                 *string                     `json:"defaultValueUri,omitempty"`
	DefaultValueUrl                 *string                     `json:"defaultValueUrl,omitempty"`
	DefaultValueUuid                *string                     `json:"defaultValueUuid,omitempty"`
	DefaultValueAddress             *Address                    `json:"defaultValueAddress,omitempty"`
	DefaultValueAge                 *Age                        `json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation          *Annotation                 `json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment          *Attachment                 `json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept     *CodeableConcept            `json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding              *Coding                     `json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint        *ContactPoint               `json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount               *Count                      `json:"defaultValueCount,omitempty"`
	DefaultValueDistance            *Distance                   `json:"defaultValueDistance,omitempty"`
	DefaultValueDuration            *Duration                   `json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName           *HumanName                  `json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier          *Identifier                 `json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney               *Money                      `json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod              *Period                     `json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity            *Quantity                   `json:"defaultValueQuantity,omitempty"`
	DefaultValueRange               *Range                      `json:"defaultValueRange,omitempty"`
	DefaultValueRatio               *Ratio                      `json:"defaultValueRatio,omitempty"`
	DefaultValueReference           *Reference                  `json:"defaultValueReference,omitempty"`
	DefaultValueSampledData         *SampledData                `json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature           *Signature                  `json:"defaultValueSignature,omitempty"`
	DefaultValueTiming              *Timing                     `json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail       *ContactDetail              `json:"defaultValueContactDetail,omitempty"`
	DefaultValueContributor         *Contributor                `json:"defaultValueContributor,omitempty"`
	DefaultValueDataRequirement     *DataRequirement            `json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression          *Expression                 `json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition *ParameterDefinition        `json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact     *RelatedArtifact            `json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition   *TriggerDefinition          `json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext        *UsageContext               `json:"defaultValueUsageContext,omitempty"`
	DefaultValueDosage              *Dosage                     `json:"defaultValueDosage,omitempty"`
	DefaultValueMeta                *Meta                       `json:"defaultValueMeta,omitempty"`
	Element                         *string                     `json:"element,omitempty"`
	ListMode                        *StructureMapSourceListMode `json:"listMode,omitempty"`
	Variable                        *string                     `json:"variable,omitempty"`
	Condition                       *string                     `json:"condition,omitempty"`
	Check                           *string                     `json:"check,omitempty"`
	LogMessage                      *string                     `json:"logMessage,omitempty"`
}

// Content to create because of this mapping rule.
type StructureMapGroupRuleTarget struct {
	ID                *string                                `json:"ID,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Context           *string                                `json:"context,omitempty"`
	ContextType       *StructureMapContextType               `json:"contextType,omitempty"`
	Element           *string                                `json:"element,omitempty"`
	Variable          *string                                `json:"variable,omitempty"`
	ListMode          []StructureMapTargetListMode           `json:"listMode,omitempty"`
	ListRuleId        *string                                `json:"listRuleId,omitempty"`
	Transform         *StructureMapTransform                 `json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
}

// Parameters to the transform.
type StructureMapGroupRuleTargetParameter struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueId           string      `json:"valueId"`
	ValueString       string      `json:"valueString"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDecimal      json.Number `json:"valueDecimal"`
}

// Which other rules to apply in the context of this rule.
type StructureMapGroupRuleDependent struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Variable          []string    `json:"variable"`
}

// This function returns resource reference information
func (r StructureMap) ResourceRef() (string, *string) {
	return "StructureMap", r.ID
}

// This function returns resource reference information
func (r StructureMap) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherStructureMap StructureMap

// MarshalJSON marshals the given StructureMap as JSON into a byte slice
func (r StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureMap
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureMap: OtherStructureMap(r),
		ResourceType:      "StructureMap",
	})
}

// UnmarshalStructureMap unmarshals a StructureMap.
func UnmarshalStructureMap(b []byte) (StructureMap, error) {
	var structureMap StructureMap
	if err := json.Unmarshal(b, &structureMap); err != nil {
		return structureMap, err
	}
	return structureMap, nil
}
