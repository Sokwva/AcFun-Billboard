package saveDougaInfoToDb

import (
	"context"
	"errors"
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/fetch/dougaInfo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Handler *mongo.Client

var collection *mongo.Collection

func InitClient() error {
	if common.ConfHandle.DougaInfoSave.MongoSvrConnURI == "" {
		return errors.New("conf file did no set mongo db conf")
	}
	if common.ConfHandle.DougaInfoSave.DbName == "" {
		return errors.New("conf file did not set mongo db dbname")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(common.ConfHandle.DougaInfoSave.MongoSvrConnURI))
	if err != nil {
		return err
	}
	Handler = client
	collection = Handler.Database(common.ConfHandle.DougaInfoSave.DbName).Collection(common.ConfHandle.DougaInfoSave.ACIDInfoCollectionName)
	return nil
}

func CheckACIDExist(acid string) bool {
	if collection == nil {
		common.Log.Debug("mongodb collection is nil", "acid", acid)
		return false
	}
	res := collection.FindOne(context.Background(), bson.M{"dougaId": acid})
	return res.Err() == nil
}

func WriteInDb(acid string) {
	info, err := dougaInfo.GetVideoInfo(acid)
	if err != nil {
		return
	}
	bsonData, err := bson.Marshal(info)
	if err != nil {
		common.Log.Debug("invalid dougainfo save to mongodb", "acid", acid)
		return
	}
	collection.InsertOne(context.Background(), bsonData)
}
