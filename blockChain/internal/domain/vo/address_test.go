package vo

import "testing"

func TestAddress_FullAddress(t *testing.T) {
	tests := []struct {
		name   string
		fields Address
		want   string
	}{
		{name: "FullAddress 1", fields: Address{Host: "localhost", Port: "8080"}, want: "localhost:8080"},
		{name: "FullAddress 2", fields: Address{Host: "localhost", Port: "8081"}, want: "localhost:8081"},
		{name: "FullAddress 3", fields: Address{Host: "localhost", Port: "8082"}, want: "localhost:8082"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Address{
				PublicKey:  tt.fields.PublicKey,
				Identifier: tt.fields.Identifier,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
			}
			if got := a.FullAddress(); got != tt.want {
				t.Errorf("FullAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_Validate(t *testing.T) {
	tests := []struct {
		name   string
		fields Address
		want   bool
	}{
		{name: "Validate 1", fields: NewAddress("publicKey", "identifier", "localhost", "8080"), want: true},
		{name: "Validate 2", fields: NewAddress("", "identifier", "localhost", "8080"), want: false},
		{name: "Validate 3", fields: NewAddress("publicKey", "", "localhost", "8080"), want: false},
		{name: "Validate 4", fields: NewAddress("publicKey", "identifier", "", "8080"), want: false},
		{name: "Validate 5", fields: NewAddress("publicKey", "identifier", "localhost", ""), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Address{
				PublicKey:  tt.fields.PublicKey,
				Identifier: tt.fields.Identifier,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
			}
			if got := a.Validate(); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAddress_URL(t *testing.T) {
	type args struct {
		protocol string
	}
	tests := []struct {
		name   string
		fields Address
		args   args
		want   string
	}{
		{name: "URL 1", fields: Address{Host: "localhost", Port: "8080"}, args: args{protocol: "http"}, want: "http://localhost:8080"},
		{name: "URL 2", fields: Address{Host: "localhost", Port: "8081"}, args: args{protocol: "http"}, want: "http://localhost:8081"},
		{name: "URL 3", fields: Address{Host: "localhost", Port: "8082"}, args: args{protocol: "http"}, want: "http://localhost:8082"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Address{
				PublicKey:  tt.fields.PublicKey,
				Identifier: tt.fields.Identifier,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
			}
			if got := a.URL(tt.args.protocol); got != tt.want {
				t.Errorf("URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
