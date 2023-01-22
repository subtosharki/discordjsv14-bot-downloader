package main

import "os/exec"

func CleanUp() {
	err := exec.Command("rm", "-rf", "djs-v14-bot-template-main").Run()
	if err != nil {
		panic(err)
	}
	err = exec.Command("rm", "bot.zip").Run()
	if err != nil {
		panic(err)
	}
	println("Cleaned Up")
}
