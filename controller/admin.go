package controller

import (
	"mhb_heart_day_quiz_2023/config"
	models "mhb_heart_day_quiz_2023/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoleSelect(c *gin.Context) {
	c.HTML(http.StatusOK, "userRoleSelector.html", nil)
}

func AdminHome(c *gin.Context) {
	c.HTML(http.StatusOK, "adminHome.html", nil)
}

func DeptManagement(c *gin.Context) {
	var department []models.Department
	result := config.DB.Find(&department)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.HTML(http.StatusOK, "department.html", gin.H{"departments": department})
}

func AddDepartment(c *gin.Context) {
	var dept models.Department
	if err := c.ShouldBind(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		return
	}

	//checking the department already exists or not
	checkDept := config.DB.Where("name = ?", dept.Name).First(&models.Department{})
	if checkDept.RowsAffected != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Department name already exists", "status": "false"})
		return
	}
	result := config.DB.Create(&dept)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/department?success=true")
}

func DeleteDept(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Where("id = ?", id).Delete(&models.Department{})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/department?delete=true")
}

func UserManagerment(c *gin.Context) {
	var users []models.User
	result := config.DB.Preload("Dept").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.HTML(http.StatusOK, "userManagement.html", gin.H{"users": users})
}
