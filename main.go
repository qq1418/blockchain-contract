package main

import (
	"os"
	"strconv"

	_ "unicontract/src/api/routers"
	"unicontract/src/common/quartz"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	"unicontract/src/core/control"
	"unicontract/src/core/control/headNodeMonitor"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/engine"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/engine/gRPCClient"
	"unicontract/src/core/engine/scanengine"
	"unicontract/src/pipelines"

	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.EnableGzip = true
	beego.LoadAppConfig("ini", "../conf/app.conf")
	beego.BConfig.Log.AccessLogs = true
	// logInit must following the beego.LoadAppConfig
	logInit()

	argsCount := len(os.Args)
	if argsCount == 2 && os.Args[1] == "start" {
		runStart()
	} else if argsCount == 2 && os.Args[1] == "initdb" {
		runInitDB()
	} else if argsCount == 2 && os.Args[1] == "dropdb" {
		runDropDB()
	} else if argsCount == 4 && os.Args[1] == "reconfigdb" {
		shards, error := strconv.Atoi(os.Args[2])
		if error != nil {
			uniledgerlog.Error("shards should be int")
		}
		replicas, error := strconv.Atoi(os.Args[3])
		if error != nil {
			uniledgerlog.Error("replicas should be int")
		}
		runReconfigDB(shards, replicas)
	} else if argsCount == 2 && os.Args[1] == "config" {
		runConfig()
	} else if argsCount == 2 && os.Args[1] == "help" {
		runHelp()
	} else {
		uniledgerlog.Error("cmd should be " +
			"unicontract start|initdb|dropdb|reconfigdb $shards $replicas|config|help")
		os.Exit(2)
	}
}

func runStart() {
	config.Init()
	uniledgerlog.Info("config Init")
	pipelines.Init()
	uniledgerlog.Info("pipelines Init")
	engine.Init() // 1
	uniledgerlog.Info("engine Init")
	engineCommon.Init() // 2
	uniledgerlog.Info("scan engine db Init")
	quartz.Init() // 3
	uniledgerlog.Info("quartz Init")
	gRPCClient.Init() // 4
	uniledgerlog.Info("GRPC params Init")
	scanengine.Init() // 5      TODO: 1 2 3 的顺序一定不能错，而且 2 3 4 5 一定要在 1 后面
	uniledgerlog.Info("scanengine Init")
	go scanengine.Start()
	uniledgerlog.Info("UCVM ScanEngine Start")
	control.Init()
	uniledgerlog.Info("control Init")
	headNodeMonitor.Init()
	uniledgerlog.Info("headNodeMonitor Init")
	go headNodeMonitor.Start()
	uniledgerlog.Info("headNodeMonitor Start")
	beego.Run()
	uniledgerlog.Info("beego Run")
}

func runInitDB() {
	config.Init()
	uniledgerlog.Info("Database Init")
	rethinkdb.InitDatabase()

	engine.Init()               // 1
	engineCommon.Init()         // 2
	engineCommon.InitDatabase() // 3      TODO:初始化扫描引擎数据库，一定要按照 1 2 3 的顺序调用
}

func runDropDB() {
	config.Init()
	uniledgerlog.Info("Database Dropped")
	rethinkdb.DropDatabase()
}

func runReconfigDB(shards int, replicas int) {
	config.Init()
	uniledgerlog.Info("Database Reconfigured")
	rethinkdb.Reconfig(shards, replicas)
}

func runConfig() {
	config.WriteConToFile()
}

func runHelp() {
	uniledgerlog.Info("cmd should be " +
		"unicontract start|initdb|dropdb|reconfigdb $shards $replicas|config|help")
}

func logInit() {
	uniledgerlog.Init()
}
