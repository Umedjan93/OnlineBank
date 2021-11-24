package controllers

import (
	"github.com/natefinch/lumberjack"
	_ "github.com/natefinch/lumberjack"
	"log"
)

// InitLogger - запускает логирование в файле logs.txt/ starts logging errors in the logs.txt file
func InitLogger() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "C:/Users/Win10_Game_OS/GolandProjects/OnlineBank/logs/logs.txt",
		MaxSize:    20, // megabytes
		MaxBackups: 5,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default
	})
}
