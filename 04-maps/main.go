package main

import "fmt"

func main() {
	languageLevel := make(map[string]int) //map ordenado
	//tem uma chave string e um valor int

	languageLevel["Java"] = 6
	languageLevel["PHP"] = 5
	languageLevel["JavaScript"] = 10
	languageLevel["GoLang"] = 6

	fmt.Println(languageLevel)

	languageLevel["JavaScript"] = 8

	fmt.Println(languageLevel)

	delete(languageLevel, "GoLang")

	fmt.Println(languageLevel, len(languageLevel))


	// #######################

	consoleAndGames := make(map[string][]string)

	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"], "Donkey Kong")
	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"], "Super Mario World")
	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"], "Star Fox")


	fmt.Println(consoleAndGames)
	fmt.Println(consoleAndGames["Super Nintendo"])
	fmt.Println(consoleAndGames["Super Nintendo"][1])

	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"][:1], consoleAndGames["Super Nintendo"][2:3]...)

	//deletando super mario
	fmt.Println("Meus jogos:", consoleAndGames["Super Nintendo"])
}
