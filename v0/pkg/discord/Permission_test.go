package discord_test

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

var permissionList = []Permission{
	PermissionCreateInstantInvite,
	PermissionKickMembers,
	PermissionBanMembers,
	PermissionAdministrator,
	PermissionManageChannels,
	PermissionManageGuild,
	PermissionAddReactions,
	PermissionViewAuditLog,
	PermissionPrioritySpeaker,
	PermissionStream,
	PermissionViewChannel,
	PermissionSendMessages,
	PermissionSendTTSMessages,
	PermissionManageMessages,
	PermissionEmbedLinks,
	PermissionAttachFiles,
	PermissionReadMessageHistory,
	PermissionMentionEveryone,
	PermissionUseExternalEmojis,
	PermissionViewGuildInsights,
	PermissionConnect,
	PermissionSpeak,
	PermissionMuteMembers,
	PermissionDeafenMembers,
	PermissionMoveMembers,
	PermissionUseVAD,
	PermissionChangeNickname,
	PermissionManageNicknames,
	PermissionManageRoles,
	PermissionManageWebhooks,
	PermissionManageEmojis,
}

var permissionNames = []string{
	"PermissionCreateInstantInvite",
	"PermissionKickMembers",
	"PermissionBanMembers",
	"PermissionAdministrator",
	"PermissionManageChannels",
	"PermissionManageGuild",
	"PermissionAddReactions",
	"PermissionViewAuditLog",
	"PermissionPrioritySpeaker",
	"PermissionStream",
	"PermissionViewChannel",
	"PermissionSendMessages",
	"PermissionSendTTSMessages",
	"PermissionManageMessages",
	"PermissionEmbedLinks",
	"PermissionAttachFiles",
	"PermissionReadMessageHistory",
	"PermissionMentionEveryone",
	"PermissionUseExternalEmojis",
	"PermissionViewGuildInsights",
	"PermissionConnect",
	"PermissionSpeak",
	"PermissionMuteMembers",
	"PermissionDeafenMembers",
	"PermissionMoveMembers",
	"PermissionUseVAD",
	"PermissionChangeNickname",
	"PermissionManageNicknames",
	"PermissionManageRoles",
	"PermissionManageWebhooks",
	"PermissionManageEmojis",
}

func TestPermission_BufferSize(t *testing.T) {
	Convey("Permission.JSONSize", t, func() {
		sizes := []uint32{
			/* 1          */ 3,
			/* 2          */ 3,
			/* 4          */ 3,
			/* 8          */ 3,
			/* 16         */ 4,
			/* 32         */ 4,
			/* 64         */ 4,
			/* 128        */ 5,
			/* 256        */ 5,
			/* 512        */ 5,
			/* 1024       */ 6,
			/* 2048       */ 6,
			/* 4096       */ 6,
			/* 8192       */ 6,
			/* 16384      */ 7,
			/* 32768      */ 7,
			/* 65536      */ 7,
			/* 131072     */ 8,
			/* 262144     */ 8,
			/* 524288     */ 8,
			/* 1048576    */ 9,
			/* 2097152    */ 9,
			/* 4194304    */ 9,
			/* 8388608    */ 9,
			/* 16777216   */ 10,
			/* 33554432   */ 10,
			/* 67108864   */ 10,
			/* 134217728  */ 11,
			/* 268435456  */ 11,
			/* 536870912  */ 11,
			/* 1073741824 */ 12,
			/* nil        */ 4,
		}

		var name string
		var value *Permission
		for i := range sizes {
			Convey("returns " + strconv.Itoa(int(sizes[i])), func() {

				if i < len(permissionNames) {
					name = permissionNames[i]
				} else {
					name = "<nil>"
				}

				if i < len(permissionList) {
					value = &permissionList[i]
				} else {
					value = nil
				}

				Convey("when value is " + name, func() {
					So(value.JSONSize(), ShouldEqual, sizes[i])
				})
			})
		}
	})
}

