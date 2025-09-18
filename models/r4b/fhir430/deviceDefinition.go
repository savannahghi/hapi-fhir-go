
package fhir430

import "encoding/json"
// DeviceDefinition is documented here http://hl7.org/fhir/StructureDefinition/DeviceDefinition
// The characteristics, operational status and capabilities of a medical-related component of a medical device.
type DeviceDefinition struct {
	ID                      *string                               `json:"ID,omitempty"`
	Meta                    *Meta                                 `json:"meta,omitempty"`
	ImplicitRules           *string                               `json:"implicitRules,omitempty"`
	Language                *string                               `json:"language,omitempty"`
	Text                    *Narrative                            `json:"text,omitempty"`
	Contained               []json.RawMessage                     `json:"contained,omitempty"`
	Extension               []Extension                           `json:"extension,omitempty"`
	ModifierExtension       []Extension                           `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                          `json:"identifier,omitempty"`
	UdiDeviceIdentifier     []DeviceDefinitionUdiDeviceIdentifier `json:"udiDeviceIdentifier,omitempty"`
	ManufacturerString      *string                               `json:"manufacturerString,omitempty"`
	ManufacturerReference   *Reference                            `json:"manufacturerReference,omitempty"`
	DeviceName              []DeviceDefinitionDeviceName          `json:"deviceName,omitempty"`
	ModelNumber             *string                               `json:"modelNumber,omitempty"`
	Type                    *CodeableConcept                      `json:"type,omitempty"`
	Specialization          []DeviceDefinitionSpecialization      `json:"specialization,omitempty"`
	Version                 []string                              `json:"version,omitempty"`
	Safety                  []CodeableConcept                     `json:"safety,omitempty"`
	ShelfLifeStorage        []ProductShelfLife                    `json:"shelfLifeStorage,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic                   `json:"physicalCharacteristics,omitempty"`
	LanguageCode            []CodeableConcept                     `json:"languageCode,omitempty"`
	Capability              []DeviceDefinitionCapability          `json:"capability,omitempty"`
	Property                []DeviceDefinitionProperty            `json:"property,omitempty"`
	Owner                   *Reference                            `json:"owner,omitempty"`
	Contact                 []ContactPoint                        `json:"contact,omitempty"`
	Url                     *string                               `json:"url,omitempty"`
	OnlineInformation       *string                               `json:"onlineInformation,omitempty"`
	Note                    []Annotation                          `json:"note,omitempty"`
	Quantity                *Quantity                             `json:"quantity,omitempty"`
	ParentDevice            *Reference                            `json:"parentDevice,omitempty"`
	Material                []DeviceDefinitionMaterial            `json:"material,omitempty"`
}

// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
type DeviceDefinitionUdiDeviceIdentifier struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DeviceIdentifier  string      `json:"deviceIdentifier"`
	Issuer            string      `json:"issuer"`
	Jurisdiction      string      `json:"jurisdiction"`
}

// A name given to the device to identify it.
type DeviceDefinitionDeviceName struct {
	ID                *string        `json:"ID,omitempty"`
	Extension         []Extension    `json:"extension,omitempty"`
	ModifierExtension []Extension    `json:"modifierExtension,omitempty"`
	Name              string         `json:"name"`
	Type              DeviceNameType `json:"type"`
}

// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
type DeviceDefinitionSpecialization struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	SystemType        string      `json:"systemType"`
	Version           *string     `json:"version,omitempty"`
}

// Device capabilities.
type DeviceDefinitionCapability struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Description       []CodeableConcept `json:"description,omitempty"`
}

// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties.
type DeviceDefinitionProperty struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}

// A substance used to create the material(s) of which the device is made.
type DeviceDefinitionMaterial struct {
	ID                  *string         `json:"ID,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	Substance           CodeableConcept `json:"substance"`
	Alternate           *bool           `json:"alternate,omitempty"`
	AllergenicIndicator *bool           `json:"allergenicIndicator,omitempty"`
}

// This function returns resource reference information
func (r DeviceDefinition) ResourceRef() (string, *string) {
	return "DeviceDefinition", r.ID
}

// This function returns resource reference information
func (r DeviceDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDeviceDefinition DeviceDefinition

// MarshalJSON marshals the given DeviceDefinition as JSON into a byte slice
func (r DeviceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceDefinition: OtherDeviceDefinition(r),
		ResourceType:          "DeviceDefinition",
	})
}

// UnmarshalDeviceDefinition unmarshals a DeviceDefinition.
func UnmarshalDeviceDefinition(b []byte) (DeviceDefinition, error) {
	var deviceDefinition DeviceDefinition
	if err := json.Unmarshal(b, &deviceDefinition); err != nil {
		return deviceDefinition, err
	}
	return deviceDefinition, nil
}
