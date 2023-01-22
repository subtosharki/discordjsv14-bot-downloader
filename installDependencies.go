package main

import "os/exec"

func InstallDependencies() {
	err := exec.Command("npm", "i", "discord.js", "dotenv", "@discordjs/builders").Run()
	if err != nil {
		panic(err)
	}
	println("Dependencies Installed")
}
