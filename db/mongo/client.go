package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/shahzadhaider7/task-scheduler/config"
	"github.com/shahzadhaider7/task-scheduler/db"
	domainErr "github.com/shahzadhaider7/task-scheduler/errors"
	"github.com/shahzadhaider7/task-scheduler/models"
)

const (
	stuCollection = "Task"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mongo database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	return &client{conn: cli}, nil
}

func (c *client) AddTask(task *models.Task) (string, error) {
	if task.ID != "" {
		return "", errors.New("id is not empty")
	}
	task.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.InsertOne(context.TODO(), task); err != nil {
		return "", errors.Wrap(err, "failed to add task")
	}

	return task.ID, nil
}

func (c *client) GetTask(id string) (*models.Task, error) {
	var stu *models.Task
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&stu); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", id))
		}
		return nil, err
	}
	return stu, nil
}

func (c *client) DeleteTask(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete task")
	}

	return nil
}

func (c *client) UpdateTask(task *models.Task) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": task.ID}, bson.M{"$set": task}); err != nil {
		return errors.Wrap(err, "failed to update task")
	}

	return nil
}
