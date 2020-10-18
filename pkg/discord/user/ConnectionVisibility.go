package user

type ConnectionVisibility uint8

const (
	ConnectionVisibilityNone ConnectionVisibility = iota
	ConnectionVisibilityEveryone
)

func (c ConnectionVisibility) IsValid() bool {
	return c == ConnectionVisibilityNone ||
		c == ConnectionVisibilityEveryone
}