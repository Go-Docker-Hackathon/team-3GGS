package main

import (
	"hera"
	_ "mate"
	"os"
	"path"
	"runtime"
)

func main() {
	defer func() {
		println("main exit, catch painc")
	}()
	initEnv()
	startSvc()
}

func initEnv() {
	os.Chdir(path.Dir(os.Args[0]))
	config := hera.NewConfig("../conf/mate.yaml")
	hera.MakeServerVar(config)
	hera.NewLogger(hera.SERVER["LOG_PATH"], hera.LevelDebug)
	hera.NewRedisSvc()
}

func startSvc() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n := hera.Classic()
	n.Run(":8083")
}
