
package achievement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/achievement"
    achievementReq "github.com/flipped-aurora/gin-vue-admin/server/model/achievement/request"
)

type AchievementService struct {}
// CreateAchievement 创建成就记录
// Author [yourname](https://github.com/yourname)
func (achievement_infoService *AchievementService) CreateAchievement(achievement_info *achievement.Achievement) (err error) {
	err = global.GVA_DB.Create(achievement_info).Error
	return err
}

// DeleteAchievement 删除成就记录
// Author [yourname](https://github.com/yourname)
func (achievement_infoService *AchievementService)DeleteAchievement(ID string) (err error) {
	err = global.GVA_DB.Delete(&achievement.Achievement{},"id = ?",ID).Error
	return err
}

// DeleteAchievementByIds 批量删除成就记录
// Author [yourname](https://github.com/yourname)
func (achievement_infoService *AchievementService)DeleteAchievementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]achievement.Achievement{},"id in ?",IDs).Error
	return err
}

// UpdateAchievement 更新成就记录
// Author [yourname](https://github.com/yourname)
func (achievement_infoService *AchievementService)UpdateAchievement(achievement_info achievement.Achievement) (err error) {
	err = global.GVA_DB.Model(&achievement.Achievement{}).Where("id = ?",achievement_info.ID).Updates(&achievement_info).Error
	return err
}

// GetAchievement 根据ID获取成就记录
// Author [yourname](https://github.com/yourname)
func (achievement_infoService *AchievementService)GetAchievement(ID string) (achievement_info achievement.Achievement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&achievement_info).Error
	return
}
// GetAchievementInfoList 分页获取成就记录
// Author [yourname](https://github.com/yourname)
func (achievement_infoService *AchievementService)GetAchievementInfoList(info achievementReq.AchievementSearch) (list []achievement.Achievement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&achievement.Achievement{})
    var achievement_infos []achievement.Achievement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?","%"+*info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&achievement_infos).Error
	return  achievement_infos, total, err
}
func (achievement_infoService *AchievementService)GetAchievementPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
