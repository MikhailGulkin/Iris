package vo

import "fmt"

type Address struct {
	PublicKey  string `json:"publicKey"`
	Identifier string `json:"identifier"`
	Host       string `json:"host"`
	Port       string `json:"port"`
}

func NewAddress(publicKey string, identifier string, host string, port string) Address {
	return Address{PublicKey: publicKey, Identifier: identifier, Host: host, Port: port}
}

func (a *Address) FullAddress() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}
func (a *Address) Validate() bool {
	return a.PublicKey != "" && a.Identifier != "" && a.Host != "" && a.Port != ""
}
func (a *Address) URL(protocol string) string {
	return protocol + "://" + a.FullAddress()
}
