package mongodocs

// Theme is an object that stores the users app theme prefs
type Theme struct {
	Primary   string
	Secondary string
	Overall   string
}

// User is a structure for the mongodb document
type User struct {
	Email    string
	Password string
	Username string
	Userid   string
	Avatar   string
	Theme    Theme
}
