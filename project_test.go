package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestProjectsGetAll(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/projects"

	raw, err := ioutil.ReadFile("./mocks/projects.json")

	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	people, _, err := testClient.Project.Projects(nil)

	if people == nil {
		t.Error("Expected projects list. Project list is nil")
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestProjectsGetAllWithParameters(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/projects"

	raw, err := ioutil.ReadFile("./mocks/projects.json")

	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	projectOptions := &ProjectOptions{
		Client: "3554414",
		UpdatedSince: "1985-09-30+9:00",
	}

	people, _, err := testClient.Project.Projects(projectOptions)

	if people == nil {
		t.Error("Expected project list. Project list is nil")
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
