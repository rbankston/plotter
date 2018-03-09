package flow

import (
	"os"
	"runtime"
)

type MakeActiveFile struct {
}

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

// func (flow *MakeActiveFile) RemoveSymlink() string {
// 	dir := flow.UserHomeDir()
// 	stringArray := []string{dir, ".kube", "config"}
// 	kubeconfigfile := strings.Join(stringArray, "/")
// 	f, err := os.Remove(kubeconfigfile)
// 	if err != nil {
// 		log.Fatal("Cannot remove symlink", err)
// 	}
// 	defer f.Close()

// }

// func (flow *MakeActiveFile) perform() error {

// }
