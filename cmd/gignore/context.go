// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"context"

	"github.com/onyx-and-iris/gignore"
	log "github.com/sirupsen/logrus"
)

type contextKey string

const clientKey contextKey = "client"

func getClientFromContext(ctx context.Context) *gignore.Client {
	client, ok := ctx.Value(clientKey).(*gignore.Client)
	if !ok {
		log.Fatal("Client not found in context")
	}
	return client
}
