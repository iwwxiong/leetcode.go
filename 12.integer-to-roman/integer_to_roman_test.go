package algorithms12

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"4->IV", args{4}, "IV"},
		{"9->IX", args{9}, "IX"},
		{"40->XL", args{40}, "XL"},
		{"400->CD", args{400}, "CD"},
		{"900->CM", args{900}, "CM"},
		{"3000->MMM", args{3000}, "MMM"},
		{"3900->MMMCM", args{3900}, "MMMCM"},
		{"1994->MCMXCIV", args{1994}, "MCMXCIV"},
		{"3999->MMMCMXCIX", args{3999}, "MMMCMXCIX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
