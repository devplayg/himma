package himma

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"strings"
)

const (
	Bootstrap4       = "bootstrap4"
	BootstrapSelect  = "bootstrapSelect"
	BootstrapTable   = "bootstrapTable"
	Holder           = "holder"
	JqueryMask       = "jqueryMask"
	JqueryValidation = "jqueryValidation"
	JsCookie         = "jsCookie"
	Moment           = "moment"
	ProgressbarJs    = "progressbarJs"
	VideoJs          = "videoJs"
	WaitMe           = "waitMe"
	Images           = "images"
)

func GetAssetMap(assets ...string) (AssetMap, error) {
	rawAssetMap := map[string]*string{
		Bootstrap4:       &bootstrap4,
		BootstrapSelect:  &bootstrapSelect_1_13_14,
		BootstrapTable:   &bootstrapTable_1_16_0,
		Holder:           &holder_2_9,
		JqueryMask:       &jqueryMask_1_14_16,
		JqueryValidation: &jqueryValidation_1_19_2,
		JsCookie:         &jsCookie_2_2_1,
		Moment:           &moment_2_27,
		ProgressbarJs:    &progressbar_1_0_1,
		VideoJs:          &videoJs_7_8_0,
		WaitMe:           &waitMe_31_10_17,
		Images:           &images,
	}

	assetMap := make(AssetMap, 0)
	for i := range assets {
		_, have := rawAssetMap[assets[i]]
		if !have {
			continue
		}

		if err := merge(assetMap, rawAssetMap[assets[i]]); err != nil {
			return nil, err
		}
	}
	return assetMap, nil
}

func merge(assetMap AssetMap, data *string) error {
	m, err := decode(*data)
	if err != nil {
		return err
	}

	for k, v := range m {
		assetMap[k] = v
	}

	return nil
}

func decode(encoded string) (map[string][]byte, error) {
	encoded = strings.ReplaceAll(encoded, "\n", "")
	var fileMap map[string][]byte

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(bytes.NewReader(decoded))
	return fileMap, dec.Decode(&fileMap)
}
