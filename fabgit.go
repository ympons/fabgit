package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ympons/go-fab"
)

func main() {
	fab := fab.NewSuperFab()
	args := os.Args[1:]
	if len(args) == 0 {
		args = append(args, "--help")
	}
	if output, err := exec.Command("git", args...).CombinedOutput(); err != nil {
		fmt.Print(fab.Paint(fmt.Sprintf("[fabgit] => %v is not a valid fabgit input. See 'fabgit --help'.", args)))
	} else {
		fmt.Print(fab.Paint(string(output)))
	}
}
