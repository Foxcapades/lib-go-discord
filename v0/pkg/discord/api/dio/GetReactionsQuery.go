package dio

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	"strings"
)

type GetReactionsQuery interface {
	APIQuery

	SetUserID(dlib.Snowflake) GetReactionsQuery
	UserID() dlib.Snowflake

	SetMessageID(dlib.Snowflake) GetReactionsQuery
	MessageID() dlib.Snowflake

	SetEmoji(discord.Emoji) GetReactionsQuery
	GetEmoji() discord.Emoji

	// Before returns the current value of this query's `before` param.
	//
	// The `before` param limits the results to users before this user ID.
	//
	// If this method is called on a param that is unset, this method will panic.
	// Use BeforeIsSet to check if the param is present before use.
	Before() dlib.Snowflake

	// BeforeIsSet returns whether this query's `before` param is currently
	// present.
	BeforeIsSet() bool

	// SetBefore overwrites the current value of this query's `before` param.
	SetBefore(dlib.Snowflake) GetReactionsQuery

	// UnsetBefore removes this query's `before` param.
	UnsetBefore() GetReactionsQuery

	// After returns the current value of this query's `after` param.
	//
	// The `after` param limits the results to users after this user ID.
	//
	// If this method is called on a param that is unset, this method will panic.
	// Use AfterIsSet to check if the param is present before use.
	After() dlib.Snowflake

	// AfterIsSet returns whether this query's `after` param is currently present.
	AfterIsSet() bool

	// SetAfter overwrites the current value of this query's `after` param.
	SetAfter(dlib.Snowflake) GetReactionsQuery

	// UnsetAfter removes this query's `after` param.
	UnsetAfter() GetReactionsQuery

	// Limit returns the current value of this query's `limit` param.
	//
	// The `limit` param controls the number of results returned by the request.
	//
	// If this method is called on a param that is unset, this method will panic.
	// Use LimitIsSet to check if the param is present before use.
	Limit() RecordLimit

	// LimitIsSet returns whether this query's `limit` param is currently present.
	LimitIsSet() bool

	// SetLimit overwrites the current value of this query's `limit` param.
	SetLimit(RecordLimit) GetReactionsQuery

	// UnsetLimit removes this query's `limit` param.
	UnsetLimit() GetReactionsQuery
}

func NewGetReactionsQuery(msgId, userId dlib.Snowflake, emoji discord.Emoji) GetReactionsQuery {
	return &getReactionsQuery{
		userId: userId,
		msgId:  msgId,
		emoji:  emoji,
	}
}

type getReactionsQuery struct {
	userId dlib.Snowflake
	msgId  dlib.Snowflake
	before dlib.SnowflakeOptionalField
	after  dlib.SnowflakeOptionalField
	limit  OptionalRecordLimit
	emoji  discord.Emoji
}

func (g *getReactionsQuery) ToQuery() string {
	out := new(strings.Builder)
	out.Grow(32)
	out.Reset()

	if g.before.IsSet() {
		out.WriteByte(queryStart)
		out.WriteString(queryBefore)
		out.WriteByte(queryValSep)
		out.WriteString(g.before.Get().String())
	}

	if g.after.IsSet() {
		if out.Len() > 0 {
			out.WriteByte(queryJoin)
		} else {
			out.WriteByte(queryStart)
		}

		out.WriteString(queryAfter)
		out.WriteByte(queryValSep)
		out.WriteString(g.after.Get().String())
	}

	if g.limit.IsSet() {
		if out.Len() > 0 {
			out.WriteByte(queryJoin)
		} else {
			out.WriteByte(queryStart)
		}

		out.WriteString(queryLimit)
		out.WriteByte(queryValSep)
		out.WriteString(g.limit.Get().String())
	}

	return out.String()
}

func (g *getReactionsQuery) SetUserID(id dlib.Snowflake) GetReactionsQuery {
	g.userId = id
	return g
}

func (g *getReactionsQuery) UserID() dlib.Snowflake {
	return g.userId
}

func (g *getReactionsQuery) SetMessageID(id dlib.Snowflake) GetReactionsQuery {
	g.msgId = id
	return g
}

func (g *getReactionsQuery) MessageID() dlib.Snowflake {
	return g.msgId
}

func (g *getReactionsQuery) SetEmoji(emoji discord.Emoji) GetReactionsQuery {
	g.emoji = emoji
	return g
}

func (g *getReactionsQuery) GetEmoji() discord.Emoji {
	return g.emoji
}

func (g *getReactionsQuery) Before() dlib.Snowflake {
	return g.before.Get()
}

func (g *getReactionsQuery) BeforeIsSet() bool {
	return g.before.IsSet()
}

func (g *getReactionsQuery) SetBefore(id dlib.Snowflake) GetReactionsQuery {
	g.before.Set(id)
	return g
}

func (g *getReactionsQuery) UnsetBefore() GetReactionsQuery {
	g.before.Unset()
	return g
}

func (g *getReactionsQuery) After() dlib.Snowflake {
	return g.after.Get()
}

func (g *getReactionsQuery) AfterIsSet() bool {
	return g.after.IsSet()
}

func (g *getReactionsQuery) SetAfter(id dlib.Snowflake) GetReactionsQuery {
	g.after.Set(id)
	return g
}

func (g *getReactionsQuery) UnsetAfter() GetReactionsQuery {
	g.after.Unset()
	return g
}

func (g *getReactionsQuery) Limit() RecordLimit {
	return g.limit.Get()
}

func (g *getReactionsQuery) LimitIsSet() bool {
	return g.limit.IsSet()
}

func (g *getReactionsQuery) SetLimit(limit RecordLimit) GetReactionsQuery {
	g.limit.Set(limit)
	return g
}

func (g *getReactionsQuery) UnsetLimit() GetReactionsQuery {
	g.limit.Unset()
	return g
}
