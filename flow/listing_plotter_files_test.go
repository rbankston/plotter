package flow

import (
	"fmt"
	"os/user"
	"strings"
	"testing"
)

func TestUserHomeDir(t *testing.T) {
	flow := &ListingPlotterFiles{}

	dir := flow.UserHomeDir()
	user, err := user.Current()
	if err != nil {
		t.Errorf("Could not get the current user")
	}

	fmt.Println(user)

	if !strings.ContainsAny(dir, user.Username) {
		t.Errorf("Expected homedir to contain current user but it does not")
	}
}

func TestUserPlotterDir(t *testing.T) {

}

func Testperform(t *testing.T) {

}
