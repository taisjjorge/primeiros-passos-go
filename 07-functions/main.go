package main

import "fmt"

func main() {

	
	sendMessages("Teste 1", "Teste 2")
	fmt.Println()

	msg, ok := checkAge("Junior", 30)
	fmt.Println(msg, ok)
}

//func sendMessages(msgs []string){ slice de string

func sendMessages(msgs ...string) bool {
	for _, message := range msgs {
		fmt.Println("Enviando mensagem ", message)
	}
	return true
}

//func privada e publica -> começou com letra minúscula é privada(pacote), com maiúscula é pública

func checkAge(name string, age uint) (string, bool) {
	if age >= 18 {
		return fmt.Sprint(name, "permitido"), true //retornar o formatado - Sprint
	}
	return fmt.Sprint(name, "não permitido"), false
}

