package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var diretorio string
	fmt.Println("Por gentileza digite o diretório que gostaria de listar")
	_, errScan := fmt.Scanln(&diretorio)

	diretorio = strings.Replace(diretorio, "%", " ", -1)

	if errScan != nil {
		main()
	} else if diretorio == "1" {
		fmt.Println("Saindo...")
		os.Exit(0)
	} else {

		var continuar uint8
		dir, err := os.ReadDir(diretorio)

		if err != nil {
			log.Println(err.Error())
			main()
		}

		fmt.Println("---------- MOSTRANDO DIRETORIOS ----------")

		for _, item := range dir {
			fmt.Println(item.Name())
		}

		fmt.Println("------------------------------------------")
		fmt.Println("Deseja ver outro diretório ? 0 [SIM] / 1 [NAO]")
		fmt.Scanln(&continuar)

		if continuar == 0 {
			main()
		} else if continuar == 1 {
			fmt.Println("Saindo...")
			os.Exit(0)
		}

	}

}
