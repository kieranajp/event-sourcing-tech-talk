package main

import (
	"encoding/json"
	"github.com/nbari/violetear"
	"github.com/valyala/fastjson"
	"io"
	"net/http"
)

func CreateTestsHandler(w http.ResponseWriter, r *http.Request) {
	ids := []string{"DE83123412341234", "DE83987698769876", "DE83543254325432"}
	for _, id := range ids {
		testkit := TestKit{}
		_ = testkit.SetID(id)
		testkit.TrackChange(&testkit, &TestWasCreated{Type: DNASlimV2})

		err := repo.Save(&testkit)
		if err != nil {
			panic("failed to save aggregate: " + err.Error())
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func TestSyncHandler(w http.ResponseWriter, r *http.Request) {
	testID := violetear.GetParam("testID", r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic("failed to read request body: " + err.Error())
	}
	r.Body.Close()

	accountID := fastjson.GetString(body, "account_id")

	testkit := TestKit{}
	err = repo.Get(testID, &testkit)
	if err != nil {
		panic("failed to get aggregate: " + err.Error())
	}

	testkit.TrackChange(&testkit, &TestWasSynced{AccountID: accountID})

	err = repo.Save(&testkit)
	if err != nil {
		panic("failed to save aggregate: " + err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
}

func TestArrivedAtLabHandler(w http.ResponseWriter, r *http.Request) {
	testID := violetear.GetParam("testID", r)

	testkit := TestKit{}
	err := repo.Get(testID, &testkit)
	if err != nil {
		panic("failed to get aggregate: " + err.Error())
	}

	testkit.TrackChange(&testkit, &TestArrivedAtLab{})

	err = repo.Save(&testkit)
	if err != nil {
		panic("failed to save aggregate: " + err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetTestHandler(w http.ResponseWriter, r *http.Request) {
	testID := violetear.GetParam("testID", r)

	testkit := TestKit{}
	err := repo.Get(testID, &testkit)
	if err != nil {
		panic("failed to get aggregate: " + err.Error())
	}

	out, _ := json.Marshal(testkit)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}
