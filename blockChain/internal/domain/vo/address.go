package vo

type Address struct {
	PublicKey  string `json:"publicKey"`
	Identifier string `json:"identifier"`
	Host       string `json:"host"`
	Port       string `json:"port"`
}
