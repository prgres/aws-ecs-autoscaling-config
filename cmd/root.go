package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	fCluster   string
	fSearchKey []string

	_tableHeaders = []string{"Service", "min", "max"}
)

var rootCmd = &cobra.Command{
	Use: "aws-ecs-services-autoscaling-config",
	Run: rootCmdFunc,
}

func Execute() {
	rootCmd.Flags().StringVarP(&fCluster, "cluster", "c", "", "name of the ECS cluster")
	rootCmd.Flags().StringArrayVarP(&fSearchKey, "keys", "k", []string{}, "keys used for filter results")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootCmdFunc(cmd *cobra.Command, args []string) {
	scaleConfig, err := getScaleConfig(fCluster)
	if err != nil {
		log.Fatal(err)
	}

	filteredResults := filterScaleConfigList(scaleConfig, fSearchKey)
	printTable(scaleConfigListToRows(filteredResults), _tableHeaders)
}
