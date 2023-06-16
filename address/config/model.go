package config

import (
	"github.com/r2day/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "user_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "address"
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

	// 收货人
	AcceptName string `json:"accept_name" bson:"accept_name"`
	Mobile     string `json:"mobile"`
	Sex        int    `json:"sex"`

	Address AddressInfo `json:"district"`
	Inner   bool        `json:"inner"`
	// 每个用户只能选择一个默认地址
	IsDefault bool `json:"is_default" bson:"is_default"`
}

type AddressInfo struct {
	AreaName   string `json:"area_name" bson:"area_name"`
	CityName   string `json:"city_name" bson:"city_name"`
	Poiname    string `json:"poiname" bson:"poiname"`
	Area       string `json:"area" bson:"area"`
	City       string `json:"city" bson:"city"`
	Province   string `json:"province" bson:"province"`
	Street     string `json:"street"`
	Lat        string `json:"lat"`
	DoorNumber string `json:"door_number" bson:"door_number"`
}
