// pacote que fica entre a requisição e a resposta, podendo modificar a requisição ou a resposta, ou até mesmo encerrar a requisição antes de chegar na rota final.
package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//escreve informacoes de log sobre cada requisição NO TERMINAL
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requisição recebida: %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

// verifica se o usuário está autenticado antes de permitir o acesso a uma rota protegida. Se não estiver autenticado, retorna um erro 401 Unauthorized.
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}
}

