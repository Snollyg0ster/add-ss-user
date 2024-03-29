package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"

	config "github.com/Snollyg0ster/add-ss-user/src/config"
	logger "github.com/Snollyg0ster/add-ss-user/src/log"
	password "github.com/Snollyg0ster/add-ss-user/src/password"
)

var configDir string = "test.json"
var configOutDir string = "test.json"

func addUser(login string, pass string) {
	if login == "" {
		log.Fatal("Specify user login")
	}

	port, ok := config.Config.Users[login]

	if pass == "" {
		pass = password.GeneratePassword(20, true, true, true)
	}

	if ok {
		config.Config.PortPassword[port] = pass
		return
	}

	ports := []int{}

	for port := range config.Config.PortPassword {
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

		newPort = port
	}

	newPort++

	config.Config.Users[login] = strconv.Itoa(newPort)
	config.Config.PortPassword[strconv.Itoa(newPort)] = pass
}

func removeUser(login string) {
	if login == "" {
		log.Fatal("Specify user login")
	}

	port, ok := config.Config.Users[login]

	if !ok {
		log.Fatal("no such user")
	}

	delete(config.Config.Users, login)
	delete(config.Config.PortPassword, port)
}

func writeConfigType() {
	res, _ := json.MarshalIndent(config.Config, "", "    ")
	os.WriteFile(configOutDir, res, 0644)
}

func main() {
	add := flag.NewFlagSet("add", flag.ExitOnError)
	remove := flag.NewFlagSet("remove", flag.ExitOnError)

	addLogin := add.String("login", "", "login for new user")
	addPass := add.String("pass", "", "password for new user")
	removeLogin := remove.String("login", "", "use login to remove from config")

	configStr, err := os.ReadFile(configDir)

	if err != nil {
		log.Fatal("no such file")
	}

	json.Unmarshal(configStr, &config.Config)

	switch os.Args[1] {
	case "add":
		add.Parse(os.Args[2:])
		addUser(*addLogin, *addPass)
		logger.LogUser(*addLogin, config.Config)
		writeConfigType()

	case "remove":
		remove.Parse(os.Args[2:])
		removeUser(*removeLogin)
		writeConfigType()

	case "list":
		logger.LogUsers(config.Config)
	}

}
