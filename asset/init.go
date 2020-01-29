package asset

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

var rawAssetMap map[string]*string

func init() {
    rawAssetMap = map[string]*string{
        Bootstrap4:                &bootstrap4,
        BootstrapDatepicker_1_9_0: &bootstrapDatepicker_1_9_0,
        BootstrapSelect_1_13_9:    &bootstrapSelect_1_13_9,
        BootstrapTable_1_15_5:     &bootstrapTable_1_15_5,
        Holder_2_9:                &holder_2_9,
        JqueryMask_1_14_16:        &jqueryMask_1_14_16,
        JqueryValidation_1_19_1:   &jqueryValidation_1_19_1,
        JsCookie_2_2_1:            &jsCookie_2_2_1,
        Moment_2_24_0:             &moment_2_24_0,
        ProgressbarJs_1_0_1:       &progressbarJs_1_0_1,
        VideoJs_7_7_4:             &videoJs_7_7_4,
        WaitMe_31_10_17:           &waitMe_31_10_17,
    }
}

func GetRaw(key string) *string {
    if val, have := rawAssetMap[key]; have {
        return val
    }
    return nil
}
