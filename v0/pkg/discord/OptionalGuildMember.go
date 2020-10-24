package discord

type OptionalGuildMember struct {
	validate bool
	member   GuildMember
}

func (o OptionalGuildMember) MarshalJSON() ([]byte, error) {
	return o.member.MarshalJSON()
}

func (o *OptionalGuildMember) UnmarshalJSON(bytes []byte) error {
	tmp := NewGuildMember(o.validate)
	if err := tmp.UnmarshalJSON(bytes); err != nil {
		return err
	}

	return nil
}

func (o OptionalGuildMember) IsSet() bool {
	return o.member != nil
}

func (o OptionalGuildMember) IsUnset() bool {
	return o.member == nil
}

func (o *OptionalGuildMember) Unset() OptionalField {
	o.member = nil
	return o
}

func (o OptionalGuildMember) Get() GuildMember {
	if o.member == nil {
		panic(ErrUnsetField)
	}

	return o.member
}

func (o *OptionalGuildMember) Set(member GuildMember) {
	if member == nil {
		panic(ErrSetNilOptionalVal)
	}

	if o.validate {
		if err := member.Validate(); err != nil {
			panic(err)
		}
	}

	o.member = member
}
