package discord

import "errors"

var (
	ErrBadFeature = errors.New("unrecognized guild feature value")
)

type GuildFeature string

const (
	// Guild has access to set an invite splash background
	GuildFeatureInviteSplash GuildFeature = "INVITE_SPLASH"

	// Guild has access to set 384kbps bitrate in voice (previously VIP voice servers)
	GuildFeatureVIPRegions GuildFeature = "VIP_REGIONS"

	// Guild has access to set a vanity URL
	GuildFeatureVanityURL GuildFeature = "VANITY_URL"

	// Guild is verified
	GuildFeatureVerified GuildFeature = "VERIFIED"

	// Guild is partnered
	GuildFeaturePartnered GuildFeature = "PARTNERED"

	// Guild can enable welcome screen and discovery, and receives community updates
	GuildFeatureCommunity GuildFeature = "COMMUNITY"

	// Guild has access to use commerce features (i.e. create store channels)
	GuildFeatureCommerce GuildFeature = "COMMERCE"

	// Guild has access to create news channels
	GuildFeatureNews GuildFeature = "NEWS"

	// Guild is lurkable and able to be discovered in the directory
	GuildFeatureDiscoverable GuildFeature = "DISCOVERABLE"

	// Guild is able to be featured in the directory
	GuildFeatureFeaturable GuildFeature = "FEATURABLE"

	// Guild has access to set an animated guild icon
	GuildFeatureAnimatedIcon GuildFeature = "ANIMATED_ICON"

	// Guild has access to set a guild banner image
	GuildFeatureBanner GuildFeature = "BANNER"

	// Guild has enabled the welcome screen)
	GuildFeatureWelcomeScreenEnabled GuildFeature = "WELCOME_SCREEN_ENABLED"
)

var features = []GuildFeature{
	GuildFeatureInviteSplash,
	GuildFeatureVIPRegions,
	GuildFeatureVanityURL,
	GuildFeatureVerified,
	GuildFeaturePartnered,
	GuildFeatureCommunity,
	GuildFeatureCommerce,
	GuildFeatureNews,
	GuildFeatureDiscoverable,
	GuildFeatureFeaturable,
	GuildFeatureAnimatedIcon,
	GuildFeatureBanner,
	GuildFeatureWelcomeScreenEnabled,
}

func (f GuildFeature) IsValid() bool {
	return nil == f.Validate()
}

func (f GuildFeature) Validate() error {
	for i := range features {
		if f == features[i] {
			return nil
		}
	}

	return ErrBadFeature
}

func (f GuildFeature) JSONSize() int {
	return uint32(len(f) + 2)
}
