package harvest

import (
	"fmt"
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
