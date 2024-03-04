package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	configFile, err := os.Open("data/login.txt")
	if err != nil {
		fmt.Println("Arquivo de configuração não encontrado. Por favor, insira as informações de login.")
		fmt.Print("Usuário: ")
		user := getUserInput()
		fmt.Print("Senha: ")
		password := getUserInput()
		
		saveConfig(user, password)


	} else {
		scanner := bufio.NewScanner(configFile)
		scanner.Scan()
		user := strings.Split(scanner.Text(), ":")[1]
		scanner.Scan()
		password := strings.Split(scanner.Text(), ":")[1]

		fmt.Print("Usuário: ")
		inputUser := getUserInput()
		fmt.Print("Senha: ")
		inputPassword := getUserInput()

		for i := 0; i < 3; i++ {
			if inputUser == user && inputPassword == password {
				break
			} else {
				fmt.Println("Usuário ou senha inválidos. Tente novamente.")
				fmt.Print("Usuário: ")
				inputUser = getUserInput()
				fmt.Print("Senha: ")
				inputPassword = getUserInput()
			}
		}
		if inputUser != user || inputPassword != password {
			fmt.Println("Usuário ou senha inválidos. Tente novamente.")
			return
		}
			
	}

	for {
		fmt.Print("Digite um comando SQL (ou 'exit' para sair): ")
		command := getUserInput()

		if strings.ToLower(command) == "exit" {
			break
		}
	}
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func saveConfig(user, password string) {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Println("Erro ao criar o diretório 'data':", err)
			return
		}
	}

	configFile, err := os.Create("data/login.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de configuração:", err)
		return
	}
	defer configFile.Close()

	_, err = configFile.WriteString("user:" + user + "\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo de configuração:", err)
		return
	}

	_, err = configFile.WriteString("password:" + password + "\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo de configuração:", err)
		return
	}

	fmt.Println("Configurações salvas com sucesso.")
}


type Database struct {
	Name     string
	DataPath string
}

func NewDatabase(name, dataPath string) *Database {
	return &Database{
		Name:     name,
		DataPath: dataPath,
	}
}

func (db *Database) ExecuteQuery(query string) string {

	command := strings.Fields(query)[0]
	switch strings.ToLower(command) {
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
