package model

type Credentials struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

func (c *Credentials) IsValid() bool {
	return c.ConsumerKey != "" ||
		c.ConsumerSecret != "" ||
		c.AccessToken != "" ||
		c.AccessSecret != ""
}
