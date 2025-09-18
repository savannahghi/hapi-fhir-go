
package fhir430

import "encoding/json"
// ElementDefinition is documented here http://hl7.org/fhir/StructureDefinition/ElementDefinition
// Base StructureDefinition for ElementDefinition Type: Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition struct {
	ID                              *string                       `json:"ID,omitempty"`
	Extension                       []Extension                   `json:"extension,omitempty"`
	ModifierExtension               []Extension                   `json:"modifierExtension,omitempty"`
	Path                            string                        `json:"path"`
	Representation                  []PropertyRepresentation      `json:"representation,omitempty"`
	SliceName                       *string                       `json:"sliceName,omitempty"`
	SliceIsConstraining             *bool                         `json:"sliceIsConstraining,omitempty"`
	Label                           *string                       `json:"label,omitempty"`
	Code                            []Coding                      `json:"code,omitempty"`
	Slicing                         *ElementDefinitionSlicing     `json:"slicing,omitempty"`
	Short                           *string                       `json:"short,omitempty"`
	Definition                      *string                       `json:"definition,omitempty"`
	Comment                         *string                       `json:"comment,omitempty"`
	Requirements                    *string                       `json:"requirements,omitempty"`
	Alias                           []string                      `json:"alias,omitempty"`
	Min                             *int                          `json:"min,omitempty"`
	Max                             *string                       `json:"max,omitempty"`
	Base                            *ElementDefinitionBase        `json:"base,omitempty"`
	ContentReference                *string                       `json:"contentReference,omitempty"`
	Type                            []ElementDefinitionType       `json:"type,omitempty"`
	DefaultValueBase64Binary        *string                       `json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean             *bool                         `json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical           *string                       `json:"defaultValueCanonical,omitempty"`
	DefaultValueCode                *string                       `json:"defaultValueCode,omitempty"`
	DefaultValueDate                *string                       `json:"defaultValueDate,omitempty"`
	DefaultValueDateTime            *string                       `json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal             *json.Number                  `json:"defaultValueDecimal,omitempty"`
	DefaultValueId                  *string                       `json:"defaultValueId,omitempty"`
	DefaultValueInstant             *string                       `json:"defaultValueInstant,omitempty"`
	DefaultValueInteger             *int                          `json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown            *string                       `json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                 *string                       `json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt         *int                          `json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString              *string                       `json:"defaultValueString,omitempty"`
	DefaultValueTime                *string                       `json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt         *int                          `json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                 *string                       `json:"defaultValueUri,omitempty"`
	DefaultValueUrl                 *string                       `json:"defaultValueUrl,omitempty"`
	DefaultValueUuid                *string                       `json:"defaultValueUuid,omitempty"`
	DefaultValueAddress             *Address                      `json:"defaultValueAddress,omitempty"`
	DefaultValueAge                 *Age                          `json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation          *Annotation                   `json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment          *Attachment                   `json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept     *CodeableConcept              `json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding              *Coding                       `json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint        *ContactPoint                 `json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount               *Count                        `json:"defaultValueCount,omitempty"`
	DefaultValueDistance            *Distance                     `json:"defaultValueDistance,omitempty"`
	DefaultValueDuration            *Duration                     `json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName           *HumanName                    `json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier          *Identifier                   `json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney               *Money                        `json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod              *Period                       `json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity            *Quantity                     `json:"defaultValueQuantity,omitempty"`
	DefaultValueRange               *Range                        `json:"defaultValueRange,omitempty"`
	DefaultValueRatio               *Ratio                        `json:"defaultValueRatio,omitempty"`
	DefaultValueReference           *Reference                    `json:"defaultValueReference,omitempty"`
	DefaultValueSampledData         *SampledData                  `json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature           *Signature                    `json:"defaultValueSignature,omitempty"`
	DefaultValueTiming              *Timing                       `json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail       *ContactDetail                `json:"defaultValueContactDetail,omitempty"`
	DefaultValueContributor         *Contributor                  `json:"defaultValueContributor,omitempty"`
	DefaultValueDataRequirement     *DataRequirement              `json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression          *Expression                   `json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition *ParameterDefinition          `json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact     *RelatedArtifact              `json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition   *TriggerDefinition            `json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext        *UsageContext                 `json:"defaultValueUsageContext,omitempty"`
	DefaultValueDosage              *Dosage                       `json:"defaultValueDosage,omitempty"`
	DefaultValueMeta                *Meta                         `json:"defaultValueMeta,omitempty"`
	MeaningWhenMissing              *string                       `json:"meaningWhenMissing,omitempty"`
	OrderMeaning                    *string                       `json:"orderMeaning,omitempty"`
	FixedBase64Binary               *string                       `json:"fixedBase64Binary,omitempty"`
	FixedBoolean                    *bool                         `json:"fixedBoolean,omitempty"`
	FixedCanonical                  *string                       `json:"fixedCanonical,omitempty"`
	FixedCode                       *string                       `json:"fixedCode,omitempty"`
	FixedDate                       *string                       `json:"fixedDate,omitempty"`
	FixedDateTime                   *string                       `json:"fixedDateTime,omitempty"`
	FixedDecimal                    *json.Number                  `json:"fixedDecimal,omitempty"`
	FixedId                         *string                       `json:"fixedId,omitempty"`
	FixedInstant                    *string                       `json:"fixedInstant,omitempty"`
	FixedInteger                    *int                          `json:"fixedInteger,omitempty"`
	FixedMarkdown                   *string                       `json:"fixedMarkdown,omitempty"`
	FixedOid                        *string                       `json:"fixedOid,omitempty"`
	FixedPositiveInt                *int                          `json:"fixedPositiveInt,omitempty"`
	FixedString                     *string                       `json:"fixedString,omitempty"`
	FixedTime                       *string                       `json:"fixedTime,omitempty"`
	FixedUnsignedInt                *int                          `json:"fixedUnsignedInt,omitempty"`
	FixedUri                        *string                       `json:"fixedUri,omitempty"`
	FixedUrl                        *string                       `json:"fixedUrl,omitempty"`
	FixedUuid                       *string                       `json:"fixedUuid,omitempty"`
	FixedAddress                    *Address                      `json:"fixedAddress,omitempty"`
	FixedAge                        *Age                          `json:"fixedAge,omitempty"`
	FixedAnnotation                 *Annotation                   `json:"fixedAnnotation,omitempty"`
	FixedAttachment                 *Attachment                   `json:"fixedAttachment,omitempty"`
	FixedCodeableConcept            *CodeableConcept              `json:"fixedCodeableConcept,omitempty"`
	FixedCoding                     *Coding                       `json:"fixedCoding,omitempty"`
	FixedContactPoint               *ContactPoint                 `json:"fixedContactPoint,omitempty"`
	FixedCount                      *Count                        `json:"fixedCount,omitempty"`
	FixedDistance                   *Distance                     `json:"fixedDistance,omitempty"`
	FixedDuration                   *Duration                     `json:"fixedDuration,omitempty"`
	FixedHumanName                  *HumanName                    `json:"fixedHumanName,omitempty"`
	FixedIdentifier                 *Identifier                   `json:"fixedIdentifier,omitempty"`
	FixedMoney                      *Money                        `json:"fixedMoney,omitempty"`
	FixedPeriod                     *Period                       `json:"fixedPeriod,omitempty"`
	FixedQuantity                   *Quantity                     `json:"fixedQuantity,omitempty"`
	FixedRange                      *Range                        `json:"fixedRange,omitempty"`
	FixedRatio                      *Ratio                        `json:"fixedRatio,omitempty"`
	FixedReference                  *Reference                    `json:"fixedReference,omitempty"`
	FixedSampledData                *SampledData                  `json:"fixedSampledData,omitempty"`
	FixedSignature                  *Signature                    `json:"fixedSignature,omitempty"`
	FixedTiming                     *Timing                       `json:"fixedTiming,omitempty"`
	FixedContactDetail              *ContactDetail                `json:"fixedContactDetail,omitempty"`
	FixedContributor                *Contributor                  `json:"fixedContributor,omitempty"`
	FixedDataRequirement            *DataRequirement              `json:"fixedDataRequirement,omitempty"`
	FixedExpression                 *Expression                   `json:"fixedExpression,omitempty"`
	FixedParameterDefinition        *ParameterDefinition          `json:"fixedParameterDefinition,omitempty"`
	FixedRelatedArtifact            *RelatedArtifact              `json:"fixedRelatedArtifact,omitempty"`
	FixedTriggerDefinition          *TriggerDefinition            `json:"fixedTriggerDefinition,omitempty"`
	FixedUsageContext               *UsageContext                 `json:"fixedUsageContext,omitempty"`
	FixedDosage                     *Dosage                       `json:"fixedDosage,omitempty"`
	FixedMeta                       *Meta                         `json:"fixedMeta,omitempty"`
	PatternBase64Binary             *string                       `json:"patternBase64Binary,omitempty"`
	PatternBoolean                  *bool                         `json:"patternBoolean,omitempty"`
	PatternCanonical                *string                       `json:"patternCanonical,omitempty"`
	PatternCode                     *string                       `json:"patternCode,omitempty"`
	PatternDate                     *string                       `json:"patternDate,omitempty"`
	PatternDateTime                 *string                       `json:"patternDateTime,omitempty"`
	PatternDecimal                  *json.Number                  `json:"patternDecimal,omitempty"`
	PatternId                       *string                       `json:"patternId,omitempty"`
	PatternInstant                  *string                       `json:"patternInstant,omitempty"`
	PatternInteger                  *int                          `json:"patternInteger,omitempty"`
	PatternMarkdown                 *string                       `json:"patternMarkdown,omitempty"`
	PatternOid                      *string                       `json:"patternOid,omitempty"`
	PatternPositiveInt              *int                          `json:"patternPositiveInt,omitempty"`
	PatternString                   *string                       `json:"patternString,omitempty"`
	PatternTime                     *string                       `json:"patternTime,omitempty"`
	PatternUnsignedInt              *int                          `json:"patternUnsignedInt,omitempty"`
	PatternUri                      *string                       `json:"patternUri,omitempty"`
	PatternUrl                      *string                       `json:"patternUrl,omitempty"`
	PatternUuid                     *string                       `json:"patternUuid,omitempty"`
	PatternAddress                  *Address                      `json:"patternAddress,omitempty"`
	PatternAge                      *Age                          `json:"patternAge,omitempty"`
	PatternAnnotation               *Annotation                   `json:"patternAnnotation,omitempty"`
	PatternAttachment               *Attachment                   `json:"patternAttachment,omitempty"`
	PatternCodeableConcept          *CodeableConcept              `json:"patternCodeableConcept,omitempty"`
	PatternCoding                   *Coding                       `json:"patternCoding,omitempty"`
	PatternContactPoint             *ContactPoint                 `json:"patternContactPoint,omitempty"`
	PatternCount                    *Count                        `json:"patternCount,omitempty"`
	PatternDistance                 *Distance                     `json:"patternDistance,omitempty"`
	PatternDuration                 *Duration                     `json:"patternDuration,omitempty"`
	PatternHumanName                *HumanName                    `json:"patternHumanName,omitempty"`
	PatternIdentifier               *Identifier                   `json:"patternIdentifier,omitempty"`
	PatternMoney                    *Money                        `json:"patternMoney,omitempty"`
	PatternPeriod                   *Period                       `json:"patternPeriod,omitempty"`
	PatternQuantity                 *Quantity                     `json:"patternQuantity,omitempty"`
	PatternRange                    *Range                        `json:"patternRange,omitempty"`
	PatternRatio                    *Ratio                        `json:"patternRatio,omitempty"`
	PatternReference                *Reference                    `json:"patternReference,omitempty"`
	PatternSampledData              *SampledData                  `json:"patternSampledData,omitempty"`
	PatternSignature                *Signature                    `json:"patternSignature,omitempty"`
	PatternTiming                   *Timing                       `json:"patternTiming,omitempty"`
	PatternContactDetail            *ContactDetail                `json:"patternContactDetail,omitempty"`
	PatternContributor              *Contributor                  `json:"patternContributor,omitempty"`
	PatternDataRequirement          *DataRequirement              `json:"patternDataRequirement,omitempty"`
	PatternExpression               *Expression                   `json:"patternExpression,omitempty"`
	PatternParameterDefinition      *ParameterDefinition          `json:"patternParameterDefinition,omitempty"`
	PatternRelatedArtifact          *RelatedArtifact              `json:"patternRelatedArtifact,omitempty"`
	PatternTriggerDefinition        *TriggerDefinition            `json:"patternTriggerDefinition,omitempty"`
	PatternUsageContext             *UsageContext                 `json:"patternUsageContext,omitempty"`
	PatternDosage                   *Dosage                       `json:"patternDosage,omitempty"`
	PatternMeta                     *Meta                         `json:"patternMeta,omitempty"`
	Example                         []ElementDefinitionExample    `json:"example,omitempty"`
	MinValueDate                    *string                       `json:"minValueDate,omitempty"`
	MinValueDateTime                *string                       `json:"minValueDateTime,omitempty"`
	MinValueInstant                 *string                       `json:"minValueInstant,omitempty"`
	MinValueTime                    *string                       `json:"minValueTime,omitempty"`
	MinValueDecimal                 *json.Number                  `json:"minValueDecimal,omitempty"`
	MinValueInteger                 *int                          `json:"minValueInteger,omitempty"`
	MinValuePositiveInt             *int                          `json:"minValuePositiveInt,omitempty"`
	MinValueUnsignedInt             *int                          `json:"minValueUnsignedInt,omitempty"`
	MinValueQuantity                *Quantity                     `json:"minValueQuantity,omitempty"`
	MaxValueDate                    *string                       `json:"maxValueDate,omitempty"`
	MaxValueDateTime                *string                       `json:"maxValueDateTime,omitempty"`
	MaxValueInstant                 *string                       `json:"maxValueInstant,omitempty"`
	MaxValueTime                    *string                       `json:"maxValueTime,omitempty"`
	MaxValueDecimal                 *json.Number                  `json:"maxValueDecimal,omitempty"`
	MaxValueInteger                 *int                          `json:"maxValueInteger,omitempty"`
	MaxValuePositiveInt             *int                          `json:"maxValuePositiveInt,omitempty"`
	MaxValueUnsignedInt             *int                          `json:"maxValueUnsignedInt,omitempty"`
	MaxValueQuantity                *Quantity                     `json:"maxValueQuantity,omitempty"`
	MaxLength                       *int                          `json:"maxLength,omitempty"`
	Condition                       []string                      `json:"condition,omitempty"`
	Constraint                      []ElementDefinitionConstraint `json:"constraint,omitempty"`
	MustSupport                     *bool                         `json:"mustSupport,omitempty"`
	IsModifier                      *bool                         `json:"isModifier,omitempty"`
	IsModifierReason                *string                       `json:"isModifierReason,omitempty"`
	IsSummary                       *bool                         `json:"isSummary,omitempty"`
	Binding                         *ElementDefinitionBinding     `json:"binding,omitempty"`
	Mapping                         []ElementDefinitionMapping    `json:"mapping,omitempty"`
}

// Indicates that the element is sliced into a set of alternative definitions (i.e. in a structure definition, there are multiple different constraints on a single element in the base resource). Slicing can be used in any resource that has cardinality ..* on the base resource, or any resource with a choice of types. The set of slices is any elements that come after this in the element sequence that have the same path, until a shorter path occurs (the shorter path terminates the set).
// The first element in the sequence, the one that carries the slicing, is the definition that applies to all the slices. This is based on the unconstrained element, but can apply any constraints as appropriate. This may include the common constraints on the children of the element.
type ElementDefinitionSlicing struct {
	ID            *string                                 `json:"ID,omitempty"`
	Extension     []Extension                             `json:"extension,omitempty"`
	Discriminator []ElementDefinitionSlicingDiscriminator `json:"discriminator,omitempty"`
	Description   *string                                 `json:"description,omitempty"`
	Ordered       *bool                                   `json:"ordered,omitempty"`
	Rules         SlicingRules                            `json:"rules"`
}

// Designates which child elements are used to discriminate between the slices when processing an instance. If one or more discriminators are provided, the value of the child elements in the instance data SHALL completely distinguish which slice the element in the resource matches based on the allowed values for those elements in each of the slices.
// If there is no discriminator, the content is hard to process, so this should be avoided.
type ElementDefinitionSlicingDiscriminator struct {
	ID        *string           `json:"ID,omitempty"`
	Extension []Extension       `json:"extension,omitempty"`
	Type      DiscriminatorType `json:"type"`
	Path      string            `json:"path"`
}

// Information about the base definition of the element, provided to make it unnecessary for tools to trace the deviation of the element through the derived and related profiles. When the element definition is not the original definition of an element - i.g. either in a constraint on another type, or for elements from a super type in a snap shot - then the information in provided in the element definition may be different to the base definition. On the original definition of the element, it will be same.
// The base information does not carry any information that could not be determined from the path and related profiles, but making this determination requires both that the related profiles are available, and that the algorithm to determine them be available. For tooling simplicity, the base information must always be populated in element definitions in snap shots, even if it is the same.
type ElementDefinitionBase struct {
	ID        *string     `json:"ID,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Path      string      `json:"path"`
	Min       int         `json:"min"`
	Max       string      `json:"max"`
}

