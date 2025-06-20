package input

type ParamPost struct {
	Id int `uri:"id" binding:"required"`
}
