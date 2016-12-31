package controllers

import "time"

type ReportController struct {
	BaseController
}

type Warning struct {
	DomName string
	Time    time.Time
	Info    string
	Level   string
}

func (this *ReportController) Warn() (warning Warning) {
	//发出告警
	return
}
