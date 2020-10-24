package lib

type Validatable interface {
	IsValid() bool
	Validate() error
}
