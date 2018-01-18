package handlers

import (
	"context"
	"net/http"

	"github.com/fabbricadigitale/go-scim/shared"
)

func GetAllResourceTypeHandler(r shared.WebRequest, server ScimServer, ctx context.Context) (ri *ResponseInfo) {
	ri = newResponse()

	repo := server.Repository(shared.ResourceTypeResourceType)
	userResourceType, err := repo.Get(shared.UserResourceType, "")
	ErrorCheck(err)
	groupResourceType, err := repo.Get(shared.GroupResourceType, "")
	ErrorCheck(err)

	jsonBytes, err := server.MarshalJSON([]interface{}{
		userResourceType.GetData(),
		groupResourceType.GetData(),
	}, nil, nil, nil)
	ErrorCheck(err)

	ri.Status(http.StatusOK)
	ri.Body(jsonBytes)
	return
}
