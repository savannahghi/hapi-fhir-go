package fhir430

import "encoding/json"

// Subscription is documented here http://hl7.org/fhir/StructureDefinition/Subscription
// The subscription resource is used to define a push-based subscription from a server to another system. Once a subscription is registered with the server, the server checks every resource that is created or updated, and if the resource matches the given criteria, it sends a message on the defined "channel" so that another system can take an appropriate action.
type Subscription struct {
	ID                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Contained         []json.RawMessage   `json:"contained,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Status            SubscriptionStatus  `json:"status"`
	Contact           []ContactPoint      `json:"contact,omitempty"`
	End               *string             `json:"end,omitempty"`
	Reason            string              `json:"reason"`
	Criteria          string              `json:"criteria"`
	Error             *string             `json:"error,omitempty"`
	Channel           SubscriptionChannel `json:"channel"`
}

// Details where to send notifications when resources are received that meet the criteria.
type SubscriptionChannel struct {
	ID                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Type              SubscriptionChannelType `json:"type"`
	Endpoint          *string                 `json:"endpoint,omitempty"`
	Payload           *string                 `json:"payload,omitempty"`
	Header            []string                `json:"header,omitempty"`
}

// This function returns resource reference information
func (r Subscription) ResourceRef() (string, *string) {
	return "Subscription", r.ID
}

// This function returns resource reference information
func (r Subscription) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSubscription Subscription

// MarshalJSON marshals the given Subscription as JSON into a byte slice
func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscription
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscription: OtherSubscription(r),
		ResourceType:      "Subscription",
	})
}

// UnmarshalSubscription unmarshals a Subscription.
func UnmarshalSubscription(b []byte) (Subscription, error) {
	var subscription Subscription
	if err := json.Unmarshal(b, &subscription); err != nil {
		return subscription, err
	}
	return subscription, nil
}
