package models

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/savannahghi/scalarutils"
)

// FHIRAddress definition: an address expressed using postal conventions (as
// opposed to gps or other location definition formats).  this data type may
// be used to convey addresses for use in delivering mail as well as for visiting
// locations which might not be valid for mail delivery.  there are a variety of postal address
// formats defined around the world.
type FHIRAddress struct {
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
	PostalCode *scalarutils.Code `json:"postalCode,omitempty"`

	// Country - a nation as commonly understood or generally accepted.
	Country *string `json:"country,omitempty"`

	// Time period when address was/is in use.
	Period *FHIRPeriod `json:"period,omitempty"`
}

// FHIRAge definition: a duration of time during which an organism (or a
// process) has existed.
type FHIRAge struct {
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
	System *scalarutils.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *scalarutils.Code `json:"code,omitempty"`
}

// FHIRAnnotation definition: a  text note which also  contains information
// about who made the statement and when.
type FHIRAnnotation struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The individual responsible for making the annotation.
	AuthorReference *FHIRReference `json:"authorReference,omitempty"`

	// The individual responsible for making the annotation.
	AuthorString *string `json:"authorString,omitempty"`

	// Indicates when this particular annotation was made.
	Time *time.Time `json:"time,omitempty"`

	// The text of the annotation in markdown format.
	Text *scalarutils.Markdown `json:"text,omitempty"`
}

// FHIRAttachment definition: for referring to data content defined in other
// formats.
type FHIRAttachment struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies the type of the data in the attachment and allows a method to be
	// chosen to interpret or render the data. Includes mime type parameters such
	// as charset where appropriate.
	ContentType *scalarutils.Code `json:"contentType,omitempty"`

	// The human language of the content. The value can be any valid value
	// according to BCP 47.
	Language *scalarutils.Code `json:"language,omitempty"`

	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	Data *scalarutils.Base64Binary `json:"data,omitempty"`

	// A location where the data can be accessed.
	URL *scalarutils.URL `json:"url,omitempty"`

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
type FHIRCodeableConcept struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A reference to a code defined by a terminology system.
	Coding []*FHIRCoding `json:"coding,omitempty"`

	// A human language representation of the concept as seen/selected/uttered by
	// the user who entered the data and/or which represents the intended meaning
	// of the user.
	Text string `json:"text,omitempty"`
}

// FHIRCoding definition: a reference to a code defined by a terminology
// system.
type FHIRCoding struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The identification of the code system that defines the meaning of the
	// symbol in the code.
	System *scalarutils.URI `json:"system,omitempty"`

	// The version of the code system which was used when choosing this code. Note
	// that a well-maintained code system does not need the version reported,
	// because the meaning of codes is consistent across versions.
	// However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent,
	// the version SHOULD be exchanged.
	Version *string `json:"version,omitempty"`

	// A symbol in syntax defined by the system. The symbol may be a predefined
	// code or an expression in a syntax defined by the coding system (e.g.
	// post-coordination).
	Code *scalarutils.Code `json:"code,omitempty"`

	// A representation of the meaning of the code in the system, following the
	// rules of the system.
	Display string `json:"display,omitempty"`

	// Indicates that this coding was chosen by a user directly - e.g. off a pick
	// list of available items (codes or displays).
	UserSelected *bool `json:"userSelected,omitempty"`
}

// ToString returns the Display field of the Coding struct as a string.
func (c *FHIRCoding) ToString() string {
	return c.Display
}

// FHIRContactPoint definition: details for all kinds of technology mediated
// contact points for a person or organization, including telephone, email,
// etc.
type FHIRContactPoint struct {
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
	Period *FHIRPeriod `json:"period,omitempty"`
}

// FHIRHumanName definition: a human's name with the ability to identify parts
// and usage.
type FHIRHumanName struct {
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
	Period *FHIRPeriod `json:"period,omitempty"`
}

