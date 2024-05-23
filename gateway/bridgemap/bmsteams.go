// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/klaoslacerda/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
