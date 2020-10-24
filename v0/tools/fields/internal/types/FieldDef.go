package types

type FieldDef struct {
	// Type name prefix
	Name string

	// Go type
	Type string

	Testing Testing

	Constructor string

	AssignCast string
}
