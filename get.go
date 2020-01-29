package himma

import (
    "github.com/devplayg/eggcrate"
    "github.com/devplayg/himma/asset"
)

func GetAssetMap(assets ...string) (map[string][]byte, error) {
    assetMap := make(map[string][]byte, 0)
    for i := range assets {

    	raw := asset.GetRaw(assets[i])
		if raw == nil {
			continue
		}

		if err := merge(assetMap, raw); err != nil {
			return nil, err
		}
    }
    return assetMap, nil
}

func merge(assetMap map[string][]byte, data *string) error {
    m, err := eggcrate.Decode(*data)
    if err != nil {
        return err
    }

    for k, v := range m {
        assetMap[k] = v
    }

    return nil
}
