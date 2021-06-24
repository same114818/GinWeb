package knowledgeRepository

import (
	"GinWeb/models"

	"xorm.io/xorm"
)

func Create(engine *xorm.Engine) error {
	knowledgeMain := new(models.KnowledgeMain)
	knowledgeMain.Name = "test"
	knowledgeMain.CreateBy = 1
	knowledgeMain.UpdateBy = 1
	knowledgeMain.IsValid = 1
	_, err := engine.Insert(knowledgeMain)
	if err != nil {
		return err
	}
	return nil
}
