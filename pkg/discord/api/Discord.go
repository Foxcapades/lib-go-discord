package api

import "github.com/foxcapades/lib-go-discord/pkg/discord/comm"

type DiscordAPI interface {
	Guilds() GuildAPI

	// GetVoiceRegions fetches and returns an array of voice region objects that
	// can be used when creating servers.
	GetVoiceRegions() ([]comm.VoiceRegion, error)
}

type GuildAPI interface {

}
