// +build !nosteam

package bridgemap

import (
	bsteam "github.com/klaoslacerda/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
