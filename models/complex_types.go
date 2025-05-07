package models

import (
	"encoding/json"
	"time"

	"github.com/savannahghi/scalarutils"
)

// FHIRAddress definition: an address expressed using postal conventions (as
// opposed to gps or other location definition formats).  this data type may
// be used to convey addresses for use in delivering mail as well as for visiting
// locations which might not be valid for mail delivery.  there are a variety of postal address
// formats defined around the world.
type Address struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The purpose of this address.
	Use *AddressUseEnum `json:"use,omitempty"`

	// Distinguishes between physical addresses (those you can visit) and mailing
	// addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Type *AddressTypeEnum `json:"type,omitempty"`

	// Specifies the entire address as it should be displayed e.g. on a postal
	// label. This may be provided instead of or as well as the specific parts.
	Text string `json:"text,omitempty"`

	// This component contains the house number, apartment number, street name,
	// street direction,  P.O. Box number, delivery hints, and similar address
	// information.
	Line []*string `json:"line,omitempty"`

	// The name of the city, town, suburb, village or other community or delivery
	// center.
	City *string `json:"city,omitempty"`

	// The name of the administrative area (county).
	District *string `json:"district,omitempty"`

	// Sub-unit of a country with limited sovereignty in a federally organized
	// country. A code may be used if codes are in common use (e.g. US 2 letter
	// state codes).
	State *string `json:"state,omitempty"`

	// A postal code designating a region defined by the postal service.
	PostalCode *string `json:"postalCode,omitempty"`

	// Country - a nation as commonly understood or generally accepted.
	Country *string `json:"country,omitempty"`

	// Time period when address was/is in use.
	Period *Period `json:"period,omitempty"`
}

// FHIRAge definition: a duration of time during which an organism (or a
// process) has existed.
type Age struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision
	// in the presentation of the value.
	Value *scalarutils.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual
	// value is greater or less than the stated value due to measurement issues;
	// e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *AgeComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *string `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *string `json:"code,omitempty"`
}

// FHIRAnnotation definition: a  text note which also  contains information
// about who made the statement and when.
type Annotation struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The individual responsible for making the annotation.
	AuthorReference *Reference `json:"authorReference,omitempty"`

	// The individual responsible for making the annotation.
	AuthorString *string `json:"authorString,omitempty"`

	// Indicates when this particular annotation was made.
	Time *time.Time `json:"time,omitempty"`

	// The text of the annotation in markdown format.
	Text *scalarutils.Markdown `json:"text,omitempty"`
}

// FHIRAttachment definition: for referring to data content defined in other
// formats.
type Attachment struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies the type of the data in the attachment and allows a method to be
	// chosen to interpret or render the data. Includes mime type parameters such
	// as charset where appropriate.
	ContentType *string `json:"contentType,omitempty"`

	// The human language of the content. The value can be any valid value
	// according to BCP 47.
	Language *string `json:"language,omitempty"`

	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	Data *scalarutils.Base64Binary `json:"data,omitempty"`

	// A location where the data can be accessed.
	URL *string `json:"url,omitempty"`

	// The number of bytes of data that make up this attachment (before base64
	// encoding, if that is done).
	Size *int `json:"size,omitempty"`

	// The calculated hash of the data using SHA-1. Represented using base64.
	Hash *scalarutils.Base64Binary `json:"hash,omitempty"`

	// A label or set of text to display in place of the data.
	Title *string `json:"title,omitempty"`

	// The date that the attachment was first created.
	Creation *scalarutils.DateTime `json:"creation,omitempty"`
}

// FHIRCodeableConcept definition: a concept that may be defined by a formal
// reference to a terminology or ontology or may be provided by text.
type CodeableConcept struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A reference to a code defined by a terminology system.
	Coding []*Coding `json:"coding,omitempty"`

	// A human language representation of the concept as seen/selected/uttered by
	// the user who entered the data and/or which represents the intended meaning
	// of the user.
	Text string `json:"text,omitempty"`
}

// Coding definition: a reference to a code defined by a terminology
// system.
type Coding struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The identification of the code system that defines the meaning of the
	// symbol in the code.
	System *string `json:"system,omitempty"`

	// The version of the code system which was used when choosing this code. Note
	// that a well-maintained code system does not need the version reported,
	// because the meaning of codes is consistent across versions.
	// However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent,
	// the version SHOULD be exchanged.
	Version *string `json:"version,omitempty"`

	// A symbol in syntax defined by the system. The symbol may be a predefined
	// code or an expression in a syntax defined by the coding system (e.g.
	// post-coordination).
	Code *string `json:"code,omitempty"`

	// A representation of the meaning of the code in the system, following the
	// rules of the system.
	Display string `json:"display,omitempty"`

	// Indicates that this coding was chosen by a user directly - e.g. off a pick
	// list of available items (codes or displays).
	UserSelected *bool `json:"userSelected,omitempty"`
}

