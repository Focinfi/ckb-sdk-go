package ckbtypes

type Input struct {
	PreviousOutput OutPoint `json:"previous_output"`
	Since          string   `json:"since"`
}
