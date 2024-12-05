package position

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/position"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
)

// swagger:route post /position/list/personal position GetPositionListPersonal
//
// Get position list personal | 获取职位列表-个人用户版本
//
// Get position list personal | 获取职位列表-个人用户版本
//
// Responses:
//  200: PositionListResp

func GetPositionListPersonalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := position.NewGetPositionListPersonalLogic(r.Context(), svcCtx)
		resp, err := l.GetPositionListPersonal()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
