package message

// Activity
// TODO: Document me
type Activity interface {
	// type of message activity
	Type() ActivityType
	SetType(ActivityType) Activity

	// party_id from a Rich Presence event
	PartyId() string
	PartyIdIsSet() bool
	SetPartyId(string) Activity
	UnsetPartyId() Activity
}
