package main

import (
	"github.com/hallgren/eventsourcing"
)

type (
	TestKit struct {
		eventsourcing.AggregateRoot

		ID            string        `json:"id"`
		AccountID     string        `json:"account_id,omitempty"`
		Type          TestType      `json:"type"`
		Status        Status        `json:"status"`
		MetabolicType MetabolicType `json:"metabolic_type,omitempty"`
	}
)

func (testkit *TestKit) Transition(event eventsourcing.Event) {
	switch e := event.Data().(type) {
	case *TestWasCreated:
		testkit.ID = event.AggregateID()
		testkit.Type = e.Type
		testkit.Status = WaitingForSync

	case *TestWasSynced:
		testkit.AccountID = e.AccountID
		testkit.Status = WaitingForImport

	case *TestArrivedAtLab:
		testkit.Status = WaitingForInterpret

	case *TestResultsReceived:
		testkit.MetabolicType = e.MetabolicType
		testkit.Status = Success
	}
}

func (testkit *TestKit) Register(r eventsourcing.RegisterFunc) {
	r(&TestWasCreated{}, &TestWasSynced{}, &TestArrivedAtLab{}, &TestResultsReceived{})
}
