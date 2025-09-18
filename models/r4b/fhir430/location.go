
package fhir430

import "encoding/json"
// Location is documented here http://hl7.org/fhir/StructureDefinition/Location
// Details and position information for a physical place where services are provided and resources and participants may be stored, found, contained, or accommodated.
type Location struct {
	ID                     *string                    `json:"ID,omitempty"`
	Meta                   *Meta                      `json:"meta,omitempty"`
	ImplicitRules          *string                    `json:"implicitRules,omitempty"`
	Language               *string                    `json:"language,omitempty"`
	Text                   *Narrative                 `json:"text,omitempty"`
	Contained              []json.RawMessage          `json:"contained,omitempty"`
	Extension              []Extension                `json:"extension,omitempty"`
	ModifierExtension      []Extension                `json:"modifierExtension,omitempty"`
	Identifier             []Identifier               `json:"identifier,omitempty"`
	Status                 *LocationStatus            `json:"status,omitempty"`
	OperationalStatus      *Coding                    `json:"operationalStatus,omitempty"`
	Name                   *string                    `json:"name,omitempty"`
	Alias                  []string                   `json:"alias,omitempty"`
	Description            *string                    `json:"description,omitempty"`
	Mode                   *LocationMode              `json:"mode,omitempty"`
	Type                   []CodeableConcept          `json:"type,omitempty"`
	Telecom                []ContactPoint             `json:"telecom,omitempty"`
	Address                *Address                   `json:"address,omitempty"`
	PhysicalType           *CodeableConcept           `json:"physicalType,omitempty"`
	Position               *LocationPosition          `json:"position,omitempty"`
	ManagingOrganization   *Reference                 `json:"managingOrganization,omitempty"`
	PartOf                 *Reference                 `json:"partOf,omitempty"`
	HoursOfOperation       []LocationHoursOfOperation `json:"hoursOfOperation,omitempty"`
	AvailabilityExceptions *string                    `json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                `json:"endpoint,omitempty"`
}

// The absolute geographic location of the Location, expressed using the WGS84 datum (This is the same co-ordinate system used in KML).
type LocationPosition struct {
	ID                *string      `json:"ID,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Longitude         json.Number  `json:"longitude"`
	Latitude          json.Number  `json:"latitude"`
	Altitude          *json.Number `json:"altitude,omitempty"`
}

// What days/times during a week is this location usually open.
/*
This type of information is commonly found published in directories and on websites informing customers when the facility is available.

Specific services within the location may have their own hours which could be shorter (or longer) than the locations hours.
*/
type LocationHoursOfOperation struct {
	ID                *string      `json:"ID,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	DaysOfWeek        []DaysOfWeek `json:"daysOfWeek,omitempty"`
	AllDay            *bool        `json:"allDay,omitempty"`
	OpeningTime       *string      `json:"openingTime,omitempty"`
	ClosingTime       *string      `json:"closingTime,omitempty"`
}

// This function returns resource reference information
func (r Location) ResourceRef() (string, *string) {
	return "Location", r.ID
}

// This function returns resource reference information
func (r Location) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherLocation Location

// MarshalJSON marshals the given Location as JSON into a byte slice
func (r Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLocation
		ResourceType string `json:"resourceType"`
	}{
		OtherLocation: OtherLocation(r),
		ResourceType:  "Location",
	})
}

// UnmarshalLocation unmarshals a Location.
func UnmarshalLocation(b []byte) (Location, error) {
	var location Location
	if err := json.Unmarshal(b, &location); err != nil {
		return location, err
	}
	return location, nil
}
