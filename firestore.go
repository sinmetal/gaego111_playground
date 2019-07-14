package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"go.opencensus.io/trace"
)

func NewFirestoreClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	ctx, span := trace.StartSpan(ctx, "/firestore/NewFirestoreClient")
	defer span.End()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}
