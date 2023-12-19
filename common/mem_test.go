package common

import (
	"reflect"
	"testing"
)

func Test_getPrivKeyFromMnemonic(t *testing.T) {
	type args struct {
		mnemonic string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test mem to priv",
			args: args{
				mnemonic: "blast about old claw current first paste risk involve victory edit current",
			},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("getPrivKeyFromMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := getPrivKeyFromMnemonic(tt.args.mnemonic)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPrivKeyFromMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrivKeyFromMnemonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
