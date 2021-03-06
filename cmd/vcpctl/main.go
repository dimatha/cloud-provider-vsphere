/*
Copyright 2018 The Kubernetes Authors.

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

// The vcpctl tool is responsible for facilitating cloud controller manager provisioning

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cloud-provider-vsphere/cmd/vcpctl/provision"
)

func main() {

	provision.AddProvision(cmd)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nCompleted!\n")
}

var cmd = &cobra.Command{
	Use:   "vcpctl",
	Short: "The vcpctl tool is responsible for facilitating cloud controller manager provisioining.",
	Long: `Deploying a cloud provider on vSphere is a task that has many prerequisites, this tool provides these needs:
* Perform vSphere configuration health check.
* Create vSphere role with a minimal set of permissioins.
* Create vSphere solution user, to be used with CCM
* Convert old in-tree vsphere.conf configuration files to new configMap

`,

	Run: RunMain,
}

// RunMain is the "Run" function callback for a cobra command object.
func RunMain(cmd *cobra.Command, args []string) {
	cmd.Help()
}
