package service

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

/*
{
	type: "input",
	name: "name",
	message: "请输入项目名:",
	default: "go_example1"
  }
*/

type Prompt struct {
	Type    string
	Name    string
	Message string
	Default string
}

func NewPrompt(typ, name, msg, defaultVal string) *Prompt {
	ret := &Prompt{
		Type:    typ,
		Name:    name,
		Message: msg,
		Default: defaultVal,
	}
	return ret
}

func (_this *Prompt) Run() {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("Invalid input, please enter a string")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | blue }} ",
		Invalid: "{{ . | yellow }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Enter a float: ",
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	_this.Name = result
}
