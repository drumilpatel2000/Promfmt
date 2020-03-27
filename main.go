package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

func main(){
	
	var b yaml.Node
	
	yamlFile, err := ioutil.ReadFile("comments.yaml")
	if err != nil {
		fmt.Print(err)
	}
	
	err = yaml.Unmarshal(yamlFile, &b)

	d, err := yaml.Marshal(&b)

	fmt.Println(string(d))

}