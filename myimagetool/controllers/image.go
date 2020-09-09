package controllers

import (
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {

	this.TplName = "upload.html"
}

func (this *UploadController) Post() {
	file, information, err := this.GetFile("file")
	defer file.Close()
	if err != nil {
		this.Ctx.WriteString("File retrieval failure")
		return
	} else {
		filename := information.Filename
		picture := strings.Split(filename, ".")
		layout := strings.ToLower(picture[len(picture)-1])

		if layout != "jpg" && layout != "png" && layout != "gif" {
			this.Ctx.WriteString("请上传符合格式的图片（png、jpg、gif）")
			return //结束整个程序，不执行保存文件
		}

		err = this.SaveToFile("file", path.Join("static/upload", filename))
		if err != nil {
			this.Ctx.WriteString("File upload failed！")
		} else {
			this.Ctx.WriteString("File upload succeed!")
		}
	}
	this.TplName = "upload.html"
}
