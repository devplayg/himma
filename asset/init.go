package asset

import "fmt"

var rawAssetMap map[string]*string

func init() {
    fmt.Println("HIMMA STARTED")
	rawAssetMap = map[string]*string{
        "bootstrap4":                &bootstrap4,
        "bootstrapDatepicker_1_9_0": &bootstrapDatepicker_1_9_0,
        "bootstrapSelect_1_13_9":    &bootstrapSelect_1_13_9,
        "bootstrapTable_1_15_5":     &bootstrapTable_1_15_5,
        "holder_2_9":                &holder_2_9,
        "jqueryMask_1_14_16":        &jqueryMask_1_14_16,
        "jqueryValidation_1_19_1":   &jqueryValidation_1_19_1,
        "jsCookie_2_2_1":            &jsCookie_2_2_1,
        "moment_2_24_0":             &moment_2_24_0,
        "progressbarJs_1_0_1":       &progressbarJs_1_0_1,
        "videoJs_7_7_4":             &videoJs_7_7_4,
        "waitMe_31_10_17":           &waitMe_31_10_17,
    }
}

func GetRaw(key string) *string {
	if val, have := rawAssetMap[key]; have {
		return val
	}
	return nil
}