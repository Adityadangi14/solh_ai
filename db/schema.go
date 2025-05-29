package db

import (
	"context"
	"fmt"

	"github.com/Adityadangi14/solh_ai/constants"
	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

func InitSchema() {
	err := DeclareChatSchema(initializers.WeaviateClient)
	err = DeclareContentSchema(initializers.WeaviateClient)

	fmt.Println(err)
}

func DeclareChatSchema(client *weaviate.Client) error {
	// Check if class already exists
	existingClass, _ := client.Schema().ClassGetter().
		WithClassName("Chat").
		Do(context.Background())

	if existingClass != nil {
		fmt.Println("Class 'Chat' already exists. Skipping creation.")
		return nil
	}

	// Define the Chat class
	chatClass := &models.Class{
		Class: constants.ClassChat.String(),
		ModuleConfig: map[string]any{
			"text2vec-transformers": map[string]any{
				"vectorizeClassName": false,
			},
		},
		Properties: []*models.Property{
			{
				Name:        "query",
				DataType:    []string{"text"},
				Description: "User question or message",
			},
			{
				Name:        "answer",
				DataType:    []string{"text"},
				Description: "Assistant response",
			},
			{
				Name:        "userID",
				DataType:    []string{"string"},
				Description: "User identifier",
			},
			{
				Name:        "timestamp",
				DataType:    []string{"date"},
				Description: "When the chat occurred",
			},
		},
	}

	// Create the class
	err := client.Schema().ClassCreator().WithClass(chatClass).Do(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	fmt.Println("Class 'Chat' created successfully.")
	return nil
}

func DeclareContentSchema(client *weaviate.Client) error {
	// Check if class already exists
	existingClass, _ := client.Schema().ClassGetter().
		WithClassName(constants.ClassContent.String()).
		Do(context.Background())

	if existingClass != nil {
		fmt.Println("Class 'Content' already exists. Skipping creation.")
		return nil
	}

	// Define the Chat class
	chatClass := &models.Class{
		Class: constants.ClassContent.String(),
		ModuleConfig: map[string]any{
			"text2vec-transformers": map[string]any{
				"vectorizeClassName": false,
			},
		},
		Properties: []*models.Property{
			{
				Name:        "title",
				DataType:    []string{"text"},
				Description: "Title of the video",
			},
			{
				Name:        "description",
				DataType:    []string{"text"},
				Description: "Short summary or description of the video",
			},
			{
				Name:        "url",
				DataType:    []string{"text"},
				Description: "URL to the YouTube video",
			},
			{
				Name:        "type",
				DataType:    []string{"text"},
				Description: "Type/category of the video",
			},
		},
	}

	// Create the class
	err := client.Schema().ClassCreator().WithClass(chatClass).Do(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	fmt.Println("Class 'Content' created successfully.")
	return nil
}
