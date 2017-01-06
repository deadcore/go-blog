package ping

type pingController struct {}

func Controller() *pingController {
	return &pingController{}
}