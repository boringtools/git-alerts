package logger

import "github.com/sirupsen/logrus"

func Log(text string) {
	logrus.Info(text)
}

func LogP(text string, args any) {
	logrus.Info(text, args)
}

func LogWRN(text string) {
	logrus.Warn(text)
}

func LogERR(text string) {
	logrus.Error(text)
}

func LogERRP(text string, args any) {
	logrus.Error(text, args)
}
