package log

import "github.com/sirupsen/logrus"

func I(args ...interface{}) {
	logrus.Infoln(args)
}

func D(args ...interface{}) {
	logrus.Debugln(args)
}

func E(args ...interface{}) {
	logrus.Errorln(args)
}
