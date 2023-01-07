package util

import "LudumDare/entities"

func DoesCollide(a entities.Entity, b entities.Entity) bool {
	aS := int(a.Scale)
	bS := int(b.Scale)
	cX := int(a.Pos[0]*a.Scale)+a.W*aS >= int(b.Pos[0]*b.Scale) && int(b.Pos[0]*b.Scale)+b.W*bS >= int(a.Pos[0]*a.Scale)
	cY := int(a.Pos[1]*a.Scale)+a.H*aS >= int(b.Pos[1]*b.Scale) && int(b.Pos[1]*b.Scale)+b.H*bS >= int(a.Pos[1]*a.Scale)
	return cX && cY
}
