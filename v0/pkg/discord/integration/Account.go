package integration

type Account interface {
	// id of the account
	ID() string
	SetID(string) Account

	// name of the account
	Name() string
	SetName() Account
}
