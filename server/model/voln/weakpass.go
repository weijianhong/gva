// 自动生成模板Weakpass
package voln

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// weakpass表 结构体  Weakpass
type Weakpass struct {
      global.GVA_MODEL
      Username  string `json:"username" form:"username" gorm:"column:username;comment:用户名;size:50;"`  //用户名 
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;size:150;"`  //密码 
      FingerprintId  *int `json:"fingerprintId" form:"fingerprintId" gorm:"column:fingerprint_id;comment:指纹ID;"`  //指纹ID 
      FingerprintName  string `json:"fingerprintName" form:"fingerprintName" gorm:"column:fingerprint_name;comment:指纹名称;size:200;"`  //指纹名称 
      Port  string `json:"port" form:"port" gorm:"column:port;comment:常用端口;size:20;"`  //常用端口 
      Protocol  string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:协议;size:20;"`  //协议 
      CorpId  *int `json:"corpId" form:"corpId" gorm:"column:corp_id;comment:厂商/公司/组织;"`  //厂商/公司/组织 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:状态 1有效，0无效，默认1有效;"`  //状态 1有效，0无效，默认1有效 
      Creator  string `json:"creator" form:"creator" gorm:"column:creator;comment:;size:20;"`  //creator字段 
      CreatorId  *int `json:"creatorId" form:"creatorId" gorm:"column:creator_id;comment:创建人id;size:10;"`  //创建人id 
      Updater  string `json:"updater" form:"updater" gorm:"column:updater;comment:更新人;size:50;"`  //更新人 
      UpdaterId  *int `json:"updaterId" form:"updaterId" gorm:"column:updater_id;comment:更新人id;size:10;"`  //更新人id 
}


// TableName weakpass表 Weakpass自定义表名 weakpass
func (Weakpass) TableName() string {
  return "weakpass"
}

