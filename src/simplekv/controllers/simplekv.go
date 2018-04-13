package controllers

import (
	"simplekv/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about simplekv
type SimpleKVController struct {
	beego.Controller
}

// @Title Post
// @Description Put key-value to cloud
// @Param	body		body 	models.KVObject	true		"The KVObject content"
// @Success 200 {string} models.KVObject.Key
// @Failure 403 body is empty
// @router / [post]
func (o *SimpleKVController) Post() {
	var ob models.KVObject
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	//simplekvid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"Key":ob.Key}
	models.KV[ob.Key] = ob.Value;
	o.ServeJSON()
}

// @Title Get
// @Description find key-value by key
// @Param	key		path 	string	true		"the key of kv you want to get"
// @Success 200 {simplekv} models.KVObject
// @Failure 403 : empty object
// @router /:key [get]
func (o *SimpleKVController) Get() {
	key := o.Ctx.Input.Param(":key")
	if key != "" {
		o.Data["json"] = &models.KVObject{
			Key:key,
			Value: models.KV[key] }
	}
	o.ServeJSON()
}



// // @Title GetAll
// // @Description get all simplekvs
// // @Success 200 {simplekv} models.KVObject
// // @Failure 403 :simplekvId is empty
// // @router / [get]
// func (o *SimpleKVController) GetAll() {
// 	obs := models.GetAll()
// 	o.Data["json"] = obs
// 	o.ServeJSON()
// }

// // @Title Update
// // @Description update the simplekv
// // @Param	simplekvId		path 	string	true		"The simplekvid you want to update"
// // @Param	body		body 	models.KVObject	true		"The body"
// // @Success 200 {simplekv} models.KVObject
// // @Failure 403 :simplekvId is empty
// // @router /:simplekvId [put]
// func (o *SimpleKVController) Put() {
// 	simplekvId := o.Ctx.Input.Param(":simplekvId")
// 	var ob models.KVObject
// 	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

// 	err := models.Update(simplekvId, ob.Score)
// 	if err != nil {
// 		o.Data["json"] = err.Error()
// 	} else {
// 		o.Data["json"] = "update success!"
// 	}
// 	o.ServeJSON()
// }

// // @Title Delete
// // @Description delete the simplekv
// // @Param	simplekvId		path 	string	true		"The simplekvId you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 simplekvId is empty
// // @router /:simplekvId [delete]
// func (o *SimpleKVController) Delete() {
// 	simplekvId := o.Ctx.Input.Param(":simplekvId")
// 	models.Delete(simplekvId)
// 	o.Data["json"] = "delete success!"
// 	o.ServeJSON()
// }

