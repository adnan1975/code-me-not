package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)


	out, err := exec.Command("echo","wtf" ).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("output is %s\n", out)

    //cmd := exec.Command("/bin/sh", mongoToCsvSH)

}