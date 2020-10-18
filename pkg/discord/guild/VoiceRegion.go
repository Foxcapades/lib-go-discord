package guild

// VoiceRegion
// TODO: Documentation
type VoiceRegion interface {
	// unique ID for the region
	ID() string
	SetID(id string) VoiceRegion

	// name of the region
	Name() string
	SetName(name string) VoiceRegion

	// true if this is a vip-only server
	VIP() bool
	SetVIP(vip bool) VoiceRegion

	// true for a single server that is closest to the current user's client
	Optimal() bool
	SetOptimal(opt bool) VoiceRegion

	// whether this is a deprecated voice region (avoid switching to these)
	Deprecated() bool
	SetDeprecated(dep bool) VoiceRegion

	// whether this is a custom voice region (used for events/etc)
	Custom() bool
	SetCustom(cus bool) VoiceRegion
}
