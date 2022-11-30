package dpfm_api_output_formatter

import (
	"encoding/json"
	"fmt"
	"os"
)

func ConvertToSDC(data []byte) SDC {
	sdc := SDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}
