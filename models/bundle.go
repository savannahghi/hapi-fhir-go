package models

import (
	"encoding/json"

	"github.com/savannahghi/scalarutils"
)

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
	ID            *string         `json:"id,omitempty"`
	Meta          *FHIRMeta       `json:"meta,omitempty"`
	ImplicitRules *string         `json:"implicitRules,omitempty"`
	Language      *string         `json:"language,omitempty"`
	Identifier    *FHIRIdentifier `json:"identifier,omitempty"`
	Type          BundleType      `json:"type"`
	Timestamp     *string         `json:"timestamp,omitempty"`
	Total         *int            `json:"total,omitempty"`
	Link          []BundleLink    `json:"link,omitempty"`
	Entry         []BundleEntry   `json:"entry,omitempty"`
}

type BundleLink struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Relation          string      `json:"relation"`
	URL               string      `json:"url"`
}

type BundleEntry struct {
	ID                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Link              []BundleLink         `json:"link,omitempty"`
	FullURL           *string              `json:"fullUrl,omitempty"`
	Resource          json.RawMessage      `json:"resource,omitempty"`
	Search            *BundleEntrySearch   `json:"search,omitempty"`
	Request           *BundleEntryRequest  `json:"request,omitempty"`
	Response          *BundleEntryResponse `json:"response,omitempty"`
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
	Method            string      `json:"method"`
	URL               string      `json:"url"`
	IfNoneMatch       *string     `json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `json:"ifMatch,omitempty"`
	IfNoneExist       *string     `json:"ifNoneExist,omitempty"`
}

type BundleEntryResponse struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Status            string           `json:"status"`
	Location          *string          `json:"location,omitempty"`
	Etag              *string          `json:"etag,omitempty"`
	LastModified      *string          `json:"lastModified,omitempty"`
	Outcome           OperationOutcome `json:"outcome,omitempty"`
}

// SuccessfulCreate checks whether a transaction Bundle response
// contains an OperationOutcome with the code SUCCESSFUL_CREATE.
func (r *Bundle) SuccessfulCreate() bool {
	successfulCreateStatus := "SUCCESSFUL_CREATE"

	if r == nil || len(r.Entry) == 0 {
		return false
	}

	for _, entry := range r.Entry {
		for _, issue := range entry.Response.Outcome.Issue {
			if issue.Details == nil {
				continue
			}

			for _, code := range issue.Details.Coding {
				if code.Code != nil && *code.Code == scalarutils.Code(successfulCreateStatus) {
					return true
				}
			}
		}
	}

	return false
}
