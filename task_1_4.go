package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Task 1_4")
	iniData := []string{
		"; Cut down copy of Mozilla application.ini file", "",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0", "EnableExtensionManager=1",
	}
	fmt.Println(ParseIni(iniData))
}

func ParseIni(iniData []string) map[string] map[string]string{
	result :=make(map[string] map[string]string)
	tmp :=make(map[string]string)
	for _, el := range iniData {
		if "" != el && !strings.Contains(el, ";"){

			isItGroup := strings.Contains(el, "[")
			if isItGroup {
				tmp = make(map[string]string)
				result[el] = tmp
			}

			if el!="" && !isItGroup{
				ch := strings.Split(el,"=")
				tmp[ch[0]] = ch[1]
			}
		}
	}
	return  result
}
