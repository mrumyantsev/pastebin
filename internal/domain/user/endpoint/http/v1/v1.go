package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin/internal/domain/user"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
	"github.com/mrumyantsev/pastebin/internal/pkg/core"
	"github.com/mrumyantsev/pastebin/internal/pkg/core/response"
	urlquery "github.com/mrumyantsev/pastebin/internal/pkg/core/url-query"
)

type UserV1HttpEndpoint struct {
	config      *config.Config
	userUseCase user.UserUseCase
}

func NewUserV1HttpEndpoint(cfg *config.Config, userUC user.UserUseCase) *UserV1HttpEndpoint {
	return &UserV1HttpEndpoint{config: cfg, userUseCase: userUC}
}

func (e *UserV1HttpEndpoint) CreateUser(ctx *gin.Context) {
	var input user.UserInput

	err := ctx.BindJSON(&input)
	if err != nil {
		response.Fail(ctx, err, core.ErrBindJson)
		return
	}

	output, err := e.userUseCase.CreateUser(input)
	if err != nil {
		response.Fail(ctx, err, core.ErrCreateUser)
		return
	}

	response.Object(ctx, output, core.MsgUserCreated)
}

func (e *UserV1HttpEndpoint) GetUsers(ctx *gin.Context) {
	pg := urlquery.Pagination(ctx, e.config)

	users, err := e.userUseCase.GetUsers(pg)
	if err != nil {
		response.Fail(ctx, err, core.ErrGetUsers)
		return
	}

	response.Object(ctx, users, core.MsgUsersGot)
}

func (e *UserV1HttpEndpoint) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(user.ColId))
	if err != nil {
		response.Fail(ctx, err, core.ErrParseContextParameter)
		return
	}

	usr, err := e.userUseCase.GetUser(id)
	if err != nil {
		response.Fail(ctx, err, core.ErrGetUser)
		return
	}

	response.Object(ctx, usr, core.MsgUserGot)
}

func (e *UserV1HttpEndpoint) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(user.ColId))
	if err != nil {
		response.Fail(ctx, err, core.ErrParseContextParameter)
		return
	}

	var input user.UserInput

	if err = ctx.BindJSON(&input); err != nil {
		response.Fail(ctx, err, core.ErrBindJson)
		return
	}

	if err = e.userUseCase.UpdateUser(id, input); err != nil {
		response.Fail(ctx, err, core.ErrUpdateUser)
		return
	}

	response.Success(ctx, core.MsgUserUpdated)
}

func (e *UserV1HttpEndpoint) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(user.ColId))
	if err != nil {
		response.Fail(ctx, err, core.ErrParseContextParameter)
		return
	}

	if err = e.userUseCase.DeleteUser(id); err != nil {
		response.Fail(ctx, err, core.ErrDeleteUser)
		return
	}

	response.Success(ctx, core.MsgUserDeleted)
}

func (e *UserV1HttpEndpoint) IsUserExists(ctx *gin.Context) {
	username := ctx.Param(user.ColUsername)

	output, err := e.userUseCase.IsUserExists(username)
	if err != nil {
		response.Fail(ctx, err, core.ErrCheckUserExistance)
		return
	}

	response.Object(ctx, output, core.MsgUserExistanceChecked)
}

func (e *UserV1HttpEndpoint) IsEmailExists(ctx *gin.Context) {
	email := ctx.Param(user.ColEmail)

	output, err := e.userUseCase.IsEmailExists(email)
	if err != nil {
		response.Fail(ctx, err, core.ErrCheckEmailExistance)
		return
	}

	response.Object(ctx, output, core.MsgEmailExistanceChecked)
}

func (e *UserV1HttpEndpoint) UserCount(ctx *gin.Context) {
	output, err := e.userUseCase.UserCount()
	if err != nil {
		response.Fail(ctx, err, core.ErrGetUserCount)
		return
	}

	response.Object(ctx, output, core.MsgUserCountGot)
}
