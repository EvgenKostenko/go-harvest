package harvest

import (
	"fmt"
	"github.com/EvgenKostenko/go-harvest/models"
	"io/ioutil"
	"net/http"
	"testing"
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

func TestGetProject_Sucsess(t *testing.T) {
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

func TestGetProject_NoProjects(t *testing.T) {
	setup()
	defer teardown()
	projectId := 2222

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

	if resp.StatusCode == 404 {
		t.Errorf("Expected status 404. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetProject_ServerError(t *testing.T) {
	projectId := 2222

	testClient, _ = NewClient(nil, "https://harvest.com/test")
	u, _, err := testClient.Project.GetProject(projectId)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

// Create a project

func TestCreateProject_Success(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/projects"

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, testAPIEndpoint)
	})

	project := models.Project{ClientID: 4868513, Name: "NEW PROJECT 26", Active: true, Notes: "Hello"}

	resp, err := testClient.Project.CreateProject(&project)

	if resp.StatusCode != 200 {
		t.Errorf("Expected Status code 200. Given %d", resp.StatusCode)
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateProject_WrongData(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/projects"

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	project := models.Project{ClientID: 4868513, Name: "NEW PROJECT 26", Active: true, Notes: "Hello"}

	resp, err := testClient.Project.CreateProject(&project)

	if resp.StatusCode != 400 {
		t.Errorf("Expected status 400. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

// Update a project

func TestUpdateProject_Success(t *testing.T) {
	setup()
	defer teardown()
	project := models.Project{ID: 11832718, ClientID: 4868513, Name: "NEW PROJECT 26", Active: true, Notes: "Hello"}

	testAPIEndpoint := fmt.Sprintf("/projects/%d", project.ID)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, "{}")
	})

	resp, err := testClient.Project.UpdateProject(&project)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200. Got %s", resp.Status)
	}

	if err != nil {
		t.Error(err.Error())
	}
}

// Delete a project

func TestDeleteProject_Success(t *testing.T) {
	setup()
	defer teardown()
	projectId := 11832718

	testAPIEndpoint := fmt.Sprintf("/projects/%d", projectId)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, "{}")
	})

	resp, err := testClient.Project.DeleteProject(projectId)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200. Got %s", resp.Status)
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestDeleteProjects_NoProjects(t *testing.T) {
	setup()
	defer teardown()
	projectId := 2222

	testAPIEndpoint := fmt.Sprintf("/projects/%d", projectId)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	resp, err := testClient.Project.DeleteProject(projectId)

	if resp.StatusCode != 400 {
		t.Errorf("Expected status 400. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}
