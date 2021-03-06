package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	"os"
)

func IconLoadFromBox(filename string, size walk.Size) *walk.Icon {
	body, err := BoxFile().Bytes(filename)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	dir := DEFAULT_HOME + "\\icon\\"
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 644)
		if err != nil {
			logs.Error(err.Error())
			return nil
		}
	}
	filepath := dir + filename
	err = SaveToFile(filepath, body)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	icon, err := walk.NewIconFromFileWithSize(filepath, size)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	return icon
}


func IconLoadImageFromBox(filename string) walk.Image {
	body, err := BoxFile().Bytes(filename)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	dir := DEFAULT_HOME + "\\image\\"
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 644)
		if err != nil {
			logs.Error(err.Error())
			return nil
		}
	}
	filepath := dir + filename
	err = SaveToFile(filepath, body)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	image, err := walk.NewImageFromFile(filepath)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	return image
}

var ICON_Main          *walk.Icon
var ICON_Main_Mini     *walk.Icon
var ICON_Network_Flow  *walk.Icon

var ICON_TOOL_ADD      *walk.Icon
var ICON_TOOL_DEL      *walk.Icon
var ICON_TOOL_DOWNLOAD *walk.Icon
var ICON_TOOL_LOADING  *walk.Icon
var ICON_TOOL_SETTING  *walk.Icon
var ICON_TOOL_STOP     *walk.Icon
var ICON_TOOL_RESERVE  *walk.Icon
var ICON_TOOL_RESERVEDL *walk.Icon

var ICON_STATUS_STOP  *walk.Icon
var ICON_STATUS_DONE  *walk.Icon
var ICON_STATUS_WAIT  *walk.Icon
var ICON_STATUS_RESERVER  *walk.Icon
var ICON_STATUS_LOAD  *walk.Icon

var ICON_Max_Size = walk.Size{
	Width: 128, Height: 128,
}

var ICON_Tool_Size = walk.Size{
	Width: 64, Height: 64,
}

var ICON_Min_Size = walk.Size{
	Width: 16, Height: 16,
}

func IconInit() error {
	ICON_Main = IconLoadFromBox("main.ico", ICON_Max_Size)
	ICON_Main_Mini = IconLoadFromBox("mainmini.ico", ICON_Min_Size)
	ICON_Network_Flow = IconLoadFromBox("flow.ico", ICON_Min_Size)

	ICON_TOOL_ADD = IconLoadFromBox("add.ico", ICON_Tool_Size)
	ICON_TOOL_DEL = IconLoadFromBox("delete.ico", ICON_Tool_Size)
	ICON_TOOL_DOWNLOAD = IconLoadFromBox("download.ico", ICON_Tool_Size)
	ICON_TOOL_LOADING = IconLoadFromBox("loading.ico", ICON_Tool_Size)
	ICON_TOOL_SETTING = IconLoadFromBox("setting.ico", ICON_Tool_Size)
	ICON_TOOL_STOP = IconLoadFromBox("stop.ico", ICON_Tool_Size)
	ICON_TOOL_RESERVE = IconLoadFromBox("reserve.ico", ICON_Tool_Size)
	ICON_TOOL_RESERVEDL = IconLoadFromBox("reservedl.ico", ICON_Tool_Size)

	ICON_STATUS_STOP = IconLoadFromBox("stop.ico", ICON_Min_Size)
	ICON_STATUS_DONE = IconLoadFromBox("done.ico", ICON_Min_Size)
	ICON_STATUS_WAIT = IconLoadFromBox("wait.ico", ICON_Min_Size)
	ICON_STATUS_RESERVER = IconLoadFromBox("reserve.ico", ICON_Min_Size)
	ICON_STATUS_LOAD = IconLoadFromBox("loading.ico", ICON_Min_Size)

	return nil
}

