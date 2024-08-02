package artgen

import (
	"reflect"
	"testing"
)

func TestAsciiMapping(t *testing.T) {
	type args struct {
		patternFile string
	}
	tests := []struct {
		name    string
		args    args
		want    map[rune][]string
		wantErr bool
	}{
		{"Mapping",args{""},nil,true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AsciiMapping(tt.args.patternFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("AsciiMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsciiMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
