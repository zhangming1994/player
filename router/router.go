package router

import (
	"net/http"
	"player/src"
)

func Router(){
	//首页路由
	http.HandleFunc("/",src.Index)

}