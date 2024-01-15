package database

import (
	"backend/internal/models"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

// FormatPgUuid Formats a pgtype.UUID into a string
// Adapted from https://stackoverflow.com/a/71134336
func FormatPgUuid(uuid pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid.Bytes[0:4],
		uuid.Bytes[4:6],
		uuid.Bytes[6:8],
		uuid.Bytes[8:10],
		uuid.Bytes[10:16])
}

// FormatPgComment Formats a database.Comment into a models.Comment
func FormatPgComment(pgComment Comment) models.Comment {
	return models.Comment{
		ID:          FormatPgUuid(pgComment.ID),
		Body:        pgComment.Body,
		Creator:     pgComment.Creator,
		ThreadID:    FormatPgUuid(pgComment.ThreadID),
		CreatedTime: pgComment.CreatedTime.Time,
		UpdatedTime: pgComment.UpdatedTime.Time,
	}
}

// FormatPgComments Formats a slice of database.Comment into a slice of models.Comment
func FormatPgComments(pgComments []Comment) []models.Comment {
	comments := []models.Comment{}
	for _, pgComment := range pgComments {
		comments = append(comments, FormatPgComment(pgComment))
	}
	return comments
}

// FormatPgThread Formats a database.GetThreadDetailsRow into a models.Thread
func FormatPgThread(pgThread GetThreadDetailsRow) models.Thread {
	return models.Thread{
		ID:          FormatPgUuid(pgThread.ID),
		Title:       pgThread.Title,
		Body:        pgThread.Body,
		Creator:     pgThread.Creator,
		CreatedTime: pgThread.CreatedTime.Time,
		UpdatedTime: pgThread.UpdatedTime.Time,
		NumComments: pgThread.NumComments,
		Tags:        pgThread.Tags,
	}
}

// FormatPgThreads Formats a slice of database.GetThreadsByCriteriaRow into a slice of models.Thread
func FormatPgThreads(pgThread []GetThreadsByCriteriaRow) []models.Thread {
	threads := []models.Thread{}
	for _, pgThread := range pgThread {
		// Conversion is possible as both types have the same fields
		threads = append(threads, FormatPgThread(GetThreadDetailsRow(pgThread)))
	}
	return threads
}
