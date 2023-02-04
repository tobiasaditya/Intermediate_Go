package main

import (
	"10-unit-test/repository"
	mock_repository "10-unit-test/test/mock"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestLuasPersegi(t *testing.T) {
	type args struct {
		sisi int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"persegi 4", args{sisi: 4}, 16},
		{"persegi 2", args{sisi: 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LuasPersegi(tt.args.sisi); got != tt.want {
				t.Errorf("LuasPersegi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"username kosong", args{username: "", password: "password"}, true},
		{"password kosong", args{username: "username", password: ""}, true},
		{"sukses", args{username: "username", password: "password"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Register(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegisterToDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockUserRepo.EXPECT().Register("username", "password").Return(nil)
	mockUserRepo.EXPECT().Register("", "password").Return(errors.New("error"))
	type args struct {
		userRepo repository.IUserRepository
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"sukses", args{userRepo: mockUserRepo, username: "username", password: "password"}, false},
		{"username kosong", args{userRepo: mockUserRepo, username: "", password: "password"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterToDB(tt.args.userRepo, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("RegisterToDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
