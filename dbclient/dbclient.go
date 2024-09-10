package dbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"goblog/accountservice/model"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

// 桶名称
var accountBucket string = "AccountBucket"

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
}

type BoltClient struct {
	boltDB *bolt.DB
}

// 连接数据库
func (boltClient *BoltClient) OpenBoltDb() {
	var err error
	boltClient.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//boltClietI = boltClient
	log.Println("open bolt db sucess!")
}

func (bc *BoltClient) Seed() {

	bc.initializeBucket()
	bc.seedAccounts()
}

// 制造账号信息
func (bc *BoltClient) seedAccounts() {

	var total = 100
	for i := 0; i < total; i++ {
		var key = strconv.Itoa(10000 + i)
		acc := model.Account{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}
		jsonBytes, _ := json.Marshal(acc)
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(accountBucket))
			err := bucket.Put([]byte(key), []byte(jsonBytes))
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}

// 初始化用户桶
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(accountBucket))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// 根据账户ID查询账号
func (boltClient *BoltClient) QueryAccount(id string) (model.Account, error) {
	var account = model.Account{}

	err := boltClient.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(accountBucket))
		accountBytes := bucket.Get([]byte(id))
		if accountBytes == nil {

			return errors.New("No account found")
			//fmt.Errorf("No account found for %s", id)
		}
		err := json.Unmarshal(accountBytes, &account)
		if err != nil {
			return errors.New("parse account fail!")
		}
		return nil

	})

	return account, err
}
