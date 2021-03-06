package main

import (
	"1898/dll"
	"1898/middleware/logg"
	"1898/utils/log"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"1898/utils/doc"
)

func HttpRun(addr string) {

	r := mux.NewRouter()

	// user
	r.HandleFunc("/login", dll.Login).Methods("POST")
	r.HandleFunc("/rootlogin", dll.RootLogin).Methods("POST")
	r.HandleFunc("/register", dll.Register).Methods("POST")
	r.HandleFunc("/regkey", dll.CheckRegKey).Methods("POST")
	r.HandleFunc("/upload", dll.Upload).Methods("POST")

	r.HandleFunc("/user/newroot", dll.CreateRoot).Methods("POST")
	r.HandleFunc("/user/editpwd", dll.EditPassword).Methods("POST")
	r.HandleFunc("/user/findpwd", dll.FindPassword).Methods("POST")
	r.HandleFunc("/user/edit", dll.EditUser).Methods("POST")
	r.HandleFunc("/user/info", dll.GetUserByID).Methods("POST")
	r.HandleFunc("/user/byphone", dll.GetUserByPhone).Methods("POST")
	r.HandleFunc("/user/avatar", dll.Avatar).Methods("POST")
	r.HandleFunc("/user/sendsms", dll.SendSMS).Methods("POST", "GET")

	// friends
	r.HandleFunc("/friend/add", dll.AddFriend).Methods("POST")
	r.HandleFunc("/friend/list", dll.FriendsList).Methods("POST")
	r.HandleFunc("/friend/del", dll.DelFriend).Methods("POST")

	//message
	r.HandleFunc("/msg/push", dll.PushMsg).Methods("POST")
	r.HandleFunc("/msg/pull", dll.PullMsg).Methods("POST")
	r.HandleFunc("/msg/read", dll.ReadMsg).Methods("POST")

	// key
	r.HandleFunc("/key/getkey", dll.GetKey).Methods("POST")

	// event
	r.HandleFunc("/event/new", dll.NewEvent).Methods("POST")
	r.HandleFunc("/event/edit", dll.EditEvent).Methods("POST")
	r.HandleFunc("/event/info", dll.EventInfo).Methods("POST")
	r.HandleFunc("/event/reg", dll.RegEvent).Methods("POST")
	r.HandleFunc("/event/list", dll.EventList).Methods("POST")
	r.HandleFunc("/event/del", dll.DelEvent).Methods("POST")
	r.HandleFunc("/event/cancel", dll.CancelEvent).Methods("POST")
	r.HandleFunc("/event/publish", dll.MyPublishEvent).Methods("POST")
	r.HandleFunc("/event/join", dll.MyJoinEvent).Methods("POST")

	// news
	r.HandleFunc("/news/add", dll.AddNews).Methods("POST")
	r.HandleFunc("/news/one", dll.FindNews).Methods("POST")
	r.HandleFunc("/news/edit", dll.EditNews).Methods("POST")
	r.HandleFunc("/news/list", dll.NewsList).Methods("POST")
	r.HandleFunc("/news/del", dll.DelNews).Methods("POST")

	// 文件服务器

	//http.Handle("/doc/", http.StripPrefix("/doc/", http.FileServer(http.Dir("/usr/share/doc"))))

	r.Handle("/image/{name}", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	// 开启文档服务
	doc := doc.New("./dll").SetHost("192.168.1.150", "9900")
	r.HandleFunc("/doc", doc.Handler).Methods("GET")

	n := negroni.New()
	n.Use(logg.New())
	n.UseHandler(r)

	log.CLog("[TRAC] Server start listen on # %s #\n", addr)

	if err := http.ListenAndServe(":"+addr, n); err != nil {
		panic(err)
	}

}
