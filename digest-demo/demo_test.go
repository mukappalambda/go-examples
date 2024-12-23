package main

import (
	"testing"
)

func TestReturnSha256sum(t *testing.T) {
	testCases := []struct {
		Data string
		Want string
	}{
		{
			Data: "hello sha256sum",
			Want: "fe7f4136968f405bfea8cdd25470a8bc40bf4a59dbdf64eb311965527947da03",
		},
		{
			Data: "asdf",
			Want: "f0e4c2f76c58916ec258f246851bea091d14d4247a2fc3e18694461b1816e13b",
		},
		{
			Data: "i love go",
			Want: "4a27c9995d0929670e1370b70e29578537cd534578cad21fe1cc5249f0b6d466",
		},
	}
	for _, tc := range testCases {
		gotData := ReturnSha256sum([]byte(tc.Data))
		got := string(gotData)
		if got != tc.Want {
			t.Fatalf("got %s but want %s", got, tc.Want)
		}
	}
}
