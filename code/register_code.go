package code

import (
	"github.com/google/uuid"
	"time"
)

//
// @Author yfy2001
// @Date 2024/9/8 22 47
//

type RegisterCode struct {
	ID             string    `json:"id"`
	Code           string    `json:"code"`
	CreateTime     time.Time `json:"createTime"`
	ValidTime      int       `json:"validTime"`
	IsActive       bool      `json:"isActive"`
	ActiveTime     time.Time `json:"activeTime"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func NewRegisterCode(validTime int) *RegisterCode {
	id := uuid.NewString()
	return &RegisterCode{
		ID:             id,
		Code:           id,
		CreateTime:     time.Now(),
		ValidTime:      validTime,
		IsActive:       false,
		ActiveTime:     time.Time{},
		ExpirationTime: time.Time{},
	}
}

func (c RegisterCode) Active() {
	if c.IsActive {
		return
	}
	c.ActiveTime = time.Now()
	c.ExpirationTime = c.ActiveTime.Add(time.Duration(c.ValidTime) * time.Hour)
	c.IsActive = true
}

func (c RegisterCode) Verify() bool {
	currentTime := time.Now()
	valid := currentTime.Before(c.ExpirationTime)
	return valid
}
