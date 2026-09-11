package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// conexao com o banco
var (
	StringConexaoBanco = ""

//PORTA ONDE A API VAI RODAR
	Porta = 0

//SECRET_KEY PARA ASSINAR O TOKEN
	SecretKey []byte
)

//inicia as variaveis de ambiente

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatalf("Erro ao carregar variaveis de ambiente: %s", erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
