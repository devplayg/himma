package himma

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"strings"
)

const (
	Bootstrap4                = "bootstrap4"
	BootstrapDatepicker_1_9_0 = "bootstrapDatepicker_1_9_0"
	BootstrapSelect_1_13_9    = "bootstrapSelect_1_13_9"
	BootstrapTable_1_16_0     = "BootstrapTable_1_16_0"
	Holder_2_9                = "holder_2_9"
	JqueryMask_1_14_16        = "jqueryMask_1_14_16"
	JqueryValidation_1_19_1   = "jqueryValidation_1_19_1"
	JsCookie_2_2_1            = "jsCookie_2_2_1"
	Moment_2_24_0             = "moment_2_24_0"
	ProgressbarJs_1_0_1       = "progressbarJs_1_0_1"
	VideoJs_7_7_4             = "videoJs_7_7_4"
	WaitMe_31_10_17           = "waitMe_31_10_17"
)

func GetAssetMap(assets ...string) (AssetMap, error) {
	rawAssetMap := map[string]*string{
		Bootstrap4:                &bootstrap4,
		BootstrapDatepicker_1_9_0: &bootstrapDatepicker_1_9_0,
		BootstrapSelect_1_13_9:    &bootstrapSelect_1_13_9,
		BootstrapTable_1_16_0:     &bootstrapTable_1_16_0,
		Holder_2_9:                &holder_2_9,
		JqueryMask_1_14_16:        &jqueryMask_1_14_16,
		JqueryValidation_1_19_1:   &jqueryValidation_1_19_1,
		JsCookie_2_2_1:            &jsCookie_2_2_1,
		Moment_2_24_0:             &moment_2_24_0,
		ProgressbarJs_1_0_1:       &progressbarJs_1_0_1,
		VideoJs_7_7_4:             &videoJs_7_7_4,
		WaitMe_31_10_17:           &waitMe_31_10_17,
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
