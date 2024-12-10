
// 自动生成模板Job
package job
import (
)

// 职位 结构体  Job
type Job struct {
    CompanyId  *int `json:"companyId" form:"companyId" gorm:"column:company_id;comment:;size:10;"`  //公司id 
    Description  *int `json:"description" form:"description" gorm:"column:description;comment:;size:10;"`  //职位描述 
    SalarmBottom  *int `json:"salarmBottom" form:"salarmBottom" gorm:"column:salarm_bottom;comment:;size:10;"`  //薪资下限 
    SalaryCeil  *int `json:"salaryCeil" form:"salaryCeil" gorm:"column:salary_ceil;comment:;size:10;"`  //薪资上限 
    BlobId  *string `json:"blobId" form:"blobId" gorm:"column:blob_id;comment:;size:100;"`  //WalrusBlobId 
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`  //id 
}


// TableName 职位 Job自定义表名 job
func (Job) TableName() string {
    return "job"
}





