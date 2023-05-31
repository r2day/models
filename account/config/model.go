package config

import (
	"github.com/r2day/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "affiliate_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_info"
	// 这个需要用户根据具体业务完成设定
	modelName = "account"
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
	// 第三方账号信息
	External ExternalInfo `json:"external" bson:"external"`
	// 手机号
	Phone string `json:"phone"`
	// 更多信息
	// 账号名称
	Nickname string     `json:"nickname"`
	Avatar   string     `json:"avatar"`
	Member   MemberInfo `json:"member" bson:"member"`
	Assets   AssetsInfo `json:"assets" bson:"assets"`
}

type ExternalInfo struct {
	OpenID string `json:"open_id"  bson:"open_id"`
}

type MemberInfo struct {
	MemberLevel int    `json:"member_level"`
	CardName    string `json:"card_name" bson:"card_name"`
	CardUrl     string `json:"cardUrl"`
}

type AssetsInfo struct {
	PointNum     int `json:"pointNum"`
	CouponNum    int `json:"couponNum"`
	Balance      int `json:"balance"`
	GiftBalance  int `json:"giftBalance"`
	CurrentValue int `json:"currentValue"`
	Level        int `json:"level"`
	NeedValue    int `json:"needValue"`
}
