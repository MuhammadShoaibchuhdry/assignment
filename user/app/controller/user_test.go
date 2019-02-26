package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"assignment/user/app/model"
	"assignment/user/app/routes"
	"assignment/user/app/service"

	"github.com/go-chi/chi"
	"gopkg.in/mgo.v2"
)

func Init() (*chi.Mux, *mgo.Session, error) {
	goc, err := service.GetSession()
	if err != nil {
		return nil, nil, err
	}
	userDB := service.NewUserDbClient(goc)
	ctrl := NewUserRepository(userDB)
	r := routes.NewRouter(ctrl)
	return r, goc, nil
}

func TestUser_CreateUserHandler(t *testing.T) {
	r, goc, err := Init()
	if err != nil {
		t.Fatal(err)
	}
	user := model.UserSignUp{
		Name: "test",
		Password: "1234",
		Email: "test@gmail.com",
	}
	payload, _ := json.Marshal(user)
	ts := httptest.NewServer(r)
	defer ts.Close()
	req, err := http.NewRequest("POST", ts.URL+"/register", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer goc.Close()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var response model.Success
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}
	if response.Message != "success" {
		t.Fatal("message should be success")
	}
}

func TestUser_GetUserHandler(t *testing.T) {
	r, goc, err := Init()
	if err != nil {
		t.Fatal(err)
	}
	user := model.UserSignIn{
		Email: "test@gmail.com",
		Password: "1234",
	}
	payload, _ := json.Marshal(user)
	ts := httptest.NewServer(r)
	req, err := http.NewRequest("POST", ts.URL+"/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer goc.Close()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var response model.UserResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}
	if response.AccessToken == "" {
		t.Fatal("access token should be there")
	}
}

func TestUser_UpdateUserHandler(t *testing.T) {
	r, goc, err := Init()
	if err != nil {
		t.Fatal(err)
	}
	user := model.UserSignIn{
		Email: "test@gmail.com",
		Password: "1234",
	}
	payload, _ := json.Marshal(user)
	ts := httptest.NewServer(r)
	req, err := http.NewRequest("POST", ts.URL+"/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var response model.UserResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}
	if response.AccessToken == "" {
		t.Fatal("access token should be there")
	}
	userUpdate := model.UserUpdate{
		Name: "updatedName",
		PhoneNumber: "1234",
	}
	payloadUpdate, _ := json.Marshal(userUpdate)
	defer ts.Close()
	reqUpdate, err := http.NewRequest("PUT", ts.URL+"/updateuserprofile", bytes.NewBuffer(payloadUpdate))
	if err != nil {
		t.Fatal(err)
	}
	reqUpdate.Header.Set("Authorization", fmt.Sprintf("Bearer %v", response.AccessToken))
	defer goc.Close()
	respUpdate, err := http.DefaultClient.Do(reqUpdate)
	if err != nil {
		t.Fatal(err)
	}
	respBodyUpdate, err := ioutil.ReadAll(respUpdate.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var responseUpdate model.Success
	if err := json.Unmarshal(respBodyUpdate, &responseUpdate); err != nil {
		t.Fatal(err)
	}
	if responseUpdate.Message != "success" {
		t.Fatal("response should be success")
	}
}

func TestUser_DeleteUserHandler(t *testing.T) {
	r, goc, err := Init()
	if err != nil {
		t.Fatal(err)
	}
	user := model.UserSignIn{
		Email: "test@gmail.com",
		Password: "1234",
	}
	payload, _ := json.Marshal(user)
	ts := httptest.NewServer(r)
	req, err := http.NewRequest("POST", ts.URL+"/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	defer goc.Close()
	var response model.UserResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}
	if response.AccessToken == "" {
		t.Fatal("access token should be there")
	}
	reqDelete, err := http.NewRequest("DELETE", ts.URL+"/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqDelete.Header.Add("Authorization", fmt.Sprintf("Bearer %v", response.AccessToken))
	respUpdate, err := http.DefaultClient.Do(reqDelete)
	if err != nil {
		t.Fatal(err)
	}
	respBodyDelete, err := ioutil.ReadAll(respUpdate.Body)
	if err != nil {
		t.Fatal(err)
	}
	var responseDelete model.Success
	if err := json.Unmarshal(respBodyDelete, &responseDelete); err != nil {
		t.Fatal(err)
	}
	if responseDelete.Message != "success" {
		t.Fatal("response should be success")
	}
}

