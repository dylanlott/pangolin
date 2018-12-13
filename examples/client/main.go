package main

import (
	"ctx"

	"github.com/golang/protobuf/ptypes/any"

	"github.com/dylanlott/pangolin/pkg/pb"
)

func main() {
	ctx := context.Background()

	drama := map[string]interface{}{
		"here": "is",
		"some": "drama",
	}
	d := drama["some"].(*any.Any)
	insertReq := &pb.InsertRequest{
		ApiKey: "pangomango",
		Data:   d,
	}

	// TODO: give me a conn
	pango := pb.NewPangolinDBClient()

	data, err := pango.Insert(ctx, insertReq)
	if err != nil {
		log.Fatalln("error inserting", err)
		os.Exit(1)
	}
	fmt.Println(data)
}
