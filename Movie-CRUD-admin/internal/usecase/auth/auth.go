package auth

import (
	"context"
	"movie_crud/internal/config"
	"movie_crud/pkg/common"
	"movie_crud/pkg/response"

	"net/http"
)

type TokenValues struct {
	Session string
}
type Service interface {
	Authorize(next http.Handler) http.Handler
}

type service struct {
	//logger            logging.Service
	//config            *config.Config
	//sessionRepository ports.SessionRepository
}

func NewService() *service {
	return &service{}
}

func (s *service) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uniqueID, _ := common.GenerateUUID()
		ctx := context.WithValue(r.Context(), config.CtxKey{},
			&config.CtxKey{
				RequestID: uniqueID,
			})

		cookie, err := r.Cookie("session")
		if err != nil {
			apiResponse := response.NewAPIResponse("Unauthorized", http.StatusUnauthorized)
			response.ResponseJSON(w, r, apiResponse)
			return
		}

		token := cookie.Value
		values := TokenValues{}

		values.Session = token

		ctx = context.WithValue(ctx, TokenValues{}, values)
		ctxKey := ctx.Value(config.CtxKey{}).(*config.CtxKey)
		ctxKey.Session = token

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
