package main

import (
	"fmt"
)

func main() {
	var nome string = "Micael" // a variavel 'var' não é necessaria para fazer uma função // tambem não é preciso por a forma que a minha variavel vai receber por exemplo
	var versao float32 = 1.1   // não é necessario por 'string', 'int', "float", se não colocar o float o go vai sempre jogar ara o maior float que tem o 64
	fmt.Println("ola, sr.", nome)
	fmt.Println("Este programa esta na versão", versao)
	fmt.Println("------------------------------")

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando) // o %d indica algo inteiro, e o comando seria uma lista vazia que vai receber o input do usuario e converter para o sistemas
	// O '&' serve para ele poder escolher entre um comando que a gente listou, por que senão tivesse ele iria puar direto para o 0
	// '&' = endereço

	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando esolhido foi", comando)
}

// função para capturar a escrita do usuario o input ==> 'fmt.Scanf'
// colocar ':=' ele identifica uma variavel
