// Copyright 2013 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// An open source project for official blog and documentation website for Beego web framework.
package main

import (
	"os"
	// "runtime"

	"github.com/astaxie/beego"
	"github.com/beego/beeweb/routers"
)

const (
	APP_VER = "0.4.0.0823"
)

func init() {
	//runtime.GOMAXPROCS(runtime.NumCPU())

	// Set App version and log level.
	if beego.AppConfig.String("runmode") == "pro" {
		beego.SetLevel(beego.LevelInfo)

		beego.Info("Beego Web", APP_VER)

		os.Mkdir("./log", os.ModePerm)
		filew := beego.NewFileWriter("log/log", true)
		err := filew.StartLogger()
		if err != nil {
			beego.Critical("NewFileWriter ->", err)
		}
	} else {
		// beewatch.Start(beewatch.Trace)
	}
}

func main() {
	beego.AppName = "Beego Web"
	beego.Info("Beego Web", APP_VER)

	// Register routers.
	beego.Router("/", &routers.HomeRouter{})
	beego.Router("/about", &routers.AboutRouter{})
	beego.Router("/community", &routers.CommunityRouter{})
	beego.Router("/quickstart", &routers.QuickStartRouter{})
	beego.Router("/docs", &routers.DocsRouter{})
	beego.Router("/docs/:all", &routers.DocsRouter{})
	beego.Router("/samples", &routers.SamplesRouter{})
	beego.Router("/samples/:all", &routers.SamplesRouter{})

	// Register template functions.

	beego.Run()
}
