package main

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"go.opencensus.io/trace"
)

type FirestoreStore struct {
	db *firestore.Client
}

func NewFirestoreStore(ctx context.Context, client *firestore.Client) (*FirestoreStore, error) {
	return &FirestoreStore{
		db: client,
	}, nil
}

type Sample struct {
	ID        string `firestore:"-"`
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *FirestoreStore) Create(ctx context.Context, sample *Sample) (*Sample, error) {
	ctx, span := trace.StartSpan(ctx, "/firestoreStore/Create")
	defer span.End()

	now := time.Now()
	sample.CreatedAt = now
	sample.UpdatedAt = now

	docRef, _, err := s.db.Collection("Sample").Add(ctx, sample)
	sample.ID = docRef.ID
	if err != nil {
		return nil, err
	}

	return sample, nil
}
