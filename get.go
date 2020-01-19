package himma

import (
	"github.com/devplayg/eggcrate"
)

const (
	Bootstrap4 = "bs4"
	Plugins    = "plugins"
)

func GetAssetMap(assets ...string) (map[string][]byte, error) {
	assetMap := make(map[string][]byte, 0)
	for i := range assets {
		if assets[i] == Bootstrap4 {
			if err := merge(bootstrap4, assetMap); err != nil {
				return nil, err
			}
			continue
		}

		if assets[i] == Plugins {
			if err := merge(plugins, assetMap); err != nil {
				return nil, err
			}
			continue
		}

	}
	return assetMap, nil
}

func merge(data string, assetMap map[string][]byte) error {
	m, err := eggcrate.Decode(data)
	if err != nil {
		return err
	}

	for k, v := range m {
		assetMap[k] = v
	}

	return nil
}
