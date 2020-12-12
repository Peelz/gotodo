package controllers

type TaskHandler interface {
	Get()
	Create()
	List()
	Update()
	Delete()
}
