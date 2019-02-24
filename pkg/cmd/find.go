package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:       "find",
	Short:     "Find the node running the provided pod",
	ValidArgs: []string{"pod-name"},
	RunE: func(cmd *cobra.Command, args []string) error {
		GlobalSettings.InitClient()
		if len(args) == 0 {
			return fmt.Errorf("you have to specify the pod name")
		}
		podName := args[0]
		nodeName, err := GlobalSettings.GeNodeForPod(podName)
		if err != nil {
			return err
		}
		fmt.Printf("%s is running on node %s\n", podName, nodeName)
		return nil
	},
}
