package handler

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/tap202207/tap/gen/models"
	"github.com/tap202207/tap/gen/restapi/handlers"
)

type schema struct {
	MemberId      string `firestore:"memberId"`
	MemberBarcode string `firestore:"member_barcode"`
	CompanyCode   string `firestore:"company_code"`
	Pin           string `firestore:"pin"`
}

func GetMembershipCard(params handlers.GetMembershipCardParams) middleware.Responder {
	log.Print(params.MemberID)
	firebaseApp, err := firebase.NewApp(
		context.TODO(), &firebase.Config{
			ProjectID: "okaimono-app-prototype",
		},
	)
	if err != nil {
		log.Fatalf("failed to init firebase: %s", err.Error())
	}
	firestoreClient, err := firebaseApp.Firestore(context.TODO())
	if err != nil {
		log.Fatalf("failed to init firestore client: %s", err.Error())
	}
	doc, err := firestoreClient.Doc(
		"membership_card/0123456789012345678901234567890123456789012345678901234567890123",
	).Get(context.TODO())
	var v schema
	if err = (doc.DataTo(&v)); err != nil {
		return nil
	}
	barcode := v.MemberBarcode
	token := v.Pin
	return handlers.NewGetMembershipCardOK().
		WithPayload(&models.Barcode{
			Barcode: &barcode,
			Token:   &token,
		})
}
