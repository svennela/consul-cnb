/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"
	//"path/filepath"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/consul-cnb/consul"
	//"github.com/cloudfoundry/libcfbuildpack/helper"

	"github.com/cloudfoundry/libcfbuildpack/detect"
)

func main() {
	detectionContext, err := detect.DefaultDetect()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to run detection: %s", err)
		os.Exit(101)
	}

	code, err := runDetect(detectionContext)

	fmt.Fprintf(os.Stdout, "code %s",code)
	fmt.Fprintf(os.Stdout, "err %s",err)

	if err != nil {
		detectionContext.Logger.Info(err.Error())
	}

	os.Exit(code)
}

func runDetect(context detect.Detect) (int, error) {

	fmt.Println(os.Stdout, "runDetect consul detecting")
	 //return context.Fail(), nil
	// consulConfExists, err := helper.FileExists(filepath.Join(context.Application.Root, "config.json"))
	// fmt.Println(os.Stdout, "context.Application.Root: %s",context.Application.Root)
	// fmt.Println(os.Stdout, "consulConfExists %t",consulConfExists)
	// fmt.Println(os.Stdout, "consulConfExists :err: %s",err)

	// if err != nil {
	// 	return context.Fail(), err
	// }

	//buildpackYAML, err := consul.LoadBuildpackYAML(context.Application.Root)
	//fmt.Fprintf(os.Stdout, "buildpackYAML %s",buildpackYAML)


	// if err != nil {
	// 	return context.Fail(), err
	// }

	plan := buildplan.Plan{
		Provides: []buildplan.Provided{{Name: consul.Dependency}},
	}

	fmt.Println(os.Stdout, "plan : %s",plan)

	//if consulConfExists {
		plan.Requires = []buildplan.Required{
			{
				Name:     consul.Dependency,
				Metadata: buildplan.Metadata{"launch": true},
			},
		}
	//}
	fmt.Println(os.Stdout, "--consul.Dependency %s--",consul.Dependency)
	fmt.Println(os.Stdout, "--plan %s--",plan)

	return context.Pass(plan)
}
