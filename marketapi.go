package marketapi

import (
    "errors"
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "reflect"
    "strconv"
)

func (i *Item) reload() {
    json.Unmarshal([]byte(i.IDescriptionsString), &i.IDescriptions)
    json.Unmarshal([]byte(i.ITagsString), &i.ITags)
}

func (r *APIResponse) Success() bool {
    if r.RespError != nil {
        return false
    }
    return true
}

func (r *APIResponse) Error() string {
    result, err := "", ""
    if r.RespError != nil {
        switch reflect.ValueOf(r.RespError).Type().String() {
        case "float64":
            err = strconv.FormatFloat(r.RespError.(float64), 'f', 0, 32)
        default:
            err = r.RespError.(string)
        }
        err = fmt.Sprintf("Error: %s", err)
    }
    if r.RespResult != nil {
        result = fmt.Sprintf("\tResult: %s", r.RespResult.(string))
    }
    return err + result
}

func makeGet(url string) ([]byte, error){
    resp, err := http.Get(url)
    if err != nil {
        return []byte{}, err
    }
    defer resp.Body.Close()
    if resp.StatusCode == http.StatusGatewayTimeout {
        return []byte{}, errors.New(ErrAPITimeout)
    }
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return []byte{}, err
    }
    var apiResponse APIResponse
    json.Unmarshal(bytes, &apiResponse)
    if !apiResponse.Success() {
        return []byte{}, errors.New(apiResponse.Error())
    }
    return bytes, nil
}

//ItemInfo - Информация и предложения о продаже конкретной вещи.
func (a *API) ItemInfo(classid string, instanceid string) (APIItemInfo, error) {
    bytes, err := makeGet(fmt.Sprintf(URLItemInfo, a.URL, classid, instanceid, a.Lang, a.Key))
    if err != nil {
        return APIItemInfo{}, err
    }
    var apiItemInfo APIItemInfo
    json.Unmarshal(bytes, &apiItemInfo)
    return apiItemInfo, nil
}

//ItemHistory - Информация о ценах и о последних 500 покупках конкретной вещи.
func (a *API) ItemHistory(classid string, instanceid string) (APIItemHistory, error) {
    bytes, err := makeGet(fmt.Sprintf(URLItemHistory, a.URL, classid, instanceid, a.Key))
    if err != nil {
        return APIItemHistory{}, err
    }
    var apiItemHistory APIItemHistory
    json.Unmarshal(bytes, &apiItemHistory)
    return apiItemHistory, nil
}

//MarketTrades - Список трейдов, которые маркет отправил вам и они активны в данный момент.
func (a *API) MarketTrades() (APIMarketTrades, error){
    bytes, err := makeGet(fmt.Sprintf(URLMarketTrades, a.URL, a.Key))
    if err != nil {
        return APIMarketTrades{}, err
    }
    var apiMarketTrades APIMarketTrades
    json.Unmarshal(bytes, &apiMarketTrades)
    return apiMarketTrades, nil
}

//Trades - Cписок предметов со страницы "Мои вещи".
// "UIStatus" = 1 - Вещь выставлена на продажу.
// "UIStatus" = 2 - Вы продали вещь и должны ее передать боту.
// "UIStatus" = 3 - Ожидание передачи боту купленной вами вещи от продавца.
// "UIStatus" = 4 - Вы можете забрать купленную вещь.
func (a *API) Trades() (APITrades, error){
    bytes, err := makeGet(fmt.Sprintf(URLTrades, a.URL, a.Key))
    if err != nil {
        return APITrades{}, err
    }
    var apiTrades APITrades
    json.Unmarshal(bytes, &apiTrades)
    return apiTrades, nil
}

