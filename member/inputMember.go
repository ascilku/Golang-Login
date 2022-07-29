package member

type InputMember struct {
	Nama     string `json:"nama" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Nama     string `json:"nama" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
