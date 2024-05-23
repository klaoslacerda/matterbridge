// +build !nogitter

package bridgemap

import (
	bgitter "github.com/klaoslacerda/matterbridge/bridge/gitter"
)

func init() {
	FullMap["gitter"] = bgitter.New
}
