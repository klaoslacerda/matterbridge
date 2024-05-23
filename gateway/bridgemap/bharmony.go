//go:build !noharmony
// +build !noharmony

package bridgemap

import (
	bharmony "github.com/klaoslacerda/matterbridge/bridge/harmony"
)

func init() {
	FullMap["harmony"] = bharmony.New
}
