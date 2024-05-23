// +build !noirc

package bridgemap

import (
	birc "github.com/klaoslacerda/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