// ToString returns the Display field of the Coding struct as a string.
func (c *Coding) ToString() string {
	return c.Display
}

// FHIRContactPoint definition: details for all kinds of technology mediated
// contact points for a person or organization, including telephone, email,
// etc.
type ContactPoint struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Telecommunications form for contact point - what communications system is
	// required to make use of the contact.
	System *ContactPointSystemEnum `json:"system,omitempty"`

	// The actual contact point details, in a form that is meaningful to the
	// designated communication system (i.e. phone number or email address).
	Value *string `json:"value,omitempty"`

	// Identifies the purpose for the contact point.
	Use *ContactPointUseEnum `json:"use,omitempty"`

	// Specifies a preferred order in which to use a set of contacts.
	// ContactPoints with lower rank values are more preferred than those with
	// higher rank values.
	Rank *int64 `json:"rank,omitempty"`

	// Time period when the contact point was/is in use.
	Period *Period `json:"period,omitempty"`
}

// FHIRHumanName definition: a human's name with the ability to identify parts
// and usage.
type HumanName struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies the purpose for this name.
	Use HumanNameUseEnum `json:"use,omitempty"`

	// Specifies the entire name as it should be displayed e.g. on an application
	// UI. This may be provided instead of or as well as the specific parts.
	Text string `json:"text,omitempty"`

	// The part of a name that links to the genealogy. In some cultures (e.g.
	// Eritrea) the family name of a son is the first name of his father.
	Family *string `json:"family,omitempty"`

	// Given name.
	Given []*string `json:"given,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal,
	// employment or nobility status, etc. and that appears at the start of the
	// name.
	Prefix []*string `json:"prefix,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal,
	// employment or nobility status, etc. and that appears at the end of the name.
	Suffix []*string `json:"suffix,omitempty"`

	// Indicates the period of time when this name was valid for the named person.
	Period *Period `json:"period,omitempty"`
}

// FHIRIdentifier definition: an identifier - identifies some entity uniquely
// and unambiguously. typically this is used for business identifiers.
type Identifier struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The purpose of this identifier.
	Use IdentifierUseEnum `json:"use,omitempty"`

	// A coded type for the identifier that can be used to determine which
	// identifier to use for a specific purpose.
	Type CodeableConcept `json:"type,omitempty"`

	// Establishes the namespace for the value - that is, a URL that describes a
	// set values that are unique.
	System *string `json:"system,omitempty"`

	// The portion of the identifier typically relevant to the user and which is
	// unique within the context of the system.
	Value string `json:"value,omitempty"`

	// Time period during which identifier is/was valid for use.
	Period *Period `json:"period,omitempty"`

	// Organization that issued/manages the identifier.
	Assigner *Reference `json:"assigner,omitempty"`
}

// FHIRNarrative definition: a human-readable summary of the resource
// conveying the essential clinical and business information for the resource.
type Narrative struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The status of the narrative - whether it's entirely generated (from just
	// the defined data or the extensions too), or whether a human authored it and
	// it may contain additional data.
	Status *NarrativeStatusEnum `json:"status,omitempty"`

	// The actual narrative content, a stripped down version of XHTML.
	Div scalarutils.XHTML `json:"div,omitempty"`
}

// FHIRPeriod definition: a time period defined by a start and end date and
// optionally time.
type Period struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The start of the period. The boundary is inclusive.
	Start scalarutils.DateTime `json:"start,omitempty"`

	// The end of the period. If the end of the period is missing, it means no end
	// was known or planned at the time the instance was created. The start may be
	// in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	End scalarutils.DateTime `json:"end,omitempty"`
}

// FHIRReference definition: a reference from one resource to another.
type Reference struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A reference to a location at which the other resource is found. The
	// reference may be a relative reference, in which case it is relative to the
	// service base URL, or an absolute URL that resolves to the location where the resource is found.
	// The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific.
	// Internal fragment references (start with '#') refer to contained resources.
	Reference *string `json:"reference,omitempty"`

	//     The expected type of the target of the reference. If both
	// Reference.type and Reference.reference are populated and
	// Reference.reference is a FHIR URL, both SHALL be consistent.
	//
	// The type is the Canonical URL of Resource Definition that is the type this
	// reference refers to. References are URLs that are relative to
	// http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to http://hl7.org/fhir/StructureDefinition/Patient.
	// Absolute URLs are only allowed for logical models (and can only be used in references in logical models, not resources).
	Type *string `json:"type,omitempty"`

	// An identifier for the target resource. This is used when there is no way to
	// reference the other resource directly, either because the entity it
	// represents is not available through a FHIR server, or because there is no way for the author of the resource to convert
	//  a known identifier to an actual location.
	// There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance,
	// but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance,
	// and that instance would need to be of a FHIR resource type allowed by the reference.
	Identifier *Identifier `json:"identifier,omitempty"`

	// Plain text narrative that identifies the resource in addition to the
	// resource reference.
	Display string `json:"display,omitempty"`
}

