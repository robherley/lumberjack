package util

import "github.com/sirupsen/logrus"

// Log is just a single logrus instance, incase we want to expand the fields
// later
var Log = logrus.New()
