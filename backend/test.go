package main

import (
	"backend/tutorial"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// list all authors
	authors, err := queries.GetThreads(ctx, tutorial.GetThreadsParams{
		Column1: "created_time_asc",
		Limit:   10,
		Offset:  0,
	})
	if err != nil {
		return err
	}
	log.Println(authors)
	for _, author := range authors {
		var myUUID, _ = author.ID.Value()
		log.Println("q1", myUUID)
	}

	// query 2
	threads, err := queries.GetThreadsByMultipleKeyword(ctx, tutorial.GetThreadsByMultipleKeywordParams{
		ToTsquery: "how & homework",
		Column2:   "created_time_asc",
		Limit:     10,
		Offset:    0,
	})

	if err != nil {
		return err
	}

	log.Println("q2", threads)

	//query 3
	//threads_by_tag, err := queries.GetThreadsByMultipleTags(ctx, tutorial.GetThreadsByMultipleTagsParams{
	//	TagName: "'my-first-post'",
	//	Column3: "created_time_asc",
	//	Limit:   10,
	//	Offset:  0,
	//})

	threads_by_tag, err := queries.GetThreadsByMultipleTagsv2(ctx, tutorial.GetThreadsByMultipleTagsv2Params{
		TagName: []string{"my-first-post", "homework"},
		Column2: "created_time_asc",
		Limit:   10,
		Offset:  0,
	})

	if err != nil {
		return err
	}
	log.Println("q3", threads_by_tag)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
