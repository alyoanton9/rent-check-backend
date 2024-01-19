package handler

import (
	"github.com/labstack/echo/v4"
	"rent-checklist-backend/internal/dto"
)

func (h handler) AddGroupToFlat(ctx echo.Context) error {
	userId := ctx.Get("userId").(uint64)

	flatId, err := ParsePathParamUInt64(ctx, "flatId")
	if err != nil {
		return err
	}

	var groupIdFlatIdCopyRequest *dto.GroupIdFlatIdCopyRequest
	err = ParseBody(ctx, &groupIdFlatIdCopyRequest, "group id flat id copy request")
	if err != nil {
		return err
	}

	err = h.groupRepository.AddGroupToFlat(groupIdFlatIdCopyRequest.GroupId, flatId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error adding group to flat")
	}

	if groupIdFlatIdCopyRequest.FlatIdCopy != nil {
		err = h.itemRepository.CopyItemsFromFlatGroup(
			groupIdFlatIdCopyRequest.GroupId, flatId, *groupIdFlatIdCopyRequest.FlatIdCopy, userId)
		if err != nil {
			return HandleDbError(ctx, err, "error copying items from flat group")
		}
	}

	return nil
}
