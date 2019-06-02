package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/smnalex/twilio-go"
	"github.com/smnalex/twilio-go/chat"
)

func main() {
	tctx := twilio.NewContext()
	client, err := chat.New(tctx)
	if err != nil {
		log.Fatal(err)
	}
	sid := "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	role, err := client.Roles.Create(context.Background(), sid, chat.RoleCreateParams{
		FriendlyName: "test-1234",
		Type:         "deployment",
		Permission:   []string{"createChannel", "joinChannel"},
	})
	if err != nil {
		log.Fatal(err)
	}
	data, _ := json.MarshalIndent(role, "", "\t")
	fmt.Printf("%s\n", data)

	role, err = client.Roles.Update(context.Background(), role.ServiceSID, role.SID, chat.RoleUpdateParams{
		Permission: []string{"destroyChannel"},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.Roles.Delete(context.Background(), role.ServiceSID, role.SID)
	if err != nil {
		log.Fatal(err)
	}
}
