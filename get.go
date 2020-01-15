package himma

import (
	"github.com/devplayg/eggcrate"
)

const Bootstrap4 = "bs4"

func GetAssetMap(assetType string) (map[string][]byte, error) {
	if assetType == Bootstrap4 {
		return eggcrate.Decode(bs4)
	}
	return nil, nil
}
