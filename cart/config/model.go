package config

import (
	"github.com/r2day/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "client_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_flow"
	// 这个需要用户根据具体业务完成设定
	modelName = "cart"
)

// 每一个应用表示一个大的模块，通常其子模块是一个个接口
// 是有系统默认设定，用户无需修改
// 用户只需要在创建角色的时候选择好需要的应用即可
// 用户选择所需要的应用后->完成角色创建->系统自动拷贝应用具体信息到角色下
// 此时用户可以针对当前的角色中具体的项再自行选择是否移除部分接口，从而进行更精细的权限管理

// Model 模型
type Model struct {
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	Meta models.MetaModel `json:"meta" bson:"meta"`

	Status           int    `json:"status"`
	CompletedTime    string `json:"completed_time"`
	MultiStore       string `json:"multi_store"`
	ProductionedTime string `json:"productioned_time"`
	TypeCate         int    `json:"typeCate"`
	SendStatus       int    `json:"send_status"`

	Items []Item `json:"items"`
}

//id: good.id,
//cate_id: cate.id,
//name: good.name,
//price: good.price,
//number: num,
//image: good.images,
//use_property: good.use_property,
//props_text: good.props_text ? good.props_text: '',
//props: good.props

type Item struct {
	ID     int    `json:"id" bson:"id"`
	CateID string `json:"cate_id"`
	Name   string `json:"name" bson:"name"`
	Price  int    `json:"price"`
	Number int    `json:"number" bson:"number"`
	Image  string `json:"image"`
	// 订单是否有属性
	UseProperty int `json:"use_property"`
	// 直接显示在订单上的 规格属性
	PropsText string `json:"props_text"`
	// 规格编号
	Props []int `json:"props"`
}
