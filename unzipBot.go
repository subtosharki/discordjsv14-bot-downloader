package main

import (
	"os"
	"os/exec"
)

func UnzipBot() {
	err := exec.Command("unzip", "bot.zip").Run()
	if err != nil {

		panic(err)
	}
	files, err := os.ReadDir("djs-v14-bot-template-main")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			err = exec.Command("mv", "djs-v14-bot-template-main/"+file.Name(), "./").Run()
			if err != nil {
				panic(err)
			}
		} else {
			err = exec.Command("cp", "-r", "djs-v14-bot-template-main/"+file.Name(), "./").Run()
			if err != nil {
				panic(err)
			}
		}
	}
	println("Bot Unzipped")
}
