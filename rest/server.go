package rest

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	logging "github.com/ipfs/go-log/v2"
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
}

func NewServer(addr string, storage types.Storage) *Server {
	engine := gin.Default()
	server := &Server{
		httpServer: &http.Server{Addr: addr, Handler: engine},
		storage:    storage,
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
