package alidns

// Client for alidns config
type Client struct {
	SecretID  string
	SecretKey string
}

// New return a alidns client
func New(secretID, secretKey string) *Client {
	return &Client{
		SecretID:  secretID,
		SecretKey: secretKey,
	}
}
