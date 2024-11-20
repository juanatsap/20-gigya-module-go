package helpers

import (
	"encoding/json"
	"fmt"
)

type DynamicStringArray []string

func (d *DynamicStringArray) UnmarshalJSON(data []byte) error {
	// Si los datos son un string
	var singleValue string
	if err := json.Unmarshal(data, &singleValue); err == nil {
		*d = []string{singleValue}
		return nil
	}

	// Si los datos son un array de strings
	var arrayValue []string
	if err := json.Unmarshal(data, &arrayValue); err == nil {
		*d = arrayValue
		return nil
	}

	// Retornar un error si no es ninguno de los dos
	return fmt.Errorf("DynamicStringArray: cannot unmarshal %s", string(data))
}

func (a DynamicStringArray) RemoveNulls() []string {
	var result []string
	for _, value := range a {
		if value != "null" {
			result = append(result, value)
		}
	}
	return result
}
