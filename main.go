package main

import (
	"fmt"
	"os/exec"
)

type aliases struct {
	command, alias string
}

var helper = []aliases{
	{
		command: "fetch",
		alias:   "",
	},
	{
		command: "pull",
		alias:   "",
	},
	{
		command: "status -sb",
		alias:   "",
	},
	{
		command: "commit -m",
		alias:   "",
	},
	{
		command: "push",
		alias:   "",
	},
	{
		command: "checkout",
		alias:   "",
	},
	{
		command: "checkout -b",
		alias:   "",
	},
	{
		command: "submodule foreach git pull origin master",
		alias:   "",
	},
	{
		command: "pull origin master:master",
		alias:   "",
	},
}

func main() {
	cls()
	showSetAliases()
	fmt.Println("lets setup some aliases...")

	for i, alias := range helper {
		helper[i].alias = getAliasKey(alias.command)
	}

	for _, alias := range helper {
		setAlias(alias)
	}
	showSetAliases()
}

func showSetAliases() {
	fmt.Println("current aliases:")
	out, err := exec.Command("git", "config", "--global", "-l").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out[:]))
}

func setAlias(aliasStruct aliases) {
	err := exec.Command("git", "config", "--global", "alias."+aliasStruct.alias, "'"+aliasStruct.command+"'").Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getAliasKey(command string) string {
	fmt.Println("get " + command + " alias will be")
	alias := ""
	fmt.Scanln(&alias)
	cls()
	return alias
}

func cls() {
	fmt.Print("\033[H\033[2J")
}
