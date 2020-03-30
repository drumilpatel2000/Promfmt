package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

func contents(content *yaml.Node){
	fmt.Println(content.Value)
	if content.HeadComment != "" {
		fmt.Print("Head ", content.HeadComment)
	}
	if content.LineComment != "" {
		fmt.Print("Line ", content.LineComment)
	}
	if content.FootComment != "" {
		fmt.Print("Foot ", content.FootComment)
	}
}

func main(){
	
	var b yaml.Node
	
	yamlFile, err := ioutil.ReadFile("comments.yaml")
	if err != nil {
		fmt.Print(err)
	}
	
	err = yaml.Unmarshal(yamlFile, &b)

	for _, content := range(b.Content){
		contents(content)
		for _, content := range(content.Content) {
			contents(content)
			for _, content := range(content.Content) {
				contents(content)
				for _, content := range(content.Content) {
					contents(content)
					for _, content := range(content.Content) {
						contents(content)
						for _, content := range(content.Content) {
							if content.Value == "expr" {
								fmt.Println("##################################################################")
								fmt.Println("My Expressions")
							} else {
								contents(content)
								for _, content := range(content.Content) {
									contents(content)
								}	
							}
						}
					}
				}
			}
		}	
	}

}