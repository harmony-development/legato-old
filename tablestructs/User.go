package tablestructs

type User struct {
	Id string `db:"id"`
	Username string `db:"username"`
	Email string `db:"email"`
	Password string `db:"password"`
}