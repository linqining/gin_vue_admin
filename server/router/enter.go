package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/router/company"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/job"
	"github.com/flipped-aurora/gin-vue-admin/server/router/job_certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user_certificate"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System           system.RouterGroup
	Example          example.RouterGroup
	User             user.RouterGroup
	Company          company.RouterGroup
	Job              job.RouterGroup
	Certificate      certificate.RouterGroup
	Job_certificate  job_certificate.RouterGroup
	User_certificate user_certificate.RouterGroup
}
