package modelos


// Senha representa a estrutura de dados para a senha de um usuário.
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}
