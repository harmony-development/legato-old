package mongodocs

// Theme is an object that stores the users app theme prefs
type Theme struct {
	primary   string
	secondary string
	overall   string
}

// User is a structure for the mongodb document
type User struct {
	email    string
	password string
	username string
	userid   string
	avatar   string
	theme    Theme
}
