//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type BotMovementLedger struct {
	ID                int64 `sql:"primary_key"`
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	BotID             int64
	UserID            int64
	TimeActionStarted time.Time
	NewX              float64
	NewY              float64
}
