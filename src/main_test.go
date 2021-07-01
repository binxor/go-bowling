package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCalculateFrameScore(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"5,4,/", 19},   // Test: Normal frame
		{"X", 10},       // Test: Normal frame
		{"0,/", 10},     // Test: Normal frame
		{"X,X,X,X", 30}, // Test: Max frame score for strikes is 30
		{"/,/,/", 20},   // Test: Max frame score for spares is 20
		{"", 0},         // Test: Empty frame
	}
	for _, c := range cases {
		fmt.Print(" * Testing GetFrameScore(" + c.in + ") == " + strconv.Itoa(c.want))
		got := CalculateFrameScore(c.in)
		if got != c.want {
			fmt.Println(" FAIL")
			t.Errorf("CalculateFrameScore(%q) == %d, want %d", c.in, got, c.want)
		} else {
			fmt.Println(" PASS")
		}
	}
}

func TestSumRolls(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"5,4,/", 19},   // Test: Normal frame
		{"X", 10},       // Test: Normal frame
		{"0,/", 10},     // Test: Normal frame
		{"X,X,X,X", 30}, // Test: Max frame score for strikes is 30
		{"/,/,/", 20},   // Test: Max frame score for spares is 20
		{"", 0},         // Test: Empty frame
	}
	for _, c := range cases {
		fmt.Print(" * Testing SumRolls(" + c.in + ") == " + strconv.Itoa(c.want))
		got := SumRolls(c.in)
		if got != c.want {
			fmt.Println(" FAIL")
			t.Errorf("SumRolls(%q) == %d, want %d", c.in, got, c.want)
		} else {
			fmt.Println(" PASS")
		}
	}
}
