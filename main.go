package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func executeInput(input string) error{

	// remove the \n from the string
	input = strings.TrimSuffix(input,"\n")
	args := strings.Split(input, " ")
	switch args[0]{
	case "cd":
		if len(args) < 2{
			return errors.New("path required")
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}
	// prepare the command
	cmd := exec.Command(args[0],args[1:]...)
	fmt.Println("command---",cmd)
	
	// assign the correct input and output devices 
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// run the command 
	return cmd.Run()
}
func getOutPut(cmd string)string {
	output,err:=exec.Command(cmd).Output()
	if err!=nil {
		 fmt.Fprintln(os.Stderr, err)
	}

	return strings.TrimSpace(string(output))
}

func getHeader() string{
	// return this string [majjikishore@fedora:~/t/Shell_In_Go]─[08:15:00 AM]
	userName := getOutPut("whoami")
	hostName:= getOutPut("hostname")
	currDir := getOutPut("pwd")
	currentTime := time.Now().Format("03:04:05 PM")

	currDir = strings.Replace(string(currDir),"/home/majjikishore","~",1)
	header := fmt.Sprintf("%s@%s:%s─%s ->", userName, hostName, currDir, currentTime)

	return header
}

func main(){
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		

		fmt.Print(getHeader())
		input ,err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stderr,err)
		}
		if err = executeInput(input); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
	}
	
}


