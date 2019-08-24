package main

import (
	"dianying/src/share/config"
	"dianying/src/share/pb"
	"dianying/src/share/utils/log"
	"dianying/src/user-srv/db"
	"dianying/src/user-srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
)

func main()  {
	log.Init("user")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameUser),
		micro.Version("latest"),
	)

	//定义service动作操作
	service.Init(
		micro.Action(func(context *cli.Context) {
			logger.Info("info",zap.Any("user-srv","user-srv is start..."))

			//初始化数据库
			db.Init(config.MysqlDNS)
			pb.RegisterUserServiceExtHandler(service.Server(),handler.NewUserServiceExtHandler(),server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("info",zap.Any("user-srv","user-srv is stop..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run();err != nil{
		logger.Panic("user-srv 服务启动失败...")
	}


}
