package cmd

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func ifSearchKey(sc *ScaleConfig, keys []string) bool {
	for _, k := range keys {
		if strings.Contains(sc.Name, k) {
			return true
		}
	}

	return false
}

func printTable(rows [][]string, headers []string) {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(headers)
	tw.AppendBulk(rows)
	tw.Render()
}

func getNameFromResourceId(resourceId string) string {
	parts := strings.Split(resourceId, "/")
	return parts[len(parts)-1]
}
