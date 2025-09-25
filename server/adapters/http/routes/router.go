package routes

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rubenalves-dev/taskuik/adapters"
)

type Router struct {
	Engine *gin.Engine
	Port   string
	appDB  adapters.Database
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		Engine: engine,
	}
}

func (r *Router) AddDbClient(dbClient adapters.Database) {
	if r.appDB != nil {
		log.Printf("Database client already set: %v", r.appDB)
		return
	}
	r.appDB = dbClient
}

func (r *Router) RegisterRoutes() {
	r.RegisterTaskRoutes()
}

func (r *Router) Run(port int) error {
	return r.Engine.Run(":" + strconv.Itoa(port))
}
