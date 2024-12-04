/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package router

import (
	"net/http"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/config"
	"openeuler.org/PilotGo/atune-plugin/controller"
	"openeuler.org/PilotGo/atune-plugin/plugin"
)

func HttpServerInit(conf *config.HttpServer) error {

	go func() {
		r := setupRouter()

		logger.Info("start http service on: http://%s", conf.Addr)
		if err := r.Run(conf.Addr); err != nil {
			logger.Error("start http server failed:%v", err)
		}

	}()

	return nil
}
func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logger.RequestLogger())
	router.Use(gin.Recovery())

	registerAPIs(router)
	StaticRouter(router)

	return router
}
func registerAPIs(router *gin.Engine) {
	logger.Debug("router register")
	plugin.GlobalClient.RegisterHandlers(router)
	api := router.Group("/plugin/" + plugin.GlobalClient.PluginInfo.Name)

	atune := api.Group("")
	{
		atune.POST("atune_install", controller.AtuneClientInstall)
		atune.POST("atune_uninstall", controller.AtuneClientRemove)
		atune.GET("all", controller.GetAtuneAll)
		atune.GET("info", controller.GetAtuneInfo)
	}

	dbtune := api.Group("")
	{
		dbtune.GET("tunes", controller.QueryTunes)
		dbtune.GET("tunes_noPaged", controller.QueryTunesNoPaged)
		dbtune.POST("save_tune", controller.SaveTune)
		dbtune.POST("update_tune", controller.UpdateTune)
		dbtune.DELETE("delete_tune", controller.DeleteTune)
		dbtune.GET("search_tune", controller.SearchTune)
		dbtune.GET("tune_machine", controller.GetTuneMachines)
	}

	task := api.Group("")
	{
		task.POST("task_run", controller.StartTask)
		task.POST("task_new", controller.CreatTask)
		task.GET("tasks", controller.TaskLists)
		task.DELETE("task_delete", controller.DeleteTask)
		task.GET("task_search", controller.SearchTask)
	}

	restune := api.Group("")
	{
		restune.GET("result", controller.QueryResults)
		restune.GET("tune_result", controller.QueryTuneResults)
		restune.DELETE("result_delete", controller.DeleteResult)
	}
}

func StaticRouter(router *gin.Engine) {
	router.Static("/plugin/atune/static", "../web/dist/static")
	router.StaticFile("/plugin/atune", "../web/dist/index.html")

	// 解决页面刷新404的问题
	router.NoRoute(func(c *gin.Context) {
		logger.Debug("process noroute: %s", c.Request.URL)
		if !strings.HasPrefix(c.Request.RequestURI, "/plugin/atune/*path") {
			c.File("../web/dist/index.html")
			return
		}
		c.Status(http.StatusNotFound)
	})
}
