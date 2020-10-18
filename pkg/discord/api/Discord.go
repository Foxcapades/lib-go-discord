package api

import (
	"github.com/foxcapades/lib-go-discord/pkg/discord"
	"github.com/foxcapades/lib-go-discord/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/pkg/dlib"
)

type DiscordAPI interface {
	Guilds() GuildAPI

	// GetVoiceRegions fetches and returns an array of voice region objects that
	// can be used when creating servers.
	GetVoiceRegions() ([]comm.VoiceRegion, error)
}

type GuildAPI interface {

}

type ChannelAPI interface {
	// Get a channel by ID.
	//
	//Returns a channel object.
	Get(id dlib.Snowflake) (discord.Channel, error)
}