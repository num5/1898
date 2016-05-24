package dll

import (
	"net/http"
	"1898/dal"
	"time"
)

//@name 发送消息
func PushMsg(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	toid := r.FormValue("toid")

	if toid == "" {
		Errors(w, ErrMissParam("toid", ErrCode_MissParamToId))

		return
	}

	if !IsObjectId(toid) {
		Errors(w, ErrForbidden("toid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	// 检测是否是朋友关系
	f := new(dal.Friends)
	f.UId = uid
	f.Fid = toid
	err := f.FindByUFID()

	if err != nil || f.Agree.IsZero() {
		Errors(w, ErrForbidden("not frients", ErrCode_NotFriend))
		return
	}

	msg := r.FormValue("msg")
	if msg == "" {
		Errors(w, ErrMissParam("msg", ErrCode_MissParamMsg))

		return
	}

	if len(msg) > 140 {
		Errors(w, ErrForbidden("title must be at least 140 characters", ErrCode_StringLenErr))
		return
	}

	m := new(dal.Message)

	m.SendUId = uid
	m.GetUId = toid
	m.Msg = msg
	m.Read = 0
	m.Created = time.Now()

	err = m.AddMsg()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get news success", m)
}

//@name 拉取消息
func PullMsg(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	m := new(dal.Message)
	ms, err := m.FindGetByUid(uid)

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get news success", ms)
}