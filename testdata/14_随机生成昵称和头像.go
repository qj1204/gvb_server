package main

import (
	"github.com/DanPlayer/randomname"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
	"github.com/sirupsen/logrus"
	"image/png"
	"os"
	"path"
	"unicode/utf8"
)

func main() {
	// 随机昵称
	//name := randomname.GenerateName()
	//r := []rune(name)
	// 绘制头像
	GenerateNameAvatar("static/chat_avatar")
}

func GenerateNameAvatar(dir string) {
	for _, s := range randomname.AdjectiveSlice {
		r := []rune(s)
		DrawImage(string(r[0]), dir)
	}
	for _, s := range randomname.PersonSlice {
		r := []rune(s)
		DrawImage(string(r[0]), dir)
	}
}

func DrawImage(name string, dir string) {
	fontFile, _ := os.ReadFile("static/system/方正清刻本悦宋简体.TTF")
	font, _ := freetype.ParseFont(fontFile)
	options := &letteravatar.Options{
		Font: font,
	}
	// 绘制文字
	firstLetter, _ := utf8.DecodeRuneInString(name)

	img, err := letteravatar.Draw(140, firstLetter, options)
	if err != nil {
		logrus.Fatal(err)
	}

	// 保存
	filePath := path.Join(dir, name+".png")
	file, err := os.Create(filePath)
	if err != nil {
		logrus.Fatal(err)
	}

	err = png.Encode(file, img)
	if err != nil {
		logrus.Fatal(err)
	}
}
