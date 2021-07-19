package lesson1

import (
	"os"
	"testing"
)

func TestMyFileFunc(t *testing.T) {

	tests := []struct {
		name      string
		arg       string
		wantErr   bool
		wantPanic bool
	}{
		// arg is exist and FileName is correct
		{
			name:      "1",
			arg:       "testfile.txt",
			wantErr:   false,
			wantPanic: false,
		},
		// arg is exist and File is not correct
		{
			name: "2",
			arg:  "\\",

			wantErr:   true,
			wantPanic: false,
		},
		// arg is not exist
		{
			name:      "3",
			arg:       "",
			wantErr:   true,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = []string{"", tt.arg}
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("MyFile() paniced when wantPanic = %v", tt.wantPanic)
				}
			}()

			err := MyFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("MyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
