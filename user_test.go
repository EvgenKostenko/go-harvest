package harvest

import (
	"fmt"
	"github.com/EvgenKostenko/go-harvest/models"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPeopleGetAll(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/people"

	raw, err := ioutil.ReadFile("./mocks/people.json")

	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	people, _, err := testClient.User.People(nil)

	if people == nil {
		t.Error("Expected people list. People list is nil")
	}

	if err != nil {
		t.Errorf("Error gilen: %s", err)
	}
}

func TestPeopleGetAllWithParameters(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/people"

	raw, err := ioutil.ReadFile("./mocks/people.json")

	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	peopleOptions := &PeopleOptions{
		UpdatedSince: "1985-09-30+9:00",
	}

	people, _, err := testClient.User.People(peopleOptions)

	if people == nil {
		t.Error("Expected people list. People list is nil")
	}

	if err != nil {
		t.Errorf("Error gilen: %s", err)
	}
}

func TestGetUserSucsess(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/people/1"

	raw, err := ioutil.ReadFile("./mocks/user.json")

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	u, _, err := testClient.User.GetUser(1)

	if u == nil {
		t.Error("Expected user with 1, get nil")
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

func TestGetUser_NoUser(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/people/2"

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, nil)
	})

	u, resp, err := testClient.User.GetUser(1)

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

func TestCreateUser_Success(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/people"

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, testAPIEndpoint)
	})

	user := models.UserParameters{Email: "test@test.com",
		FirstName: "TestName",
		LastName:  "TestSoname",
	}
	resp, err := testClient.User.CreateUser(&user)

	if resp.StatusCode != 200 {
		t.Errorf("Expected Status code 200. Given %d", resp.StatusCode)
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateUser_WrongData(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/people"

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	user := models.UserParameters{Email: "wrong@test.com"}
	resp, err := testClient.User.CreateUser(&user)

	if resp.StatusCode != 400 {
		t.Errorf("Expected status 400. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}
