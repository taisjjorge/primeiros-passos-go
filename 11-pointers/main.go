package main

import "fmt"

type game struct{
	name string
	year uint
}


func (g *game) changeName(newName string) { //* -> quem chamar changeName vai passar exatamente obj que está alocado em memória naquela chamada
	g.name = newName
	fmt.Println("Setei novo valor", g) //passado como uma copia
	
}

func setSkill(newSkill *string) {
	fmt.Println("skill recebida", newSkill)
	fmt.Println(newSkill)
	*newSkill = "Super força" //pasando endereço de memoria 
}

func main() {
	
	mario := game {
		name: "Super Mario",
		year: 1990,
	}

	mario.changeName("Super Mario World")

	fmt.Println("Pós", mario)

	skill := "Super pulo"
	fmt.Println(&skill)
	
	setSkill(&skill) // recebendo ponteiro com endereçamento de memoria '&skill'

	fmt.Println("Nova skill", skill) //imprime super pulo, alterando uma copia

}