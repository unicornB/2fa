package service

import (
	"gongniu/model"
	"gongniu/serializer"

	"github.com/gin-gonic/gin"
)

type BizDataService struct {
	BizKey  string `form:"biz_key" json:"biz_key"`
	BizData string `form:"biz_data" json:"biz_data"`
}

func (service *BizDataService) GetBizData(c *gin.Context) serializer.Response {
	var bizData model.BizData
	result := model.DB.Where("biz_key = ?", c.Param("bizKey")).First(&bizData)
	if result.Error != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "没有找到数据",
			Error: result.Error.Error(),
		}
	}
	return serializer.Success("成功", bizData.BizData)
}

func (service *BizDataService) UpdateBizData(c *gin.Context) serializer.Response {
	var bizDataModel model.BizData
	result := model.DB.Where("biz_key =?", service.BizKey).First(&bizDataModel)
	if result.Error != nil {
		return serializer.Response{
			Code:  0,
			Msg:   "没有找到数据",
			Error: result.Error.Error(),
		}
	}
	bizDataModel.BizData = service.BizData
	result = model.DB.Save(&bizDataModel)
	if result.Error != nil {
		return serializer.Response{
			Code:  0,
			Msg:   "更新失败",
			Error: result.Error.Error(),
		}
	}
	return serializer.Success("更新成功", nil)
}
