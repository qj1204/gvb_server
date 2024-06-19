package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "xiaoxin",
		NickName: "xxx",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZlbmdmZW5nIiwibmlja19uYW1lIjoieHh4Iiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNjc2NzA4MzE2LjU1NjE4NywiaXNzIjoiMTIzNCJ9.bwbPAOVG5kxUeZZOkdfaHMOA86l_Vsu3UfkDt3Bi90A")
	fmt.Println(claims, err)

}
