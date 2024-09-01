package constants

// Message define a standard structure for messages.
type Message struct {
	Status  int
	Message string
}

func NewMessage(status int, message string) Message {
	return Message{
		Status:  status,
		Message: message,
	}
}

var (
	ErrInvalidCredentials = NewMessage(401, "Invalid credentials")

	ErrUserNotFound = NewMessage(404, "User not found")

	ErrUserAlreadyExists = NewMessage(409, "User already exists")

	ErrBadRequest = NewMessage(400, "Bad request")

	ErrUnprocessableEntity = NewMessage(422, "Validation failed")

	ErrHashingPassword = NewMessage(500, "Failed to hash password")

	ErrCreatingUser = NewMessage(500, "Failed to create user")

	ErrInternalServer = NewMessage(500, "Internal server error")

	ErrUnauthorized = NewMessage(401, "Unauthorized")

	SuccessOperationCompleted = NewMessage(200, "Operation completed successfully")

	SuccessUserRegistered = NewMessage(201, "User registered successfully")

	SuccessLoginSuccessful = NewMessage(200, "Login successful")

	SuccessLogoutSuccessful = NewMessage(200, "Logout successful")

	SuccessPasswordResetEmailSent = NewMessage(200, "Password reset email sent")

	ErrPasswordResetToken = NewMessage(400, "Invalid password reset token")

	SuccessPasswordReset = NewMessage(200, "Password reset successfully")

	ErrEmailAlreadyVerified = NewMessage(400, "Email already verified")

	ErrInvalidVerificationLink = NewMessage(400, "Invalid verification link")

	ErrEmailVerificationFailed = NewMessage(500, "Email verification failed")

	SuccessEmailVerified = NewMessage(200, "Email verified successfully")

	ErrSendingVerificationEmail = NewMessage(500, "Failed to send verification email")
)
