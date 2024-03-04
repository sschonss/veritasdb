package table

import (
	"fmt"
	"os"
	"strings"
)

type Table struct {
	Name     string
	DataPath string
	
}


func NewTable(name, dataPath string) *Table {
	dataPath = dataPath + ".csv"
	return &Table{
		Name:     name,
		DataPath: dataPath,
	}
}

func SetTable(name, dataPath string) *Table {
	// Cria o diretório se ele não existir
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		fmt.Println("Tabela não encontrado. Criando...")
		err := os.Mkdir(dataPath, 0755)
		if err != nil {
			fmt.Println("Erro ao criar tabela 'data/table':", err)
			return nil
		}
	}

	return &Table{
		Name:     name,
		DataPath: dataPath,
	}
}

func ListTable() []string {
	var tables_name []string
	if _, err := os.Stat("data/table"); os.IsNotExist(err) {
		fmt.Println("Nenhum banco de dados encontrado")
		return tables_name
	}
	files, err := os.ReadDir("data/table")
	if err != nil {
		fmt.Println("Erro ao listar bancos de dados:", err)
		return tables_name
	}

	for _, file := range files {
		tables_name = append(tables_name, file.Name())
	}
	return tables_name
}

func (db *Table) ExecuteQuery(query string) string {
	command := strings.Fields(query)[0]
	switch strings.ToLower(command) {
	case "clear":
		fmt.Print("\033[H\033[2J")
	case "select":
		fmt.Println("Executando comando SELECT")
	case "insert":
		fmt.Println("Executando comando INSERT")
	case "update":
		fmt.Println("Executando comando UPDATE")
	case "delete":
		fmt.Println("Executando comando DELETE")
	case "create":
		fmt.Println("Executando comando CREATE")
	case "drop":
		fmt.Println("Executando comando DROP")
	case "alter":
		fmt.Println("Executando comando ALTER")
	default:
		fmt.Println("Comando não reconhecido")
	}

	return "Resultado da execução do comando: " + query
}