package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BasicWork() bool {
	return true
}

func AddOne(s []int) []int {
	out := make([]int, len(s))
	for i, elem := range s {
		out[i] = elem + 1
	}
	return out
}

func TestBasicWork(t *testing.T) {
	assert.True(t, BasicWork(), "return true")
}

func TestAddOne(t *testing.T) {
	testCases := []struct {
		S    []int
		Want []int
	}{
		{
			S:    []int{1, 1, 1},
			Want: []int{2, 2, 2},
		},
		{
			S:    []int{1, 2, 3},
			Want: []int{2, 3, 4},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, AddOne(tc.S), tc.Want)
	}
}

func TestMap(t *testing.T) {
	myMap := map[string]string{}
	myMap["Name"] = "alpha"
	myMap["Email"] = "alpha@email.com"
	assert.Contains(t, myMap, "Name", "attribute Name should exist")
	assert.Contains(t, myMap, "Email", "attribute Email should exist")
}

func TestStruct(t *testing.T) {
	type User struct {
		Id      int
		Name    string
		IsValid bool
	}

	user1 := &User{
		Id:      0,
		Name:    "alpha",
		IsValid: false,
	}

	user2 := &User{
		Id:      1,
		Name:    "beta",
		IsValid: true,
	}

	assert.EqualExportedValues(t, user1, user1)
	assert.NotEqualValues(t, user1, user2)
}

func TestSubsets(t *testing.T) {
	s1 := []string{"alpha", "beta", "gamma", "delta"}
	s2 := []string{"alpha", "beta"}
	assert.Subset(t, s1, s2, "s1 does not contain s2")
}
