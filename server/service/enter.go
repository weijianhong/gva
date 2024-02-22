package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ttt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/voln"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	VolnServiceGroup    voln.ServiceGroup
	TttServiceGroup     ttt.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
