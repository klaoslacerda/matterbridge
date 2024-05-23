// +build !noapi

package bridgemap

import (
	"github.com/klaoslacerda/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
