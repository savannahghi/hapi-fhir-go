package hapifhirgo

import (
	"encoding/json"
	"fmt"
)

// structToMap converts an object (struct) to a map.
//
// WARNING: int inputs are converted to floats in the output map. This is an
// unintended consequence of converting through JSON.
func structToMap(item interface{}) (map[string]interface{}, error) {
	bs, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal to JSON: %w", err)
	}

	res := map[string]interface{}{}

	err = json.Unmarshal(bs, &res)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal from JSON to map: %w", err)
	}

	return res, nil
}

// remarshal re-encodes src and decodes it into dst. It is used to convert a
// generically-assembled bundle (map[string]interface{}) into the caller's
// version-specific bundle type.
func remarshal(src interface{}, dst interface{}) error {
	bs, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("unable to marshal merged bundle: %w", err)
	}

	if err := json.Unmarshal(bs, dst); err != nil {
		return fmt.Errorf("unable to unmarshal merged bundle: %w", err)
	}

	return nil
}
