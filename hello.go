package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Micael" // a variavel 'var' não é necessaria para fazer uma função
	var idade int = 22         // tambem não é preciso por a forma que a minha variavel vai receber por exemplo
	var versao float32 = 1.1   // não é necessario por 'string', 'int', "float", se não colocar o float o go vai sempre jogar ara o maior float que tem o 64
	fmt.Println("ola, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa esta na versão", versao)

	fmt.Println("o tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("o tipo da variavel iade é", reflect.TypeOf(idade))
}
// colocar ':=' ele identifica uma variavel 