// FHIRIdentifier definition: an identifier - identifies some entity uniquely
// and unambiguously. typically this is used for business identifiers.
type FHIRIdentifier struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The purpose of this identifier.
	Use IdentifierUseEnum `json:"use,omitempty"`

	// A coded type for the identifier that can be used to determine which
	// identifier to use for a specific purpose.
	Type FHIRCodeableConcept `json:"type,omitempty"`

	// Establishes the namespace for the value - that is, a URL that describes a
	// set values that are unique.
	System *scalarutils.URI `json:"system,omitempty"`

	// The portion of the identifier typically relevant to the user and which is
	// unique within the context of the system.
	Value string `json:"value,omitempty"`

	// Time period during which identifier is/was valid for use.
	Period *FHIRPeriod `json:"period,omitempty"`

	// Organization that issued/manages the identifier.
	Assigner *FHIRReference `json:"assigner,omitempty"`
}

// FHIRNarrative definition: a human-readable summary of the resource
// conveying the essential clinical and business information for the resource.
type FHIRNarrative struct {
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

// NarrativeStatusEnum is a FHIR enum.
type NarrativeStatusEnum string

const (
	// NarrativeStatusEnumGenerated ...
	NarrativeStatusEnumGenerated NarrativeStatusEnum = "generated"
	// NarrativeStatusEnumExtensions ...
	NarrativeStatusEnumExtensions NarrativeStatusEnum = "extensions"
	// NarrativeStatusEnumAdditional ...
	NarrativeStatusEnumAdditional NarrativeStatusEnum = "additional"
	// NarrativeStatusEnumEmpty ...
	NarrativeStatusEnumEmpty NarrativeStatusEnum = "empty"
)

// IsValid ...
func (e NarrativeStatusEnum) IsValid() bool {
	switch e {
	case NarrativeStatusEnumGenerated, NarrativeStatusEnumExtensions, NarrativeStatusEnumAdditional, NarrativeStatusEnumEmpty:
		return true
	}

	return false
}

// String ...
func (e NarrativeStatusEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e NarrativeStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = NarrativeStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NarrativeStatusEnum", str)
	}

	return nil
}

