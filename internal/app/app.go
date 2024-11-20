package app

import (
	"fmt"
	"time"
	"password-generator/internal/password"
	"password-generator/internal/file"
)

func Run() {
	var opcao int
	var tipo string
	var comprimento int

	fmt.Println("Qual o comprimento da senha? ")
	fmt.Scan(&comprimento)
	if comprimento <= 0 {
		fmt.Println("O comprimento da senha deve ser maior que zero.")
		return
	}

	for opcao != 5 {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - Letras")
		fmt.Println("2 - Símbolos Especiais")
		fmt.Println("3 - Números")
		fmt.Println("4 - Letras, Números e Símbolos Especiais")
		fmt.Println("5 - Sair do Programa!")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			tipo = "letras"
		case 2:
			tipo = "símbolos especiais"
		case 3:
			tipo = "números"
		case 4:
			tipo = "letras, números e símbolos especiais"
		case 5:
			fmt.Println("Você escolheu sair do programa.")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha entre 1 e 4.")
			return
		}
		if tipo != "" {
			fmt.Printf("Você escolheu o comprimento da senha: %d associado a %s.\n", comprimento, tipo)
			break
		}
	}

	generatedPassword, err := password.RandonPassword(tipo, comprimento)
	if err != nil {
		fmt.Println("Erro ao gerar a senha:", err)
		return
	}
	
	fmt.Println("Gerando senha...")
	time.Sleep(3 * time.Second)
	fmt.Printf("Senha gerada com sucesso: %s\n",generatedPassword)
	fmt.Println("---------------------------------------")


	fmt.Println("Gerando Arquivo...")
	time.Sleep(5 * time.Second)
	_ , err = file.GeneratedFile("/Users/mayara.perez/courses/password-generator/output/fileData.txt", generatedPassword)
	if err != nil {
		fmt.Println("Erro ao gerar o arquivo:", err)
		return
	}

	fmt.Println("Arquivo gerado com sucesso")
	fmt.Println("---------------------------------------")




}
