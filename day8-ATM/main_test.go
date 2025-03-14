package main

import (
	"reflect"
	"testing"
)

func TestNewATM(t *testing.T) {
	for _, tc := range []struct {
		initial float64
		want    float64
	}{
		{1000.0, 1000.0},
		{-1000.0, 0.0},
		{0.0, 0.0},
		{1e10, 1e10},
	} {
		atm := NewATM(tc.initial)
		got := atm.Balance()
		if got != tc.want {
			t.Errorf("Balance() = %v, want %v", got, tc.want)
		}
	}
}

func TestATMWithdraw(t *testing.T) {
	tests := []struct {
		initialAmount float64
		withdrawal    float64
		want          float64
	}{
		{initialAmount: 1000.0, withdrawal: 500.0, want: 500.0},
		{initialAmount: 1000.0, withdrawal: 1500.0, want: 0.0}, // assume balance can't be negative
		{initialAmount: 0.0, withdrawal: 100.0, want: 0.0},
	}

	for _, test := range tests {
		atm := NewATM(test.initialAmount)
		atm.Withdraw(test.withdrawal)
		got := atm.Balance()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Balance() = %v, want %v", got, test.want)
		}
	}
}

func TestATMDeposit(t *testing.T) {
	tests := []struct {
		initialAmount float64
		deposit       float64
		want          float64
	}{
		{initialAmount: 1000.0, deposit: 500.0, want: 1500.0},
		{initialAmount: 0.0, deposit: 100.0, want: 100.0},
		{initialAmount: 1000.0, deposit: -500.0, want: 1000.0}, // assume deposit can't be negative
	}

	for _, test := range tests {
		atm := NewATM(test.initialAmount)
		atm.Deposit(test.deposit)
		got := atm.Balance()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Balance() = %v, want %v", got, test.want)
		}
	}
}
func (a *ATM) Balance() float64 {
	return a.balance
}
func NewATM(f float64) interface {
	panic("unimplemented")
}