package main

import (
	"helm.sh/helm/v3/pkg/action"
)

func PullFromRegistry(actionConfig *action.Configuration) {

	client := action.NewChartPull(actionConfig)

	client.Run(nil, "hellp")


}
