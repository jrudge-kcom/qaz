package commands

// -- Contains helper functions

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// configTemplate - Returns template byte string for init() function
func configTemplate(project string, region string) []byte {
	return []byte(fmt.Sprintf(`
# AWS Region
region: %s

# Project Name
project: %s

# Global Stack Variables
global:

# Stacks
stacks:

`, region, project))
}

// all - returns true if all items in array the same as the given string
func all(a []string, s string) bool {
	for _, str := range a {
		if s != str {
			return false
		}
	}
	return true
}

// stringIn - returns true if string in array
func stringIn(s string, a []string) bool {
	Log(fmt.Sprintf("Checking If [%s] is in: %s", s, a), level.debug)
	for _, str := range a {
		if str == s {
			return true
		}
	}
	return false
}

// getInput - reads input from stdin - request & default (if no input)
func getInput(request string, def string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [%s]:", request, def)
	t, _ := r.ReadString('\n')

	// using len as t will always have atleast 1 char, "\n"
	if len(t) > 1 {
		return strings.Trim(t, "\n")
	}
	return def
}

// Get - HTTP Get request of given url and returns string
func Get(url string) (string, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)

	if resp == nil {
		return "", errors.New(fmt.Sprintln("Error, GET request timeout @:", url))
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("GET request failed, url: %s - Status:[%s]", url, strconv.Itoa(resp.StatusCode))
	}

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// defaultConfig - sets config based on ENV variable or default config.yml
func defaultConfig() string {
	env := os.Getenv(configENV)
	if env == "" {
		return "config.yml"
	}
	return env
}
