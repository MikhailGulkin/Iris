package vo

import "testing"

func TestLocationMetadata_DistanceTo(t *testing.T) {
	type args struct {
		other LocationMetadata
	}
	tests := []struct {
		name   string
		fields LocationMetadata
		args   args
		want   string
	}{
		{name: "DistanceTo 1", fields: LocationMetadata{Latitude: 1, Longitude: 1}, args: args{other: NewLocationMetadata(1, 1)}, want: "0"},
		{name: "DistanceTo 2", fields: LocationMetadata{Latitude: 1, Longitude: 1}, args: args{other: NewLocationMetadata(2, 2)}, want: "157.23"},
		{name: "DistanceTo 3", fields: LocationMetadata{Latitude: 1, Longitude: 1}, args: args{other: NewLocationMetadata(3, 3)}, want: "314.40"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLocationMetadata(tt.fields.Latitude, tt.fields.Longitude)
			if got := l.DistanceTo(tt.args.other); got != tt.want {
				t.Errorf("DistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestLocationMetadata_IsValid(t *testing.T) {
	type fields struct {
		Latitude  float64
		Longitude float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "IsValid 1", fields: fields{Latitude: -90, Longitude: -180}, want: true},
		{name: "IsValid 2", fields: fields{Latitude: 0, Longitude: 0}, want: true},
		{name: "IsValid 3", fields: fields{Latitude: 90, Longitude: 180}, want: true},
		{name: "IsValid 4", fields: fields{Latitude: -91, Longitude: -180}, want: false},
		{name: "IsValid 5", fields: fields{Latitude: 91, Longitude: -180}, want: false},
		{name: "IsValid 6", fields: fields{Latitude: -90, Longitude: -181}, want: false},
		{name: "IsValid 7", fields: fields{Latitude: -90, Longitude: 181}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLocationMetadata(tt.fields.Latitude, tt.fields.Longitude)
			if got := l.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
