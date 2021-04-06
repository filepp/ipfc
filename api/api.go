package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
	"ipfc/api/status"
	"ipfc/utils/xfile"
	"net/http"
	"os"
	"time"
)

func (m *Server) Add(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respondError(c, status.StatusInvalidParam, err)
		return
	}
	tempFile := m.repository.GetTempFilePath()
	err = c.SaveUploadedFile(file, tempFile)
	if err != nil {
		respondError(c, status.StatusFileOperationError, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	cid, err := m.storage.AddFile(ctx, tempFile)
	if err != nil {
		respondError(c, status.StatusFileOperationError, err)
		return
	}
	// 删除文件
	// os.Remove(tempFile)

	ret := struct {
		Cid string `json:"cid"`
	}{
		Cid: cid.String(),
	}
	respondSuccess(c, ret)
}

func (m *Server) Get(c *gin.Context) {
	strCid := c.Param("cid")
	//对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	filePath := m.repository.GetCacheFilePath(strCid)
	fileCid, err := cid.Decode(strCid)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid cid")
		return
	}

	// 如果文件不存在，则先从storage获取
	if !xfile.PathExists(filePath) {
		err := m.storage.RetrieveFile(context.TODO(), fileCid, filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.File(filePath)

	// 文件缓存1天, 这里需要优化,用定时任务清缓存
	time.AfterFunc(time.Hour*24, func() {
		os.Remove(filePath)
	})
}
