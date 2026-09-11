package seguranca

import "golang.org/x/crypto/bcrypt"


// Hash gera um hash a partir da senha recebida
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara a senha recebida com a senha armazenada no banco de dados

func VerificarSenha(senha string, senhaComHash []byte) error {
	return bcrypt.CompareHashAndPassword(senhaComHash, []byte(senha))
}
