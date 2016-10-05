package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"github.com/EvgenKostenko/go-harvest/models"
)

func TestGetProjects_GetAll(t *testing.T) {
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

	projects, _, err := testClient.Project.Projects(nil)

	if projects == nil {
		t.Error("Expected projects list. Project list is nil")
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetProjects_GetAllWithParameters(t *testing.T) {
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

	projects, _, err := testClient.Project.Projects(projectOptions)

	if projects == nil {
		t.Error("Expected project list. Project list is nil")
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetProjects_Sucsess(t *testing.T) {
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

func TestGetProjects_WrongAPIEndpoint(t *testing.T) {
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

func TestGetProjects_WrongDataFromResponse(t *testing.T) {
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


// Create project

func TestCreateProject_Success(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/projects"

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, testAPIEndpoint)
	})

	project := models.Project{ClientID: 4868513, Name: "NEW PROJECT 26", Active: true, Notes: "Hello"}
	structProject := ProjectDetail{Project: project}

	resp, err := testClient.Project.CreateProject(&structProject)

	if resp.StatusCode != 200 {
		t.Errorf("Expected Status code 200. Given %d", resp.StatusCode)
	}

	if err != nil {
		t.Error(err.Error())
	}
}

// Update project

func TestUpdateProject_Success(t *testing.T) {
	setup()
	defer teardown()
	project := models.Project{ID: 11832718, ClientID: 4868513, Name: "NEW PROJECT 26", Active: true, Notes: "Hello"}
	structProject := ProjectDetail{Project: project}

	testAPIEndpoint := fmt.Sprintf("/projects/%d", project.ID)
	raw, err := ioutil.ReadFile("./mocks/project.json")

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})


	resp, err := testClient.Project.UpdateProject(&structProject)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200. Got %s", resp.Status)
	}

	if err != nil {
		t.Error(err.Error())
	}
}
