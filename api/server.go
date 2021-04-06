package api

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	logging "github.com/ipfs/go-log/v2"
	"ipfc/storage/repo"
	"ipfc/storage/types"
	"net/http"
	"sync"
	"time"
)

var log = logging.Logger("rest")

type Server struct {
	httpServer *http.Server
	wg         sync.WaitGroup
	storage    types.Storage
	repository *repo.Repository
}

func NewServer(addr string, storage types.Storage, repository *repo.Repository) *Server {
	engine := gin.Default()
	server := &Server{
		httpServer: &http.Server{Addr: addr, Handler: engine},
		storage:    storage,
		repository: repository,
	}
	server.setCors(engine)
	server.initRoute(engine)
	return server
}

func (m *Server) Run() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		err := m.httpServer.ListenAndServe()
		if err != nil {
			log.Infof("%v", err.Error())
		}
	}()
}

func (m *Server) Stop() {
	m.httpServer.Shutdown(context.Background())
	m.wg.Wait()
}

func (m *Server) initRoute(r gin.IRouter) {
	api := r.Group("/v1")
	api.PUT("/files", m.Add)      // upload a file
	api.GET("/files/:cid", m.Get) // download a file
}

func (m *Server) setCors(r gin.IRouter) {
	corsCfg := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Access-Control-Allow-Origin", "Accept",
			"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsCfg))
}
