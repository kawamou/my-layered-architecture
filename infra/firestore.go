package infra

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func NewFirestoreClient(ctx context.Context) *firestore.Client {
	conf := &firebase.Config{ProjectID: projectID()}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func projectID() string {
	return os.Getenv("PROJECT_ID")
}
