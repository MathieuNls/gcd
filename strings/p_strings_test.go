package strings

import (
	"testing"
)

func Test_minimum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}}, 1},
		{"3", args{[]int{3, 5, 7}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimum(tt.args.nums...); got != tt.want {
				t.Errorf("minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevenshteinDistance(t *testing.T) {
	type args struct {
		s    string
		lenS int
		t    string
		lenT int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"kitten", 6, "sitting", 7}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevenshteinDistance(tt.args.s, tt.args.lenS, tt.args.t, tt.args.lenT); got != tt.want {
				t.Errorf("LevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCS(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{
			"1",
			args{"XMJYAUZ", "MZJAWXU"},
			4,
			"MJAU",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LCS(tt.args.first, tt.args.second)
			if got != tt.want {
				t.Errorf("LCS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LCS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBijectiveMorphisme_find(t *testing.T) {
	type fields struct {
		morphismes          []*BijectiveReplacement
		source              string
		target              string
		transformed         string
		LevenshteinDistance int
		lcs                 int
		lcsString           string
	}
	tests := []struct {
		name   string
		fields fields
	}{

		{"Test ints",
			fields{
				nil,
				`int a = 2`,
				`int b = 3`,
				"",
				0,
				0,
				"",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bm := &BijectiveMorphisme{
				morphismes:          tt.fields.morphismes,
				source:              tt.fields.source,
				target:              tt.fields.target,
				transformed:         tt.fields.transformed,
				LevenshteinDistance: tt.fields.LevenshteinDistance,
				lcs:                 tt.fields.lcs,
				lcsString:           tt.fields.lcsString,
			}
			bm.find()
		})
	}
}
