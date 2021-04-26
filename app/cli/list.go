package main

import (
	"helm.sh/helm/v3/pkg/action"
	"log"
	"strings"
)

type HelmInstallation struct {
	FullName string
	Name     string
}

func ExistInHelm(s string, h []HelmInstallation) bool {
	for _, v := range h {
		if v.Name == s {
			return false
		}
	}

	return true
}

func List(actionConfig *action.Configuration) ([]HelmInstallation, error) {
	//client := action.NewInstall(actionConfig)

	client := action.NewList(actionConfig)
	client.AllNamespaces = true
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	all := make([]HelmInstallation, 0)
	for _, v := range results {
		if strings.HasPrefix(v.Name, INSTALLATION_PREFIX+".") {
			all = append(all, HelmInstallation{
				FullName: v.Name,
				Name:     strings.TrimPrefix(v.Name, INSTALLATION_PREFIX+"."),
			})
		}
	}
	return nil, nil
}
