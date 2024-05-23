// +build !notelegram

package bridgemap

import (
	btelegram "github.com/klaoslacerda/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
