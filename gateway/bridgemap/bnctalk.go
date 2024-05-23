// +build !nonctalk

package bridgemap

import (
	btalk "github.com/klaoslacerda/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
