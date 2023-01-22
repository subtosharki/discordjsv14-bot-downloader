package main

import "os/exec"

func DownloadBot() {
	downloadUrl := "https://github.com/subtosharki/djs-v14-bot-template/archive/refs/heads/main.zip"
	err := exec.Command("curl", "-L", downloadUrl, "-o", "bot.zip").Run()
	if err != nil {
		panic(err)
	}
	println("Bot Downloaded")
}
