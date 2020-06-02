package himma

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
)

type AssetMap map[string][]byte

func (m AssetMap) Get(key string) []byte {
	return m[key]
}

func (m AssetMap) AddRaw(key, value string) error {
	b, err := zip([]byte(value))
	if err != nil {
		return err
	}
	m[key] = b
	return nil
}

func (m AssetMap) AddZipped(key string, value []byte) error {
	m[key] = value
	return nil
}

func (m AssetMap) AddZippedAndBase64Encoded(key, value string) error {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	m[key] = decoded
	return nil
}

func zip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	if _, err := zw.Write(data); err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
