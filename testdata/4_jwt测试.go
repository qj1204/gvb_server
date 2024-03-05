package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwt"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	//token, err := jwt.GenerateToken(jwt.JwtPayLoad{
	//	UserID:   2,
	// NickName: "小新qj",
	//	Role:     1,
	//})
	//fmt.Println(token, err)

	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJuaWNrX25hbWUiOiLlsI_mlrBxaiIsInJvbGUiOjEsImV4cCI6MTcwOTQ1NTkyOC4wNTkyMSwiaXNzIjoicWlhbmppbiJ9.Jye9nMPXNYBS1Y8OrufTisuD_ZHzkSaCr66w2_g2hT4")
	fmt.Println(claims, err)
}
