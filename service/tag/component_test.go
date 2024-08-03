package tag

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	//"github.com/alaa2amz/g1/service"

	//	"github.com/gin-gonic/gin"
	"fmt"
	"log"

	//"os"
	"github.com/stretchr/testify/assert"
)

var lastID uint

// type APIResBody struct {data ,error any}
func pv(v any) { log.Printf("%+v\n", v) }
func TestCreateAPI(t *testing.T) {
	w := httptest.NewRecorder()
	m := Proto()
	m.Title = "test title"
	m.Content = "test content"
	j, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api"+Path, strings.NewReader(string(j)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	R.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	/*
		rm := struct {
			Data  any
			Error any
		}{nil, nil}
	*/
	rm := map[string]any{}
	err = json.Unmarshal(w.Body.Bytes(), &rm)
	if err != nil {
		t.Fatal(err)
	}
	pv(rm)
	pv(string(w.Body.Bytes()))
	lastID = uint((rm["data"].(map[string]any)["id"].(float64)))
	pv(lastID)

}

func TestRetrieveAPI(t *testing.T) {

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api"+Path, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	R.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetAPI(t *testing.T) {

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api"+Path+"/"+fmt.Sprintf("%d", lastID), nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	R.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateAPI(t *testing.T) {
	m := Proto()
	m.Title = "updated title"
	m.Content = "updated Content"
	j, err := json.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api"+Path+"/"+fmt.Sprintf("%d", lastID), bytes.NewReader(j))

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	R.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	pv(w.Body.String())
}

func TestDeleteAPI(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/api"+Path+"/"+fmt.Sprintf("%d", lastID), nil)
	if err != nil {
		t.Fatal(err)
	}
	R.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
func TestGetDeletedAPI(t *testing.T) {

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api"+Path+"/"+fmt.Sprintf("%d", lastID), nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	R.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
