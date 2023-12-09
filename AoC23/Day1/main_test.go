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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processCalibrationData(tt.args.rawData); got != tt.want {
				t.Errorf("processCalibrationData() = %v, want %v", got, tt.want)
			}
		})
	}
}
