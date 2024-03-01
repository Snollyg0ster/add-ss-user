package logger

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/Snollyg0ster/add-ss-user/src/config"
)

var columns []string = []string{"login", "port", "link"}

func getUserLink(login string, config config.ConfigType) string {
	port := config.Users[login]
	pass := config.PortPassword[port]
	encodedPass := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Method, pass)))

	return fmt.Sprintf("ss://%s@%s:%s#%v", encodedPass, config.Server, port, login)
}

func getUserData(login string, config config.ConfigType) []string {
	port := config.Users[login]
	link := getUserLink(login, config)

	return []string{login, port, link}
}

func drawTable(table [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, '	', 0)

	fmt.Println(strings.Join(columns, "\t"))

	for _, row := range table {
		fmt.Println(strings.Join(row, "\t"))
	}

	w.Flush()
}

func LogUser(login string, config config.ConfigType) {
	drawTable([][]string{getUserData(login, config)})
}

func LogUsers(config config.ConfigType) {
	usersData := make([][]string, 0)

	for login := range config.Users {
		usersData = append(usersData, getUserData(login, config))
	}

	drawTable(usersData)
}
