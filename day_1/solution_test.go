package day1

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkTrebuchetPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trebuchetPart1(input)
	}
}

func BenchmarkTrebuchetPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trebuchetPart2(input)
	}
}

func Test_trebuchetPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet"},
		}, want: 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trebuchetPart1(tt.args.input); got != tt.want {
				t.Errorf("trebuchetPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trebuchetPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]string{"1"}}, want: 11},
		{name: "12", args: args{[]string{"12"}}, want: 12},
		{name: "one", args: args{[]string{"one"}}, want: 11},
		{name: "oneone", args: args{[]string{"oneone"}}, want: 11},
		{name: "oneight", args: args{[]string{"oneeight"}}, want: 18},
		{name: "oneightwoneight", args: args{[]string{"oneightwoneight"}}, want: 18},
		{name: "two1nine", args: args{[]string{"two1nine"}}, want: 29},
		{name: "eightwothree", args: args{[]string{"eightwothree"}}, want: 83},
		{name: "abcone2threexyz", args: args{[]string{"abcone2threexyz"}}, want: 13},
		{name: "xtwone3four", args: args{[]string{"xtwone3four"}}, want: 24},
		{name: "4nineeightseven2", args: args{[]string{"4nineeightseven2"}}, want: 42},
		{name: "zoneight234", args: args{[]string{"zoneight234"}}, want: 14},
		{name: "7pqrstsixteen", args: args{[]string{"7pqrstsixteen"}}, want: 76},
		{name: "example", args: args{[]string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}}, want: 281},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trebuchetPart2(tt.args.input); got != tt.want {
				t.Errorf("trebuchet_part_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
