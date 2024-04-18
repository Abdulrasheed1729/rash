package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	//
	// userDetail, err := user.Current()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	//
	// }
	//
	// name := userDetail.Name
	// shellPrompt := fmt.Sprintf("%s@rash> ", name)
	//
	// for {
	// 	fmt.Print(shellPrompt)
	// 	input, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 	}
	// 	if err := RashIt(input); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 	}
	// }
	CC()
}

func CC() {
	cmd := exec.Command("cat", "go.mod")
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}

func RashIt(input string) error {
	command := strings.TrimSuffix(input, "\n")
	history := make([]string, 0)
	args := strings.Split(command, " ")
	history = append(history, command)
	cmd := exec.Command(args[0], args[1:]...)

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("no path specified")
		}
		return os.Chdir(args[1])
	case "":
		return nil
	case "^C":
		//		processState := cmd.Process.
	case "exit":
		os.Exit(0)
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
