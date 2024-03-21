package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/api/internal/logic/admin"
	"zero/api/internal/svc"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
