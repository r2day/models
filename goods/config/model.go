package config

import (
	"github.com/r2day/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "product_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "goods"
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

	Name            string      `json:"name"`
	IsShowBackstage int         `json:"is_show_backstage"`
	Sort            int         `json:"sort"`
	Icon            string      `json:"icon"`
	Goods           []GoodsInfo `json:"goods_list" bson:"goods_list"`
}

type UserInfo struct {
	Mobile   string `json:"mobile"`
	UserName string `json:"user_name"`
}

type PropertyInfo struct {
	IsOpenCheckbox bool `json:"is_open_checkbox"`
	Id             int  `json:"id"`
	Values         []struct {
		IsDefault int    `json:"is_default,omitempty"`
		Id        int    `json:"id"`
		Code      string `json:"code"`
		Value     string `json:"value"`
	} `json:"values"`
	Name string  `json:"name"`
	Desc *string `json:"desc,omitempty"`
}

type EntityInfo struct {
	SpecId          string      `json:"spec_id"`
	TradeMark       string      `json:"trade_mark"`
	Id              string      `json:"id"`
	Stock           string      `json:"stock"`
	SpecText        interface{} `json:"spec_text"`
	Spec            interface{} `json:"spec"`
	Image           string      `json:"image"`
	Num             int         `json:"num"`
	Price           float64     `json:"price"`
	MembershipPrice int         `json:"membership_price"  bson:"membership_price"`
}
type GoodsInfo struct {
	SellTimeStatus  int            `json:"sell_time_status" bson:"sell_time_status"`
	Id              int            `json:"id"`
	IsSell          bool           `json:"is_sell" bson:"is_sell"`
	PackCost        string         `json:"pack_cost" bson:"pack_cost"`
	Sales           int            `json:"sales"`
	GoodsType       int            `json:"goods_type" bson:"goods_type"`
	CoverImg        string         `json:"cover_img" bson:"cover_img"`
	Property        []PropertyInfo `json:"property"`
	GoodsMealsInfo  []interface{}  `json:"goods_meals_info,omitempty" bson:"goods_meals_info"`
	IsAdd           int            `json:"is_add,omitempty" bson:"is_add"`
	UseSpec         bool           `json:"use_spec" bson:"use_spec"`
	Entity          []EntityInfo   `json:"entity"`
	StallCode       string         `json:"stall_code" bson:"stall_code"`
	Sort            int            `json:"sort"`
	Price           float64        `json:"price"`
	Unit            string         `json:"unit"`
	ImageArr        []string       `json:"imageArr"`
	MembershipPrice int            `json:"membership_price" bson:"membership_price"`
	UseProperty     int            `json:"use_property" bson:"use_property"`
	UnitType        int            `json:"unit_type" bson:"unit_type"`
	MinBuyNum       int            `json:"min_buy_num" bson:"min_buy_num"`
	Specs           []struct {
		Values []struct {
			Id    int         `json:"id"`
			Image interface{} `json:"image"`
			Value string      `json:"value"`
		} `json:"values"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"specs"`
	Content      string `json:"content"`
	IsFollowSuit int    `json:"is_follow_suit,omitempty"`
	Stock        string `json:"stock"`
	Type         int    `json:"type"`
	IsLabel      int    `json:"is_label"`
	Name         string `json:"name"`
	Images       string `json:"images"`
}
