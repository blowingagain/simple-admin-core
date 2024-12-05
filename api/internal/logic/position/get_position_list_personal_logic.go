package position

import (
	"context"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"strconv"
	"strings"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionListPersonalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPositionListPersonalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListPersonalLogic {
	return &GetPositionListPersonalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPositionListPersonalLogic) GetPositionListPersonal() (resp *types.PositionListResp, err error) {
	tokenPIds, _ := l.ctx.Value("positionIds").(string)
	// 将当前角色的id转为数组
	pIds := strings.Split(tokenPIds, ",")
	positionIds := []uint64{}
	for _, pid := range pIds {
		num, _ := strconv.ParseUint(pid, 10, 64)
		positionIds = append(positionIds, num)
	}

	data, err := l.svcCtx.CoreRpc.GetPositionList(l.ctx,
		&core.PositionListReq{
			Ids: positionIds,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.PositionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.PositionInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Trans:  l.svcCtx.Trans.Trans(l.ctx, *v.Name),
				Status: v.Status,
				Sort:   v.Sort,
				Name:   v.Name,
				Code:   v.Code,
				Remark: v.Remark,
			})
	}
	return resp, nil
}
