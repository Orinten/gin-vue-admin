package myPkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.
	MyApi
	// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
}

var (
	myApiService			= service.ServiceGroupApp.MypkgServiceGroup.MyApi
)
