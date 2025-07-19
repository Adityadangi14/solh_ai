package db

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Adityadangi14/solh_ai/constants"
	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/fault"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func SaveData(obj map[string]any) (*data.ObjectWrapper, error) {
	res, err := initializers.WeaviateClient.Misc().LiveChecker().Do(context.Background())
	if err != nil {
		initializers.AppLogger.Error("Weaviate live check failed", "error", err)
		return nil, err
	}

	initializers.AppLogger.Info("Weaviate live check successful", "live", res)

	creator := initializers.WeaviateClient.Data().Creator()

	created, err := creator.
		WithClassName(constants.ClassChat.String()).
		WithProperties(obj).
		Do(context.Background())

	if err != nil {
		initializers.AppLogger.Error("Error while creating object in Weaviate", "error", err)
		return nil, fmt.Errorf("error while creating object: %w", err)
	}

	initializers.AppLogger.Info("Object created successfully in Weaviate", "class", constants.ClassChat.String())
	return created, nil
}

func GetPreviousChat() (*models.GraphQLResponse, error) {
	fields := []graphql.Field{
		{Name: "query"},
		{Name: "answer"},
		{Name: "userID"},
		{Name: "timestamp"},
	}

	resp, err := initializers.WeaviateClient.GraphQL().Get().
		WithClassName(constants.ClassChat.String()).
		WithFields(fields...).
		WithLimit(10).
		Do(context.Background())

	if err != nil {
		initializers.AppLogger.Error("Failed to get previous chats", "error", err)
		return nil, fmt.Errorf("failed to get previous chat: %w", err)
	}

	initializers.AppLogger.Info("Fetched previous chats successfully", "count", len(resp.Data))
	return resp, nil
}

func DeleteAllChat() error {
	err := initializers.WeaviateClient.Schema().ClassDeleter().
		WithClassName(constants.ClassChat.String()).
		Do(context.Background())

	if err != nil {
		if status, ok := err.(*fault.WeaviateClientError); ok && status.StatusCode != http.StatusBadRequest {
			initializers.AppLogger.Error("Failed to delete class", "error", err)
			return err
		}
	}

	initializers.AppLogger.Info("All chats deleted successfully")
	return nil
}

func DeleteChatByUserId(userId string) error {
	_, err := initializers.WeaviateClient.Batch().ObjectsBatchDeleter().
		WithClassName(constants.ClassChat.String()).
		WithWhere(filters.Where().
			WithPath([]string{"userID"}).
			WithOperator(filters.Equal).
			WithValueText(userId),
		).
		Do(context.Background())

	if err != nil {
		initializers.AppLogger.Error("Failed to delete chat by userID", "userID", userId, "error", err)
		return err
	}

	initializers.AppLogger.Info("Deleted chat(s) by userID", "userID", userId)
	return nil
}

func ReadChatsByUserId(userId string) (*models.GraphQLResponse, error) {
	fields := []graphql.Field{
		{Name: "query"},
		{Name: "answer"},
		{Name: "userID"},
		{Name: "timestamp"},
	}

	response, err := initializers.WeaviateClient.GraphQL().Get().
		WithClassName(constants.ClassChat.String()).
		WithFields(fields...).
		WithWhere(filters.Where().
			WithPath([]string{"userID"}).
			WithOperator(filters.Equal).
			WithValueText(userId),
		).
		WithSort(
			graphql.Sort{
				Path:  []string{"timestamp"},
				Order: graphql.Desc,
			},
		).
		WithLimit(90).
		Do(context.Background())

	if err != nil {
		initializers.AppLogger.Error("GraphQL query failed", "userID", userId, "error", err)
		return nil, fmt.Errorf("GraphQL query failed: %w", err)
	}

	if raw, err := json.Marshal(response.Data); err == nil {
		initializers.AppLogger.Info("Fetched chats for user", "userID", userId)
		// You can log raw string if needed:
		// initializers.AppLogger.Debug("Chat raw JSON", "json", string(raw))
		_ = raw
	}

	return response, nil
}
