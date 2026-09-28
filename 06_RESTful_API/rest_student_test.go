package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStudentAPI(t *testing.T) {
	handler := MakeWebHandler()
	request := func(method, path, body string) *httptest.ResponseRecorder {
		t.Helper()
		res := httptest.NewRecorder()
		req := httptest.NewRequest(method, path, strings.NewReader(body))
		handler.ServeHTTP(res, req)
		return res
	}

	res := request(http.MethodGet, "/students", "")
	if res.Code != http.StatusOK || res.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("GET /students: status=%d, content-type=%q", res.Code, res.Header().Get("Content-Type"))
	}
	var list []Student
	if err := json.NewDecoder(res.Body).Decode(&list); err != nil {
		t.Fatal(err)
	}
	if len(list) != 2 || list[0].Id != 1 || list[1].Id != 2 {
		t.Fatalf("GET /students: unexpected sorted list: %+v", list)
	}

	res = request(http.MethodGet, "/students/2", "")
	var student Student
	if res.Code != http.StatusOK {
		t.Fatalf("GET /students/2: status=%d", res.Code)
	}
	if err := json.NewDecoder(res.Body).Decode(&student); err != nil {
		t.Fatal(err)
	}
	if student.Name != "bbb" {
		t.Fatalf("GET /students/2: student=%+v", student)
	}
	if res := request(http.MethodGet, "/students/999", ""); res.Code != http.StatusNotFound {
		t.Fatalf("GET /students/999: status=%d", res.Code)
	}
	if res := request(http.MethodPost, "/students", "{"); res.Code != http.StatusBadRequest {
		t.Fatalf("POST malformed JSON: status=%d", res.Code)
	}
	if res := request(http.MethodPost, "/students", `{"Name":""}`); res.Code != http.StatusBadRequest {
		t.Fatalf("POST empty name: status=%d", res.Code)
	}

	res = request(http.MethodPost, "/students", `{"Id":99,"Name":"ccc","Age":15,"Score":78}`)
	if res.Code != http.StatusCreated {
		t.Fatalf("POST /students: status=%d", res.Code)
	}
	res = request(http.MethodGet, "/students/3", "")
	if res.Code != http.StatusOK {
		t.Fatalf("GET /students/3: status=%d", res.Code)
	}
	if err := json.NewDecoder(res.Body).Decode(&student); err != nil {
		t.Fatal(err)
	}
	if student.Id != 3 || student.Name != "ccc" {
		t.Fatalf("GET /students/3: student=%+v", student)
	}

	if res := request(http.MethodDelete, "/students/1", ""); res.Code != http.StatusOK {
		t.Fatalf("DELETE /students/1: status=%d", res.Code)
	}
	if res := request(http.MethodGet, "/students/1", ""); res.Code != http.StatusNotFound {
		t.Fatalf("GET deleted student: status=%d", res.Code)
	}
	if res := request(http.MethodDelete, "/students/1", ""); res.Code != http.StatusNotFound {
		t.Fatalf("DELETE deleted student: status=%d", res.Code)
	}
}
