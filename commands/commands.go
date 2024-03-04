package commands

import (
	"fmt"
	"os"
	"strings"
)

func CreateTable(command string) {

	table := strings.Split(command, " ")[2]

	columns := strings.Split(command, "(")[1]
	columns = strings.Split(columns, ")")[0]
	columns = strings.Replace(columns, ",", ";", -1)
	columns = strings.Replace(columns, " ", "", -1)
	columns = columns + "\n"

	if _, err := os.Stat("data/" + table + ".csv"); err == nil {
		fmt.Println("Tabela já existe")
		return
	}

	file, err := os.Create("data/" + table + ".csv")
	if err != nil {
		fmt.Println("Erro ao criar arquivo")
		return
	}
	defer file.Close()

	_, err = file.WriteString(columns)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo")
		return
	}

	fmt.Println("Tabela criada com sucesso")
}
