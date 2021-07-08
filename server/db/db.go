package db

type IAuthDB interface {
	ValidateToken(token string) (bool, error)
	ExpireToken(token string) error
}

type IHarmonyDB interface {
	IAuthDB
	Migrate() error
}
