package banners

import "testing"

func TestIsValidBanner(t *testing.T) {
	type args struct {
		patternFile string
		textFile    []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid file",args{"shadow",[]byte{88, 40, 71, 6, 150, 77, 77, 228, 230, 34, 145, 25, 177, 146, 82, 213, 194, 159, 105, 87}},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBanner(tt.args.patternFile, tt.args.textFile); got != tt.want {
				t.Errorf("IsValidBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}
