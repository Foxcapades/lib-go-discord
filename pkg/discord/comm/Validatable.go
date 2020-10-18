package comm

type Validatable interface {
	IsValid() bool
	Validate() error
}
