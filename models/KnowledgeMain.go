package models

import "time"

type KnowledgeMain struct {
	Id       int64
	Name     string `xorm:"nvarchar(255) comment('知识库名称')"`
	IsValid  int64
	CreateBy int64
	CreateOn time.Time `xorm:"created"`
	UpdateBy int64
	UpdateOn time.Time `xorm:"updated"`
}
