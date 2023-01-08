package util

import (
	"LudumDare/entities"
)

func DoesCollide(a entities.Entity, b entities.Entity) bool {
	aS := int(a.Scale)
	bS := int(b.Scale)
	cX := int(a.Pos[0])+a.W*aS >= int(b.Pos[0]) && int(b.Pos[0])+b.W*bS >= int(a.Pos[0])
	cY := int(a.Pos[1])+a.H*aS >= int(b.Pos[1]) && int(b.Pos[1])+b.H*bS >= int(a.Pos[1])
	return cX && cY
}
