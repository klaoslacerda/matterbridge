// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/klaoslacerda/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
