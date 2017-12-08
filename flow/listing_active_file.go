package flow

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type ListingActiveFile struct {
}

func (flow *ListingActiveFile) UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func (flow *ListingActiveFile) perform() error {
	dir := flow.UserHomeDir()

	stringArray := []string{dir, ".kube", "config"}
	kubeconfigfile := strings.Join(stringArray, "/")

	activefile, err := os.Readlink(kubeconfigfile)

	if err != nil {
		// TODO: Not use log.fatal but implement error handling with RunE in cobra command
		log.Fatal(err)
	}
	fmt.Println(activefile)

	return nil
}
