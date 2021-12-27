package repository

import (
	"github.com/guregu/dynamo"

	"gopractice/domain"
)

type DynamoDBRepo struct {
	dynamoDB              *dynamo.DB
	tableName, primaryKey string
}

func NewDynamoDBRepo(d *dynamo.DB, isMock bool) *DynamoDBRepo {
	if isMock {
		return nil
	}
	return &DynamoDBRepo{
		dynamoDB:   d,
		tableName:  "test_list",
		primaryKey: "name",
	}
}

func (d *DynamoDBRepo) GetByName(name string) (domain.Data, error) {
	var data domain.Data
	if err := d.dynamoDB.Table(d.tableName).Get(d.primaryKey, name).One(&data); err != nil {
		return domain.Data{}, err
	}
	return data, nil
}
