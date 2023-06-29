package service

import (
	"errors"
	"log"

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
	Type    string `json:"type"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Default string `json:"default"`
	Value   string `json:"value"`
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
	prompt := _this.RenderPromptByType("")
	result, err := prompt.Run()
	if err != nil {
		log.Printf("Run Prompt failed err:%v", err)
		return
	}
	_this.Value = result
}

// RenderPromptByType 主要类型有：prompt、select、confirm、password
func (_this *Prompt) RenderPromptByType(typ string) promptui.Prompt {
	var ret promptui.Prompt
	switch typ {
	default:
		validate := func(input string) error {
			if len(input) < 1 {
				return errors.New("invalid input, please enter a value")
			}
			return nil
		}
		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | blue }} ",
			Invalid: "{{ . | yellow }} ",
			Success: "{{ . | bold }} ",
		}
		ret = promptui.Prompt{
			Label:     _this.Message,
			Templates: templates,
			Validate:  validate,
		}
	}
	return ret
}
