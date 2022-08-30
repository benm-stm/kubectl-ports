/*
Copyright © 2022 BEN MANSOUR Mohamed Rafik

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ports",
	Short: "Show resources ports in clear tabular mode",
	Long: `Show resources ports in clear tabular mode, for the moment
	i only the service resource is taken in charge For example:

kubectl ports svc
or
kubectl ports svc -n namespace
or
kubectl ports svc -n namespace svc1 svc2`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
