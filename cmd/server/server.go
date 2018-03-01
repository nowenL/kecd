package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nowenl/kecd/pkg"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:   "server",
	Usage:  "start the kecd server daemon",
	Action: server,
	Flags: []cli.Flag{
		cli.StringFlag{
			EnvVar: "KECD_SERVER_ADDR",
			Name:   "server-addr",
			Usage:  "server address",
			Value:  ":8876",
		},
		cli.StringFlag{
			EnvVar: "KECD_REGION",
			Name:   "region",
			Usage:  "region",
			Value:  "xq",
		},
		cli.StringFlag{
			EnvVar: "KECD_NAMESPACE",
			Name:   "namespace",
			Usage:  "namespace",
			Value:  "xq",
		},
		cli.StringFlag{
			EnvVar: "KECD_ACCOUNT",
			Name:   "account",
			Usage:  "account",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "KECD_PASSWORD",
			Name:   "password",
			Usage:  "password",
			Value:  "",
		},
	},
}

func server(c *cli.Context) (err error) {
	router := gin.New()

	logrus.Info("kecd starting...")

	config := pkg.Config{
		Username:  c.String("account"),
		Password:  c.String("password"),
		Region:    c.String("region"),
		Namespace: c.String("namespace"),
	}

	pkg.Init(config, router)

	logrus.Info("===Config===")
	logrus.Info("Account:", config.Username)
	logrus.Info("Password:", "******")
	logrus.Info("Namespace:", config.Namespace)
	logrus.Info("============")

	return http.ListenAndServe(
		c.String("server-addr"),
		router,
	)
}
