// +build !novk

package bridgemap

import (
	bvk "github.com/klaoslacerda/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
