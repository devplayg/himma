package himma

import (
    "github.com/devplayg/eggcrate"
	"github.com/devplayg/himma/asset"
)

const (
    Bootstrap4                = "bootstrap4"
    BootstrapDatepicker_1_9_0 = "bootstrapDatepicker_1_9_0"
    BootstrapSelect_1_13_9    = "bootstrapSelect_1_13_9"
    BootstrapTable_1_15_5     = "bootstrapTable_1_15_5"
    Holder_2_9                = "holder_2_9"
    JqueryMask_1_14_16        = "jqueryMask_1_14_16"
    JqueryValidation_1_19_1   = "jqueryValidation_1_19_1"
    JsCookie_2_2_1            = "jsCookie_2_2_1"
    Moment_2_24_0             = "moment_2_24_0"
    ProgressbarJs_1_0_1       = "progressbarJs_1_0_1"
    VideoJs_7_7_4             = "videoJs_7_7_4"
    WaitMe_31_10_17           = "waitMe_31_10_17"
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