// FHIRExpression is documented here
// http://hl7.org/fhir/StructureDefinition/Expression
type Expression struct {
	ID          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	Description *string     `json:"description,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Language    string      `json:"language,omitempty"`
	Expression  *string     `json:"expression,omitempty"`
	Reference   *string     `json:"reference,omitempty"`
}

// FHIRMeta is a set of metadata that provides technical and workflow context
// to a resource.
type Meta struct {
	VersionID   string    `json:"versionId,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Source      string    `json:"source,omitempty"`
	Tag         []Coding  `json:"tag,omitempty"`
	Security    []Coding  `json:"security,omitempty"`
}

// Quantity definition: a measured amount (or an amount that can
// potentially be measured). note that measured amounts include amounts that
// are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type Quantity struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision
	// in the presentation of the value.
	Value float64 `json:"value"`

	// How the value should be understood and represented - whether the actual
	// value is greater or less than the stated value due to measurement issues;
	// e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *QuantityComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit string `json:"unit"`

	// The identification of the system that provides the coded form of the unit.
	System string `json:"system"`

	// A computer processable form of the unit in some unit representation system.
	Code *string `json:"code"`
}

// FHIRRange definition: a set of ordered quantities defined by a low and high limit.
type Range struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The low limit. The boundary is inclusive.
	Low Quantity `json:"low,omitempty"`

	// The high limit. The boundary is inclusive.
	High Quantity `json:"high,omitempty"`
}

// FHIRRatio definition: a relationship of two quantity values - expressed as a numerator and a denominator.
type Ratio struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the numerator.
	Numerator Quantity `json:"numerator,omitempty"`

	// The value of the denominator.
	Denominator Quantity `json:"denominator,omitempty"`
}

// Extension is an optional element that provides additional information not
// captured in the basic resource definition.
// Extensions allow the definition of new data elements or the modification of
// existing data elements in the FHIR data model.
type Extension struct {
	URL                  string           `json:"url,omitempty"`
	ValueBoolean         bool             `json:"valueBoolean,omitempty"`
	ValueInteger         *int             `json:"valueInteger,omitempty"`
	ValueDecimal         *float64         `json:"valueDecimal,omitempty"`
	ValueBase64Binary    string           `json:"valueBase64Binary,omitempty"`
	ValueInstant         string           `json:"valueInstant,omitempty"`
	ValueString          string           `json:"valueString,omitempty"`
	ValueURI             string           `json:"valueURI,omitempty"`
	ValueDate            string           `json:"valueDate,omitempty"`
	ValueDateTime        string           `json:"valueDateTime,omitempty"`
	ValueTime            string           `json:"valueTime,omitempty"`
	ValueCode            string           `json:"valueCode,omitempty"`
	ValueOid             string           `json:"valueOid,omitempty"`
	ValueUUID            string           `json:"valueUUID,omitempty"`
	ValueID              string           `json:"valueID,omitempty"`
	ValueUnsignedInt     int              `json:"valueUnsignedInt,omitempty"`
	ValuePositiveInt     int              `json:"valuePositiveInt,omitempty"`
	ValueMarkdown        string           `json:"valueMarkdown,omitempty"`
	ValueAnnotation      *Annotation      `json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
	ValueIdentifier      *Identifier      `json:"valueIdentifier,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `json:"valueCoding,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValuePeriod          *Period          `json:"valuePeriod,omitempty"`
	ValueRatio           *Ratio           `json:"valueRatio,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
	ValueExpression      *Expression      `json:"valueExpression,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
}

// FHIRCodeableReference is documented here http://hl7.org/fhir/StructureDefinition/CodeableReference
type CodeableReference struct {
	ID        *string          `json:"id,omitempty"`
	Extension []Extension      `json:"extension,omitempty"`
	Concept   *CodeableConcept `json:"concept,omitempty"`
	Reference *Reference       `json:"reference,omitempty"`
}

