package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// retorna router com as rotas da aplicação

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
