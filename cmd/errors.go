package cmd

import (
	"fmt"
	"os"
)

func errorMsg(err string) {
	fmt.Println(err)
	os.Exit(1)
}
