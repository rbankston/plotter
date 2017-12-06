package flow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

type ListingPlotterFiles struct {
}

func (flow *ListingPlotterFiles) UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

// TODO: Functioning for windows users.
func (flow *ListingPlotterFiles) UserPlotterDir() string {
	dir := flow.UserHomeDir()
	stringArray := []string{dir, ".kube", "plotter"}
	plotterDir := strings.Join(stringArray, "/")
	return plotterDir
}

func (flow *ListingPlotterFiles) perform() error {
	dir := flow.UserPlotterDir()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		// TODO: Not use log.fatal but implement error handling with RunE in cobra command
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
