package dmeta

type Validatable interface {
	IsValid() bool
	Validate() error
}
