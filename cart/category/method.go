package category

import (
	"context"
	"time"

	db "github.com/r2day/auth"
	rtime "github.com/r2day/base/time"
	"github.com/r2day/rest"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}

// Create 创建
// create	POST http://my.api.url/posts
func (m *Model) Create(ctx context.Context) (string, error) {
	coll := db.MDB.Collection(m.CollectionName())

	// 保存时间设定
	m.Meta.CreatedAt = rtime.FomratTimeAsReader(time.Now().Unix())
	// 更新时间设定
	m.Meta.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	// 插入记录
	result, err := coll.InsertOne(ctx, m)
	if err != nil {
		log.WithField("m", m).Error(err)
		return "", err
	}
	stringObjectID := result.InsertedID.(primitive.ObjectID).Hex()
	return stringObjectID, nil
}

// Delete 删除
// delete	DELETE http://my.api.url/posts/123
func (m *Model) Delete(ctx context.Context, id string) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	logCtx := log.WithField("id", id)
	coll := db.MDB.Collection(m.CollectionName())
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	// 执行删除
	result, err := coll.DeleteOne(ctx, filter)

	if err != nil {
		logCtx.Error(err)
		return err
	}

	if result.DeletedCount < 1 {
		logCtx.Warning("result.DeletedCount < 1")
		return nil
	}
	return nil
}

// GetOne 详情
// getOne	GET http://my.api.url/posts/123
func (m *Model) GetOne(ctx context.Context, id string) (*Model, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(m.CollectionName())
	// 绑定查询结果
	result := &Model{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	logCtx := log.WithField("filter", filter)

	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		logCtx.Error(err)
		return nil, err
	}
	return result, nil
}

// GetMany 获取条件查询的结果
// getMany	GET http://my.api.url/posts?filter={"ids":[123,456,789]}
func (m *Model) GetMany(ctx context.Context, ids []*primitive.ObjectID) ([]*Model, int64, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(m.CollectionName())
	// 绑定查询结果
	results := make([]*Model, 0)
	logCtx := log.WithField("ids", ids)
	cursor, err := coll.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})

	if err != nil {
		logCtx.Error(err)
		return nil, 0, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		logCtx.Error(err)
		return nil, 0, err
	}
	return results, int64(len(ids)), nil
}

// Update 更新
// update	PUT http://my.api.url/posts/123
func (m *Model) Update(ctx context.Context, id string) error {
	coll := db.MDB.Collection(m.CollectionName())
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	// 设定更新时间
	m.Meta.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	result, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: m}})
	if err != nil {
		log.WithField("id", id).Error(err)
		return err
	}

	if result.MatchedCount < 1 {
		log.WithField("id", id).Warning("no matched record")
		return nil
	}

	return nil
}

// IncrementReference 更新
// https://www.mongodb.com/docs/manual/reference/operator/update/inc/
func (m *Model) IncrementReference(ctx context.Context, id string) error {
	coll := db.MDB.Collection(m.CollectionName())
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	// 设定更新时间
	m.Meta.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	result, err := coll.UpdateOne(ctx, filter,
		bson.D{{Key: "$set", Value: bson.D{{"reference", 1}}}})
	if err != nil {
		log.WithField("id", id).Error(err)
		return err
	}

	if result.MatchedCount < 1 {
		log.WithField("id", id).Warning("no matched record")
		return nil
	}
	return nil
}

// DecrementReference 更新
// https://www.mongodb.com/docs/manual/reference/operator/update/inc/
func (m *Model) DecrementReference(ctx context.Context, id string) error {
	coll := db.MDB.Collection(m.CollectionName())
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	// 设定更新时间
	m.Meta.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	result, err := coll.UpdateOne(ctx, filter,
		bson.D{{Key: "$set", Value: bson.D{{"reference", -1}}}})
	if err != nil {
		log.WithField("id", id).Error(err)
		return err
	}

	if result.MatchedCount < 1 {
		log.WithField("id", id).Warning("no matched record")
		return nil
	}
	return nil
}

// GetList 获取列表
// getList	GET http://my.api.url/posts?sort=["title","ASC"]&range=[0, 24]&filter={"title":"bar"}
func (m *Model) GetList(ctx context.Context, merchantID string, accountID string, p *rest.Params) ([]*Model, int64, error) {
	coll := db.MDB.Collection(m.CollectionName())
	// 声明需要返回的列表
	results := make([]*Model, 0)
	// 声明日志基本信息
	logCtx := log.WithField("merchantID", merchantID).WithField("urlParams", p)
	// 声明数据库过滤器
	// 定义基本过滤规则
	// 以商户id为基本命名空间
	// 并且只能看到小于等于自己的级别的数据
	opt := p.ToMongoOptions()
	filters := p.ToMongoFilter(merchantID, m.Meta.AccessLevel)
	if p.HasFilterIn {
		return m.GetMany(ctx, p.MongoIDList)
	}

	logCtx.WithField("filer -->", filters).WithField("client_filter", p.Filter).
		WithField("opt", opt).Info("~~~~~~~~~~~~~~~~~~~")

	//// 获取总数（含过滤规则）
	totalCounter, err := coll.CountDocuments(context.TODO(), filters)
	if err == mongo.ErrNoDocuments {
		logCtx.Error(err)
		return nil, 0, err
	}
	// 获取数据列表
	cursor, err := coll.Find(ctx, filters, opt)
	if err == mongo.ErrNoDocuments {
		logCtx.Error(err)
		return nil, totalCounter, err
	}

	if err != nil {
		logCtx.Error(err)
		return nil, totalCounter, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		logCtx.Error(err)
		return nil, totalCounter, err
	}
	return results, totalCounter, nil
}
