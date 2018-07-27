package api

import (
	"net/http"
	"time"

	"github.com/fy026/weather/gateway/pkg/e"
	"github.com/fy026/weather/gateway/pkg/setting"
	"github.com/fy026/weather/pkg/client"
	"github.com/fy026/weather/pkg/registry"
	"github.com/fy026/weather/pkg/registry/etcd"
	"github.com/fy026/weather/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

var wcClient *client.Client
var wcServiceName = "ts"

func init() {
	reg, err := etcd.NewEtcdRegisty(registry.WithAddress(setting.RegisterUrl))
	if err != nil {
		return
	}

	wcClient = client.NewClient(
		client.WithServiceName(wcServiceName),
		client.WithTimeout(time.Second*10),
		client.WithRegistry(reg),
	)

}

func GetTest(c *gin.Context) {
	name := c.Query("name")

	grpcConn, err := wcClient.GetConn()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_CONN_SERVICE_FAIL,
			"msg":  e.GetMsg(e.ERROR_CONN_SERVICE_FAIL),
			"data": "",
		})
		return
	}

	client := pb.NewGreeterClient(grpcConn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_CONN_SERVICE_FAIL,
			"msg":  e.GetMsg(e.ERROR_CONN_SERVICE_FAIL),
			"data": "",
		})
		return
	}

	//返回信息
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": resp.Message,
	})

}
