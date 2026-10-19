package router

import (
     "gorm.io/gorm"

	"Pelatihan_7/todo"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	todo.RegisterRoutes(v1, &todo.Handler{DB: db})
	return r
}