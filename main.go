package main

import (
	"fmt"

	"veritasdb/database"
)

func main() {

	var db database.Database

	dirData := database.CheckExisteDirData()
	if !dirData {
		fmt.Println("Diretório data não encontrado.")
		fmt.Println("Criando diretório data...")
		database.CreateDirData()
		fmt.Println("Diretório data criado com sucesso.")
	}

	if config {
		fmt.Println("Configurações encontradas.")
		fmt.Println("Digite o usuário e a senha para carregar as configurações.")
		fmt.Print("Usuário: ")
		user := database.GetUserInput()
		fmt.Print("Senha: ")
		password := database.GetUserInput()
		valid := db.CheckUser(user, password)
		count_try := 0
		for !valid {
			count_try++
			if count_try > 2 {
				fmt.Println("Número de tentativas excedido.")
				return
			}
			fmt.Println("Usuário ou senha inválidos.")
			fmt.Print("Usuário: ")
			user = database.GetUserInput()
			fmt.Print("Senha: ")
			password = database.GetUserInput()
			valid = db.CheckUser(user, password)
		}
		db = database.Database{User: user, Password: password}

	}else {
		fmt.Println("Nenhuma configuração encontrada.")
		fmt.Println("Digite o usuário e a senha para salvar as configurações.")
		fmt.Print("Usuário: ")
		user := database.GetUserInput()
		fmt.Print("Senha: ")
		password := database.GetUserInput()
		db = database.Database{User: user, Password: password}
		db.CreateUser(user, password)
	}



	fmt.Println("Bem-vindo ao VeritasDB!")

	for {
		fmt.Print("veritasdb> ")
		query := database.GetUserInput()
		if query == "exit" {
			break
		}
		fmt.Println(db.ExecuteQuery(query))

	}
}
