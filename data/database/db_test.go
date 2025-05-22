package database

import (
	"testing"

	"gorm.io/gorm"
)

func TestInit(t *testing.T) {
	type args struct {
		test bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Init with test = true",
			args: args{test: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.test)
		})
	}
}

func TestDB(t *testing.T) {
	Init(true)
	tests := []struct {
		name string
		want *gorm.DB
	}{
		{
			name: "Test DB",
			want: DB(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DB(); !(got == tt.want) {
				t.Errorf("DB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClose(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test Close",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
