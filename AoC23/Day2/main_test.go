package main

import (
	"testing"
)

func Test_parseGameNumber(t *testing.T) {
	type args struct {
		gameHeader string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Happy path", args{"Game 1"}, 1},
		{"Happy path", args{"Game 27"}, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGameNumber(tt.args.gameHeader); got != tt.want {
				t.Errorf("parseGameNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkMaxColorValuesForImpossibility(t *testing.T) {
	type args struct {
		games string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid Game", args{"1 green, 6 red, 5 blue; 3 green, 2 red, 4 blue; 3 green, 1 red, 3 blue"}, true},
		{"Invalid Game", args{"6 green, 1 red, 9 blue; 11 red, 4 blue, 12 green; 6 green, 9 red, 19 blue; 2 green, 6 blue; 10 green, 1 red, 16 blue; 4 green, 14 blue"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMaxColorValuesForImpossibility(tt.args.games); got != tt.want {
				t.Errorf("checkMaxColorValuesForImpossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_powerFunction(t *testing.T) {
	type args struct {
		games string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First Game", args{"1 green, 6 red, 5 blue; 3 green, 2 red, 4 blue; 3 green, 1 red, 3 blue"}, 3 * 6 * 5},
		{"Second Game", args{"6 green, 1 red, 9 blue; 11 red, 4 blue, 12 green; 6 green, 9 red, 19 blue; 2 green, 6 blue; 10 green, 1 red, 16 blue; 4 green, 14 blue"}, 12 * 11 * 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powerFunction(tt.args.games); got != tt.want {
				t.Errorf("powerFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
