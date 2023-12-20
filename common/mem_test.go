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
				mnemonic: "delay wall sentence supreme sunny fat cream write polar claw candy exile",
			},
			want:    "1a592f33fd5fd80528722f0f84dd93dcf10409d85f1e3a94f17a4309fafd6cc1",
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
				t.Errorf("getPrivKeyFromMnemonic() = %v, want %v", hex.EncodeToString(got), tt.want)
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
				mnemonic: "delay wall sentence supreme sunny fat cream write polar claw candy exile",
			},
			want:    "9d4023a00ffba730e3b4c214bb03ee6ec9d7d3f107460be0a1e2b21fadadf0c39099c0bbb3d6b78c5c72a4fc95ca860fa7570efa99b079c5972412391a19e6fa",
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
