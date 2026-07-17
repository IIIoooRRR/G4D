package templates

var Init = `package commands

import "proSofi/backend/services"

var logger = services.AbstractLogger.Named("commands")
// use for configuration/incapsulate data in commands/
`