//Buy - Покупка предмета.
//classid и instanceid - идентификаторы предмета, который можно найти в ссылке на предмет.
//BASE_URL/item/57939770-57939888-Treasure+Key/ 57939770 - classid, 57939888 - instanceid.
//price - цена в копейках(целое число), уже какого-то выставленного лота, или можно указать любую сумму больше цены самого дешевого лота, во втором случае купится предмет по самой низкой цене.
//hash - md5 от описания предмета. Вы можете найти его в ответе метода ItemInfo. Это введено, чтобы вы были уверены в покупке именно той вещи, которую покупаете. Если для вас это не интересно, просто пришлите пустую строку.
func (a *API) Buy(classid string, instanceid string, price int64, hash string) (APIBuy, error){
    bytes, err := makeGet(fmt.Sprintf(URLBuy, a.URL, classid, instanceid, price, hash, a.Key))
    if err != nil {
        return APIBuy{}, err
    }
    var apiBuy APIBuy
    json.Unmarshal(bytes, &apiBuy)
    if apiBuy.ID == "" {
        return APIBuy{}, errors.New(apiBuy.Result)
    }
    return apiBuy, nil
}

func (a *API) SetPriceNew(classid string, instanceid string, price int64) (APISetPrice, error){
    if price < MinAmount {
        return APISetPrice{}, errors.New(ErrAPIMinAmount)
    }
    bytes, err := makeGet(fmt.Sprintf(URLSetPriceNew, a.URL, classid, instanceid, price, a.Key))
    if err != nil {
        return APISetPrice{}, err
    }
    var apiSetPrice APISetPrice
    json.Unmarshal(bytes, &apiSetPrice)
    return apiSetPrice, nil
}

func (a *API) RemoveAll() (APIRemoveAll, error){
    bytes, err := makeGet(fmt.Sprintf(URLRemoveAll, a.URL, a.Key))
    if err != nil {
        return APIRemoveAll{}, err
    }
    var apiRemoveAll APIRemoveAll
    json.Unmarshal(bytes, &apiRemoveAll)
    return apiRemoveAll, nil
}

func (a *API) SetPrice(itemid string, price int64) (APISetPrice, error){
    if price < MinAmount {
        return APISetPrice{}, errors.New(ErrAPIMinAmount)
    }
    bytes, err := makeGet(fmt.Sprintf(URLSetPrice, a.URL, itemid, price, a.Key))
    fmt.Println(string(bytes))
    if err != nil {
        return APISetPrice{}, err
    }
    var apiSetPrice APISetPrice
    json.Unmarshal(bytes, &apiSetPrice)
    return apiSetPrice, nil
}

func (a *API) PingPong() (APIPingPong, error){
    bytes, err := makeGet(fmt.Sprintf(URLPingPong, a.URL, a.Key))
    fmt.Println(string(bytes))
    if err != nil {
        return APIPingPong{}, err
    }
    var apiPingPong APIPingPong
    json.Unmarshal(bytes, &apiPingPong)
    return apiPingPong, nil
}

func (a *API) ItemRequest(act string, botid string) (APIItemRequest, error){
    // act in or out
    fmt.Println(fmt.Sprintf(URLItemRequest, a.URL, act, botid, a.Key))
    bytes, err := makeGet(fmt.Sprintf(URLItemRequest, a.URL, act, botid, a.Key))
    if err != nil {
        return APIItemRequest{}, err
    }
    fmt.Println(string(bytes))
    var apiItemRequest APIItemRequest
    json.Unmarshal(bytes, &apiItemRequest)
    return apiItemRequest, nil
}

func (a *API) OperationHistory(startTime int64, endTime int64) (APIOperationHistory, error){
    bytes, err := makeGet(fmt.Sprintf(URLOperationHistory, a.URL, startTime, endTime, a.Key))
    if err != nil {
        return APIOperationHistory{}, err
    }
    var apiOperationHistory APIOperationHistory
    json.Unmarshal(bytes, &apiOperationHistory)
    return apiOperationHistory, nil
}

func (a *API) GetMoney() (APIGetMoney, error){
    bytes, err := makeGet(fmt.Sprintf(URLGetMoney, a.URL, a.Key))
    if err != nil {
        return APIGetMoney{}, err
    }
    var apiGetMoney APIGetMoney
    json.Unmarshal(bytes, &apiGetMoney)
    return apiGetMoney, nil
}

func (a *API) Test() (APITest, error){
    bytes, err := makeGet(fmt.Sprintf(URLTest, a.URL, a.Key))
    if err != nil {
        return APITest{}, err
    }
    var apiTest APITest
    json.Unmarshal(bytes, &apiTest)
    return apiTest, nil
}