// The data type or resource that the value of this element is permitted to be.
// The Type of the element can be left blank in a differential constraint, in which case the type is inherited from the resource. Abstract types are not permitted to appear as a type when multiple types are listed.  (I.e. Abstract types cannot be part of a choice).
type ElementDefinitionType struct {
	ID            *string                `json:"ID,omitempty"`
	Extension     []Extension            `json:"extension,omitempty"`
	Code          string                 `json:"code"`
	Profile       []string               `json:"profile,omitempty"`
	TargetProfile []string               `json:"targetProfile,omitempty"`
	Aggregation   []AggregationMode      `json:"aggregation,omitempty"`
	Versioning    *ReferenceVersionRules `json:"versioning,omitempty"`
}

// A sample value for this element demonstrating the type of information that would typically be found in the element.
// Examples will most commonly be present for data where it's not implicitly obvious from either the data type or value set what the values might be.  (I.e. Example values for dates or quantities would generally be unnecessary.)  If the example value is fully populated, the publication tool can generate an instance automatically.
type ElementDefinitionExample struct {
	ID                       *string             `json:"ID,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	Label                    string              `json:"label"`
	ValueBase64Binary        string              `json:"valueBase64Binary"`
	ValueBoolean             bool                `json:"valueBoolean"`
	ValueCanonical           string              `json:"valueCanonical"`
	ValueCode                string              `json:"valueCode"`
	ValueDate                string              `json:"valueDate"`
	ValueDateTime            string              `json:"valueDateTime"`
	ValueDecimal             json.Number         `json:"valueDecimal"`
	ValueId                  string              `json:"valueId"`
	ValueInstant             string              `json:"valueInstant"`
	ValueInteger             int                 `json:"valueInteger"`
	ValueMarkdown            string              `json:"valueMarkdown"`
	ValueOid                 string              `json:"valueOid"`
	ValuePositiveInt         int                 `json:"valuePositiveInt"`
	ValueString              string              `json:"valueString"`
	ValueTime                string              `json:"valueTime"`
	ValueUnsignedInt         int                 `json:"valueUnsignedInt"`
	ValueUri                 string              `json:"valueUri"`
	ValueUrl                 string              `json:"valueUrl"`
	ValueUuid                string              `json:"valueUuid"`
	ValueAddress             Address             `json:"valueAddress"`
	ValueAge                 Age                 `json:"valueAge"`
	ValueAnnotation          Annotation          `json:"valueAnnotation"`
	ValueAttachment          Attachment          `json:"valueAttachment"`
	ValueCodeableConcept     CodeableConcept     `json:"valueCodeableConcept"`
	ValueCoding              Coding              `json:"valueCoding"`
	ValueContactPoint        ContactPoint        `json:"valueContactPoint"`
	ValueCount               Count               `json:"valueCount"`
	ValueDistance            Distance            `json:"valueDistance"`
	ValueDuration            Duration            `json:"valueDuration"`
	ValueHumanName           HumanName           `json:"valueHumanName"`
	ValueIdentifier          Identifier          `json:"valueIdentifier"`
	ValueMoney               Money               `json:"valueMoney"`
	ValuePeriod              Period              `json:"valuePeriod"`
	ValueQuantity            Quantity            `json:"valueQuantity"`
	ValueRange               Range               `json:"valueRange"`
	ValueRatio               Ratio               `json:"valueRatio"`
	ValueReference           Reference           `json:"valueReference"`
	ValueSampledData         SampledData         `json:"valueSampledData"`
	ValueSignature           Signature           `json:"valueSignature"`
	ValueTiming              Timing              `json:"valueTiming"`
	ValueContactDetail       ContactDetail       `json:"valueContactDetail"`
	ValueContributor         Contributor         `json:"valueContributor"`
	ValueDataRequirement     DataRequirement     `json:"valueDataRequirement"`
	ValueExpression          Expression          `json:"valueExpression"`
	ValueParameterDefinition ParameterDefinition `json:"valueParameterDefinition"`
	ValueRelatedArtifact     RelatedArtifact     `json:"valueRelatedArtifact"`
	ValueTriggerDefinition   TriggerDefinition   `json:"valueTriggerDefinition"`
	ValueUsageContext        UsageContext        `json:"valueUsageContext"`
	ValueDosage              Dosage              `json:"valueDosage"`
	ValueMeta                Meta                `json:"valueMeta"`
}

// Formal constraints such as co-occurrence and other constraints that can be computationally evaluated within the context of the instance.
// Constraints should be declared on the "context" element - the lowest element in the hierarchy that is common to all nodes referenced by the constraint.
type ElementDefinitionConstraint struct {
	ID           *string            `json:"ID,omitempty"`
	Extension    []Extension        `json:"extension,omitempty"`
	Key          string             `json:"key"`
	Requirements *string            `json:"requirements,omitempty"`
	Severity     ConstraintSeverity `json:"severity"`
	Human        string             `json:"human"`
	Expression   *string            `json:"expression,omitempty"`
	Xpath        *string            `json:"xpath,omitempty"`
	Source       *string            `json:"source,omitempty"`
}

// Binds to a value set if this element is coded (code, Coding, CodeableConcept, Quantity), or the data types (string, uri).
// For a CodeableConcept, when no codes are allowed - only text, use a binding of strength "required" with a description explaining that no coded values are allowed and what sort of information to put in the "text" element.
type ElementDefinitionBinding struct {
	ID          *string         `json:"ID,omitempty"`
	Extension   []Extension     `json:"extension,omitempty"`
	Strength    BindingStrength `json:"strength"`
	Description *string         `json:"description,omitempty"`
	ValueSet    *string         `json:"valueSet,omitempty"`
}

// Identifies a concept from an external specification that roughly corresponds to this element.
// Mappings are not necessarily specific enough for safe translation.
type ElementDefinitionMapping struct {
	ID        *string     `json:"ID,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Identity  string      `json:"identity"`
	Language  *string     `json:"language,omitempty"`
	Map       string      `json:"map"`
	Comment   *string     `json:"comment,omitempty"`
}
