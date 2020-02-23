package himma

import "encoding/base64"

type AssetMap map[string][]byte

func (m AssetMap) Get(key string) []byte {
	return m[key]
}

func (m AssetMap) Add(key, value string) error {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	m[key] = decoded
	return nil
}
