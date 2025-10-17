package main

import (
	"log"

	"github.com/manifoldco/promptui"
)

func confirmPrompt(label string) {
	prompt := promptui.Prompt{
		IsConfirm: true,
		Label:     label,
	}

	_, err := prompt.Run()
	if err != nil {
		log.Fatalf("Program exiting due to prompt response.")
	}

}
