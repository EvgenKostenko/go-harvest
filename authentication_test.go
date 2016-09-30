package harvest

import (
	"net/http"
	"testing"
)

func TestAcquire(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/account/who_am_i", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/account/who_am_i")

		// Emulate error
		w.WriteHeader(http.StatusInternalServerError)
	})

	res, err := testClient.Authentication.Acquire("foo", "bar")
	if err == nil {
		t.Errorf("Expected error, but no error given")
	}
	if res == true {
		t.Error("Expected error, but result was true")
	}

	if testClient.Authentication.Authenticated() != false {
		t.Error("Expected false, but result was true")
	}
}

//func TestAcquire(t *testing.T) {
//	TODO
//
//}
