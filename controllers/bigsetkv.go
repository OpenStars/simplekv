package controllers

import (
	"strings"
	"fmt"
	"context"
	"simplekv/models"
	"encoding/json"
	"github.com/astaxie/beego"
	thriftpool "github.com/OpenStars/thriftpool"
	bs "github.com/OpenStars/backendclients/go/bigset/thrift/gen-go-0.11.0/openstars/core/bigset/generic"
	   "github.com/OpenStars/backendclients/go/bigset/transports"
)

var (	
	BSHost string
	BSPort string
)

func ConfigSetBSHostPort(host string, port string){
	BSHost = host;
	BSPort = port;
}

func getBSClient() *thriftpool.ThriftSocketClient{
	return transports.GetBSClient(BSHost, BSPort)
}

// Operations about simplekv
type BigsetKVController struct {
	beego.Controller
}
 
// @Title PutItem
// @Description Put key-value to cloud
// @Param	sessionID		query 	string	true		"Session ID"
// @Param	appID		query 	string	true		"App ID"
// @Param	key		query  	string	true		"The Key"
// @Param	val		query  	string	true		"The Value"
// @Success 200 {string} models.KVObject.Key
// @Failure 403 body is empty
// @router /putitem [post]
func (o *BigsetKVController) PutItem() {

	sessionID := o.GetString("sessionID")
	fmt.Println("PutItem ",o)

	//Todo: check sessionID 

	appID := o.GetString("appID")
	fmt.Printf("sessionID: %s, appID: %s, object: ",sessionID , appID)
	key := o.GetString("key") 
	val := o.GetString("val") 
	
	o.Data["json"] = map[string]string{"Key":key}

	client := getBSClient(); 
	
	if client != nil {
		defer client.BackToPool();
		fmt.Print(  client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(),  
				bs.TStringKey(appID)  , &bs.TItem{[]byte( key ), []byte(val)}) )
	}
	//models.KV[ob.Key] = ob.Value;
	o.ServeJSON()
}


// @Title Post
// @Description Put key-value to cloud
// @Param	sessionID		query 	string	true		"Session ID"
// @Param	appID		query 	string	true		"App ID"
// @Param	body		body 	models.KVObject	true		"The KVObject content"
// @Success 200 {string} models.KVObject.Key
// @Failure 403 body is empty
// @router / [post]
func (o *BigsetKVController) Post() {

	sessionID := o.GetString("sessionID")

	//Todo: check sessionID 

	appID := o.GetString("appID")
	fmt.Printf("sessionID: %s, appID: %s, object: ",sessionID , appID)

	var ob models.KVObject
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	fmt.Println(ob);

	//simplekvid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"Key":ob.Key}

	client := getBSClient();
	
	if client != nil {
		defer client.BackToPool();
		fmt.Print(  client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(),  
				bs.TStringKey(appID)  , &bs.TItem{[]byte( ob.Key ), []byte(ob.Value)  }) )
	}
	//models.KV[ob.Key] = ob.Value;
	o.ServeJSON()
}


// @Title Get
// @Description find key-value by key
// @Param	sessionID		query 	string	true		"Session ID of user/app"
// @Param	appID		query 	string	true		"appID"
// @Param	key		query 	string	true		"the key of kv you want to get"
// @Success 200 {simplekv} models.KVObject
// @Failure 403 : empty object
// @router /get [get]
func (o *BigsetKVController) Get() {
	sessionID := o.GetString("sessionID")
	appID := o.GetString("appID")
	key := o.GetString("key") //o.Ctx.Input.Param(":key")

	fmt.Printf("sessionID: %s, appID: %s, key: %s",sessionID , appID , key)
	
	if key != "" {
		client := getBSClient();
		if client != nil {
			defer client.BackToPool();
			res, _ := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(context.Background(),  bs.TStringKey(appID), []byte(key) )
			fmt.Print("result: ")
			fmt.Print(res)
			fmt.Print("\n")
			if (res !=nil && res.Item != nil && res.Item.Value != nil) {
				
				fmt.Println((string)(res.Item.Value[:]) )

				o.Data["json"] = &models.KVObject{
					Key:key,
					Value:  (string)(res.Item.Value[:] )} 
			} else {
				o.Data["json"] = &models.KVObject{
					Key:key,
					Value: "notfound",
				}
			}
		}
	}
	o.ServeJSON()
}


// @Title MultiGet
// @Description find key-value by key
// @Param	sessionID		query 	string	true		"Session ID of user/app"
// @Param	appID		query 	string	true		"appID"
// @Param	keys		query 	[]string	true		"the keys of kv you want to get, saparated by comma, "
// @Success 200 {bigsetkv} []models.KVObject
// @Failure 403 : empty object
// @router /multiget [get]
func (o *BigsetKVController) MultiGet() {
	sessionID := o.GetString("sessionID")
	appID := o.GetString("appID")
	
	skeys := o.GetString("keys")

	keys:=strings.Split(skeys, ",")

	fmt.Printf("Multiget  sessionID: %s, appID: %s , keys: ",sessionID , appID, keys, len(keys) )
	
	
	if len(keys) != 0 {
		client := getBSClient();
		if client != nil {
			defer client.BackToPool();
		}

		kvs := []*models.KVObject{}
		for _, key := range keys {
			fmt.Println("get item ...", key)
			res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(context.Background(),  bs.TStringKey(appID), []byte(key) )
			fmt.Println("get item ", key, res, err);
			if (res !=nil && res.Item != nil && res.Item.Value != nil) {
				fmt.Println("get item ok ", res);
				kvs = append(kvs,&models.KVObject{
					Key: key,
					Value:  (string)(res.Item.Value[:] )} )				
			}
		}
		fmt.Println("kvs: ",kvs)
		o.Data["json"] = kvs
	}
	o.ServeJSON()
}


// @Title Remove
// @Description find key-value by key
// @Param	sessionID		query 	string	true		"Session ID of user/app"
// @Param	appID		query 	string	true		"appID"
// @Param	keys		query 	[]string	true		"the keys of kv you want to get, saparated by comma, "
// @Success 200 {deletedKeys} []string
// @Failure 403 : empty object
// @router /remove [get]
func (o *BigsetKVController) Remove() {
	sessionID := o.GetString("sessionID")
	appID := o.GetString("appID")
	skeys := o.GetString("keys")

	keys:=strings.Split(skeys, ",")

	fmt.Printf("Remove sessionID: %s, appID: %s ",sessionID , appID )
	fmt.Println(keys)
	
	if len(keys) != 0 {
		client := getBSClient();
		if client != nil {
			defer client.BackToPool();
		}
		for _, key := range keys {
			fmt.Println("get item ...", key)
			client.Client.(*bs.TStringBigSetKVServiceClient).BsRemoveItem(context.Background(),  bs.TStringKey(appID), []byte(key) )

		}

	}
	o.Data["json"] = keys
	o.ServeJSON()
}
