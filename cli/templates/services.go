package templates

var Services = `package services

import (
	"go.uber.org/zap"
)

var localLogger, _ = zap.NewProduction().Named("services")
/* example
var Card = newCardManager(localLogger)
var Cache = newCacheManager(localLogger)
var User = newUserService(localLogger)
var Inventory = newInventoryService(localLogger)
var Link = NewLinkService(localLogger)
var Upload = NewUploadService(localLogger)
*/ 
`
