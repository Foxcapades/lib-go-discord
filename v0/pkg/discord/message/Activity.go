package message

// Activity
// TODO: Document me
type Activity interface {
	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of message activity.
	Type() ActivityType

	// SetType overwrites the current value of this record's `type` field.
	SetType(ActivityType) Activity

	// PartyID returns the current value of this record's `party_id` field.
	//
	// The `party_id` field contains the party_id from a Rich Presence event.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PartyIDIsSet to check if the field is present before use.
	PartyID() string

	// PartyIDIsSet returns whether this record's `party_id` field is currently present.
	PartyIDIsSet() bool

	// SetPartyID overwrites the current value of this record's `party_id` field.
	SetPartyID(string) Activity

	// UnsetPartyID removes this record's `party_id` field.
	UnsetPartyID() Activity
}
