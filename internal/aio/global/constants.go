package global

import (
	"aresbot/internal/aio/models"
	"aresbot/internal/aio/shops"
)

var WithQuickTasks = false
var SettingQuicktaskProfileName = ""
var SettingQuicktaskAmount = 1
var SettingQuicktaskWithProxy = false
var SettingQuicktaskProfile models.Profile
var ShopRunner shops.BotRunner
