// +build !nomumble

package bridgemap

import (
	bmumble "github.com/klaoslacerda/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
