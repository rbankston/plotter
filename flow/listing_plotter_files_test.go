package flow

import (
	"fmt"
	"log"
	"os/user"
	"strings"
	"testing"
)

func TestUserHomeDir(t *testing.T) {
	dir := (flow).UserHomeDir
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	if strings.ContainsAny(dir, user.Username) {
		t.Errorf("Expected homedir to contain current user but it does not")
	}
}

func TestUserPlotterDir(t *testing.T) {

}

func Testperform(t *testing.T) {

}
