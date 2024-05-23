// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/klaoslacerda/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
