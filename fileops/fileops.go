package fileops

import (
	"fmt"
	"os"
)

const saveFile = "balance.txt"

func writeToFile(value float64) {
	txtBalance := fmt.Sprint(value)
	os.WriteFile(saveFile, []byte(txtBalance), 0644)
}
