package api

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type DiscordAPI interface {
	Guilds() GuildAPI

	Channels() ChannelAPI

	// GetVoiceRegions fetches and returns an array of voice region objects that
	// can be used when creating servers.
	GetVoiceRegions() ([]discord.VoiceRegion, error)
}

