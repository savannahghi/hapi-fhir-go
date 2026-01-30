package fhir430

import "encoding/json"

// Device is documented here http://hl7.org/fhir/StructureDefinition/Device
// A type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.
type Device struct {
	ID                 *string                `json:"id,omitempty"`
	Meta               *Meta                  `json:"meta,omitempty"`
	ImplicitRules      *string                `json:"implicitRules,omitempty"`
	Language           *string                `json:"language,omitempty"`
	Text               *Narrative             `json:"text,omitempty"`
	Contained          []json.RawMessage      `json:"contained,omitempty"`
	Extension          []Extension            `json:"extension,omitempty"`
	ModifierExtension  []Extension            `json:"modifierExtension,omitempty"`
	Identifier         []Identifier           `json:"identifier,omitempty"`
	Definition         *Reference             `json:"definition,omitempty"`
	UdiCarrier         []DeviceUdiCarrier     `json:"udiCarrier,omitempty"`
	Status             *FHIRDeviceStatus      `json:"status,omitempty"`
	StatusReason       []CodeableConcept      `json:"statusReason,omitempty"`
	DistinctIdentifier *string                `json:"distinctIdentifier,omitempty"`
	Manufacturer       *string                `json:"manufacturer,omitempty"`
	ManufactureDate    *string                `json:"manufactureDate,omitempty"`
	ExpirationDate     *string                `json:"expirationDate,omitempty"`
	LotNumber          *string                `json:"lotNumber,omitempty"`
	SerialNumber       *string                `json:"serialNumber,omitempty"`
	DeviceName         []DeviceDeviceName     `json:"deviceName,omitempty"`
	ModelNumber        *string                `json:"modelNumber,omitempty"`
	PartNumber         *string                `json:"partNumber,omitempty"`
	Type               *CodeableConcept       `json:"type,omitempty"`
	Specialization     []DeviceSpecialization `json:"specialization,omitempty"`
	Version            []DeviceVersion        `json:"version,omitempty"`
	Property           []DeviceProperty       `json:"property,omitempty"`
	Patient            *Reference             `json:"patient,omitempty"`
	Owner              *Reference             `json:"owner,omitempty"`
	Contact            []ContactPoint         `json:"contact,omitempty"`
	Location           *Reference             `json:"location,omitempty"`
	Url                *string                `json:"url,omitempty"`
	Note               []Annotation           `json:"note,omitempty"`
	Safety             []CodeableConcept      `json:"safety,omitempty"`
	Parent             *Reference             `json:"parent,omitempty"`
}

// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
// UDI may identify an unique instance of a device, or it may only identify the type of the device.  See [UDI mappings](device-mappings.html#udi) for a complete mapping of UDI parts to Device.
type DeviceUdiCarrier struct {
	ID                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	DeviceIdentifier  *string       `json:"deviceIdentifier,omitempty"`
	Issuer            *string       `json:"issuer,omitempty"`
	Jurisdiction      *string       `json:"jurisdiction,omitempty"`
	CarrierAIDC       *string       `json:"carrierAIDC,omitempty"`
	CarrierHRF        *string       `json:"carrierHRF,omitempty"`
	EntryType         *UDIEntryType `json:"entryType,omitempty"`
}

// This represents the manufacturer's name of the device as provided by the device, from a UDI label, or by a person describing the Device.  This typically would be used when a person provides the name(s) or when the device represents one of the names available from DeviceDefinition.
type DeviceDeviceName struct {
	ID                *string        `json:"id,omitempty"`
	Extension         []Extension    `json:"extension,omitempty"`
	ModifierExtension []Extension    `json:"modifierExtension,omitempty"`
	Name              string         `json:"name"`
	Type              DeviceNameType `json:"type"`
}

// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
type DeviceSpecialization struct {
	ID                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	SystemType        CodeableConcept `json:"systemType"`
	Version           *string         `json:"version,omitempty"`
}

// The actual design of the device or software version running on the device.
type DeviceVersion struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Component         *Identifier      `json:"component,omitempty"`
	Value             string           `json:"value"`
}

// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties.
type DeviceProperty struct {
	ID                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}

// This function returns resource reference information
func (r Device) ResourceRef() (string, *string) {
	return "Device", r.ID
}

// This function returns resource reference information
func (r Device) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDevice Device

// MarshalJSON marshals the given Device as JSON into a byte slice
func (r Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDevice
		ResourceType string `json:"resourceType"`
	}{
		OtherDevice:  OtherDevice(r),
		ResourceType: "Device",
	})
}

// UnmarshalDevice unmarshals a Device.
func UnmarshalDevice(b []byte) (Device, error) {
	var device Device
	if err := json.Unmarshal(b, &device); err != nil {
		return device, err
	}
	return device, nil
}
