package menu

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"strconv"
	"strings"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole() (resp *types.MenuListResp, err error) {
	roleId, _ := l.ctx.Value("roleId").(string)
	// 将当前角色的id转为数组
	roleIds := strings.Split(roleId, ",")
	// 2024-10-25 新增区域id 用于对菜单的加载对应其关联的省份
	regionId := int64(0)
	if jsonUid, ok := l.ctx.Value("regionId").(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			regionId = int64Uid
		} else {
			err.Error()
		}
	}
	// 如果为管理员 && 选择的为全国，则不做处理
	if strings.Contains(roleId, "000") && regionId == 999 {

	} else {
		fmt.Println("-----------------roleId: ", roleId, ",regionId:", regionId)
		role, err := l.svcCtx.CoreRpc.GetRoleList(l.ctx, &core.RoleListReq{
			Remark: pointy.GetPointer(strconv.Itoa(int(regionId))),
		})
		if err != nil {
			return nil, err
		}
		if role.Total > 0 {
			//fmt.Println("-------- 当前角色数据大于0")
			for _, datum1 := range roleIds {
				for _, datum := range role.Data {
					if datum1 == *datum.Code {
						// 最终匹配的菜单权限
						roleId = *datum.Code
					}
				}
			}
		}
		roleId = "-1"
		fmt.Println("final-----------------roleId: ", roleId)
	}

	data, err := l.svcCtx.CoreRpc.GetMenuListByRole(l.ctx, &core.BaseMsg{Msg: roleId})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuListResp{}
	resp.Data.Total = data.Total
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.MenuInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id: v.Id,
			},
			MenuType:    v.MenuType,
			Level:       v.Level,
			Path:        v.Path,
			Name:        v.Name,
			Redirect:    v.Redirect,
			Component:   v.Component,
			Sort:        v.Sort,
			ParentId:    v.ParentId,
			ServiceName: v.ServiceName,
			Permission:  v.Permission,
			Meta: types.Meta{
				Title:              pointy.GetPointer(l.svcCtx.Trans.Trans(l.ctx, *v.Meta.Title)),
				Icon:               v.Meta.Icon,
				HideMenu:           v.Meta.HideMenu,
				HideBreadcrumb:     v.Meta.HideBreadcrumb,
				IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
				HideTab:            v.Meta.HideTab,
				FrameSrc:           v.Meta.FrameSrc,
				CarryParam:         v.Meta.CarryParam,
				HideChildrenInMenu: v.Meta.HideChildrenInMenu,
				Affix:              v.Meta.Affix,
				DynamicLevel:       v.Meta.DynamicLevel,
				RealPath:           v.Meta.RealPath,
			},
		})
	}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
