package file

import(
	"os"
	"fmt"
)
var FFile *os.File

func FileCreation() {
	var err error
	FFile, err = os.Create("errors.txt")
	if err != nil {
		fmt.Println("FileNotCreated")
	}
}