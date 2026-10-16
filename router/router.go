package router

import (
	"database/sql"

	"Pelatihan_7/todo"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	h := &todo.Handler{DB: db}

	todo.RegisterRoutes(r.Group("/"), h)

	return r
}