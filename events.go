package main

type Event interface{}

type TestWasCreated struct {
	Type TestType
}

type TestWasSynced struct {
	AccountID string
}

type TestArrivedAtLab struct{}

type TestResultsReceived struct {
	MetabolicType MetabolicType
}
