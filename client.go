package alidns

type Client struct {
	SecretID  string
	SecretKey string
}

func New(secretID, secretKey string) *Client {
	return &Client{
		SecretID:  secretID,
		SecretKey: secretKey,
	}
}
