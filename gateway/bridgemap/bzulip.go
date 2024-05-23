// +build !nozulip

package bridgemap

import (
	bzulip "github.com/klaoslacerda/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
