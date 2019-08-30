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

	"github.com/cloudfoundry/consul-cnb/consul"

	"github.com/cloudfoundry/libcfbuildpack/build"
)

func main() {
	buildContext, err := build.DefaultBuild()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to build: %s", err)
		os.Exit(101)
	}

	fmt.Println(os.Stdout, "build main ")

	code, err := runBuild(buildContext)
	if err != nil {
		buildContext.Logger.Info(err.Error())
	}

	os.Exit(code)

}

func runBuild(context build.Build) (int, error) {

	fmt.Println(os.Stdout, "runBuildrunBuild %s ",context.Buildpack)

	context.Logger.Title(context.Buildpack)

	contributor, willContribute, err := consul.NewContributor(context)
	if err != nil {
		return context.Failure(102), err
	}
 fmt.Println(os.Stdout, "willContribute %t ",willContribute)
 
	if willContribute {
		err := contributor.Contribute()
		if err != nil {
			return context.Failure(103), err
		}
	}

	return context.Success()
}
