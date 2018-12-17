package model

import (
	"encoding/json"
)

// Building represents a building on the map
type Building struct {
	X  int     `json:"x"`
	Y  int     `json:"y"`
	Dx int     `json:"dx"`
	Dy int     `json:"dy"`
	Dz float64 `json:"dz"`
}

// ParseBuilding decodes a JSON string and unmarshals it into the Building struct
func ParseBuilding(s string) (*Building, error) {
	var b Building
	if err := json.Unmarshal([]byte(s), &b); err != nil {
		return nil, err
	}
	return &b, nil
}

// IsValid checks if a polygon is valid, and returns an error string in case it's invalid
func (b *Building) IsValid() (bool, string) {
	if b.X < 0 {
		return false, "Building x less than 0"
	}
	if b.Y < 0 {
		return false, "Building y less than 0"
	}
	if b.Dx < 1 {
		return false, "Building dx less than 1"
	}
	if b.Dy < 1 {
		return false, "Building dy less than 1"
	}
	if b.Dz < 1 {
		return false, "Building dz less than 1"
	}
	if b.X+b.Dx > 99 {
		return false, "Building outside bounds (x + dx > 99)"
	}
	if b.Y+b.Dy > 99 {
		return false, "Building outside bounds (y + dy > 99)"
	}
	if b.Dz > 9 {
		return false, "Building too tall (dz > 9)"
	}
	return true, ""
}