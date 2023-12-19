package common

import (
	"encoding/hex"
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
		want    string
		wantErr bool
	}{
		{
			name: "test mem to priv",
			args: args{
				mnemonic: "blast about old claw current first paste risk involve victory edit current",
			},
			want:    "69668f2378b43009b16b5c6eb5e405d9224ca2a326a65a17919e567105fa4e5a",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPrivKeyFromMnemonic(tt.args.mnemonic)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPrivKeyFromMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(hex.EncodeToString(got), tt.want) {
				t.Errorf("getPrivKeyFromMnemonic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSeedFromMnemonic(t *testing.T) {
	type args struct {
		mnemonic string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test get seed",
			args: args{
				mnemonic: "blast about old claw current first paste risk involve victory edit current",
			},
			want:    "dd5ffa7088c0fa4c665085bca7096a61e42ba92e7243a8ad7fbc6975a4aeea1845c6b668ebacd024fd2ca215c6cd510be7a9815528016af3a5e6f47d1cca30dd",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSeedFromMnemonic(tt.args.mnemonic)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSeedFromMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getSeedFromMnemonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
