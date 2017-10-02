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

func Test_containsKeyWord(t *testing.T) {
	type args struct {
		str      string
		keywords []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"{", []string{"{", "}", "+", "=", "]", "[", ")", "("}}, true},
		{"2", args{"int", []string{"{", "}", "+", "=", "]", "[", ")", "("}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsKeyWord(tt.args.str, tt.args.keywords); got != tt.want {
				t.Errorf("minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeCode(t *testing.T) {

	str := `boolean t = true;
	while(t){
		t = false;
		for(int i = 0; i < mas.length - 1; i++){
			if (mas[i] > mas[i+1]){
				int temp = mas[i];
				mas[i] = mas[i+1];
				mas[i+1] = temp;
				t = true;
			}
		}
	}
	`

	encodeCode(str)

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

		// {"Test ints",
		// 	fields{
		// 		nil,
		// 		`int a = 2`,
		// 		`int b = 3`,
		// 		"",
		// 		0,
		// 		0,
		// 		"",
		// 	},
		// },
		{"Test bubble sorts",
			fields{
				nil,
				`
				boolean t = true;
				while(t){
					t = false;
					for(int i = 0; i < mas.length - 1; i++){
						if (mas[i] > mas[i+1]){
							int temp = mas[i];
							mas[i] = mas[i+1];
							mas[i+1] = temp;
							t = true;
						}
					}
				}
				`,
				`
				for(int j = tab.length - 1; j >= 0; j--){
					for(int i = 0; i < j; i++){
						if (tab[i] > tab[i+1]){
							int tmp = tab[i];
							tab[i] = tab[i+1];
							tab[i+1] = tmp;
						}
					}
				}
				`,
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
