package controller

type ControllerI interface {
	GetPages(args []string) ([]Response, error)
	BeautyPrint(args []string)
}
