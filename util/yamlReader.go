package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Snippet struct {
	Id       int    `yaml:"id"`
	Language string `yaml:"language"`
	Code     string `yaml:"code"`
}

func ParseYaml() []Snippet {

	var snippetSlice []Snippet

	data, err := os.ReadFile("seedData.yaml")
	if err != nil {
		fmt.Println("Read File:", err)
	}

	err = yaml.Unmarshal(data, &snippetSlice)
	if err != nil {
		fmt.Printf("Unmashal data: %d\n", err)
	}

	return snippetSlice
}
