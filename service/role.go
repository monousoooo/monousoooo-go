package service

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"net/http"
	"strconv"
)

func GetRoleList(c *gin.Context) {
	in := &GetRoleListRequest{NewQueryRequest()}
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var (
		count int64
		list  = make([]*GetRoleListReply, 0)
	)
	err = models.GetRoleList(in.Keyword).Count(&count).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功",
		"data": gin.H{
			"list":  list,
			"count": count,
		},
	})
}

func AddRole(c *gin.Context) {
	in := new(AddRoleRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var cnt int64
	err = models.DB.Model(new(models.SysRole)).Where("name = ?", in.Name).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "角色名已存在",
		})
		return
	}

	err = models.DB.Create(&models.SysRole{
		Name:    in.Name,
		Sort:    in.Sort,
		IsAdmin: in.IsAdmin,
		Remarks: in.Remarks,
	}).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func GetRoleDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	uId, err := strconv.Atoi(id)
	data := new(GetRoleDetailReply)
	sysRole, err := models.GetRoleDetail(uint(uId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取数据失败",
		})
		return
	}
	data.ID = sysRole.ID
	data.Name = sysRole.Name
	data.Sort = sysRole.Sort
	data.IsAdmin = sysRole.IsAdmin
	data.Remarks = sysRole.Remarks
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功",
		"data": data,
	})
}

func UpdateRole(c *gin.Context) {
	in := new(UpdateRoleRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var cnt int64
	err = models.DB.Model(new(models.SysRole)).Where("name = ? AND id != ?", in.Name, in.ID).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "角色名已存在",
		})
		return
	}
	err = models.DB.Model(new(models.SysRole)).Where("id = ?", in.ID).Updates(map[string]any{
		"name":     in.Name,
		"sort":     in.Sort,
		"is_admin": in.IsAdmin,
		"remarks":  in.Remarks,
	}).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	err := models.DB.Where("id = ?", id).Delete(new(models.SysRole)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
