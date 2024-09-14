package apperror

// US - user service. This is US error.
var (
	ErrNotFound = NewAppErr(nil, "not found", "entity was not found", "US-001")
)

// VS - validation service. This is VS error.
var (
	ErrValidateData = NewAppErr(nil, "failed to validate data", "failed to validate data", "VS-001")
	ErrSerializeData = NewAppErr(nil, "failed serialize/desirialize data", "failed serialize/desirialize data", "VS-002")
)

// AS - auth server. This is AS error.
var (
	ErrSignToken = NewAppErr(nil, "failed to sign jwt-token", "failed to sign jwt-token", "AS-001")
	ErrInvalidPassword = NewAppErr(nil, "invalid password provided", "invalid password provided", "AS-002")
	ErrInvalidSigningMethod = NewAppErr(nil, "invalid signing method", "invalid signing method", "AS-003")
	ErrAssertingJWT = NewAppErr(nil, "token claims are not of type jwt.MapClaims", "token claims are not of type jwt.MapClaims", "AS-004")
	ErrEmptyAuthHeader = NewAppErr(nil, "empty auth header", "empty auth header", "AS-005")
	ErrInvalidAuthHeader = NewAppErr(nil, "invalid auth header", "invalid auth header", "AS-006")
	ErrTokenExired = NewAppErr(nil, "access token is expired", "access token is expired", "AS-007")
)

