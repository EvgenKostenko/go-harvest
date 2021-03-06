package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// tests ByProject

func TestGetDayEntriesByProject_Sucsess(t *testing.T) {
	setup()
	defer teardown()
	projectId := 11832706

	testAPIEndpoint := fmt.Sprintf("/projects/%d/entries", projectId)

	raw, err := ioutil.ReadFile("./mocks/day_entries.json")

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, _, err := testClient.Report.DayEntriesByProject(projectId, reportOptions)

	if u == nil {
		t.Error("Expected day_entries list. Day_entries list is nil")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetDayEntriesByProject_BadRequest(t *testing.T) {
	setup()
	defer teardown()
	projectId := 22222

	testAPIEndpoint := fmt.Sprintf("/projects/%d/entries", projectId)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, resp, err := testClient.Report.DayEntriesByProject(projectId, reportOptions)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if resp.StatusCode != 400 {
		t.Errorf("Expected status 400. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetDayEntriesByProject_NoProjects(t *testing.T) {
	setup()
	defer teardown()
	projectId := 2222

	testAPIEndpoint := fmt.Sprintf("/projects/%d/entries", projectId)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, "{ 'foo': 'bar' }")
	})

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, resp, err := testClient.Report.DayEntriesByProject(projectId, reportOptions)

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

func TestGetDayEntriesByProject_ServerError(t *testing.T) {
	projectId := 2222

	testClient, _ := NewClient(nil, "https://harvest.com/test")

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, _, err := testClient.Report.DayEntriesByProject(projectId, reportOptions)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

// tests ByUser

func TestGetDayEntriesByUser_Sucsess(t *testing.T) {
	setup()
	defer teardown()
	userId := 1406631

	testAPIEndpoint := fmt.Sprintf("/people/%d/entries", userId)

	raw, err := ioutil.ReadFile("./mocks/day_entries.json")

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, _, err := testClient.Report.DayEntriesByUser(userId, reportOptions)

	if u == nil {
		t.Error("Expected day_entries list. Day_entries list is nil")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetDayEntriesByUser_BadRequest(t *testing.T) {
	setup()
	defer teardown()
	userId := 1406631

	testAPIEndpoint := fmt.Sprintf("/people/%d/entries", userId)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, resp, err := testClient.Report.DayEntriesByUser(userId, reportOptions)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if resp.StatusCode != 400 {
		t.Errorf("Expected status 400. Got %s", resp.Status)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGetDayEntriesByUser_NoProjects(t *testing.T) {
	setup()
	defer teardown()
	userId := 2222

	testAPIEndpoint := fmt.Sprintf("/people/%d/entries", userId)

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, "{ 'foo': 'bar' }")
	})

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, resp, err := testClient.Report.DayEntriesByUser(userId, reportOptions)

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

func TestGetDayEntriesByUser_ServerError(t *testing.T) {
	userId := 2222

	testClient, _ := NewClient(nil, "https://harvest.com/test")

	reportOptions := &ReportOptions{
		From: "1985-09-30+9:00",
		To:   "1985-09-30+9:00",
	}

	u, _, err := testClient.Report.DayEntriesByProject(userId, reportOptions)

	if u != nil {
		t.Errorf("Expected nil. Got %+v", u)
	}

	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}
