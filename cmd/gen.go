// Copyright © 2019 Andy Pan <panjf2000@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var license, template string
var licensePathTemplate, icuPathTemplate = "licenses/%s.txt", "licenses/996.icu.template.%s.txt"

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen is a 996.icu license generator-command.",
	Long: `gen is a 996.icu license generator-command,
it is used to generate various open-source licenses including MIT, Apache, etc.
More importantly, the main purpose of this tool is to incoporate those aforesaid licenses into
a brand new license: 996.icu, defined by this repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		icuTemplate, err := ioutil.ReadFile(fmt.Sprintf(icuPathTemplate, template))
		handleError(err)
		licenseContent, err := ioutil.ReadFile(fmt.Sprintf(licensePathTemplate, license))
		handleError(err)
		newLicenseContent := strings.Replace(strings.Replace(string(icuTemplate), "{other}", string(license), -1), "{content}", string(licenseContent), 1)
		handleError(ioutil.WriteFile("LICENSE", []byte(newLicenseContent), 0644))
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genCmd.Flags().StringVarP(&license, "license", "l", "", "generate a specific license")
	genCmd.Flags().StringVar(&template, "996icu", "en-us", "incoporate a specific license into 996icu license")
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
