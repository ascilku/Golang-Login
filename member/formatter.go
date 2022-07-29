package member

type formatter struct {
	Nama     string `json:"nama"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func Formatter(member Member, token string) formatter {
	formatter := formatter{}
	formatter.Nama = member.Nama
	formatter.Password = member.Password
	formatter.Token = token

	return formatter
}
