
package job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/job"
    jobReq "github.com/flipped-aurora/gin-vue-admin/server/model/job/request"
)

type JobService struct {}
// CreateJob 创建职位记录
// Author [yourname](https://github.com/yourname)
func (job_infoService *JobService) CreateJob(job_info *job.Job) (err error) {
	err = global.GVA_DB.Create(job_info).Error
	return err
}

// DeleteJob 删除职位记录
// Author [yourname](https://github.com/yourname)
func (job_infoService *JobService)DeleteJob(id string) (err error) {
	err = global.GVA_DB.Delete(&job.Job{},"id = ?",id).Error
	return err
}

// DeleteJobByIds 批量删除职位记录
// Author [yourname](https://github.com/yourname)
func (job_infoService *JobService)DeleteJobByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]job.Job{},"id in ?",ids).Error
	return err
}

// UpdateJob 更新职位记录
// Author [yourname](https://github.com/yourname)
func (job_infoService *JobService)UpdateJob(job_info job.Job) (err error) {
	err = global.GVA_DB.Model(&job.Job{}).Where("id = ?",job_info.Id).Updates(&job_info).Error
	return err
}

// GetJob 根据id获取职位记录
// Author [yourname](https://github.com/yourname)
func (job_infoService *JobService)GetJob(id string) (job_info job.Job, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&job_info).Error
	return
}
// GetJobInfoList 分页获取职位记录
// Author [yourname](https://github.com/yourname)
func (job_infoService *JobService)GetJobInfoList(info jobReq.JobSearch) (list []job.Job, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&job.Job{})
    var job_infos []job.Job
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&job_infos).Error
	return  job_infos, total, err
}
func (job_infoService *JobService)GetJobPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
