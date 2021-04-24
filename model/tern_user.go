package model

import (
	"fmt"
	"time"
)

// TernUser struct
type TernUser struct {
	ID                       int32
	UserID                   int32
	//GeneralTernUserStatusID  int32
	GeneralTernUserStatusID  string
	TernCode                 string
	Deactivated_at           *time.Time
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                *time.Time
}

func init() {
	fmt.Println("okok")
}