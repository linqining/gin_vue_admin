package employee

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type EmployeeService struct{}

// CreateEmployee 创建员工信息记录
// Author [yourname](https://github.com/yourname)
func (employeeInfoService *EmployeeService) CreateEmployee(employeeInfo *employee.Employee) (err error) {
	err = global.GVA_DB.Create(employeeInfo).Error
	return err
}

// DeleteEmployee 删除员工信息记录
// Author [yourname](https://github.com/yourname)
func (employeeInfoService *EmployeeService) DeleteEmployee(ID string) (err error) {
	err = global.GVA_DB.Delete(&employee.Employee{}, "id = ?", ID).Error
	return err
}

// DeleteEmployeeByIds 批量删除员工信息记录
// Author [yourname](https://github.com/yourname)
func (employeeInfoService *EmployeeService) DeleteEmployeeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]employee.Employee{}, "id in ?", IDs).Error
	return err
}

// UpdateEmployee 更新员工信息记录
// Author [yourname](https://github.com/yourname)
func (employeeInfoService *EmployeeService) UpdateEmployee(employeeInfo employee.Employee) (err error) {
	err = global.GVA_DB.Model(&employee.Employee{}).Where("id = ?", employeeInfo.ID).Updates(&employeeInfo).Error
	return err
}

// GetEmployee 根据ID获取员工信息记录
// Author [yourname](https://github.com/yourname)
func (employeeInfoService *EmployeeService) GetEmployee(ID string) (employeeInfo employee.Employee, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&employeeInfo).Error
	return
}

// GetEmployeeInfoList 分页获取员工信息记录
// Author [yourname](https://github.com/yourname)
func (employeeInfoService *EmployeeService) GetEmployeeInfoList(info employeeReq.EmployeeSearch) (list []employee.Employee, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.Employee{})
	var employeeInfos []employee.Employee
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&employeeInfos).Error
	return employeeInfos, total, err
}
func (employeeInfoService *EmployeeService) GetEmployeePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// UpdateContractDigest 回填合同交易摘要
// Author [yourname](https://github.com/yourname)
func (employeeInfoService *EmployeeService) UpdateContractDigest(employerAddr, employeeAddr *string, digest *string) (err error) {
	if employerAddr == nil || employeeAddr == nil || digest == nil {
		return errors.New("invalid arguments")
	}
	cmpy := &company.Company{WalletAddress: employerAddr}
	err = global.GVA_DB.Model(&company.Company{}).Where("wallet_address = ?", employerAddr).FirstOrCreate(cmpy).Error
	if err != nil {
		return err
	}
	signer := &user.User{Address: employeeAddr}
	err = global.GVA_DB.Model(&user.User{}).Where("address = ?", employeeAddr).FirstOrCreate(signer).Error
	if err != nil {
		return err
	}
	userID := int(signer.ID)
	companyID := int(cmpy.ID)
	employeeInfo := &employee.Employee{
		UserId:    &userID,
		CompanyId: &companyID,
		Digest:    digest,
	}
	// 请在这里实现自己的业务逻辑
	err = global.GVA_DB.Model(&employee.Employee{}).
		Where("user_id=? and company_id=?", employeeInfo.UserId, employeeInfo.CompanyId).FirstOrCreate(employeeInfo).Error
	if err != nil {
		return err
	}
	return employeeInfoService.UpdateEmployee(*employeeInfo)
}
