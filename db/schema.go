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
	initializers.AppLogger.Info("Initializing Weaviate schema...")

	err := DeclareChatSchema(initializers.WeaviateClient)
	if err != nil {
		initializers.AppLogger.Error("Error in DeclareChatSchema", "error", err)
	} else {
		initializers.AppLogger.Info("DeclareChatSchema completed successfully")
	}

	err = DeclareContentSchema(initializers.WeaviateClient)
	if err != nil {
		initializers.AppLogger.Error("Error in DeclareContentSchema", "error", err)
	} else {
		initializers.AppLogger.Info("DeclareContentSchema completed successfully")
	}
}

func DeclareChatSchema(client *weaviate.Client) error {
	existingClass, _ := client.Schema().ClassGetter().
		WithClassName("Chat").
		Do(context.Background())

	if existingClass != nil {
		initializers.AppLogger.Info("Class 'Chat' already exists. Skipping creation.")
		return nil
	}

	chatClass := &models.Class{
		Class: constants.ClassChat.String(),
		ModuleConfig: map[string]any{
			"text2vec-transformers": map[string]any{
				"vectorizeClassName": false,
			},
		},
		Properties: []*models.Property{
			{Name: "query", DataType: []string{"text"}, Description: "User question or message"},
			{Name: "answer", DataType: []string{"text"}, Description: "Assistant response"},
			{Name: "userID", DataType: []string{"string"}, Description: "User identifier"},
			{Name: "timestamp", DataType: []string{"date"}, Description: "When the chat occurred"},
		},
	}

	err := client.Schema().ClassCreator().WithClass(chatClass).Do(context.Background())
	if err != nil {
		initializers.AppLogger.Error("Failed to create Chat class", "error", err)
		return fmt.Errorf("failed to create Chat class: %w", err)
	}

	initializers.AppLogger.Info("Class 'Chat' created successfully")
	return nil
}

func DeclareContentSchema(client *weaviate.Client) error {
	existingClass, _ := client.Schema().ClassGetter().
		WithClassName(constants.ClassContent.String()).
		Do(context.Background())

	if existingClass != nil {
		initializers.AppLogger.Info("Class 'Content' already exists. Skipping creation.")
		return nil
	}

	contentClass := &models.Class{
		Class: constants.ClassContent.String(),
		ModuleConfig: map[string]any{
			"text2vec-transformers": map[string]any{
				"vectorizeClassName": false,
			},
		},
		Properties: []*models.Property{
			{Name: "title", DataType: []string{"text"}, Description: "Title of the video"},
			{Name: "description", DataType: []string{"text"}, Description: "Short summary or description of the video"},
			{Name: "url", DataType: []string{"text"}, Description: "URL to the video"},
			{Name: "type", DataType: []string{"text"}, Description: "Type/category of the video"},
		},
	}

	err := client.Schema().ClassCreator().WithClass(contentClass).Do(context.Background())
	if err != nil {
		initializers.AppLogger.Error("Failed to create Content class", "error", err)
		return fmt.Errorf("failed to create Content class: %w", err)
	}

	initializers.AppLogger.Info("Class 'Content' created successfully")
	return nil
}
