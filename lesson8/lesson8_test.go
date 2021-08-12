package lesson8

import (
	"os"
	"testing"

	. "github.com/otiai10/copy"
)

func TestDuplicateFileFunc(t *testing.T) {

	tests := []struct {
		name        string
		arg         string
		delFlag     bool
		confirmFlag bool
		wantErr     bool
	}{
		// no deletion
		{
			name:        "1",
			arg:         "..\\testdir",
			delFlag:     false,
			confirmFlag: false,
			wantErr:     false,
		},
		// with deletion
		{
			name:        "2",
			arg:         "..\\testdir",
			delFlag:     true,
			confirmFlag: false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = []string{"", tt.arg}
			err := Copy("..\\testdir2", "..\\testdir")
			Start(tt.delFlag, tt.confirmFlag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
