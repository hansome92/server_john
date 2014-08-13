// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApiServers struct {}
var ApiServers tApiServers


func (_ tApiServers) All(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiServers.All", args).Url
}

func (_ tApiServers) Show(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("ApiServers.Show", args).Url
}

func (_ tApiServers) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiServers.Create", args).Url
}

func (_ tApiServers) Action(
		method string,
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "method", method)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("ApiServers.Action", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


