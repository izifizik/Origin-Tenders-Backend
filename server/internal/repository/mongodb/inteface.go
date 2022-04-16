package mongodb

type Repository interface {
	SaveToken(ID string, token string) error
	ProofToken(ID string, token string) error
}
