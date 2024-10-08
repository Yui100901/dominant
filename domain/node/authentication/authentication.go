package authentication

import (
	"github.com/google/uuid"
	"time"
)

//
// @Author yfy2001
// @Date 2024/9/8 22 47
//

type Authentication struct {
	ID             string    `json:"id"`
	NodeId         string    `json:"nodeId"`
	Token          string    `json:"token"`
	CreateTime     time.Time `json:"createTime"`
	ValidTime      int       `json:"validTime"` //有效时间，最小单位小时
	IsActive       bool      `json:"isActive"`  //是否激活
	ActiveTime     time.Time `json:"activeTime"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func NewAuthentication(nodeId string, validTime int) *Authentication {
	id := uuid.NewString()
	return &Authentication{
		ID:             id,
		NodeId:         nodeId,
		CreateTime:     time.Now(),
		ValidTime:      validTime,
		IsActive:       false,
		ActiveTime:     time.Time{},
		ExpirationTime: time.Time{},
	}
}

// Active 激活认证
func (a *Authentication) Active() {
	if a.IsActive {
		return
	}
	a.ActiveTime = time.Now()
	a.ExpirationTime = a.ActiveTime.Add(time.Duration(a.ValidTime) * time.Hour)
	a.IsActive = true
}

// Verify 验证激活状态
func (a *Authentication) Verify() bool {
	if !a.IsActive {
		//未激活，则先激活
		a.Active()
	}
	currentTime := time.Now()
	isValid := currentTime.Before(a.ExpirationTime)
	return isValid
}

// AdjustValidTime 调整有效期
func (a *Authentication) AdjustValidTime(hours int) {
	a.ValidTime += hours
	a.ExpirationTime = a.ExpirationTime.Add(time.Duration(hours) * time.Hour)
}
