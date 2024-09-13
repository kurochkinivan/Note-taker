package constants

import "time"

// hash and jwt token
const (
	Salt       = "dkwf;ljewpop0f9s9pfjio2=[a]ca[fkaw]"
	Signingkey = "w0d@#a0wdWDAPWD;aw;@wa@!"
	TokenTTL   = 12 * time.Hour
)

// postgresql entities
const (
	scheme     = "public."
	UsersTable = scheme + "users"
	NotesTable = scheme + "notes"
)
