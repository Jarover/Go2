package lesson1

import (
	"reflect"
	"testing"
)

func TestOpenMyFileFunc(t *testing.T) {

	tests := []struct {
		name      string
		arg       string
		want      string
		wantErr   bool
		wantPanic bool
	}{
		// arg is exist and File is exist
		{
			name:      "1",
			arg:       `testfile.txt`,
			want:      `some text`,
			wantErr:   false,
			wantPanic: false,
		},
		// arg is exist and File is not exist
		{
			name:      "2",
			arg:       `testfile`,
			want:      ``,
			wantErr:   true,
			wantPanic: false,
		},
		// arg is not exist
		{
			name:      "3",
			arg:       ``,
			want:      "",
			wantErr:   true,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OpenMyFile() paniced when wantPanic = %v", tt.wantPanic)
				}
			}()

			got, err := ReadMyFile(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadMyFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
