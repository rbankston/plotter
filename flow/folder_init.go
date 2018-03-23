package flow

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type FolderInit struct {
}

func (flow *FolderInit) UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func (flow *FolderInit) perform() error {
	dir := flow.UserHomeDir()

	stringArray := []string{dir, ".kube", "plotter"}
	plotterconfigdirectory := strings.Join(stringArray, "/")

	activefile, err := os.Readlink(plotterconfigdirectory)

	if err != nil {
		// TODO: Not use log.fatal but implement error handling with RunE in cobra command
		log.Fatal(err)
	}
	fmt.Println(activefile)

	return nil
}
