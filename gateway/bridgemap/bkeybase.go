// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/klaoslacerda/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
