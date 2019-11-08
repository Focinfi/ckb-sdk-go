package ckbtypes

type Witness struct {
	Lock       string
	InputType  string
	OutputType string
}

func (witness Witness) Clone() *Witness {
	return &Witness{
		Lock:       witness.Lock,
		InputType:  witness.InputType,
		OutputType: witness.OutputType,
	}
}
