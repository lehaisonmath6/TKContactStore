package models

import (
	"github.com/OpenStars/GoEndpointManager"
)

type ThriftService struct {
	ServiceID  string
	EndpoinMgr GoEndpointManager.EnpointManagerIf
	Protocol   string
}
