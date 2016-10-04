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
		Client:       "3554414",
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

func TestGetProjectSucsess(t *testing.T) {
	setup()
	defer teardown()
	projectId := 11832718

	testAPIEndpoint := fmt.Sprintf("/projects/%d", projectId)

	raw, err := ioutil.ReadFile("./mocks/project.json")

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	u, _, err := testClient.Project.GetProject(projectId)

	if u == nil {
		t.Errorf("Expected project with %d, get nil", projectId)
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetProjectWrongAPIEndpoint(t *testing.T) {
	setup()
	defer teardown()
	projectId := 11832718

	testAPIEndpoint := fmt.Sprintf("projects/%d", projectId)

	raw, err := ioutil.ReadFile("./mocks/project.json")

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	u, resp, err := testClient.Project.GetProject(projectId)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if resp.Status == "404" {
		t.Errorf("Expected status 404. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetProjectWrongDataFromResponse(t *testing.T) {
	setup()
	defer teardown()
	projectId := 11832718

	testAPIEndpoint := fmt.Sprintf("/projects/%d", projectId)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, "{ 'foo': 'bar' }")
	})

	u, resp, err := testClient.Project.GetProject(projectId)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if resp.Status == "200" {
		t.Errorf("Expected status 200. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetProjects_NoProjects(t *testing.T) {
	setup()
	defer teardown()
	projectId := 2222

	testAPIEndpoint := fmt.Sprintf("/projects/%d", projectId)
	raw, err := ioutil.ReadFile("./mocks/projects.json")

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	u, resp, err := testClient.User.GetUser(projectId)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if resp.Status == "404" {
		t.Errorf("Expected status 404. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetProjects_ServerError(t *testing.T) {
	projectId := 2222

	testClient, _ = NewClient(nil, "https://harvest.com/test")
	u, _, err := testClient.User.GetUser(projectId)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}