package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo19/models"
)

type BankController struct{
	BaseController
}

func (c BankController) BankIndex(ctx *gin.Context) {
	// 开启一个事务
	tx := models.DB.Begin()
	// 回滚操作
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.Fail(ctx)
		}
	}()
	// 用户 1 扣 100 元
	u1 := models.Bank{Id: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance - 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
		c.Fail(ctx)
	}
	// 用户 2 扣 200 元
	u2 := models.Bank{Id: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance + 100
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		c.Fail(ctx)
	}
	tx.Commit()
	c.Success(ctx)
}