func (a *API) InventoryStatus() (APIInventoryStatus, error){
    bytes, err := makeGet(fmt.Sprintf(URLInventoryStatus, a.URL, a.Key))
    if err != nil {
        return APIInventoryStatus{}, err
    }
    var apiInventoryStatus APIInventoryStatus
    json.Unmarshal(bytes, &apiInventoryStatus)
    return apiInventoryStatus, nil
}

func (a *API) UpdateInventory() (APIUpdateInventory, error){
    bytes, err := makeGet(fmt.Sprintf(URLUpdateInventory, a.URL, a.Key))
    if err != nil {
        return APIUpdateInventory{}, err
    }
    var apiUpdateInventory APIUpdateInventory
    json.Unmarshal(bytes, &apiUpdateInventory)
    return apiUpdateInventory, nil
}

//GetToken - Получить установленный токен.
func (a *API) GetToken() (APIGetToken, error){
    bytes, err := makeGet(fmt.Sprintf(URLGetToken, a.URL, a.Key))
    if err != nil {
        return APIGetToken{}, err
    }
    var apiGetToken APIGetToken
    json.Unmarshal(bytes, &apiGetToken)
    return apiGetToken, nil
}

func (a *API) SetToken(newToken string) (APISetToken, error){
    bytes, err := makeGet(fmt.Sprintf(URLSetToken, a.URL, newToken, a.Key))
    if err != nil {
        return APISetToken{}, err
    }
    var apiSetToken APISetToken
    json.Unmarshal(bytes, &apiSetToken)
    return apiSetToken, nil
}

//QuickItems - Получить список предметов для моментальной покупки с страницы BASE_URL/quick/
func (a *API) QuickItems() (APIQuickItems, error){
    bytes, err := makeGet(fmt.Sprintf(URLQuickItems, a.URL, a.Key))
    if err != nil {
        return APIQuickItems{}, err
    }
    var apiQuickItems APIQuickItems
    json.Unmarshal([]byte(string(bytes)), &apiQuickItems)
    for i := range apiQuickItems.Items{
        apiQuickItems.Items[i].reload()
    }
    return apiQuickItems, nil
}

//QuickBuy - Моментально купить предмет из метода QuickItems (За цену, которая указана в параметре "LPaid" в копейках). Через секунду его можно будет забрать через метод ItemRequest.
func (a *API) QuickBuy(uiID string) (APIQuickBuy, error){
    bytes, err := makeGet(fmt.Sprintf(URLQuickBuy, a.URL, uiID, a.Key))
    if err != nil {
        return APIQuickBuy{}, err
    }
    var apiQuickBuy APIQuickBuy
    json.Unmarshal([]byte(string(bytes)), &apiQuickBuy)
    return apiQuickBuy, nil
}

//GetOrders - Получить список выставленных ордеров с страницы BASE_URL/orders/
func (a *API) GetOrders() (APIGetOrders, error){
    bytes, err := makeGet(fmt.Sprintf(URLGetOrders, a.URL, a.Key))
    if err != nil {
        return APIGetOrders{}, err
    }
    var apiGetOrders APIGetOrders
    json.Unmarshal([]byte(string(bytes)), &apiGetOrders)
    return apiGetOrders, nil
}

//InsertOrder - Создание новой заявки на покупку.
//classid и instanceid - идентификаторы предмета.
//price - цена в копейках(целое число), именно с этой ценой вы создате заявку на покупку
//hash - md5 от описания предмета. Вы можете найти его в ответе метода ItemInfo. Это введено, чтобы вы были уверены в покупке именно той вещи, которую покупаете.
func (a *API) InsertOrder(classid string, instanceid string, price int64, hash string) (APIInsertOrder, error){
    bytes, err := makeGet(fmt.Sprintf(URLInsertOrder, a.URL, classid, instanceid, price, hash, a.Key))
    if err != nil {
        return APIInsertOrder{}, err
    }
    var apiInsertOrder APIInsertOrder
    json.Unmarshal([]byte(string(bytes)), &apiInsertOrder)
    return apiInsertOrder, nil
}

