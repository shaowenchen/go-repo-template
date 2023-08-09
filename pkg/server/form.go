package server

type FormUser struct {
	Nickname string `form:"nickname" json:"nickname"`
	Code     string `form:"code" json:"code" binding:"required"`
	Password string `form:"password" json:"password" binding:"max=20,min=10"`
}
