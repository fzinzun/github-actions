package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Name string
	Version string
	sha string
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println("arg 0:", argsWithProg)
	fmt.Println("arg 1:", argsWithoutProg)

	t := Config {
		Name: "app1",
		Version: "1.0.0",
		sha: "",
	}

	d, err := yaml.Marshal(&t)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))



	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)  // for example /home/user



	s := Config{}
	data, err := ioutil.ReadFile(filepath.Join(path,argsWithoutProg[0]))

	err = yaml.Unmarshal([]byte(data), &s)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", s)

}