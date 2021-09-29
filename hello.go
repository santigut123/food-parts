package main

import (
	"fmt"
	"os"
)
func main(){
	cmd := os.Args

	fmt.Printf("Args Passed %s %s %s ", cmd[0],cmd[1],cmd[2])



}
