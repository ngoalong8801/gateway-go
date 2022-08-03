package controller

import (
	"gateway-service/common"
	"gateway-service/grpc/client"
	"gateway-service/grpc/client/account"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type AccountController struct {
	accountClient *client.AccountClient
}

func NewAccountController(accountClient *client.AccountClient) *AccountController {
	return &AccountController{
		accountClient: accountClient,
	}
}

func (controller *AccountController) GetUser(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	req := &account.UserRequest{
		Id: int32(id),
	}

	res, err := controller.accountClient.GetUser(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		log.Println("Failed when get user account", err.Error())
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		ctx.JSON(http.StatusOK, res.GetData())
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}
