package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/smnalex/twilio-go"
	"github.com/smnalex/twilio-go/chat"
)

// TwilioRoleResource interface for twilio roles creation and retrieval.
type TwilioRoleResource interface {
	Read(ctx context.Context, serviceSID, roleSID string) (chat.Role, error)
	Create(ctx context.Context, serviceSID string, body chat.RoleCreateParams) (chat.Role, error)
}

// TwilioChannelResource interface for twilio channel reads
type TwilioChannelResource interface {
	Read(ctx context.Context, serviceSID, identity string) (chat.Channel, error)
}

type twilioService struct {
	tcr TwilioChannelResource
	trr TwilioRoleResource
}

func (ts twilioService) createRole(sid string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	role, err := ts.trr.Create(ctx, sid, chat.RoleCreateParams{
		FriendlyName: "Test",
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v", role)
}

func (ts twilioService) readChannel(sid, csid string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	role, err := ts.tcr.Read(ctx, sid, csid)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v", role)
}

func main() {
	tctx := twilio.NewContext()
	client, err := chat.New(tctx)
	if err != nil {
		log.Fatal(err)
	}
	ts := twilioService{
		tcr: client.Channels,
		trr: client.Roles,
	}
	ts.createRole("ServiceSID")
	ts.readChannel("ServiceSID", "ChannelSID")
}
