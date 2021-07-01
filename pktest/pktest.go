package pktest

import "github.com/sirupsen/logrus"

func Seelog(){
	logrus.StandardLogger().Info("I am log from package")
}