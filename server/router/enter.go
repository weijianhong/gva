package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ttt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/voln"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Voln    voln.RouterGroup
	Ttt     ttt.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
