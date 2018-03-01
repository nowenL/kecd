package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/qbox/ke-base/sdk"
	"github.com/sirupsen/logrus"
)

type DeployMgr struct {
	client *sdk.ClientSet
	Config
}

type Config struct {
	Username  string
	Password  string
	Region    string
	Namespace string
}

func Init(config Config, router gin.IRouter) error {
	client := sdk.New(sdk.Config{
		Username:  "chenkaijun@qiniu.com",
		Password:  "123456",
		Host:      "https://keapi.qiniu.com",
		Transport: nil,
	})

	deployMgr := DeployMgr{
		client: client,
		Config: config,
	}

	router.POST("/deploy", deployMgr.Deploy)
	return nil
}

func (p *DeployMgr) Deploy(c *gin.Context) {
	logrus.Info("Trigger")
	logrus.Info("Start deploy...")

	logrus.Info("Deploy done.")
}
