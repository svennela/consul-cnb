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
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"
	"github.com/cloudfoundry/libcfbuildpack/helper"
	yaml "gopkg.in/yaml.v2"
)

// Dependency is the key used in the build plan by this buildpack
const Dependency = "consul"

// BuildpackYAML defines configuration options allowed to end users
type BuildpackYAML struct {
	Config Config `yaml:"consul"`
}

// Config is used by BuildpackYAML and defines HTTPD specific config options available to users
type Config struct {
	Version string `yaml:"version"`
}

// LoadBuildpackYAML reads `buildpack.yml` and HTTPD specific config options in it
func LoadBuildpackYAML(appRoot string) (BuildpackYAML, error) {

	buildpackYAML, configFile := BuildpackYAML{}, filepath.Join(appRoot, "buildpack.yml")

	fmt.Println(os.Stdout, "buildpackYAML %s ",buildpackYAML)
	fmt.Println(os.Stdout, "configFile %s ",configFile)


	if exists, err := helper.FileExists(configFile); err != nil {
		return BuildpackYAML{}, err
	} else if exists {
		fmt.Println(os.Stdout, "else  %s ",configFile)
		file, err := os.Open(configFile)
		if err != nil {
			return BuildpackYAML{}, err
		}
		defer file.Close()

		contents, err := ioutil.ReadAll(file)
		fmt.Println(os.Stdout, "buildpackYAMLcontents  %s ",contents)
		if err != nil {
			return BuildpackYAML{}, err
		}

		err = yaml.Unmarshal(contents, &buildpackYAML)
		if err != nil {
			return BuildpackYAML{}, err
		}
	}
	return buildpackYAML, nil
}
