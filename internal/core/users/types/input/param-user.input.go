package input

type ParamUser struct {
	Id int `uri:"id" binding:"required"`
}
