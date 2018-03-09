package flow

import (
	"os"
	"runtime"
	"strings"
)

type MakeActiveFile struct {
}

// UserHomeDir detects the User's Home Directory
func (flow *MakeActiveFile) UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

// TODO make the back up function be part of this command. We'll start with just what is needed.

// RemoveSymlink removes the symlink in place for the use command
func (flow *MakeActiveFile) RemoveSymlink() {
	dir := flow.UserHomeDir()
	stringArray := []string{dir, ".kube", "config"}
	oldKubeConfigFile := strings.Join(stringArray, "/")

	if _, err := os.Lstat(oldKubeConfigFile); err == nil {
		os.Remove(oldKubeConfigFile)
	}
}

// CreateSymlink Creates new symlink given the name of the plotter file to use
func (flow *MakeActiveFile) CreateSymlink(file string) {
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

	return nil
}
