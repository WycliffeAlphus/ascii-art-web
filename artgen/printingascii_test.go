package artgen

import "testing"

func TestPrintingAscii(t *testing.T) {
	type args struct {
		text        string
		patternFile string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Invalid inputs",args{"hello","shadow"},"",true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrintingAscii(tt.args.text, tt.args.patternFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintingAscii() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrintingAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleInput(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Valid input",args{"hello"},"hello",false},
		{"Invalid input",args{"hello✊"},"",true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleInput(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
