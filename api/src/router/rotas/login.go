package rotas

import (
	"net/http"
	"api/src/controllers"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
