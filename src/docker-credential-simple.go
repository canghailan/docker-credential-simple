package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"log"
	"fmt"
)

type Credential struct {
	ServerURL string
	Username  string
	Secret    string
}

func GetString() string {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panicf("GetString: %v", err)
	}
	return strings.TrimSpace(string(bs))
}

func GetUserHomeDir() string {
	if u, err := user.Current(); err == nil {
		return u.HomeDir
	} else {
		return os.Getenv("HOME")
	}
}

func GetDockerCredsPath() string {
	return filepath.Join(GetUserHomeDir(), ".docker", "creds.json")
}

func LoadCredentials() (cs []Credential) {
	bs, err := ioutil.ReadFile(GetDockerCredsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return []Credential{}
		} else {
			log.Panicf("LoadCredentials: %v", err)
		}
	}

	if err := json.Unmarshal(bs, &cs); err != nil {
		log.Panicf("LoadCredentials: %v", err)
	}
	return cs
}

func SaveCredentials(cs []Credential) {
	bs, err := json.Marshal(cs)
	if err != nil {
		log.Panicf("SaveCredentials: %v", err)
	}

	ioutil.WriteFile(GetDockerCredsPath(), bs, 0444)
}

func UpdateCredentials(cs []Credential, c Credential) []Credential {
	for i, csn := range cs {
		if csn.ServerURL == c.ServerURL {
			cs = append(cs[:i], cs[i+1:]...)
			break
		}
	}
	return append(cs, c)
}

func PrintJson(val interface{}) {
	bs, err := json.MarshalIndent(val, "", "\t")
	if err != nil {
		log.Panicf("PrintJson: %v", err)
	}

	fmt.Fprintln(os.Stdout, string(bs))
}

func store(credential string) {
	var c Credential
	if err := json.Unmarshal([]byte(credential), &c); err != nil {
		fmt.Printf("store: %v", err)
		os.Exit(1)
	}

	cs := LoadCredentials()
	cs = UpdateCredentials(cs, c)
	SaveCredentials(cs)
}

func get(serverURL string) {
	cs := LoadCredentials()
	for _, csn := range cs {
		if csn.ServerURL == serverURL {
			PrintJson(csn)
			return
		}
	}
	fmt.Printf("get: %v", serverURL)
	os.Exit(1)
}

func erase(serverURL string) {
	cs := LoadCredentials()
	for i, csn := range cs {
		if csn.ServerURL == serverURL {
			cs = append(cs[:i], cs[i+1:]...)
			SaveCredentials(cs)
			return
		}
	}
	fmt.Printf("erase: %v", serverURL)
	os.Exit(1)
}

func list() {
	PrintJson(LoadCredentials())
}

func message() string {
	return os.Args[0] + " <store|get|erase|list>"
}

func main() {
	if len(os.Args) <= 1 {
		println(message())
		return
	}
	switch os.Args[1] {
	case "store":
		store(GetString())
	case "get":
		get(GetString())
	case "erase":
		erase(GetString())
	case "list":
		list()
	default:
		println(message())
	}
}
