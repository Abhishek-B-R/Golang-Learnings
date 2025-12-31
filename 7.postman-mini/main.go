package main

import (
	"fmt"
	"net/http"
	"net/url"
	"net"
	"os"
)

func main() {
	if(len(os.Args) < 2){
		fmt.Println("Please enter the website url to continue")		
		return
	}

	websiteUrl := os.Args[1]
	resp,err := http.Get(websiteUrl)

	if(err != nil){
		fmt.Println("Error while fetching the website",err)
		return
	}
	defer resp.Body.Close()

	u, err := url.Parse(websiteUrl)
	if err != nil{
		fmt.Println("Error while parsing the URL",err)
		return
	}

	ips, err := net.LookupIP(u.Hostname())
	fmt.Println("IP: "+string(ips[0].String()))

	fmt.Println(resp)
}