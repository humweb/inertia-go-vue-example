package main

import (
	"errors"
	"fmt"
	"github.com/humweb/inertia-go-example/server"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

func main() {
	port, url, err := processENV()

	if err != nil {
		panic("Trouble setting up server: " + err.Error())
	}

	a, err := server.NewServer(port, url)
	if err != nil {
		panic("Trouble setting up server: " + err.Error())
	}

	err = a.Serve()
	if err != nil {
		panic("Trouble serving from server: " + err.Error())
	}

	fmt.Println("Server started:", url+":"+strconv.Itoa(port))
}

func processENV() (int, string, error) {
	requiredENV := []string{
		"APP_PORT",
		"APP_URL",
	}

	for _, env := range requiredENV {
		check := os.Getenv(env)
		if check == "" {
			return 0, "", errors.New("missing env: " + env)
		}
	}

	webPort := os.Getenv("APP_PORT")
	portInt, err := strconv.Atoi(webPort)
	if err != nil {
		panic("Invalid APP_PORT: " + err.Error())
	}

	return portInt, os.Getenv("APP_URL"), nil
}
