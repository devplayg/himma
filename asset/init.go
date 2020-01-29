package asset

import (
    "github.com/devplayg/himma"
)

var rawAssetMap map[string]*string

func init() {
    rawAssetMap = map[string]*string{
        himma.Bootstrap4:                &bootstrap4,
        himma.BootstrapDatepicker_1_9_0: &bootstrapDatepicker_1_9_0,
        himma.BootstrapSelect_1_13_9:    &bootstrapSelect_1_13_9,
        himma.BootstrapTable_1_15_5:     &bootstrapTable_1_15_5,
        himma.Holder_2_9:                &holder_2_9,
        himma.JqueryMask_1_14_16:        &jqueryMask_1_14_16,
        himma.JqueryValidation_1_19_1:   &jqueryValidation_1_19_1,
        himma.JsCookie_2_2_1:            &jsCookie_2_2_1,
        himma.Moment_2_24_0:             &moment_2_24_0,
        himma.ProgressbarJs_1_0_1:       &progressbarJs_1_0_1,
        himma.VideoJs_7_7_4:             &videoJs_7_7_4,
        himma.WaitMe_31_10_17:           &waitMe_31_10_17,
    }
}

func GetRaw(key string) *string {
    if val, have := rawAssetMap[key]; have {
        return val
    }
    return nil
}
