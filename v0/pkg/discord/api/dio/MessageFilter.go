package dio

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

// The before, after, and around keys are mutually exclusive, only one may be
// passed at a time.
type MessageFilter interface {
	// Around filters to messages around the given message ID.
	//
	// If this method is called on a param that is unset, this method will panic.
	// Use AroundIsSet to check if the param is present before use.
	Around() discord.Snowflake

	// AroundIsSet returns whether this filter's `around` param is currently
	// present.
	AroundIsSet() bool

	// SetAround overwrites the current value of this filter's `around` param.
	SetAround(discord.Snowflake) MessageFilter

	// UnsetAround removes this filter's `around` param.
	UnsetAround() MessageFilter

	// Before returns the current value of this record's `before` param.
	//
	// The `before` param contains the
	//
	// If this method is called on a param that is unset, this method will panic.
	// Use BeforeIsSet to check if the param is present before use.
	Before() discord.Snowflake

	// BeforeIsSet returns whether this record's `before` param is currently present.
	BeforeIsSet() bool

	// SetBefore overwrites the current value of this record's `before` param.
	SetBefore(discord.Snowflake) MessageFilter

	// UnsetBefore removes this record's `before` param.
	UnsetBefore() MessageFilter

	// After returns the current value of this record's `after` param.
	//
	// The `after` param contains the
	//
	// If this method is called on a param that is unset, this method will panic.
	// Use AfterIsSet to check if the param is present before use.
	After() discord.Snowflake

	// AfterIsSet returns whether this record's `after` param is currently present.
	AfterIsSet() bool

	// SetAfter overwrites the current value of this record's `after` param.
	SetAfter(discord.Snowflake) MessageFilter

	// UnsetAfter removes this record's `after` param.
	UnsetAfter() MessageFilter

	// Limit returns the current value of this record's `limit` param.
	//
	// The `limit` param contains the
	//
	// If this method is called on a param that is unset, this method will panic.
	// Use LimitIsSet to check if the param is present before use.
	Limit() RecordLimit

	// LimitIsSet returns whether this record's `limit` param is currently present.
	LimitIsSet() bool

	// SetLimit overwrites the current value of this record's `limit` param.
	SetLimit(RecordLimit) MessageFilter

	// UnsetLimit removes this record's `limit` param.
	UnsetLimit() MessageFilter
}