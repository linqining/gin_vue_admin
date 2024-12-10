package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/service/company"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/job"
	"github.com/flipped-aurora/gin-vue-admin/server/service/job_certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user_certificate"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup           system.ServiceGroup
	ExampleServiceGroup          example.ServiceGroup
	UserServiceGroup             user.ServiceGroup
	CompanyServiceGroup          company.ServiceGroup
	JobServiceGroup              job.ServiceGroup
	CertificateServiceGroup      certificate.ServiceGroup
	Job_certificateServiceGroup  job_certificate.ServiceGroup
	User_certificateServiceGroup user_certificate.ServiceGroup
}
