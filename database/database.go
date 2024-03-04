package database

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"veritasdb/commands"
)

type Database struct {
	User     string
	Password string
}

func (Database) ExecuteQuery(query string) string {
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
		commands.CreateTable(query)
	case "drop":
		fmt.Println("Executando comando DROP")
	case "alter":
		fmt.Println("Executando comando ALTER")
	default:
		fmt.Println("Comando não reconhecido")
	}

	return "Resultado da execução do comando: " + query
}

func (Database) SaveConfig(user, password string) {
	fmt.Println("Salvando configuração...")

	configFile, err := os.Create("data/config.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de configuração")
		return
	}
	defer configFile.Close()

	_, err = configFile.WriteString("user:" + user + "\n")
	if err != nil {
		fmt.Println("Erro ao salvar usuário no arquivo de configuração")
		return
	}

	_, err = configFile.WriteString("password:" + password + "\n")
	if err != nil {
		fmt.Println("Erro ao salvar senha no arquivo de configuração")
		return
	}

	fmt.Println("Configuração salva com sucesso!")

}

func (Database) LoadConfig() (string, string) {
	fmt.Println("Carregando configuração...")

	configFile, err := os.Open("data/config.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de configuração")
		return "", ""
	}
	defer configFile.Close()

	var user, password string
	fmt.Fscanf(configFile, "user:%s\n", &user)
	fmt.Fscanf(configFile, "password:%s\n", &password)

	fmt.Println("Configuração carregada com sucesso!")

	return user, password
}

func (Database) CheckConfig() bool {
	_, err := os.Stat("data/config.txt")
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (Database) CheckUser(user, password string) bool {
	savedUser, savedPassword := Database{}.LoadConfig()
	return user == savedUser && password == savedPassword
}

func (Database) CreateUser(user, password string) {
	Database{}.SaveConfig(user, password)
}

func (Database) UpdateUser(user, password string) {
	Database{}.SaveConfig(user, password)
}

func (Database) DeleteUser() {
	os.Remove("data/config.txt")
}

func (Database) GetUsers() []string {
	return []string{"admin"}
}

func GetUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}



