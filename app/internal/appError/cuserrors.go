package apperror

// US - user service
var (
	ErrNotFound = NewAppErr(nil, "not found", "entity was not found", "US-001")
)

// VS - validation service
var (
	ErrValidateData = NewAppErr(nil, "failed to validate data", "failed to validate data", "VS-001")
	ErrSerializeData = NewAppErr(nil, "failed serialize/desirialize data", "failed serialize/desirialize data", "VS-002")
)

// AS - auth server
var (
	ErrSignToken = NewAppErr(nil, "failed to sign jwt-token", "failed to sign jwt-token", "AS-001")
	ErrInvalidPassword = NewAppErr(nil, "invalid password provided", "invalid password provided", "AS-002")
)
