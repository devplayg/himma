package asset

var AssetMap map[string]string

func init() {
	AssetMap = make(map[string]string)
	AssetMap["BootstrapDatepicker_1_9_0"] = BootstrapDatepicker_1_9_0
}
