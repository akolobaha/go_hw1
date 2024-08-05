package httphandler

import (
	"authservice/internal/domain"
	"authservice/internal/service"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/sign_up", CORS(LogUser(http.HandlerFunc(SignUp))))
	router.Handle("/sign_in", CORS(LogUser(http.HandlerFunc(SignIn))))
	router.Handle("/reset_password", CORS(LogUser(http.HandlerFunc(ResetPassword))))

	router.Handle("/get_user_info", CORS(Auth(IsActive(LogUser(http.HandlerFunc(GetUserInfo))))))
	router.Handle("/set_user_info", CORS(Auth(IsActive(LogUser(http.HandlerFunc(SetUserInfo))))))
	router.Handle("/set_user_role", CORS(Auth(IsActive(LogUser(http.HandlerFunc(SetUserRole))))))

	router.Handle("/change_psw", CORS(Auth(IsActive(LogUser(http.HandlerFunc(ChangePsw))))))
	// admin handlers
	router.Handle("/admin/get_user_info", CORS(Auth(isAdmin(LogUser(http.HandlerFunc(AdminGetUserInfo))))))
	router.Handle("/set_user_active", CORS(Auth(isAdmin(LogUser(http.HandlerFunc(AdminSetUserIsActive))))))

	router.Handle("/v2/get_user_info", CORS(Auth(IsActive(LogUser(http.HandlerFunc(GetUserInfoV2))))))

	return router
}

func GetUserInfoV2(resp http.ResponseWriter, req *http.Request) {

	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	id := req.URL.Query().Get("user_id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		respBody.SetError(errors.New("invalid input"))
		return
	}

	authorUserID, _ := primitive.ObjectIDFromHex(req.Header.Get(HeaderUserID))
	authorFullInfo, err := service.GetUserFullInfo(authorUserID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		respBody.SetError(err)
		return
	}

	switch authorFullInfo.Role {
	case domain.UserRoleAdmin:
		info, err := service.GetUserFullInfo(userID)
		if err != nil {
			resp.WriteHeader(http.StatusNotFound)
			respBody.SetError(err)
			return
		}
		respBody.SetData(info)
		return
	case domain.UserRoleDefault:
		if authorUserID != userID {
			resp.WriteHeader(http.StatusForbidden)
			respBody.SetError(errors.New("user is not admin to get info about someone else"))
			return
		}
		info, err := service.GetUserShortInfo(userID)
		if err != nil {
			resp.WriteHeader(http.StatusNotFound)
			respBody.SetError(err)
		}

		respBody.SetData(info)
	}
}
