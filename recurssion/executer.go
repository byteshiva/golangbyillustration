package main

// @TODO - the below packag import is only for testing purpose it needs refractoring
// 1. Need to adher to proper naming convention

// Steps to build and install package to be accessible from current executer emthod
// 1. create a directory mkdir $GOPATH/golang-utils/
// 2. copy the folder that needs to be accessible from outside
// 3. make sure all the method names are exported by captilizing ex. fibo to Fibo
// 4. after copying into $GOPATH, do $go build and $go install, it makes it accessible to other packages

import (
	"fmt"
	"golang-utils/recurssion/iterative"
	"golang-utils/recurssion/recurssion"
)

func main() {
	fmt.Println(iterative.FiboR(10))
	fmt.Println(iterative.FiboI(10))
	fmt.Println(iterative.FiboTail(10))
	fmt.Println(recurssion.Fact(17))
}
