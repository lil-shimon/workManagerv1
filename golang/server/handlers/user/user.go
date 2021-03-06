package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/lil-shimon/workManader/golang/env"
	"github.com/lil-shimon/workManader/golang/errors"
	"github.com/lil-shimon/workManader/golang/models"
	"github.com/lil-shimon/workManader/golang/server/jwt"
	"github.com/lil-shimon/workManader/golang/server/write"
	"github.com/lil-shimon/workManader/golang/utils"
)

func Handler(env env.Env, w http.ResponseWriter, r *http.Request, u *models.User) http.HandlerFunc {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		if head == "verify" {
			return verify(env, w, r)
		} else {
			return signup(env, w, r)
		}
	case http.MethodGet:
		return whoami(env, w, r)
	default:
		return write.Error(errors.BadRequestMethod)
	}
}

type signupResponse struct {
	URL string
}

func signup(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var u models.User
	err := decoder.Decode(&u)
	if err != nil || &u == nil {
		return write.Error(errors.NoJSONBody)
	}

	dbUser, err := env.UserRepo().Signup(&u)
	if err != nil {
		return write.Error(err)
	}

	// TODO: this is where we should actually email the code with mailgun or something
	// for now we just pass verification code back in the response...
	return write.JSON(&signupResponse{
		URL: fmt.Sprintf("%s/verify/%s", os.Getenv("APP_ROOT"), dbUser.Verification),
	})
}

func whoami(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return write.JSONorErr(jwt.HandleUserCookie(env, w, r))
}

type verifyRequest struct {
	Code string
}

func verify(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var req verifyRequest

	err := decoder.Decode(&req)
	if err != nil || &req == nil || req.Code == "" {
		return write.Error(errors.NoJSONBody)
	}

	u, err := env.UserRepo().Verify(req.Code)
	if err != nil {
		return write.Error(err)
	}

	jwt.WriteUserCookie(w, u)
	return write.JSON(u)
}
