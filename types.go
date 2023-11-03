package main

type (
	TestType      string
	Status        int
	MetabolicType string
)

const (
	DNASlimV2   TestType      = "mydnaslim-v2"
	ProteinType MetabolicType = "protein"
)

const (
	WaitingForSync Status = iota
	WaitingForImport
	WaitingForInterpret
	Success
)
