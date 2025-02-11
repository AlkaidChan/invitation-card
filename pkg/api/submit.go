package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alkaidchen/invitation-card/pkg/entity/dto"
	"github.com/alkaidchen/invitation-card/pkg/log"
	"github.com/alkaidchen/invitation-card/pkg/model"
	"github.com/alkaidchen/invitation-card/pkg/utils/errmsg"
)

// Submit godoc
// @Summary submit form
// @Schemes http https
// @Description submit form
// @Accept application/json
// @Param req body dto.SubmitForm true "body"
// @Success 200 {object} errmsg.Response{data=nil}
// @Router /api/submit [post]
func Submit(c *gin.Context) {
	var data dto.SubmitForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, errmsg.ErrorInvalidParam)
		return
	}
	item := data.Transfer2Model()

	tx := model.BeginTx()
	// create item
	err = model.CreateItem(tx, item)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, errmsg.Error())
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}
