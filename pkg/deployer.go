package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qbox/ke-base/sdk"
	"github.com/qbox/ke-base/sdk/proto"
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

	router.POST("/deploy/:repo", deployMgr.Deploy)
	return nil
}

type WebhookData struct {
	EventID   string `json:"eventID"`
	Namespace string `json:"namespace"`
	Repo      string `json:"repoName"`
	Tag       string `json:"tag"`
	Digest    string `json:"digest"`
}

const (
	AppName = "demoapp"
	SvcName = "demosvc"
	Image   = "reg.qiniu.com/g57g/%s:%s"
)

func (p *DeployMgr) Deploy(c *gin.Context) {
	logrus.Info("Start deploy...")

	data := WebhookData{}
	repo := c.Param("repo")
	err := c.BindJSON(&data)
	if err != nil {
		logrus.Info("Deploy failed: ", err.Error())
		return
	}

	imageToDeploy := fmt.Sprintf(Image, repo, data.Tag)

	logrus.Info("===Deploy===")
	logrus.Info("Repo:", repo)
	logrus.Info("Tag:", data.Tag)
	logrus.Info("Digest:", data.Digest)
	logrus.Info("Image:", imageToDeploy)
	logrus.Info("============")

	old, err := p.client.MicroService(p.Region).GetService(nil, p.Namespace, AppName, SvcName)
	if err != nil {
		logrus.Info("Deploy failed: ", err.Error())
		return
	}

	newContainerSpec := old.Containers[0]
	newContainerSpec.Image = imageToDeploy

	_, err = p.client.MicroService(p.Region).UpgradeService(nil, p.Namespace, AppName, SvcName, proto.MicroServiceUpgradeArgs{
		ResourceSpec: old.ResourceSpec,
		Containers: []proto.Container{
			newContainerSpec,
		},
	})

	if err != nil {
		logrus.Info("Deploy failed: ", err.Error())
		return
	}

	logrus.Info("Deploy done.")
}
