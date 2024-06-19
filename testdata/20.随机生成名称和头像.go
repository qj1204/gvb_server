package main

import (
	"fmt"
	"github.com/DanPlayer/randomname"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
	"image/png"
	"os"
	"path"
	"unicode/utf8"
)

func main() {
	// 随机名称
	//name := randomname.GenerateName()
	//fmt.Println(name)
	// 头像生成
	// 生成一个100*100大小的以字母‘A’为图像的头像
	//img, _ := letteravatar.Draw(100, 'A', nil)
	//file, _ := os.Create("A.png")
	//png.Encode(file, img)
	//names := []rune(name)
	//DrawImage(string(names[0]), "uploads/chat_avatar")
	GenerateNameAvatar()
}

func GenerateNameAvatar() {
	dir := "uploads/chat_avatar"
	for _, s := range randomname.AdjectiveSlice {
		DrawImage(string([]rune(s)[0]), dir)
	}
	for _, s := range randomname.PersonSlice {
		DrawImage(string([]rune(s)[0]), dir)
	}
}

func DrawImage(name string, dir string) {
	fontFile, err := os.ReadFile("uploads/system/方正清刻本悦宋简体.TTF")
	font, err := freetype.ParseFont(fontFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	options := &letteravatar.Options{
		Font: font,
	}
	// 绘制文字
	firstLetter, _ := utf8.DecodeRuneInString(name)
	img, err := letteravatar.Draw(140, firstLetter, options)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 保存
	filePath := path.Join(dir, name+".png")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err)
		return
	}
}
