package deanio

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

var label *gtk.Label
var controller_lbl string
var memory_lbl string

func startUI() {
	controller_lbl = ""
	memory_lbl = ""
	gtk.Init(nil)
	gdk.ThreadsInit()

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")
	label = gtk.NewLabel("Controllers: Memory: ")
	window.Add(label)

	window.SetSizeRequest(500, 100)
	window.ShowAll()

	gdk.ThreadsEnter()
	gtk.Main()
	gdk.ThreadsLeave()
}

func updateController(text string) {
	controller_lbl = text
	gdk.ThreadsEnter()
	label.SetLabel("Controller: " + controller_lbl + " Memory: " + memory_lbl)
	gdk.ThreadsLeave()
}

func updateMemory(text string) {
	memory_lbl = text
	gdk.ThreadsEnter()
	label.SetLabel("Controller: " + controller_lbl + " Memory: " + memory_lbl)
	gdk.ThreadsLeave()
}
