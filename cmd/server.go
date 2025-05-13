package cmd

import (
	"errors"
	"fmt"
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/micro_contrib/postgresql"
	"github.com/CXeon/nav_demo/config"
	"github.com/CXeon/nav_demo/internal/controller/http"
	"github.com/CXeon/nav_demo/internal/dao"
	"github.com/CXeon/nav_demo/internal/model"
	"github.com/CXeon/nav_demo/internal/service"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	httpOri "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Init 初始化并启动服务
func Init() {

	//加载配置文件
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	//创建日志
	logFileName := fmt.Sprintf("./%s.log", conf.Application.ServiceName)
	var logger *log.Logger
	if strings.ToUpper(conf.Application.Env) == "Pro" {
		logger = log.NewLogger(logFileName, zapcore.InfoLevel)
	} else {
		logger = log.NewLogger(logFileName)
	}
	defer logger.Sync()

	logger.Info("server init...")

	//创建pgsql客户端
	pgsqlCloud, err := postgresql.Init(&postgresql.PostgresqlConfig{
		Host:     conf.Postgresql.Host,
		Port:     conf.Postgresql.Port,
		User:     conf.Postgresql.User,
		DbName:   conf.Postgresql.DbName,
		Password: conf.Postgresql.Password,
	})
	if err != nil {
		logger.Fatalf("connect postgresql err: %s", err.Error())
	}

	//TODO 自动创建表（默认应当禁用）
	err = model.AutoMigrate(pgsqlCloud.Db)
	if err != nil {
		logger.Fatalf("connect postgresql err: %s", err.Error())
	}

	//根据pgsql客户端创建Dao
	postgresqlCloudDao := dao.NewPostgresqlCloudDao(pgsqlCloud)

	//创建service
	err = service.NewService(conf, logger, postgresqlCloudDao)
	if err != nil {
		logger.Fatalf("init service err: %s", err.Error())
	}

	//启动web服务
	var eg errgroup.Group

	eg.Go(func() error {
		logger.Info("start http server")
		err = http.Start(conf, logger)
		if err != nil {
			if !errors.Is(httpOri.ErrServerClosed, err) {
				logger.Fatalf("start http err: %s", err.Error())
			}
			logger.Infof("servier exit gracefully")

		}
		return nil
	})

	eg.Go(func() error {
		//优雅退出
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGSEGV)
		<-quit

		err = http.Stop()
		if err != nil {
			logger.Fatalf("stop http err: %s", err.Error())
		}

		return nil
	})

	err = eg.Wait()
	if err != nil {
		logger.Fatalf("stop server err: %s", err.Error())
	}
	logger.Infof("server stopped")
}
