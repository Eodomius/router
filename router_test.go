package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

const PORT = "8080"

func StartServer(t *testing.T){

	var router = New()
	http.Handle("/", router)

	router.Use(func (w http.ResponseWriter, req *http.Request){
		fmt.Println("Middleware 1")
		time.Sleep(time.Second * 2)
	})
	router.Get("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
		w.Write([]byte("Get 0"))
	})
	router.Get("/test/{id}", func(w http.ResponseWriter, r *http.Request, result *Result){
		w.Write([]byte("Get 1: " + result.Params["id"]))
	})
	router.Get("/test/{id}/test", func(w http.ResponseWriter, r *http.Request, result *Result){
		w.Write([]byte("Get 2: " + result.Params["id"]))
	})
	router.Post("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
		w.Write([]byte("Post: Test"))
	})
	router.Patch("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
		w.Write([]byte("Patch: Test"))
	})
	router.Put("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
		w.Write([]byte("Put: Test"))
	})
	router.Delete("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
		w.Write([]byte("Delete: Test"))
	})

	http.ListenAndServe("127.0.0.1:"+PORT, nil)
}
var wg  sync.WaitGroup
// Test requests for the router
func TestRequests(t *testing.T){
	go StartServer(t);
	go TRequest(t, "http://localhost:"+PORT+"/test", "GET", "Get 0")
	go TRequest(t, "http://localhost:"+PORT+"/test/", "GET", "Get 0")
	go TRequest(t, "http://localhost:"+PORT+"/test/abc", "GET", "Get 1: abc")
	go TRequest(t, "http://localhost:"+PORT+"/test/def/test", "GET", "Get 2: def")
	go TRequest(t, "http://localhost:"+PORT+"/test", "POST", "Post: Test")
	go TRequest(t, "http://localhost:"+PORT+"/test", "PATCH", "Patch: Test")
	go TRequest(t, "http://localhost:"+PORT+"/test", "PUT", "Put: Test")
	go TRequest(t, "http://localhost:"+PORT+"/test", "DELETE", "Delete: Test")
	wg.Add(8)
	wg.Wait()
}

func TRequest(t *testing.T, url, method, exeptedResult string){
	defer wg.Done()
	client := &http.Client{}
	var (
		request *http.Request
		response *http.Response
		err error
	)
	switch method {
		case "GET":
		request, _ = http.NewRequest(http.MethodGet, url, nil)
		response, err = client.Do(request)
		case "POST":
		request, _ = http.NewRequest(http.MethodPost, url, nil)
		response, err = client.Do(request)
		case "PATCH":
		request, _ = http.NewRequest(http.MethodPatch, url, nil)
		response, err = client.Do(request)
		case "PUT":
		request, _ = http.NewRequest(http.MethodPut, url, nil)
		response, err = client.Do(request)
		case "DELETE":
		request, _ = http.NewRequest(http.MethodDelete, url, nil)
		response, err = client.Do(request)
	}

	if err != nil {
		t.Error("Error: ", err)
	}
	if response.StatusCode != 200 {
		t.Error("Status code should be 200, but is: ", response.StatusCode)
	}
	// Read an check if responce is "Get 0"
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("Error: ", err)
	}
	if string(body) != exeptedResult {
		t.Error("Response should be '",exeptedResult,"' but is: ", string(body))
	}
}