// MarshalGQL writes the given enum to the supplied writer as a quoted string.
func (e NarrativeStatusEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// FHIRPeriod definition: a time period defined by a start and end date and
// optionally time.
type FHIRPeriod struct {
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
type FHIRReference struct {
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
	Type *scalarutils.URI `json:"type,omitempty"`

	// An identifier for the target resource. This is used when there is no way to
	// reference the other resource directly, either because the entity it
	// represents is not available through a FHIR server, or because there is no way for the author of the resource to convert
	//  a known identifier to an actual location.
	// There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance,
	// but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance,
	// and that instance would need to be of a FHIR resource type allowed by the reference.
	Identifier *FHIRIdentifier `json:"identifier,omitempty"`

	// Plain text narrative that identifies the resource in addition to the
	// resource reference.
	Display string `json:"display,omitempty"`
}

// FHIRExpression is documented here
// http://hl7.org/fhir/StructureDefinition/Expression
type FHIRExpression struct {
	ID          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	Description *string     `json:"description,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Language    string      `json:"language,omitempty"`
	Expression  *string     `json:"expression,omitempty"`
	Reference   *string     `json:"reference,omitempty"`
}

// AddressTypeEnum is a FHIR enum.
type AddressTypeEnum string

const (
	// AddressTypeEnumPostal ...
	AddressTypeEnumPostal AddressTypeEnum = "postal"
	// AddressTypeEnumPhysical ...
	AddressTypeEnumPhysical AddressTypeEnum = "physical"
	// AddressTypeEnumBoth ...
	AddressTypeEnumBoth AddressTypeEnum = "both"
)

// IsValid ...
func (e AddressTypeEnum) IsValid() bool {
	switch e {
	case AddressTypeEnumPostal, AddressTypeEnumPhysical, AddressTypeEnumBoth:
		return true
	}

	return false
}

// String ...
func (e AddressTypeEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e AddressTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = AddressTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddressTypeEnum", str)
	}

	return nil
}

// MarshalGQL writes the address type enum to the supplied writer as a quoted
// string.
func (e AddressTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// AddressUseEnum is a FHIR enum.
type AddressUseEnum string

const (
	// AddressUseEnumHome ...
	AddressUseEnumHome AddressUseEnum = "home"
	// AddressUseEnumWork ...
	AddressUseEnumWork AddressUseEnum = "work"
	// AddressUseEnumTemp ...
	AddressUseEnumTemp AddressUseEnum = "temp"
	// AddressUseEnumOld ...
	AddressUseEnumOld AddressUseEnum = "old"
	// AddressUseEnumBilling ...
	AddressUseEnumBilling AddressUseEnum = "billing"
)

// IsValid ...
func (e AddressUseEnum) IsValid() bool {
	switch e {
	case AddressUseEnumHome, AddressUseEnumWork, AddressUseEnumTemp, AddressUseEnumOld, AddressUseEnumBilling:
		return true
	}

	return false
}

// String ...
func (e AddressUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e AddressUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = AddressUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddressUseEnum", str)
	}

	return nil
}

// MarshalGQL writes the address use enum to the supplied writer as a quoted
// string.
func (e AddressUseEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// AgeComparatorEnum is a FHIR enum.
type AgeComparatorEnum string

const (
	// AgeComparatorEnumLessThan ...
	AgeComparatorEnumLessThan AgeComparatorEnum = "less_than"
	// AgeComparatorEnumLessThanOrEqualTo ...
	AgeComparatorEnumLessThanOrEqualTo AgeComparatorEnum = "less_than_or_equal_to"
	// AgeComparatorEnumGreaterThanOrEqualTo ...
	AgeComparatorEnumGreaterThanOrEqualTo AgeComparatorEnum = "greater_than_or_equal_to"
	// AgeComparatorEnumGreaterThan ...
	AgeComparatorEnumGreaterThan AgeComparatorEnum = "greater_than"
)

// IsValid ...
func (e AgeComparatorEnum) IsValid() bool {
	switch e {
	case AgeComparatorEnumLessThan, AgeComparatorEnumLessThanOrEqualTo, AgeComparatorEnumGreaterThanOrEqualTo, AgeComparatorEnumGreaterThan:
		return true
	}

	return false
}

// String renders an age comparator enum as a string.
func (e AgeComparatorEnum) String() string {
	switch e {
	case AgeComparatorEnumLessThan:
		return "<"
	case AgeComparatorEnumLessThanOrEqualTo:
		return "<="
	case AgeComparatorEnumGreaterThanOrEqualTo:
		return ">="
	case AgeComparatorEnumGreaterThan:
		return ">"
	}

	return string(e)
}

// UnmarshalGQL ...
func (e AgeComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = AgeComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AgeComparatorEnum", str)
	}

	return nil
}

// MarshalGQL writes the age comparator to the supplied writer as a quoted
// string.
func (e AgeComparatorEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// ContactPointSystemEnum is a FHIR enum.
type ContactPointSystemEnum string

const (
	// ContactPointSystemEnumPhone ...
	ContactPointSystemEnumPhone ContactPointSystemEnum = "phone"
	// ContactPointSystemEnumFax ...
	ContactPointSystemEnumFax ContactPointSystemEnum = "fax"
	// ContactPointSystemEnumEmail ...
	ContactPointSystemEnumEmail ContactPointSystemEnum = "email"
	// ContactPointSystemEnumPager ...
	ContactPointSystemEnumPager ContactPointSystemEnum = "pager"
	// ContactPointSystemEnumURL ...
	ContactPointSystemEnumURL ContactPointSystemEnum = "url"
	// ContactPointSystemEnumSms ...
	ContactPointSystemEnumSms ContactPointSystemEnum = "sms"
	// ContactPointSystemEnumOther ...
	ContactPointSystemEnumOther ContactPointSystemEnum = "other"
)

// IsValid ...
func (e ContactPointSystemEnum) IsValid() bool {
	switch e {
	case ContactPointSystemEnumPhone, ContactPointSystemEnumFax,
		ContactPointSystemEnumEmail, ContactPointSystemEnumPager,
		ContactPointSystemEnumURL, ContactPointSystemEnumSms, ContactPointSystemEnumOther:
		return true
	}

	return false
}

// String ...
func (e ContactPointSystemEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e ContactPointSystemEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = ContactPointSystemEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactPointSystemEnum", str)
	}

	return nil
}

// MarshalGQL writes the given enum to the supplied writer as a quoted string.
func (e ContactPointSystemEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// ContactPointUseEnum is a FHIR enum.
type ContactPointUseEnum string

const (
	// ContactPointUseEnumHome ...
	ContactPointUseEnumHome ContactPointUseEnum = "home"
	// ContactPointUseEnumWork ...
	ContactPointUseEnumWork ContactPointUseEnum = "work"
	// ContactPointUseEnumTemp ...
	ContactPointUseEnumTemp ContactPointUseEnum = "temp"
	// ContactPointUseEnumOld ...
	ContactPointUseEnumOld ContactPointUseEnum = "old"
	// ContactPointUseEnumMobile ...
	ContactPointUseEnumMobile ContactPointUseEnum = "mobile"
)

// IsValid ...
func (e ContactPointUseEnum) IsValid() bool {
	switch e {
	case ContactPointUseEnumHome, ContactPointUseEnumWork,
		ContactPointUseEnumTemp, ContactPointUseEnumOld, ContactPointUseEnumMobile:
		return true
	}

	return false
}

// String ...
func (e ContactPointUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e ContactPointUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = ContactPointUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactPointUseEnum", str)
	}

	return nil
}

// MarshalGQL writes the contact point use to the supplied writer as a quoted
// string.
func (e ContactPointUseEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// DurationComparatorEnum is a FHIR enum.
type DurationComparatorEnum string

const (
	// DurationComparatorEnumLessThan ...
	DurationComparatorEnumLessThan DurationComparatorEnum = "less_than"
	// DurationComparatorEnumLessThanOrEqualTo ...
	DurationComparatorEnumLessThanOrEqualTo DurationComparatorEnum = "less_than_or_equal_to"
	// DurationComparatorEnumGreaterThanOrEqualTo ...
	DurationComparatorEnumGreaterThanOrEqualTo DurationComparatorEnum = "greater_than_or_equal_to"
	// DurationComparatorEnumGreaterThan ...
	DurationComparatorEnumGreaterThan DurationComparatorEnum = "greater_than"
)

// IsValid ...
func (e DurationComparatorEnum) IsValid() bool {
	switch e {
	case DurationComparatorEnumLessThan, DurationComparatorEnumLessThanOrEqualTo,
		DurationComparatorEnumGreaterThanOrEqualTo, DurationComparatorEnumGreaterThan:
		return true
	}

	return false
}

// String ...
func (e DurationComparatorEnum) String() string {
	switch e {
	case DurationComparatorEnumLessThan:
		return "<"
	case DurationComparatorEnumLessThanOrEqualTo:
		return "<="
	case DurationComparatorEnumGreaterThan:
		return ">"
	case DurationComparatorEnumGreaterThanOrEqualTo:
		return ">="
	}

	return string(e)
}

// UnmarshalGQL ...
func (e DurationComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = DurationComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DurationComparatorEnum", str)
	}

	return nil
}

// MarshalGQL writes the duration comparator to the supplied writer as a
// quoted string.
func (e DurationComparatorEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// HumanNameUseEnum is a FHIR enum.
type HumanNameUseEnum string

const (
	// HumanNameUseEnumUsual ...
	HumanNameUseEnumUsual HumanNameUseEnum = "usual"
	// HumanNameUseEnumOfficial ...
	HumanNameUseEnumOfficial HumanNameUseEnum = "official"
	// HumanNameUseEnumTemp ...
	HumanNameUseEnumTemp HumanNameUseEnum = "temp"
	// HumanNameUseEnumNickname ...
	HumanNameUseEnumNickname HumanNameUseEnum = "nickname"
	// HumanNameUseEnumAnonymous ...
	HumanNameUseEnumAnonymous HumanNameUseEnum = "anonymous"
	// HumanNameUseEnumOld ...
	HumanNameUseEnumOld HumanNameUseEnum = "old"
	// HumanNameUseEnumMaiden ...
	HumanNameUseEnumMaiden HumanNameUseEnum = "maiden"
)

// IsValid ...
func (e HumanNameUseEnum) IsValid() bool {
	switch e {
	case HumanNameUseEnumUsual, HumanNameUseEnumOfficial,
		HumanNameUseEnumTemp, HumanNameUseEnumNickname, HumanNameUseEnumAnonymous,
		HumanNameUseEnumOld, HumanNameUseEnumMaiden:
		return true
	}

	return false
}

// String ...
func (e HumanNameUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e HumanNameUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = HumanNameUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HumanNameUseEnum", str)
	}

	return nil
}

// MarshalGQL writes the given enum to the supplied writer as a quoted string.
func (e HumanNameUseEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// IdentifierUseEnum is a FHIR enum.
type IdentifierUseEnum string

const (
	// IdentifierUseEnumUsual ...
	IdentifierUseEnumUsual IdentifierUseEnum = "usual"
	// IdentifierUseEnumOfficial ...
	IdentifierUseEnumOfficial IdentifierUseEnum = "official"
	// IdentifierUseEnumTemp ...
	IdentifierUseEnumTemp IdentifierUseEnum = "temp"
	// IdentifierUseEnumSecondary ...
	IdentifierUseEnumSecondary IdentifierUseEnum = "secondary"
	// IdentifierUseEnumOld ...
	IdentifierUseEnumOld IdentifierUseEnum = "old"
)

// IsValid ...
func (e IdentifierUseEnum) IsValid() bool {
	switch e {
	case IdentifierUseEnumUsual, IdentifierUseEnumOfficial, IdentifierUseEnumTemp,
		IdentifierUseEnumSecondary, IdentifierUseEnumOld:
		return true
	}

	return false
}

// String ...
func (e IdentifierUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e IdentifierUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = IdentifierUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IdentifierUseEnum", str)
	}

	return nil
}

// MarshalGQL writes the identifier use to the supplied writer as a quoted
// string.
func (e IdentifierUseEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// PatientGenderEnum is a FHIR enum.
type PatientGenderEnum string

const (
	// PatientGenderEnumMale ...
	PatientGenderEnumMale PatientGenderEnum = "male"
	// PatientGenderEnumFemale ...
	PatientGenderEnumFemale PatientGenderEnum = "female"
	// PatientGenderEnumOther ...
	PatientGenderEnumOther PatientGenderEnum = "other"
	// PatientGenderEnumUnknown ...
	PatientGenderEnumUnknown PatientGenderEnum = "unknown"
)

// IsValid ...
func (e PatientGenderEnum) IsValid() bool {
	switch e {
	case PatientGenderEnumMale, PatientGenderEnumFemale, PatientGenderEnumOther, PatientGenderEnumUnknown:
		return true
	}

	return false
}

// String ...
func (e PatientGenderEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e PatientGenderEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = PatientGenderEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PatientGenderEnum", str)
	}

	return nil
}

// MarshalGQL writes the patient gender to the supplied writer as a quoted
// string.
func (e PatientGenderEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// PatientContactGenderEnum is a FHIR enum.
type PatientContactGenderEnum string

const (
	// PatientContactGenderEnumMale ...
	PatientContactGenderEnumMale PatientContactGenderEnum = "male"
	// PatientContactGenderEnumFemale ...
	PatientContactGenderEnumFemale PatientContactGenderEnum = "female"
	// PatientContactGenderEnumOther ...
	PatientContactGenderEnumOther PatientContactGenderEnum = "other"
	// PatientContactGenderEnumUnknown ...
	PatientContactGenderEnumUnknown PatientContactGenderEnum = "unknown"
)

// IsValid ...
func (e PatientContactGenderEnum) IsValid() bool {
	switch e {
	case PatientContactGenderEnumMale, PatientContactGenderEnumFemale,
		PatientContactGenderEnumOther, PatientContactGenderEnumUnknown:
		return true
	}

	return false
}

// String ...
func (e PatientContactGenderEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e PatientContactGenderEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = PatientContactGenderEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Patient_ContactGenderEnum", str)
	}

	return nil
}

// MarshalGQL writes the patient contact gender to the supplied writer as a
// quoted string.
func (e PatientContactGenderEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// PatientLinkTypeEnum is a FHIR enum.
type PatientLinkTypeEnum string

const (
	// PatientLinkTypeEnumReplacedBy ...
	PatientLinkTypeEnumReplacedBy PatientLinkTypeEnum = "replaced_by"
	// PatientLinkTypeEnumReplaces ...
	PatientLinkTypeEnumReplaces PatientLinkTypeEnum = "replaces"
	// PatientLinkTypeEnumRefer ...
	PatientLinkTypeEnumRefer PatientLinkTypeEnum = "refer"
	// PatientLinkTypeEnumSeealso ...
	PatientLinkTypeEnumSeealso PatientLinkTypeEnum = "seealso"
)

// IsValid ...
func (e PatientLinkTypeEnum) IsValid() bool {
	switch e {
	case PatientLinkTypeEnumReplacedBy, PatientLinkTypeEnumReplaces, PatientLinkTypeEnumRefer, PatientLinkTypeEnumSeealso:
		return true
	}

	return false
}

// String ...
func (e PatientLinkTypeEnum) String() string {
	if e == PatientLinkTypeEnumReplacedBy {
		return "replaced-by"
	}

	return string(e)
}

// UnmarshalGQL ...
func (e PatientLinkTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = PatientLinkTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Patient_LinkTypeEnum", str)
	}

	return nil
}

// MarshalGQL writes the patient link type to the supplied writer as a quoted
// string.
func (e PatientLinkTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// FHIRMeta is a set of metadata that provides technical and workflow context
// to a resource.
type FHIRMeta struct {
	VersionID   string       `json:"versionId,omitempty"`
	LastUpdated time.Time    `json:"lastUpdated,omitempty"`
	Source      string       `json:"source,omitempty"`
	Tag         []FHIRCoding `json:"tag,omitempty"`
	Security    []FHIRCoding `json:"security,omitempty"`
}

// FHIRQuantity definition: a measured amount (or an amount that can
// potentially be measured). note that measured amounts include amounts that
// are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type FHIRQuantity struct {
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
	System scalarutils.URI `json:"system"`

	// A computer processable form of the unit in some unit representation system.
	Code *scalarutils.Code `json:"code"`
}

// QuantityComparatorEnum is a FHIR enum.
type QuantityComparatorEnum string

const (
	// QuantityComparatorEnumLessThan ...
	QuantityComparatorEnumLessThan QuantityComparatorEnum = "less_than"
	// QuantityComparatorEnumLessThanOrEqualTo ...
	QuantityComparatorEnumLessThanOrEqualTo QuantityComparatorEnum = "less_than_or_equal_to"
	// QuantityComparatorEnumGreaterThanOrEqualTo ...
	QuantityComparatorEnumGreaterThanOrEqualTo QuantityComparatorEnum = "greater_than_or_equal_to"
	// QuantityComparatorEnumGreaterThan ...
	QuantityComparatorEnumGreaterThan QuantityComparatorEnum = "greater_than"
)

// IsValid ...
func (e QuantityComparatorEnum) IsValid() bool {
	switch e {
	case QuantityComparatorEnumLessThan, QuantityComparatorEnumLessThanOrEqualTo, QuantityComparatorEnumGreaterThanOrEqualTo,
		QuantityComparatorEnumGreaterThan:
		return true
	}

	return false
}

// String ...
func (e QuantityComparatorEnum) String() string {
	switch e {
	case QuantityComparatorEnumLessThan:
		return "<"
	case QuantityComparatorEnumLessThanOrEqualTo:
		return "<="
	case QuantityComparatorEnumGreaterThanOrEqualTo:
		return ">="
	case QuantityComparatorEnumGreaterThan:
		return ">"
	default:
		return string(e)
	}
}

// UnmarshalGQL ...
func (e QuantityComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	e = QuantityComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid QuantityComparatorEnum", str)
	}

	return nil
}

// MarshalGQL writes the quality comparator to the supplied writer as a quoted string.
func (e QuantityComparatorEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// FHIRRange definition: a set of ordered quantities defined by a low and high limit.
type FHIRRange struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The low limit. The boundary is inclusive.
	Low FHIRQuantity `json:"low,omitempty"`

	// The high limit. The boundary is inclusive.
	High FHIRQuantity `json:"high,omitempty"`
}

// FHIRRatio definition: a relationship of two quantity values - expressed as a numerator and a denominator.
type FHIRRatio struct {
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the numerator.
	Numerator FHIRQuantity `json:"numerator,omitempty"`

	// The value of the denominator.
	Denominator FHIRQuantity `json:"denominator,omitempty"`
}

// Extension is an optional element that provides additional information not
// captured in the basic resource definition.
// Extensions allow the definition of new data elements or the modification of
// existing data elements in the FHIR data model.
type Extension struct {
	URL                  string               `json:"url,omitempty"`
	ValueBoolean         bool                 `json:"valueBoolean,omitempty"`
	ValueInteger         *int                 `json:"valueInteger,omitempty"`
	ValueDecimal         *float64             `json:"valueDecimal,omitempty"`
	ValueBase64Binary    string               `json:"valueBase64Binary,omitempty"`
	ValueInstant         string               `json:"valueInstant,omitempty"`
	ValueString          string               `json:"valueString,omitempty"`
	ValueURI             string               `json:"valueURI,omitempty"`
	ValueDate            string               `json:"valueDate,omitempty"`
	ValueDateTime        string               `json:"valueDateTime,omitempty"`
	ValueTime            string               `json:"valueTime,omitempty"`
	ValueCode            string               `json:"valueCode,omitempty"`
	ValueOid             string               `json:"valueOid,omitempty"`
	ValueUUID            string               `json:"valueUUID,omitempty"`
	ValueID              string               `json:"valueID,omitempty"`
	ValueUnsignedInt     int                  `json:"valueUnsignedInt,omitempty"`
	ValuePositiveInt     int                  `json:"valuePositiveInt,omitempty"`
	ValueMarkdown        string               `json:"valueMarkdown,omitempty"`
	ValueAnnotation      *FHIRAnnotation      `json:"valueAnnotation,omitempty"`
	ValueAttachment      *FHIRAttachment      `json:"valueAttachment,omitempty"`
	ValueIdentifier      *FHIRIdentifier      `json:"valueIdentifier,omitempty"`
	ValueCodeableConcept *FHIRCodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueCoding          *FHIRCoding          `json:"valueCoding,omitempty"`
	ValueQuantity        *FHIRQuantity        `json:"valueQuantity,omitempty"`
	ValueRange           *FHIRRange           `json:"valueRange,omitempty"`
	ValuePeriod          *FHIRPeriod          `json:"valuePeriod,omitempty"`
	ValueRatio           *FHIRRatio           `json:"valueRatio,omitempty"`
	ValueReference       *FHIRReference       `json:"valueReference,omitempty"`
	ValueExpression      *FHIRExpression      `json:"valueExpression,omitempty"`
}

// FHIRExtension contains child elements to represent additional information
// that is not part of the basic definition of the resource.
type FHIRExtension struct {
	URL       string      `json:"url,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
}
