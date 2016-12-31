package controllers

import "fmt"

type MainController struct {
	BaseController
}

func (this *MainController) Get()  {
	fmt.Println("main")
}