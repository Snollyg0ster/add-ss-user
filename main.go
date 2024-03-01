package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"

	password "github.com/Snollyg0ster/add-ss-user/src"
)

type Config struct {
	Server        string            `json:"server"`
	Local_address string            `json:"local_address"`
	Local_port    int32             `json:"local_port"`
	Port_password map[string]string `json:"port_password"`
	Timeout       int64             `json:"timeout"`
	Method        string            `json:"method"`
	Fast_open     bool              `json:"fast_open"`
	Users         map[string]string `json:"users"`
}

var configDir string = "test.json"
var configOutDir string = "outest.json"

var config Config

func addUser(login string, pass string) {
	if login == "" {
		panic("Specify user login")
	}

	port, ok := config.Users[login]

	if pass == "" {
		pass = password.GeneratePassword(12, true, true, true)
	}

	if ok {
		config.Port_password[port] = pass
		return
	}

	ports := []int{}

	for port := range config.Port_password {
		num, err := strconv.Atoi(port)

		if err != nil {
			continue
		}

		ports = append(ports, num)
	}

	sort.Ints(ports)

	var newPort int

	for _, port := range ports {
		if newPort != 0 && port-newPort > 1 {
			break
		}

		newPort = port + 1
		// fmt.Println("newPort", newPort)
	}

	config.Users[login] = strconv.Itoa(newPort)
	config.Port_password[strconv.Itoa(newPort)] = pass
}

func main() {
	add := flag.NewFlagSet("add", flag.ExitOnError)
	remove := flag.NewFlagSet("remove", flag.ExitOnError)
	list := flag.NewFlagSet("list", flag.ExitOnError)

	addLogin := add.String("login", "", "login for new user")
	addPass := add.String("pass", "", "password for new user")
	// removeLogin := remove.String("login", "", "use login to remove from config")

	configStr, err := os.ReadFile(configDir)

	if err != nil {
		panic("no such file")
	}

	json.Unmarshal(configStr, &config)

	switch os.Args[1] {
	case "add":
		add.Parse(os.Args[2:])

		addUser(*addLogin, *addPass)

		res, _ := json.MarshalIndent(config, "", "    ")
		fmt.Println(string(res))
		os.WriteFile(configOutDir, res, 0644)

	case "remove":
		remove.Parse(os.Args[2:])

		if *addLogin == "" {
			panic("Specify user login")
		}

	case "list":
		list.Parse(os.Args[2:])
	}

}
