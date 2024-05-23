// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/klaoslacerda/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
