package dio

type APIQuery interface {
	ToQuery() string
}
