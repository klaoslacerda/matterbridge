// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/klaoslacerda/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
