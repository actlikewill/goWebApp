package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	templateString := `Lemonade Supply Service`
	t, err := template.New("title").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}