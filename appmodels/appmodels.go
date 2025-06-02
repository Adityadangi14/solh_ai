package appmodels

import "time"

type Chat struct {
	Query     string
	Answer    string
	UserID    string
	Timestamp time.Time
}

func (c *Chat) Map() map[string]any {
	return map[string]any{
		"query":     c.Query,
		"answer":    c.Answer,
		"userID":    c.UserID,
		"timestamp": c.Timestamp,
	}
}

type Content struct {
	Title       string
	Description string
	Url         string
	Image       string
	ContentType string
}

func (c *Content) Map() map[string]any {
	return map[string]any{
		"title":       c.Title,
		"description": c.Description,
		"url":         c.Url,
		"image":       c.Image,
		"contentType": c.ContentType,
	}
}
