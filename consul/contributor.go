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

package consul

import (
	"fmt"
	"path/filepath"
  "os"
	"github.com/buildpack/libbuildpack/application"
	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/helper"
	"github.com/cloudfoundry/libcfbuildpack/layers"
)

// Contributor is responsibile for deciding what this buildpack will contribute during build
type Contributor struct {
	app                application.Application
	launchContribution bool
	launchLayer        layers.Layers
	consulLayer        layers.DependencyLayer
}

// NewContributor will create a new Contributor object
func NewContributor(context build.Build) (c Contributor, willContribute bool, err error) {

	fmt.Println(os.Stdout, "Dependency %s",Dependency)


	plan, wantDependency, err := context.Plans.GetShallowMerged(Dependency)
	fmt.Println(os.Stdout, "plan %s ",plan)
	fmt.Println(os.Stdout, "err %s ",err)
	fmt.Println(os.Stdout, "wantDependency %t ",wantDependency)


	// if err != nil || !wantDependency {
	// 	return Contributor{}, false, err
	// }

	deps, err := context.Buildpack.Dependencies()

	fmt.Println(os.Stdout, "deps %s ",deps)
	fmt.Println(os.Stdout, "err %s ",err)
	if err != nil {
		return Contributor{}, false, err
	}




	fmt.Println(os.Stdout, "plan.Version %s ",plan.Version)


	dep, err := deps.Best(Dependency, plan.Version, context.Stack)


	fmt.Println(os.Stdout, "dep %s ",dep)
	fmt.Println(os.Stdout, "err %s ",err)
	if err != nil {
		return Contributor{}, false, err
	}

	fmt.Println(os.Stdout, "it will contribute ")

	contributor := Contributor{
		app:         context.Application,
		launchLayer: context.Layers,
		consulLayer:  context.Layers.DependencyLayer(dep),
	}

	contributor.launchContribution, _ = plan.Metadata["launch"].(bool)
	return contributor, true, nil
}

// Contribute will install consul, configure required env variables & set a start command
func (c Contributor) Contribute() error {

	fmt.Println(os.Stdout, "Contribute ")

	return c.consulLayer.Contribute(func(artifact string, layer layers.DependencyLayer) error {
		fmt.Println(os.Stdout, "artifact %s ",artifact)
		fmt.Println(os.Stdout, "layer.Root %s ",layer.Root)

		layer.Logger.SubsequentLine("Expanding to %s", layer.Root)
		if err := helper.ExtractZip(artifact, layer.Root, 1); err != nil {
			return err
		}

		// if err := helper.CopyFile(c.configurePath, filepath.Join(layer.Root, "bin")); err != nil {
		// 	return err
		// }

		if err := layer.AppendPathSharedEnv("PATH", filepath.Join(layer.Root, "sbin")); err != nil {
					return err
		}


		// if err := layer.OverrideLaunchEnv("APP_ROOT", c.app.Root); err != nil {
		// 	return err
		// }
		//
		//
		// if err := layer.OverrideLaunchEnv("SERVER_ROOT", layer.Root); err != nil {
		// 	return err
		// }

		return c.launchLayer.WriteApplicationMetadata(layers.Metadata{
			Processes: []layers.Process{{"web", fmt.Sprintf(`consul -f %s`, filepath.Join(c.app.Root, "config.json"))}},
		})
	}, c.flags()...)
}

func (c Contributor) flags() []layers.Flag {
	var flags []layers.Flag

	if c.launchContribution {
		flags = append(flags, layers.Launch)
	}

	return flags
}
