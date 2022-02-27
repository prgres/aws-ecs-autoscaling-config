package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	applicationautoscalingtypes "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/aws/aws-sdk-go/aws"
)

type ScaleConfig struct {
	Name string
	Min  int32
	Max  int32
}

func (s *ScaleConfig) Row() []string {
	return []string{s.Name, fmt.Sprintf("%d", s.Min), fmt.Sprintf("%d", s.Max)}
}

func scaleConfigListToRows(scaleConfig []*ScaleConfig) [][]string {
	result := make([][]string, len(scaleConfig))
	for i, r := range scaleConfig {
		result[i] = r.Row()
	}

	return result
}

func getScaleConfig(cluster string) ([]*ScaleConfig, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := applicationautoscaling.NewFromConfig(cfg)

	var nextToken *string
	result := make([]*ScaleConfig, 0)

	for {
		output, err := svc.DescribeScalableTargets(ctx, &applicationautoscaling.DescribeScalableTargetsInput{
			ServiceNamespace: applicationautoscalingtypes.ServiceNamespaceEcs,
			MaxResults:       aws.Int32(100),
			NextToken:        nextToken,
		})
		if err != nil {
			return nil, err
		}

		for _, sp := range output.ScalableTargets {
			result = append(result, &ScaleConfig{
				Name: getNameFromResourceId(*sp.ResourceId),
				Min:  *sp.MinCapacity,
				Max:  *sp.MaxCapacity,
			})
		}

		nextToken = output.NextToken
		if nextToken == nil {
			break
		}
	}

	return result, nil
}

func filterScaleConfigList(scaleConfig []*ScaleConfig, filters []string) []*ScaleConfig {
	if len(filters) == 0 {
		return scaleConfig
	}

	var results []*ScaleConfig

	for _, sc := range scaleConfig {
		if ifSearchKey(sc, filters) {
			results = append(results, sc)
		}
	}

	return results
}
