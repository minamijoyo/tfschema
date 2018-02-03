package tfschema

import "encoding/json"

func (b *Block) FormatJSON() (string, error) {
	bytes, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
