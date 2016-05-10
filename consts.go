package marketapi

const (
	ErrAPITimeout   = "timeout"
	ErrAPIMinAmount = "amount must be at least 100"
)

const (
	ActDOTA2 = "dota2"
	ActCSGO  = "csgo"
	ActTF2   = "tf2"
	ActGIFTS = "gifts"
)

const (
	CodeDOTA2 = "570"
	CodeCSGO  = "730"
	CodeTF2   = "440"
	CodeGIFTS = "753"
)

const (
	URLDota2 = "https://market.dota2.net"
	URLCsgo  = "https://market.csgo.com"
	URLTf2   = "https://tf2.tm"
	URLGifts = "https://gifts.tm"
)

const (
	URLItemDBCurrent      = "%s/itemdb/current_%s.json"
	URLItemDB             = "%s/itemdb/%s"
	URLItemInfo           = "%s/api/ItemInfo/%s_%s/%s/?key=%s"
	URLItemHistory        = "%s/api/ItemHistory/%s_%s/?key=%s"
	URLMarketTrades       = "%s/api/MarketTrades/?key=%s"
	URLTrades             = "%s/api/Trades/?key=%s"
	URLBuy                = "%s/api/Buy/%s_%s/%d/%s/?key=%s"
	URLSetPriceNew        = "%s/api/SetPrice/new_%s_%s/%d/?key=%s"
	URLRemoveAll          = "%s/api/RemoveAll/?key=%s"
	URLSetPrice           = "%s/api/SetPrice/%s/%d/?key=%s"
	URLPingPong           = "%s/api/PingPong/?key=%s"
	URLItemRequest        = "%s/api/ItemRequest/%s/%s/?key=%s"
	URLOperationHistory   = "%s/api/OperationHistory/%d/%d/?key=%s"
	URLGetMoney           = "%s/api/GetMoney/?key=%s"
	URLTest               = "%s/api/Test/?key=%s"
	URLInventoryStatus    = "%s/api/InventoryStatus/?key=%s"
	URLUpdateInventory    = "%s/api/UpdateInventory/?key=%s"
	URLGetToken           = "%s/api/GetToken/?key=%s"
	URLSetToken           = "%s/api/SetToken/%s/?key=%s"
	URLQuickItems         = "%s/api/QuickItems/?key=%s"
	URLQuickBuy           = "%s/api/QuickBuy/%s/?key=%s"
	URLGetOrders          = "%s/api/GetOrders/?key=%s"
	URLInsertOrder        = "%s/api/InsertOrder/%s/%s/%d/%s/?key=%s"
	URLUpdateOrder        = "%s/api/UpdateOrder/%s/%s/%d/?key=%s"
	URLDeleteOrders       = "%s/api/DeleteOrders/?key=%s"
	URLGetNotifications   = "%s/api/GetNotifications/?key=%s"
	URLUpdateNotification = "%s/api/UpdateNotification/%s/%s/%d/?key=%s"
	URLGetWSAuth          = "%s/api/GetWSAuth/?key=%s"
)
