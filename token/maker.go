package token

import "time"

// Maker is an interface for managin tokens
type Maker interface {
	//create a token for a specific email and duration
	CreateToken(mobile string, duration time.Duration) (string, error)
	//VerfityToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
