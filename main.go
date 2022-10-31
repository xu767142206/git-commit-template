package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

var typeInfo = []string{
	"feat:新功能",
	"fix:修复问题",
	"docs:修改文档",
	"style:修改代码格式，不影响代码逻辑",
	"refactor:重构代码，理论上不影响现有功能",
	"perf:提升性能",
	"test:增加修改测试用例",
	"chore:修改工具相关（包括但不限于文档、代码生成等）",
	"deps:升级依赖",
}

func main() {

	a := app.New()
	w := a.NewWindow("生成commit 信息")

	typeOpinos := widget.NewSelect(typeInfo, func(s string) {
		split := strings.Split(s, ":")
		fmt.Println(split[0])
	})
	typeOpinos.SetSelectedIndex(0)

	//文件范围
	scopeEntry := widget.NewEntry()
	scopeEntry.SetPlaceHolder("scope: controller,model....")

	//一句话描述
	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("subject: 这是干什么的")

	//body补充 subject，适当增加原因、目的等相关因素，也可不写。
	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.MultiLine = true

	/**
	footer
	当有非兼容修改时可在这里描述清楚
	关联相关 issue，如 Closes #1, Closes #2, #3
	如果功能点有新增或修改的，还需要关联 chair-handbook 和 chair-init 的 MR，如 chair/doc!123
	*/
	footerEntry := widget.NewMultiLineEntry()
	footerEntry.MultiLine = true

	createEntry := widget.NewMultiLineEntry()
	createEntry.MultiLine = true
	w.SetContent(container.NewVBox(
		container.NewCenter(widget.NewLabel("生成commit的信息")),
		widget.NewLabel("提交type:"),
		typeOpinos,
		widget.NewLabel("scope:修改文件的范围<包括但不限于 doc, middleware, proxy, core, config>"),
		scopeEntry,
		widget.NewLabel("一句话描述提交的功能"),
		subjectEntry,
		widget.NewLabel("补充适当增加原因、目的等相关因素，也可不写"),
		bodyEntry,
		widget.NewLabel("当有非兼容修改时可在这里描述清楚\n\t1、关联相关 issue，如 Closes #1, Closes #2, #3 \n\t2、如果功能点有新增或修改的，还需要关联 chair-handbook 和 chair-init 的 MR，如 chair/doc!123"),
		footerEntry,
		widget.NewButton("创建", func() {
			/**
			<type>(<scope>): <subject>
			<BLANK LINE>
			<body>
			<BLANK LINE>
			<footer>
			*/
			split := strings.Split(typeOpinos.Selected, ":")

			str := fmt.Sprintf("%s %s:%s\n%s\n%s",
				split[0],
				scopeEntry.Text,
				subjectEntry.Text,
				bodyEntry.Text,
				footerEntry.Text,
			)
			createEntry.SetText(str)

		}),
		layout.NewSpacer(),
		widget.NewLabel("生成的commit信息:"),
		container.NewMax(createEntry),
		layout.NewSpacer(),
	))

	w.Resize(fyne.NewSize(320, 600))
	w.ShowAndRun()
}

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}
