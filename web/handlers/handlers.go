package handlers

import (
	"net/http"

	"github.com/c-fs/cfs/client"
	pb "github.com/c-fs/cfs/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type WriteRequest struct {
	pb.WriteRequest
	Address string `json:"address,omitempty" binding:"required"`
}

func Write(c *gin.Context) {
	var req WriteRequest
	c.Bind(&req)
	cfsClient, err := client.New(req.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer cfsClient.Close()

	n, err := cfsClient.Write(context.TODO(), req.Name, req.Offset, req.Data, req.Append)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pb.WriteReply{BytesWritten: n})
}

type ReadRequest struct {
	pb.ReadRequest
	Address string `json:"address,omitempty" binding:"required"`
}

func Read(c *gin.Context) {
	var req ReadRequest
	c.Bind(&req)
	cfsClient, err := client.New(req.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer cfsClient.Close()

	n, data, err := cfsClient.Read(Context.TODO(), req.Name, req.Offset, req.Length, req.ExpChecksum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pb.ReadReply{BytesRead: n, Data: data})
}
