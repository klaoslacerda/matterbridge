// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/klaoslacerda/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
