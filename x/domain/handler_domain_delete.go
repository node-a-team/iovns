package domain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/iov-one/iovns/x/domain/keeper"
	"github.com/iov-one/iovns/x/domain/types"
)

func handleMsgDomainDelete(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteDomain) (*sdk.Result, error) {
	// check if domain exists
	domain, exists := k.GetDomain(ctx, msg.Domain)
	if !exists {
		return nil, sdkerrors.Wrapf(types.ErrDomainDoesNotExist, "%s does not exist", msg.Domain)
	}
	// check if domain has super user
	if !domain.HasSuperuser {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "can not delete domain with no superuser")
	}
	// check if domain admin matches msg owner
	if !domain.Admin.Equals(msg.Owner) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "address %s is not allowed to delete the domain owned by: %s", msg.Owner, domain.Admin)
	}
	// all checks passed delete domain
	_ = k.DeleteDomain(ctx, msg.Domain)
	// success TODO maybe emit event?
	return &sdk.Result{}, nil
}
