package authentication

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/rigdev/rig-go-api/api/v1/authentication"
	"github.com/rigdev/rig/pkg/auth"
	"github.com/rigdev/rig/pkg/errors"
	"github.com/rigdev/rig/pkg/uuid"
)

// ResetPassword updates the users password if the provided code can be validated by the hash.
func (h *Handler) ResetPassword(ctx context.Context, req *connect.Request[authentication.ResetPasswordRequest]) (*connect.Response[authentication.ResetPasswordResponse], error) {
	pID, err := uuid.Parse(req.Msg.GetProjectId())
	if err != nil {
		return nil, errors.InvalidArgumentErrorf("invalid project ID")
	}
	ctx = auth.WithProjectID(ctx, pID)
	if err := h.as.ResetPassword(ctx, req.Msg.GetIdentifier(), req.Msg.GetCode(), req.Msg.GetNewPassword()); err != nil {
		return nil, err
	}
	return &connect.Response[authentication.ResetPasswordResponse]{}, nil
}
