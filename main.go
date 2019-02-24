package main

import (
	"os"

	"github.com/mfpierre/kubectl-plugin-boilerplate/pkg/cmd"
	"github.com/spf13/pflag"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-plugin-boilerplate", pflag.ExitOnError)
	pflag.CommandLine = flags

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