// UsageContext is documented here http://hl7.org/fhir/StructureDefinition/UsageContext
type UsageContext struct {
	ID                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	Code                 Coding          `json:"code"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueReference       Reference       `json:"valueReference"`
}

// RelatedArtifact is documented here http://hl7.org/fhir/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	ID                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	Classifier        []CodeableConcept  `json:"classifier,omitempty"`
	Label             *string            `json:"label,omitempty"`
	Display           *string            `json:"display,omitempty"`
	Citation          *string            `json:"citation,omitempty"`
	Document          *Attachment        `json:"document,omitempty"`
	Resource          *string            `json:"resource,omitempty"`
	ResourceReference *Reference         `json:"resourceReference,omitempty"`
	PublicationStatus *PublicationStatus `json:"publicationStatus,omitempty"`
	PublicationDate   *string            `json:"publicationDate,omitempty"`
}

// FHIRContactDetail is documented here http://hl7.org/fhir/StructureDefinition/ContactDetail
type ContactDetail struct {
	ID        *string        `json:"id,omitempty"`
	Extension []Extension    `json:"extension,omitempty"`
	Name      *string        `json:"name,omitempty"`
	Telecom   []ContactPoint `json:"telecom,omitempty"`
}

// FHIRRatioRange is documented here http://hl7.org/fhir/StructureDefinition/RatioRange
type RatioRange struct {
	ID            *string     `json:"id,omitempty"`
	Extension     []Extension `json:"extension,omitempty"`
	LowNumerator  *Quantity   `json:"lowNumerator,omitempty"`
	HighNumerator *Quantity   `json:"highNumerator,omitempty"`
	Denominator   *Quantity   `json:"denominator,omitempty"`
}

type HTTPVerb string

const (
	HTTPVerbGET    HTTPVerb = "GET"
	HTTPVerbHEAD   HTTPVerb = "HEAD"
	HTTPVerbPOST   HTTPVerb = "POST"
	HTTPVerbPUT    HTTPVerb = "PUT"
	HTTPVerbDELETE HTTPVerb = "DELETE"
	HTTPVerbPATCH  HTTPVerb = "PATCH"
)

func (code HTTPVerb) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code HTTPVerb) String() string {
	return code.Code()
}

func (code HTTPVerb) Code() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbHEAD:
		return "HEAD"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	case HTTPVerbPATCH:
		return "PATCH"
	}

	return "unknown HTTP verb"
}

// PublicationStatus is documented here http://hl7.org/fhir/ValueSet/publication-status
type PublicationStatus string

const (
	PublicationStatusDraft   PublicationStatus = "draft"
	PublicationStatusActive  PublicationStatus = "active"
	PublicationStatusRetired PublicationStatus = "retired"
	PublicationStatusUnknown PublicationStatus = "unknown"
)

type FilterOperator string

const (
	FilterOperatorEquals         FilterOperator = "="
	FilterOperatorIsA            FilterOperator = "is-a"
	FilterOperatorDescendentOf   FilterOperator = "descendent-of"
	FilterOperatorIsNotA         FilterOperator = "is-not-a"
	FilterOperatorRegex          FilterOperator = "regex"
	FilterOperatorIn             FilterOperator = "in"
	FilterOperatorNotIn          FilterOperator = "not-in"
	FilterOperatorGeneralizes    FilterOperator = "generalizes"
	FilterOperatorChildOf        FilterOperator = "child-of"
	FilterOperatorDescendentLeaf FilterOperator = "descendent-leaf"
	FilterOperatorExists         FilterOperator = "exists"
)

// CodeSystemHierarchyMeaning is documented here http://hl7.org/fhir/ValueSet/codesystem-hierarchy-meaning
type CodeSystemHierarchyMeaning string

const (
	CodeSystemHierarchyMeaningGroupedBy      CodeSystemHierarchyMeaning = "grouped-by"
	CodeSystemHierarchyMeaningIsA            CodeSystemHierarchyMeaning = "is-a"
	CodeSystemHierarchyMeaningPartOf         CodeSystemHierarchyMeaning = "part-of"
	CodeSystemHierarchyMeaningClassifiedWith CodeSystemHierarchyMeaning = "classified-with"
)

// CodeSystemContentMode is documented here http://hl7.org/fhir/ValueSet/codesystem-content-mode
type CodeSystemContentMode string

const (
	CodeSystemContentModeNotPresent CodeSystemContentMode = "not-present"
	CodeSystemContentModeExample    CodeSystemContentMode = "example"
	CodeSystemContentModeFragment   CodeSystemContentMode = "fragment"
	CodeSystemContentModeComplete   CodeSystemContentMode = "complete"
	CodeSystemContentModeSupplement CodeSystemContentMode = "supplement"
)

type PropertyType string

const (
	PropertyTypeCode     PropertyType = "code"
	PropertyTypeCoding   PropertyType = "coding"
	PropertyTypeString   PropertyType = "string"
	PropertyTypeInteger  PropertyType = "integer"
	PropertyTypeBoolean  PropertyType = "boolean"
	PropertyTypeDateTime PropertyType = "dateTime"
	PropertyTypeDecimal  PropertyType = "decimal"
)
