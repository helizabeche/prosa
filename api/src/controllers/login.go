package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
// Login realiza a autenticação de um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
*/

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//deu um erro bem aqui
	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, []byte(usuario.Senha)); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	w.Write([]byte("Login realizado com sucesso!"))

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))

	/*
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}



		//if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, []byte(usuario.Senha)); erro != nil {
		//erro de codigo divergente do tutorial
		if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {

			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}


	*/

}
