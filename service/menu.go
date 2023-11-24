package service

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"net/http"
)

func GetMenuList(c *gin.Context) {
	Menus(c)
}

func Menus(c *gin.Context) {
	data := make([]*MenuReply, 0)
	allMenus := make([]*AllMenu, 0)
	tx := models.GetMenuList()
	err := tx.Find(&allMenus).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	data = allMenuToMenuReply(allMenus)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功",
		"data": data,
	})
}

func allMenuToMenuReply(allMenus []*AllMenu) []*MenuReply {
	data := make([]*MenuReply, 0)
	for _, v := range allMenus {
		if v.ParentId == 0 {
			data = append(data, &MenuReply{
				ID:       v.ID,
				Name:     v.Name,
				Icon:     v.Icon,
				Path:     v.Path,
				Sort:     v.Sort,
				SubMenus: getChildreMenu(v.ID, allMenus),
			})
		}
	}
	return data
}

func getChildreMenu(parentId uint, allMenus []*AllMenu) []*MenuReply {
	data := make([]*MenuReply, 0)
	for _, v := range allMenus {
		if v.ParentId == parentId {
			data = append(data, &MenuReply{
				ID:       v.ID,
				Name:     v.Name,
				Icon:     v.Icon,
				Path:     v.Path,
				Sort:     v.Sort,
				ParentId: v.ParentId,
				SubMenus: getChildreMenu(v.ID, allMenus),
			})
		}
	}
	return data
}
