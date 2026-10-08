package fhir500

import "errors"

// GetServiceRequestIDFromTask is used to extract the referral (service request) ID from a task.
func (t *Task) GetServiceRequestIDFromTask() (string, error) {
	if t == nil {
		return "", errors.New("task is nil")
	}

	var referralID string

	for _, serviceRequest := range t.BasedOn {
		if serviceRequest.Type != nil && *serviceRequest.Type == ReferralServiceRequestType.String() {
			referralID = "ServiceRequest/" + *serviceRequest.ID

			break
		}
	}

	return referralID, nil
}

// TaskRelayPayload is used to return single instances of Task.
type TaskRelayPayload struct {
	Resource *Task `json:"resource,omitempty"`
}
