
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// GroupMeasure is documented here http://hl7.org/fhir/ValueSet/group-measure
type GroupMeasure int

const (
	GroupMeasureMean GroupMeasure = iota
	GroupMeasureMedian
	GroupMeasureMeanOfMean
	GroupMeasureMeanOfMedian
	GroupMeasureMedianOfMean
	GroupMeasureMedianOfMedian
)

func (code GroupMeasure) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GroupMeasure) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal GroupMeasure code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "mean":
		*code = GroupMeasureMean
	case "median":
		*code = GroupMeasureMedian
	case "mean-of-mean":
		*code = GroupMeasureMeanOfMean
	case "mean-of-median":
		*code = GroupMeasureMeanOfMedian
	case "median-of-mean":
		*code = GroupMeasureMedianOfMean
	case "median-of-median":
		*code = GroupMeasureMedianOfMedian
	default:
		return fmt.Errorf("unknown GroupMeasure code `%s`", s)
	}
	return nil
}
func (code GroupMeasure) String() string {
	return code.Code()
}
func (code GroupMeasure) Code() string {
	switch code {
	case GroupMeasureMean:
		return "mean"
	case GroupMeasureMedian:
		return "median"
	case GroupMeasureMeanOfMean:
		return "mean-of-mean"
	case GroupMeasureMeanOfMedian:
		return "mean-of-median"
	case GroupMeasureMedianOfMean:
		return "median-of-mean"
	case GroupMeasureMedianOfMedian:
		return "median-of-median"
	}
	return "<unknown>"
}
func (code GroupMeasure) Display() string {
	switch code {
	case GroupMeasureMean:
		return "Mean"
	case GroupMeasureMedian:
		return "Median"
	case GroupMeasureMeanOfMean:
		return "Mean of Study Means"
	case GroupMeasureMeanOfMedian:
		return "Mean of Study Medins"
	case GroupMeasureMedianOfMean:
		return "Median of Study Means"
	case GroupMeasureMedianOfMedian:
		return "Median of Study Medians"
	}
	return "<unknown>"
}
func (code GroupMeasure) Definition() string {
	switch code {
	case GroupMeasureMean:
		return "Aggregated using Mean of participant values."
	case GroupMeasureMedian:
		return "Aggregated using Median of participant values."
	case GroupMeasureMeanOfMean:
		return "Aggregated using Mean of study mean values."
	case GroupMeasureMeanOfMedian:
		return "Aggregated using Mean of study median values."
	case GroupMeasureMedianOfMean:
		return "Aggregated using Median of study mean values."
	case GroupMeasureMedianOfMedian:
		return "Aggregated using Median of study median values."
	}
	return "<unknown>"
}
