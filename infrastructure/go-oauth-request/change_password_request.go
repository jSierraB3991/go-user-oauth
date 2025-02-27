package gooauthrequest

type ChangePasswordRequest struct {
	KeycloakUserId string
	NewPassword    string
	PrePassword    string
}
