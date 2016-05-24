package dll

import (
	"net/http"
	"encoding/json"
)

const (
	ErrCode_InternalServer = 500

	ErrCode_MissParamToken = iota + 40010
	ErrCode_TokenExpired
	ErrCode_MissParamUid
	ErrCode_UidNotObjectId
	//
	ErrCode_UserMissParamKey
	ErrCode_UserKeyNotFound
	ErrCode_UserKeyUsed
	ErrCode_UserMissParamPhone
	ErrCode_UserPhoneNotMatch
	ErrCode_UserMissParamPassword
	ErrCode_UpdateKeyErr
	ErrCode_NickNameErr
	//
	ErrCode_EventMissParamTitle
	ErrCode_EventMissParamDetail
	ErrCode_EventMissParamAddr
	ErrCode_EventMissParamStart
	ErrCode_EventMissParamPrice
	ErrCode_EventMissParamTotal
	ErrCode_EventDetailLenNotEnough
	ErrCode_UserAlreadySignUp
	ErrCode_UserNotSignUp
	ErrCode_EnrollmentFull
	ErrCode_EventNotCreateUser
	ErrCode_EventOrganizer
	//
	ErrCode_MissParamEid
	ErrCode_EidNotObjectId
	//
	ErrCode_AddFriendErr
	ErrCode_MissParamFid
	ErrCode_FriendRepeat

)

type ERR_CODE int

const (
	CODE_SUCCESS ERR_CODE = iota + 2000 // 成功

	CODE_FAIL ERR_CODE = iota + 4000 // 失败
)

type JSON struct {
	ID     string      `json:"id"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	ErrCode int         `json:"errcode"`
}

func Push(w http.ResponseWriter, msg string, data interface{}) {
	j := &JSON{
		ID:"ok",
		Msg:msg,
		Data:data,
		ErrCode:0,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(j)

}

type Error struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`

	ErrCode int   `json:"errcode"`
}

func (err Error) Error() string {
	return err.Detail
}

func Errors(w http.ResponseWriter, err *Error) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(err)
}

func ErrMissParam(param string, errcode int) *Error {
	return &Error{
		"miss_param",
		"Miss Param",
		"miss query param: " + param,
		errcode,
	}
}

func ErrForbidden(detail string, errcode int) *Error {
	return &Error{
		"forbidden",
		"Forbidden",
		"Forbidden: " + detail,
		errcode,
	}
}

func ErrInternalServer(detail string, errcode int) *Error {
	return &Error{
		"internal_server_error",
		"Internal Server Error",
		detail,
		errcode,
	}
}

func ErrUnauthorized(detail string, errcode int) *Error {
	return &Error{
		"unauthorized",
		"Unauthorized",
		detail,
		errcode,
	}
}
