package vo

import "testing"

func TestAmount_Add(t *testing.T) {
	type fields struct {
		Value float64
	}
	type args struct {
		amount Amount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{name: "Add 1 to 1", fields: fields{Value: 1}, args: args{amount: NewAmount(1)}, want: 2},
		{name: "Add 1 to 2", fields: fields{Value: 1}, args: args{amount: NewAmount(2)}, want: 3},
		{name: "Add 1 to 3", fields: fields{Value: 1}, args: args{amount: NewAmount(3)}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAmount(tt.fields.Value)
			a.Add(tt.args.amount)
			if a.Value != tt.want {
				t.Errorf("Add() = %v, want %v", a.Value, tt.want)
			}
		})
	}
}

func TestAmount_Compare(t *testing.T) {
	type fields struct {
		Value float64
	}
	type args struct {
		amount Amount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "Compare 1 to 1", fields: fields{Value: 1}, args: args{amount: NewAmount(1)}, want: true},
		{name: "Compare 1 to 2", fields: fields{Value: 1}, args: args{amount: NewAmount(2)}, want: false},
		{name: "Compare 1 to 3", fields: fields{Value: 1}, args: args{amount: NewAmount(3)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAmount(tt.fields.Value)
			if got := a.Compare(tt.args.amount); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_Subtract(t *testing.T) {
	type fields struct {
		Value float64
	}
	type args struct {
		amount Amount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{name: "Subtract 1 from 1", fields: fields{Value: 1}, args: args{amount: NewAmount(1)}, want: 0},
		{name: "Subtract 1 from 2", fields: fields{Value: 1}, args: args{amount: NewAmount(2)}, want: -1},
		{name: "Subtract 1 from 3", fields: fields{Value: 1}, args: args{amount: NewAmount(3)}, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAmount(tt.fields.Value)
			a.Subtract(tt.args.amount)
			if a.Value != tt.want {
				t.Errorf("Subtract() = %v, want %v", a.Value, tt.want)
			}
		})
	}
}
