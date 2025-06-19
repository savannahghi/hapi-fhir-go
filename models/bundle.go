package models

import (
	"encoding/json"
	"strings"
)

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
	ID            *string         `json:"id,omitempty"`
	Meta          *Meta           `json:"meta,omitempty"`
	ImplicitRules *string         `json:"implicitRules,omitempty"`
	Language      *string         `json:"language,omitempty"`
	Identifier    *Identifier     `json:"identifier,omitempty"`
	Type          BundleType      `json:"type"`
	Timestamp     *string         `json:"timestamp,omitempty"`
	Total         int             `json:"total,omitempty"`
	Link          []BundleLink    `json:"link,omitempty"`
	Entry         []BundleEntry   `json:"entry,omitempty"`
	Issues        json.RawMessage `json:"issues,omitempty"`
}

type BundleLink struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	URL               string      `json:"url"`
}

type BundleEntry struct {
	ID                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Link              []BundleLink `json:"link,omitempty"`
	// FullURL looks like
	//   urn:uuid:<ResourceType>-<domain-unique-identifier>
	// e.g. urn:uuid:Ingredient-VPI-00003943
	//
	// - “urn:uuid:”  — constant 9-char prefix recommended by HL7 for
	//                    transaction-local URNs
	// - <ResourceType> helps us decide which DB table to update once the
	//                    HAPI- transaction succeeds
	// - <identifier>  is our own PK coming from Postgres.
	FullURL  *BundleEntryFullURL  `json:"fullUrl,omitempty"`
	Resource json.RawMessage      `json:"resource,omitempty"`
	Search   *BundleEntrySearch   `json:"search,omitempty"`
	Request  *BundleEntryRequest  `json:"request,omitempty"`
	Response *BundleEntryResponse `json:"response,omitempty"`
}

type BundleEntryFullURL string

// GetTableAndUniqueID splits the fullUrl that was generated.
// (“urn:uuid:<ResourceType>-<identifier>”) and gives back the resource type
// and the identifier.
//
//	fullUrl := urn:uuid:Ingredient-VPI-00003943
//	tbl, id := fullUrl.GetTableAndUniqueID()
//	// tbl == "Ingredient"
//	// id  == "VPI-00003943"
func (f *BundleEntryFullURL) GetTableAndUniqueID() (string, string) {
	const prefix = "urn:uuid:"

	s := f.String()

	if !strings.HasPrefix(s, prefix) {
		return "", ""
	}

	afterPrefix := s[len(prefix):]
	parts := strings.SplitN(afterPrefix, "-", 2)

	if len(parts) != 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

func (f *BundleEntryFullURL) String() string {
	return string(*f)
}

type BundleEntrySearch struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Mode              *SearchEntryMode `json:"mode,omitempty"`
	Score             *json.Number     `json:"score,omitempty"`
}

type BundleEntryRequest struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Method            HTTPVerb    `json:"method"`
	URL               string      `json:"url"`
	IfNoneMatch       *string     `json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `json:"ifMatch,omitempty"`
	IfNoneExist       *string     `json:"ifNoneExist,omitempty"`
}

type BundleEntryResponse struct {
	ID                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Status            string            `json:"status"`
	Location          *string           `json:"location,omitempty"`
	Etag              *string           `json:"etag,omitempty"`
	LastModified      *string           `json:"lastModified,omitempty"`
	Outcome           *OperationOutcome `json:"outcome,omitempty"`
}

type OtherBundle Bundle

// MarshalJSON marshals the given Bundle as JSON into a byte slice.
func (b Bundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBundle
		ResourceType string `json:"resourceType"`
	}{
		OtherBundle:  OtherBundle(b),
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

// SuccessfulCreate checks whether a transaction Bundle response
// contains an OperationOutcome with the code SUCCESSFUL_CREATE.
func (r *BundleEntryResponse) SuccessfulCreate() bool {
	successfulCreateStatus := "SUCCESSFUL_CREATE"

	if r == nil || r.Outcome == nil || len(r.Outcome.Issue) == 0 {
		return false
	}

	for _, issue := range r.Outcome.Issue {
		if issue.Details == nil {
			continue
		}

		for _, code := range issue.Details.Coding {
			if code.Code != nil && *code.Code == successfulCreateStatus {
				return true
			}
		}
	}

	return false
}
