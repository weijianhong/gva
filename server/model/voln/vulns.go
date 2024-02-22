// 自动生成模板Vulns
package voln

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// vulns表 结构体  Vulns
type Vulns struct {
	global.GVA_MODEL
	TfvdId          string  `json:"tfvdId" form:"tfvdId" gorm:"column:tfvd_id;comment:编号;size:100;"`                                                                                              //编号
	CveId           string  `json:"cveId" form:"cveId" gorm:"column:cve_id;comment:CVE编号;size:100;"`                                                                                              //CVE编号
	CnvdId          string  `json:"cnvdId" form:"cnvdId" gorm:"column:cnvd_id;comment:cnv_id;size:100;"`                                                                                          //cnv_id
	CnnvdId         string  `json:"cnnvdId" form:"cnnvdId" gorm:"column:cnnvd_id;comment:CNNVD编号;size:100;"`                                                                                      //CNNVD编号
	VulnName        string  `json:"vulnName" form:"vulnName" gorm:"column:vuln_name;comment:漏洞名称;size:200;"`                                                                                      //漏洞名称
	Cvss            string  `json:"cvss" form:"cvss" gorm:"column:cvss;comment:CVSS值;size:100;"`                                                                                                  //CVSS值
	Verified        *int    `json:"verified" form:"verified" gorm:"column:verified;comment:是否研判;"`                                                                                                //是否研判
	Summary         string  `json:"summary" form:"summary" gorm:"column:summary;comment:简介;"`                                                                                                     //简介
	VulnCategory    string  `json:"vulnCategory" form:"vulnCategory" gorm:"column:vuln_category;comment:漏洞类型 id逗号分隔;size:100;"`                                                                   //漏洞类型 id逗号分隔
	VulnPatch       string  `json:"vulnPatch" form:"vulnPatch" gorm:"column:vuln_patch;comment:修复建议;"`                                                                                            //修复建议
	RiskLevel       *string `json:"riskLevel" form:"riskLevel" gorm:"column:risk_level;comment:风险等级：0-待定 1-低危 2-中危 3-高危- 4-超危;"`                                                                  //风险等级：0-待定 1-低危 2-中危 3-高危- 4-超危
	AlarmType       *int    `json:"alarmType" form:"alarmType" gorm:"column:alarm_type;comment:报警类别 alarmtypeset表的alarmType;"`                                                                    //报警类别 alarmtypeset表的alarmType
	Repairability   *int    `json:"repairability" form:"repairability" gorm:"column:repairability;comment:可修复性：可修复性，对应t_vpt_score_set表中type=2 ;补丁修复-2001,  其他方式修复-2002  不可修复-2003;"`              //可修复性：可修复性，对应t_vpt_score_set表中type=2 ; 补丁修复-2001,  其他方式修复-2002  不可修复-2003
	DifficultyLevel *int    `json:"difficultyLevel" form:"difficultyLevel" gorm:"column:difficulty_level;comment:难易度：漏洞利用难易程度，对应t_vpt_score_set表中type=3 ;简单-3001   中等-3002； 困难-3003 ，  未知-3004;"` //难易度：漏洞利用难易程度，对应t_vpt_score_set表中type=3 ;简单-3001   中等-3002； 困难-3003 ，  未知-3004
	Creator         string  `json:"creator" form:"creator" gorm:"column:creator;comment:创建人;size:50;"`                                                                                            //创建人
	CreatorId       *int    `json:"creatorId" form:"creatorId" gorm:"column:creator_id;comment:创建人id;size:10;"`                                                                                   //创建人id
	Updater         string  `json:"updater" form:"updater" gorm:"column:updater;comment:更新人;size:50;"`                                                                                            //更新人
	UpdaterId       *int    `json:"updaterId" form:"updaterId" gorm:"column:updater_id;comment:更新人id;size:10;"`                                                                                   //更新人id
	AffectedEntity  string  `json:"affectedEntity" form:"affectedEntity" gorm:"column:affected_entity;comment:受影响实体 逗号隔开;size:255;"`                                                              //受影响实体 逗号隔开
}

// TableName vulns表 Vulns自定义表名 vulns
func (Vulns) TableName() string {
	return "vulns"
}
