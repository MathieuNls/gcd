package gcd

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
		morphismes  []*BijectiveReplacement
		source      string
		target      string
		transformed string
		cloneType   int
		simlarity   float32
	}
	tests := []struct {
		name   string
		fields fields
	}{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bm := New(tt.fields.source, 5)
			targetBm := New(tt.fields.target, 5)

			bm.AddTarget(targetBm.source, targetBm.encodedSource, targetBm.sourceLines, targetBm.sourcePrettyfied)
			bm.check()

			if bm.cloneType != type2 || bm.simlarity != 38.095238 {
				t.Errorf("Should be type 2 got %d. Should be 38.095238 got %f",
					bm.cloneType, bm.simlarity)

			}
		})
	}
}
