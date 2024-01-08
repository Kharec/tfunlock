package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-tfe"
)

const (
	organization = "" // put your organization name here
)

func getAllWorkspaces(client tfe.Client, ctx context.Context, pageMax int) []tfe.Workspace {
	var workspaces []tfe.Workspace
	page := 1
	for page = 1; page <= pageMax; page++ {
		options := &tfe.WorkspaceListOptions{
			ListOptions: tfe.ListOptions{
				PageNumber: page,
				PageSize:   100,
			},
		}
		w, err := client.Workspaces.List(ctx, organization, options)
		if err != nil {
			panic(err)
		} else {
			for _, space := range w.Items {
				workspaces = append(workspaces, *space)
			}
		}
	}
	return workspaces
}

func main() {
	tfCloudToken := os.Getenv("TF_CLOUD_TOKEN")
	config := &tfe.Config{
		Address:           "https://app.terraform.io",
		Token:             tfCloudToken,
		RetryServerErrors: true,
	}
	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	if len(os.Args) > 1 {
		for _, workspaceName := range os.Args[1:] {
			workspace, err := client.Workspaces.Read(ctx, organization, workspaceName)
			if err != nil {
				panic(err)
			}
			if workspace.Locked {
				_, err := client.Workspaces.Unlock(ctx, workspace.ID)
				if err != nil {
					log.Default().Fatal("Cannot unlock ", workspace.Name)
				}
				fmt.Println("[UNLOCKED]", workspace.Name)
			} else {
				fmt.Println(workspace.Name, "wasn't locked.")
			}
		}
	} else {
		fmt.Println("Unlocking all locked workspaces in", organization, "organization...")
		terraformWorkspaces := getAllWorkspaces(*client, ctx, 25)
		for _, workspace := range terraformWorkspaces {
			if workspace.Locked {
				_, err := client.Workspaces.Unlock(ctx, workspace.ID)
				if err != nil {
					log.Default().Fatal("Cannot unlock ", workspace.Name)
				}
				fmt.Println("[UNLOCKED]", workspace.Name)
			}
		}
	}
}
