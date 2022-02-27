# aws-ecs-services-autoscaling-config

Simple tool for printing scale configuration (name, min, max) of AWS ECS services

## Usage

```
❯❯❯ go run . -h
Usage:
  aws-ecs-services-autoscaling-config [flags]

Flags:
  -c, --cluster string     name of the ECS cluster
  -h, --help               help for aws-ecs-services-autoscaling-config
  -k, --keys stringArray   keys used for filter results
```

### Print ScaleConfig for a default cluster
`aws-ecs-services-autoscaling-config`

### Print ScaleConfig for the specified cluster
`aws-ecs-services-autoscaling-config -c <cluster_name>`

### Print ScaleConfig for the specified cluster and services (search keys)
`aws-ecs-services-autoscaling-config -c <cluster_name> -k prefix1 prefix2 name1 suffix1 name2 suffix2 [...]`
