package controllers

type AnalysisControllers interface {
	Check() error
}
type AnalysisController struct {
	BaseController
	AnalysisControllers
}

func (this *AnalysisController) Check() (err error) {
	//todo 检查
	return
}
