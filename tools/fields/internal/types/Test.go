package types


type Test struct {
	ActualTransform func(s string) string
	ExpectTransform func(s string) string
	InputValue      string
	ExpectValue     string
	Comparison      string
	CompareParams   []string
}

func (t Test) HasActualTransform() bool {
	return t.ActualTransform != nil
}

func (t Test) HasExpectTransform() bool {
	return t.ExpectTransform != nil
}

func (t Test) HasCompParams() bool {
	return len(t.CompareParams) > 0
}