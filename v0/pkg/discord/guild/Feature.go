package guild

import "errors"

type Feature string

var (
	ErrBadFeature = errors.New("unrecognized guild feature value")
)

const (
	// Guild has access to set an invite splash background
	FeatureInviteSplash Feature = "INVITE_SPLASH"

	// Guild has access to set 384kbps bitrate in voice (previously VIP voice servers)
	FeatureVIPRegions Feature = "VIP_REGIONS"

	// Guild has access to set a vanity URL
	FeatureVanityURL Feature = "VANITY_URL"

	// Guild is verified
	FeatureVerified Feature = "VERIFIED"

	// Guild is partnered
	FeaturePartnered Feature = "PARTNERED"

	// Guild can enable welcome screen and discovery, and receives community updates
	FeatureCommunity Feature = "COMMUNITY"

	// Guild has access to use commerce features (i.e. create store channels)
	FeatureCommerce Feature = "COMMERCE"

	// Guild has access to create news channels
	FeatureNews Feature = "NEWS"

	// Guild is lurkable and able to be discovered in the directory
	FeatureDiscoverable Feature = "DISCOVERABLE"

	// Guild is able to be featured in the directory
	FeatureFeaturable Feature = "FEATURABLE"

	// Guild has access to set an animated guild icon
	FeatureAnimatedIcon Feature = "ANIMATED_ICON"

	// Guild has access to set a guild banner image
	FeatureBanner Feature = "BANNER"

	// Guild has enabled the welcome screen)
	FeatureWelcomeScreenEnabled Feature = "WELCOME_SCREEN_ENABLED"
)

var features = []Feature{
	FeatureInviteSplash,
	FeatureVIPRegions,
	FeatureVanityURL,
	FeatureVerified,
	FeaturePartnered,
	FeatureCommunity,
	FeatureCommerce,
	FeatureNews,
	FeatureDiscoverable,
	FeatureFeaturable,
	FeatureAnimatedIcon,
	FeatureBanner,
	FeatureWelcomeScreenEnabled,
}

func (f Feature) Validate() error {
	for i := range features {
		if f == features[i] {
			return nil
		}
	}

	return ErrBadFeature
}
