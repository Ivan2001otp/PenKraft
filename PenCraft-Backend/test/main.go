package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("python","../config/mongod_script.py")

	// run the command and capture the output 
	output,err:= cmd.CombinedOutput()
	if err!= nil {
		log.Fatalf("Error running the script: %v\n",err)
	}

	// output of python script 
	fmt.Println(string(output))
}