//UpdateOrder - Изменение/Удаление заявки на покупку.
//classid и instanceid - идентификаторы предмета.
//price - цена в копейках(целое число), цена указанная в заявке на покупку изменится на указанную тут. Если вы пришлете 0, то эта заявка на покупку будет удалена.
func (a *API) UpdateOrder(classid string, instanceid string, price int64) (APIUpdateOrder, error){
    bytes, err := makeGet(fmt.Sprintf(URLUpdateOrder, a.URL, classid, instanceid, price, a.Key))
    if err != nil {
        return APIUpdateOrder{}, err
    }
    var apiUpdateOrder APIUpdateOrder
    json.Unmarshal([]byte(string(bytes)), &apiUpdateOrder)
    return apiUpdateOrder, nil
}

//DeleteOrders - Удаление всех заявок на покупку.
func (a *API) DeleteOrders() (APIDeleteOrders, error){
    bytes, err := makeGet(fmt.Sprintf(URLDeleteOrders, a.URL, a.Key))
    if err != nil {
        return APIDeleteOrders{}, err
    }
    var apiDeleteOrders APIDeleteOrders
    json.Unmarshal(bytes, &apiDeleteOrders)
    return apiDeleteOrders, nil
}

//GetNotifications - Получить список включенных уведомлений о изменении цены. BASE_URL/mail/
func (a *API) GetNotifications() (APIGetNotifications, error){
    bytes, err := makeGet(fmt.Sprintf(URLGetNotifications, a.URL, a.Key))
    if err != nil {
        return APIGetNotifications{}, err
    }
    var apiGetNotifications APIGetNotifications
    json.Unmarshal(bytes, &apiGetNotifications)
    return apiGetNotifications, nil
}

//UpdateNotification - Изменение/Удаление уведомления о изменении цены на остлеживаемый предмет.
//classid и instanceid - идентификаторы предмета.
//price - цена в копейках(целое число), если появится предложение о покупке ниже этой цены, то вы получите уведомление. Если вы пришлете 0, то это уведомление будет удалено.
func (a *API) UpdateNotification(classid string, instanceid string, price int64) (APIUpdateNotification, error){
    bytes, err := makeGet(fmt.Sprintf(URLUpdateNotification, a.URL, classid, instanceid, price, a.Key))
    if err != nil {
        return APIUpdateNotification{}, err
    }
    var apiUpdateNotification APIUpdateNotification
    json.Unmarshal(bytes, &apiUpdateNotification)
    return apiUpdateNotification, nil
}

//GetWSAuth - Ключ для подписки на вебсокеты. Получение приватных оповещений. 
func (a *API) GetWSAuth() (APIGetWSAuth, error){
    bytes, err := makeGet(fmt.Sprintf(URLGetWSAuth, a.URL, a.Key))
    if err != nil {
        return APIGetWSAuth{}, err
    }
    var apiGetWSAuth APIGetWSAuth
    json.Unmarshal(bytes, &apiGetWSAuth)
    return apiGetWSAuth, nil
}

func getURLByAction(action string) (string){
    switch action {
        case "dota2":
            return URLDota2
        case "csgo":
            return URLCsgo
        case "tf2":
            return URLTf2
        case "gifts":
            return URLGifts
        default:
            panic(fmt.Sprintf("Action %s is not defined", action))
    }
}

func newAPI(key string, action string) (API, error) {
    url := getURLByAction(action)

    api := API{
        Key:  key,
        Action: action,
        URL: url,
        Lang: "ru",
    }

    _, err := api.Test()
    if err != nil {
        return API{}, err
    }

    return api, nil
}

//NewDota2API - создание нового объекта API Dota2
func NewDota2API(key string) (API, error) {
    return newAPI(key, "dota2")
}

//NewCsgoAPI - создание нового объекта API Csgo
func NewCsgoAPI(key string) (API, error) {
    return newAPI(key, "csgo")
}

//NewTf2API - создание нового объекта API Tf2
func NewTf2API(key string) (API, error) {
    return newAPI(key, "tf2")
}

//NewGiftsAPI - создание нового объекта API Gifts
func NewGiftsAPI(key string) (API, error) {
    return newAPI(key, "gifts")
}