package flow

import (
	"os"
	"runtime"
	"strings"
)

type MakeNewFile struct {
	Profile string
}

// UserHomeDir detects the User's Home Directory
func (flow *MakeNewFile) UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

// RemoveSymlink removes the symlink in place for the use command
func (flow *MakeNewFile) CheckFile() {
	dir := flow.UserHomeDir()
	stringArray := []string{dir, ".kube", "config"}
	oldKubeConfigFile := strings.Join(stringArray, "/")

	if _, err := os.Lstat(oldKubeConfigFile); err == nil {
		os.Remove(oldKubeConfigFile)
	}
}

// CreateSymlink Creates new symlink given the name of the plotter file to use
func (flow *MakeNewFile) CreateFile(file string) {
	dir := flow.UserHomeDir()
	plotterFile := strings.ToLower(file)
	newArray := []string{dir, ".kube", "plotter", plotterFile}
	configArray := []string{dir, ".kube", "config"}
	newKubeConfigFile := strings.Join(newArray, "/")
	symlinkDestination := strings.Join(configArray, "/")
	os.Symlink(newKubeConfigFile, symlinkDestination)

}
func (flow *MakeActiveFile) perform() error {
	flow.RemoveSymlink()
	flow.CreateSymlink(flow.Profile)
	return nil
}
