// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/klaoslacerda/matterbridge/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
