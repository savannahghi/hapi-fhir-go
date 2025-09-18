
package fhir430
// Address is documented here http://hl7.org/fhir/StructureDefinition/Address
// Base StructureDefinition for Address Type: An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world.
type Address struct {
	ID         *string      `json:"ID,omitempty"`
	Extension  []Extension  `json:"extension,omitempty"`
	Use        *AddressUse  `json:"use,omitempty"`
	Type       *AddressType `json:"type,omitempty"`
	Text       *string      `json:"text,omitempty"`
	Line       []string     `json:"line,omitempty"`
	City       *string      `json:"city,omitempty"`
	District   *string      `json:"district,omitempty"`
	State      *string      `json:"state,omitempty"`
	PostalCode *string      `json:"postalCode,omitempty"`
	Country    *string      `json:"country,omitempty"`
	Period     *Period      `json:"period,omitempty"`
}
