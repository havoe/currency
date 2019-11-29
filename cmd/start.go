package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	config   string
	port     string
	loglevel uint8
	cors     bool
	debug  string
	//StartCmd : set up restful api server
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start currency API server",
		Example: "currency server -c config/local.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	cors = true
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8082", "Tcp port server listening on")
	StartCmd.PersistentFlags().Uint8VarP(&loglevel, "loglevel", "l", 0, "Log level")
	StartCmd.PersistentFlags().StringVarP(&debug, "debug", "x", "true", "debug mode")

	if debug == "true"{
		config = "./config/local.yaml"
	}else {
		config = "./config/prod.yaml"
	}

}

func usage() {
	usageStr := `
  ______              
 |___  /              
    / / ___ _   _ ___ 
   / / / _ \ | | / __|
  / /_|  __/ |_| \__ \
 /_____\___|\__,_|___/
`
	fmt.Printf("%s\n", usageStr)
}

func setup() {


	//1.Set up log level
	zerolog.SetGlobalLevel(zerolog.Level(loglevel))
	//2.Set up configuration
	viper.SetConfigFile(config)
	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	//3.Set up run mode
	mode := viper.GetString("gin.mode")
	gin.SetMode(mode)
	//4.Set up database connection
	//dao.Setup()
	//5.Set up cache
	//cache.SetUp()
	//6.Set up ldap
	//ldap.Setup()
	//7.Set up permission handler
	//perm.SetUp(cluster)
	//8.DingTalk client setup
	//dingdingtalk.SetUp()
	//9.Initialize language
	//middleware.InitLang()
}

func run() error {
	engine := gin.Default()
	//router.SetUp(engine, cors)
	return engine.Run(":" + port)
}
