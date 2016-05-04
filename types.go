package marketapi

type Description struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Tag struct {
	InternalName string `json:"internal_name"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Color        string `json:"color"`
	CategoryName string `json:"category_name"`
}

type Offer struct {
	Price   string `json:"price"`
	Count   string `json:"count"`
	MyCount string `json:"my_count"`
}

type BuyOffer struct {
	OPrice  string `json:"o_price"`
	C       string `json:"c"`
	MyCount string `json:"my_count"`
}

type APIItemInfo struct {
	ClassID             string        `json:"classid"`
	InstanceID          string        `json:"instanceid"`
	OurMarketInstanceID string        `json:"our_market_instanceid"`
	MarketName          string        `json:"market_name"`
	Name                string        `json:"name"`
	MarketHashName      string        `json:"market_hash_name"`
	Rarity              string        `json:"rarity"`
	Quality             string        `json:"quality"`
	Type                string        `json:"type"`
	Mtype               string        `json:"mtype"`
	Slot                string        `json:"slot"`
	Description         []Description `json:"description"`
	Tags                []Tag         `json:"tags"`
	Hash                string        `json:"hash"`
	MinPrice            string        `json:"min_price"`
	Offers              []Offer       `json:"offers"`
	BuyOffers           []BuyOffer    `json:"buy_offers"`
}

type History struct {
	LPrice string `json:"l_price"`
	LTime  string `json:"l_time"`
}

type APIItemHistory struct {
	Success bool      `json:"success"`
	Max     int64     `json:"max"`
	Min     int64     `json:"min"`
	Average int64     `json:"average"`
	Number  int64     `json:"number"`
	History []History `json:"history"`
}

type APIMarketTrades struct {
	//TODO
}

type Trade struct {
	UIID             string  `json:"ui_id"`
	IName            string  `json:"i_name"`
	IMarketName      string  `json:"i_market_name"`
	INameColor       string  `json:"i_name_color"`
	IRarity          string  `json:"i_rarity"`
	IDescriptions    string  `json:"i_descriptions"`
	UIStatus         string  `json:"ui_status"`
	HeName           string  `json:"he_name"`
	UIPrice          float64 `json:"ui_price"`
	IClassID         string  `json:"i_classid"`
	IInstanceID      string  `json:"i_instanceid"`
	UIRealInstance   string  `json:"ui_real_instance"`
	IQuality         string  `json:"i_quality"`
	IMarketHashName  string  `json:"i_market_hash_name"`
	IMarketPrice     float64 `json:"i_market_price"`
	Position         int64   `json:"position"`
	MinPrice         float64 `json:"min_price"`
	UIBid            string  `json:"ui_bid"`
	UIAsset          string  `json:"ui_asset"`
	Type             string  `json:"type"`
	UIPriceText      string  `json:"ui_price_text"`
	MinPriceText     bool    `json:"min_price_text"`
	IMarketPriceText string  `json:"i_market_price_text"`
	OfferLiveTime    int64   `json:"offer_live_time"`
	Placed           string  `json:"placed"`
}

type APITrades []Trade

type APIBuy struct {
	Result string `json:"result"`
	ID     string `json:"id"`
}

type APISetPrice struct {
	Result    int64   `json:"result"`
	ItemID    int64   `json:"item_id"`
	Price     float64 `json:"price"`
	PriceText string  `json:"price_text"`
	Status    string  `json:"status"`
	Position  int64   `json:"position"`
	Success   bool    `json:"success"`
}

type APIRemoveAll struct {
	NumDeletedItems int64 `json:"num_deleted_items"`
	Success         bool  `json:"success"`
}

type SetPrice struct{}
type APIPingPong struct {
	Ping    string `json:"ping"`
	Success bool   `json:"success"`
}
type APIItemRequest struct {
	Success bool   `json:"success"`
	Trade   string `json:"trade"`
	Nick    string `json:"nick"`
	Botid   int64  `json:"botid"`
	Profile string `json:"profile"`
	Secret  string `json:"secret"`
	Items   interface{}
}

type OHistory struct {
	HID            string `json:"h_id"`
	HEvent         string `json:"h_event"`
	HTime          string `json:"h_time"`
	HEventID       string `json:"h_event_id"`
	Join           int64  `json:"join"`
	App            string `json:"app"`
	ID             string `json:"id"`
	ClassID        string `json:"classid"`
	InstanceID     string `json:"instanceid"`
	Quality        string `json:"quality"`
	NameColor      string `json:"name_color"`
	MarketName     string `json:"market_name"`
	MarketHashName string `json:"market_hash_name"`
	Paid           string `json:"paid"`
	Recieved       string `json:"recieved"`
	Stage          string `json:"stage"`
	Item           string `json:"item"`
	Flags          string `json:"flags"`
}

type APIOperationHistory struct {
	Success bool       `json:"success"`
	History []OHistory `json:"history"`
}

type APIGetMoney struct {
	Money int64 `json:"money"`
}

type Status struct {
	UserToken    bool `json:"user_token"`
	TradeCheck   bool `json:"trade_check"`
	SiteOnline   bool `json:"site_online"`
	SiteNotmpban bool `json:"site_notmpban"`
}

type APIResponse struct {
	RespError  interface{} `json:"error"`
	RespResult interface{} `json:"result"`
}

type APITest struct {
	Success bool   `json:"success"`
	Status  Status `json:"status"`
}

type APIInventoryStatus struct {
	Success bool   `json:"success"`
	IStatus string `json:"i_status"`
	ITime   string `json:"i_time"`
}

type APIUpdateInventory struct {
	Success bool `json:"success"`
}

type APIGetToken struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type APISetToken struct {
	Success bool `json:"success"`
}

type Item struct {
	UIID                string `json:"ui_id"`
	LPaid               string `json:"l_paid"`
	IClassID            string `json:"i_classid"`
	IInstanceID         string `json:"i_instanceid"`
	IMarketHashName     string `json:"i_market_hash_name"`
	IRarity             string `json:"i_rarity"`
	IMarket_name        string `json:"i_market_name"`
	IName               string `json:"i_name"`
	IQuality            string `json:"i_quality"`
	INameColor          string `json:"i_name_color"`
	HEName              string `json:"he_name"`
	IDescriptionsString string `json:"i_descriptions"`
	ITagsString         string `json:"i_tags"`
	IDescriptions       []Description
	ITags               []Tag
}

type APIQuickItems struct {
	Success bool   `json:"success"`
	Items   []Item `json:"items"`
}

type APIQuickBuy struct {
	Success bool `json:"success"`
}

type Order struct {
	IClassID        string `json:"i_classid"`
	IInstanceID     string `json:"i_instanceid"`
	IMarketHashName string `json:"i_market_hash_name"`
	IMarketName     string `json:"i_market_name"`
	OPrice          string `json:"o_price"`
	OState          string `json:"o_state"`
}

type APIGetOrders struct {
	Success bool `json:"success"`
	Orders  []Order
}

type APIInsertOrder struct {
	Success bool `json:"success"`
}

type APIUpdateOrder struct {
	Success bool `json:"success"`
}

type APIDeleteOrders struct {
	Success       bool  `json:"success"`
	DeletedOrders int64 `json:"deleted_orders"`
}

type Notification struct {
	IClassid        string `json:"i_classid"`
	IInstanceid     string `json:"i_instanceid"`
	IMarketHashName string `json:"i_market_hash_name"`
	IMarketName     string `json:"i_market_name"`
	NVal            string `json:"n_val"`
}

type APIGetNotifications struct {
	Success       bool           `json:"success"`
	Notifications []Notification `json:"Notifications"`
}

type APIUpdateNotification struct {
	Success bool `json:"success"`
}

type APIGetWSAuth struct {
	WSAuth  string `json:"wsAuth"`
	Success bool   `json:"success"`
}

type API struct {
	Key    string
	Action string
	URL    string
	Lang   string // ru or en
}
