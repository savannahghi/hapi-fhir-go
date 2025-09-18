
package fhir430

import "encoding/json"
// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
// A container for a collection of resources.
type Bundle struct {
	ID            *string       `json:"ID,omitempty"`
	Meta          *Meta         `json:"meta,omitempty"`
	ImplicitRules *string       `json:"implicitRules,omitempty"`
	Language      *string       `json:"language,omitempty"`
	Identifier    *Identifier   `json:"identifier,omitempty"`
	Type          BundleType    `json:"type"`
	Timestamp     *string       `json:"timestamp,omitempty"`
	Total         *int          `json:"total,omitempty"`
	Link          []BundleLink  `json:"link,omitempty"`
	Entry         []BundleEntry `json:"entry,omitempty"`
	Signature     *Signature    `json:"signature,omitempty"`
}

// A series of links that provide context to this bundle.
/*
Both Bundle.link and Bundle.entry.link are defined to support providing additional context when Bundles are used (e.g. [HATEOAS](http://en.wikipedia.org/wiki/HATEOAS)).

Bundle.entry.link corresponds to links found in the HTTP header if the resource in the entry was [read](http.html#read) directly.

This specification defines some specific uses of Bundle.link for [searching](search.html#conformance) and [paging](http.html#paging), but no specific uses for Bundle.entry.link, and no defined function in a transaction - the meaning is implementation specific.
*/
type BundleLink struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Relation          string      `json:"relation"`
	Url               string      `json:"url"`
}

// An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
type BundleEntry struct {
	ID                *string              `json:"ID,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Link              []BundleLink         `json:"link,omitempty"`
	FullUrl           *string              `json:"fullUrl,omitempty"`
	Resource          json.RawMessage      `json:"resource,omitempty"`
	Search            *BundleEntrySearch   `json:"search,omitempty"`
	Request           *BundleEntryRequest  `json:"request,omitempty"`
	Response          *BundleEntryResponse `json:"response,omitempty"`
}

// Information about the search process that lead to the creation of this entry.
type BundleEntrySearch struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Mode              *SearchEntryMode `json:"mode,omitempty"`
	Score             *json.Number     `json:"score,omitempty"`
}

// Additional information about how this entry should be processed as part of a transaction or batch.  For history, it shows how the entry was processed to create the version contained in the entry.
type BundleEntryRequest struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Method            HTTPVerb    `json:"method"`
	Url               string      `json:"url"`
	IfNoneMatch       *string     `json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `json:"ifMatch,omitempty"`
	IfNoneExist       *string     `json:"ifNoneExist,omitempty"`
}

// Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
type BundleEntryResponse struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Status            string          `json:"status"`
	Location          *string         `json:"location,omitempty"`
	Etag              *string         `json:"etag,omitempty"`
	LastModified      *string         `json:"lastModified,omitempty"`
	Outcome           json.RawMessage `json:"outcome,omitempty"`
}

// This function returns resource reference information
func (r Bundle) ResourceRef() (string, *string) {
	return "Bundle", r.ID
}

// This function returns resource reference information
func (r Bundle) ContainedResources() []json.RawMessage {
	return nil
}

type OtherBundle Bundle

// MarshalJSON marshals the given Bundle as JSON into a byte slice
func (r Bundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBundle
		ResourceType string `json:"resourceType"`
	}{
		OtherBundle:  OtherBundle(r),
		ResourceType: "Bundle",
	})
}

// UnmarshalBundle unmarshals a Bundle.
func UnmarshalBundle(b []byte) (Bundle, error) {
	var bundle Bundle
	if err := json.Unmarshal(b, &bundle); err != nil {
		return bundle, err
	}
	return bundle, nil
}
