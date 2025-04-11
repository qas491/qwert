package api

import (
	"api/handler/request"
	__ "api/proto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var req request.UserRegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1000,
			"msg":  "参数不能为空",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	register := __.NewUserServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := register.Register(ctx, &__.RegisterRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Code != 200 {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": 1000,
			"msg":  "注册失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	})
}

func Login(c *gin.Context) {
	var req request.UserRegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1000,
			"msg":  "参数不能为空",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	login := __.NewUserServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := login.Login(ctx, &__.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Code != 200 {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": 1000,
			"msg":  "登录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
		"id":   r.Id,
	})
}
