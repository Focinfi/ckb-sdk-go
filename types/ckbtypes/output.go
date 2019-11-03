package ckbtypes

type Output struct {
	Capacity string  `json:"capacity"`
	Lock     Script  `json:"lock"`
	Type     *Script `json:"type,omitempty"`
}
