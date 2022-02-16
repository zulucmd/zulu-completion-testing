/*
Copyright The Helm Authors.
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

package main

import (
	"io"

	"github.com/gowarden/zulu"
)

var disableCompDescriptions bool

func nofilecomp(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
	return nil, zulu.ShellCompDirectiveNoFileComp
}

func newCompletionCmd(out io.Writer) *zulu.Command {
	cmd := &zulu.Command{Use: "completion"}
	cmd.PersistentFlags().BoolVar(&disableCompDescriptions, "no-descriptions", false, "disable completion descriptions")

	bash := &zulu.Command{
		Use:               "bash",
		ValidArgsFunction: nofilecomp,
		RunE: func(cmd *zulu.Command, args []string) error {
			return cmd.Root().GenBashCompletion(out, !disableCompDescriptions)
		},
	}

	zsh := &zulu.Command{
		Use:               "zsh",
		ValidArgsFunction: nofilecomp,
		RunE: func(cmd *zulu.Command, args []string) error {
			if disableCompDescriptions {
				return cmd.Root().GenZshCompletionNoDesc(out)
			} else {
				return cmd.Root().GenZshCompletion(out)
			}
		},
	}

	fish := &zulu.Command{
		Use:               "fish",
		ValidArgsFunction: nofilecomp,
		RunE: func(cmd *zulu.Command, args []string) error {
			return cmd.Root().GenFishCompletion(out, !disableCompDescriptions)
		},
	}

	pwsh := &zulu.Command{
		Use:               "powershell",
		ValidArgsFunction: nofilecomp,
		RunE: func(cmd *zulu.Command, args []string) error {
			if disableCompDescriptions {
				return cmd.Root().GenPowerShellCompletion(out)
			} else {
				return cmd.Root().GenPowerShellCompletionWithDesc(out)
			}
		},
	}

	cmd.AddCommand(bash, zsh, fish, pwsh)

	return cmd
}
