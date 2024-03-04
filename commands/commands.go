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

func DropTable(command string) {

	table := strings.Split(command, " ")[2]

	if _, err := os.Stat("data/" + table + ".csv"); os.IsNotExist(err) {
		fmt.Println("Tabela não existe")
		return
	}

	err := os.Remove("data/" + table + ".csv")
	if err != nil {
		fmt.Println("Erro ao deletar tabela")
		return
	}

	fmt.Println("Tabela deletada com sucesso")
}

func InsertInto(command string) {

	//create table table (column1, column2, column3)
	//ex: insert into table (column1, column2, column3) values (value1, value2, value3)
	
	table := strings.Split(command, " ")[2]

	if _, err :=
		os.Stat("data/" + table + ".csv"); os.IsNotExist(err) {
		fmt.Println("Tabela não existe")
		return
	}

	//escrever os values
	values := strings.Split(command, "values")[1]
	values = strings.Split(values, "(")[1]
	values = strings.Split(values, ")")[0]
	values = strings.Replace(values, ",", ";", -1)
	values = strings.Replace(values, " ", "", -1)
	values = values + "\n"

	file, err := os.OpenFile("data/" + table + ".csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo")
		return
	}
	defer file.Close()

	_, err = file.WriteString(values)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo")
		return
	}

	fmt.Println("Registro inserido com sucesso")
}