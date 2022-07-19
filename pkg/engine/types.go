package engine

import (
	"github.com/ipfs/go-cid"
)

type MetaInclusion struct {
	ID             cid.Cid `json:"ID"`
	Provider       string  `json:"Provider"`
	InPando        bool    `json:"InPando"`
	InSnapShot     bool    `json:"InSnapShot"`
	SnapShotID     cid.Cid `json:"SnapShotID"`
	SnapShotHeight uint64  `json:"SnapShotHeight"`
	Context        []byte  `json:"Context"`
	TranscationID  int     `json:"TranscationID"`
}
