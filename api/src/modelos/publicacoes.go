package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa a estrutura de dados para uma publicação.
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}


//chama os metodos de validação e formatação da publicação. Se houver algum erro de validação, ele será retornado. Caso contrário, a publicação será formatada e o método retornará nil.
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o título é obrigatório")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo é obrigatório")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
