package cli

import (
	"fmt"

	"go.uber.org/zap"
)

func Initialize() string {
	fmt.Println(BlueBold, "Hello", Reset, Cyan, "\n------------------", Reset)
	fmt.Print(PurpleBold, "Enter CMID: ", Reset)
	var cmid string
	_, err := fmt.Scan(&cmid)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return cmid
}
