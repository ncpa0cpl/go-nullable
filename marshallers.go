package nullable

import "encoding/json"

func (v *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	v.isSet = true

	err := json.Unmarshal(data, &v.content)

	return err
}

func (v Nullable[T]) MarshalJSON() ([]byte, error) {
	if v.isSet {
		return json.Marshal(v.content)
	}
	return []byte("null"), nil
}
