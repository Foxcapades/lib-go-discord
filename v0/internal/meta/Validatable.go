package meta

type Validatable interface {
	IsValid() bool
	Validate() error
}
