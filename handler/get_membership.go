package handler

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/tap202207/tap/gen/models"
	"github.com/tap202207/tap/gen/restapi/handlers"
)

func GetMembershipCard(params handlers.GetMembershipCardParams) middleware.Responder {
	log.Print(params.MemberID)
	barcode := "1111111111111111"
	token := "2222"
	return handlers.NewGetMembershipCardOK().
		WithPayload(&models.Barcode{
			Barcode: &barcode,
			Token:   &token,
		})
}
