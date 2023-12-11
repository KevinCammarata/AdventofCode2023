package main

import (
	"testing"
)

// func Test_processCalibrationData2(t *testing.T) {

// 	testName = "Digits - same"
// 	want = 88
// 	got = processCalibrationData("aaa8fivthkre8on")

// 	if got != want {
// 		t.Errorf("Test: %s result = %v, want %v", testName, got, want)
// 	}

// 	testName = "Digits - different"

// }

func Test_processCalibrationData(t *testing.T) {
	type args struct {
		rawData string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Text numbers  same", args{"threeabcdthree"}, 33},
		{"Text numbers - different", args{"twothreexyzfour"}, 24},
		{"Digits - same", args{"aaa8fivthkre8on"}, 88},
		{"Digits - different", args{"aaa6fivthkre9on"}, 69},
		{"Digits and Numerals", args{"onethreexyzfour5"}, 15},
		{"Digits and Numerals2", args{"2onethreexyzfour5six"}, 26},
		{"Only Digits", args{"123"}, 13},
		{"One Digit, only at the beginning", args{"4twscpht"}, 44},
		{"One Digit, only at the end", args{"twscpht4"}, 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processCalibrationData(tt.args.rawData); got != tt.want {
				t.Errorf("processCalibrationData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processCalibrationDataDigitsOnly(t *testing.T) {
	type args struct {
		rawData string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Digits - same", args{"aaa8fivthkre8on"}, 88},
		{"Digits - different", args{"aaa6fivthkre9on"}, 69},
		{"Digits and Numerals", args{"one2threexyzfour5"}, 25},
		{"Digits and Numerals2", args{"4onethreexyzfour7six"}, 47},
		{"Zero Test", args{"4one0threexyzfour1six"}, 41},
		{"Only Digits", args{"123"}, 13},
		{"One Digit, only at the beginning", args{"4twoscpht"}, 44},
		{"One Digit, only at the end", args{"twoscpht4"}, 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processCalibrationDataDigitsOnly(tt.args.rawData); got != tt.want {
				t.Errorf("processCalibrationDataDigitsOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
