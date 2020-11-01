package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testUserError string

func (e testUserError) Error() string {
	return e.Message()
}

func (e testUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testUserError("User Error")
}

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic("err panic")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown error")
}

func errNoError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "No Error")
	return nil
}

type args struct {
	handler appHandler
}

var tests = []struct {
	name        string
	args        args
	wantcode    int
	wantmessage string
}{
	{"Panic 500", args{errPanic}, 500, "Internal Server Error"},
	{"User Error 400", args{errUserError}, 400, "User Error"},
	{"Not Found 404", args{errNotFound}, 404, "Not Found"},
	{"No Permission 403", args{errNoPermission}, 403, "Forbidden"},
	{"Unknown Error 500", args{errUnknown}, 500, "Internal Server Error"},
	{"No Error 200", args{errNoError}, 200, "No Error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := errWrapper(tt.args.handler)
			response := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "https://lroolle.com/", nil)
			f(response, request)
			assertResponse(response.Result(), tt.wantcode, tt.wantmessage, t)
		})
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := errWrapper(tt.args.handler)
			server := httptest.NewServer(http.HandlerFunc(f))
			response, _ := http.Get(server.URL)
			assertResponse(response, tt.wantcode, tt.wantmessage, t)
		})
	}
}

func assertResponse(response *http.Response, wantcode int, wantmessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")

	if response.StatusCode != wantcode || body != wantmessage {
		t.Errorf("Expect Code: %d, Message: %s; Got Code: %d, Message: %s",
			wantcode, wantmessage, response.StatusCode, body)
	}
}
