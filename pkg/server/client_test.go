package server

import (
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestUpdateSSID(t *testing.T) {
	is:=is.New(t)
	server:=httptest.NewServer(FakeServer.router)
	err:=UpdateSSID(server.URL+"/graphql", FakeData,FakeUser, FakePass)
	if err != nil {
		t.Error(err)
		return
	}

	ssid, err := GetPOESSID(server.URL+"/graphql", FakeUser, FakePass)
	if err != nil {
		t.Error(err)
		return
	}
	is.Equal(ssid, FakeData)
}



