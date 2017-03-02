package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

var logger = shim.NewLogger("Shopping_Cart")

type ShoppingCart struct {
	
}

const (
	userTable = "UserTable"
	productTable        = "ProductTable"
)


type UserRecords []UserRecord

type UserRecord struct {
	UserId  string `json:"user_id"`
	UserName string `json:"userName"`
	Type    string   `json:"type"`
}


func (t *ShoppingCart) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	
	var err error
	
	if len(args) > 0 {
		logger.Error("Incorrect number of arguments")
		return nil, errors.New("Incorrect number of arguments. No arguments required.")
	}
	
	_, err = stub.GetTable(userTable)
	if err == shim.ErrTableNotFound {
		err = stub.CreateTable(userTable, []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "UserId", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "UserName", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: false},
		})
		if err != nil {
			logger.Errorf("Error creating table:%s - %s", userTable, err.Error())
			return nil, errors.New("Failed creating userTable table.")
		}
}else {
		logger.Info("Table already exists")
	}

    _, err = stub.GetTable(productTable)
    if err == shim.ErrTableNotFound {
		err = stub.CreateTable(productTable, []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "ProductId", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "ProductName", Type: shim.ColumnDefinition_STRING, Key: true},
		})
		if err != nil {
			logger.Errorf("Error creating table:%s - %s", productTable, err.Error())
			return nil, errors.New("Failed creating productTable .")
		}
	} else {
		logger.Info("Table already exists")
	}

logger.Info("Successfully deployed chain code")

	return nil, nil
}
	
func (t *ShoppingCart) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "enroll" {
		return t.enroll(stub, args)
	}

	logger.Errorf("Unimplemented method :%s called", function)

	return nil, errors.New("Unimplemented '" + function + "' invoked")
}

func (t *ShoppingCart) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "deviceServiceRecords" {
		return t.deviceServiceRecords(stub, args)
	}
	return nil, errors.New("Invalid query function name")
}


func (t *ShoppingCart) enroll(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Info("In enroll function")
	if len(args) < 3 {
		logger.Error("Incorrect number of arguments")
		return nil, errors.New("Incorrect number of arguments. Specify device id, public key, owner, owner for check1, check2 and check3.")
	}

	UserId := args[0]
	UserName := args[1]
	Type := args[2]
	
	ok, err := stub.InsertRow(userTable, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: UserId}},
			&shim.Column{Value: &shim.Column_String_{String_: UserName}},
			&shim.Column{Value: &shim.Column_String_{String_: Type}},
		},
	})

	if !ok || err != nil {
		logger.Errorf("Error in enrolling a new user:%s", err)
		return nil, errors.New("Error in enrolling a new user")
	}
	logger.Infof("Enrolled device %s", UserId)

	return nil, nil
}

func (t *ShoppingCart) deviceServiceRecords(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Info("In deviceServiceRecord function")
	if len(args) != 1 {
		logger.Error("Incorrect number of arguments")
		return nil, errors.New("Incorrect number of arguments. Specify device id.")
	}
	userId := args[0]
	userRecords, err := t.getDeviceServiceRecords(stub, userId)
	if err != nil {
		logger.Errorf("Failed fetching device service records: [%s]", err)
		return nil, fmt.Errorf("Failed fetching device service records [%s]", err)
	}

	payload, err := json.Marshal(userRecords)
	if err != nil {
		logger.Errorf("Failed marshalling payload: [%s]", err)
		return nil, fmt.Errorf("Failed marshalling payload [%s]", err)
	}

	return payload, nil
}


func (t *ShoppingCart) getDeviceServiceRecords(stub shim.ChaincodeStubInterface, userId string) (UserRecords, error) {
	var columns []shim.Column
	if userId != "" {
		col := shim.Column{Value: &shim.Column_String_{String_: userId}}
		columns = append(columns, col)
	}
	rowChannel, err := stub.GetRows(userTable, columns)
	if err != nil {
		logger.Errorf("Error in getting rows:%s", err.Error())
		return nil, errors.New("Error in fetching rows")
	}
	userRecords := UserRecords{}
	for row := range rowChannel {
		userRecord := t.extractServiceRecord(row)
		logger.Debug(userRecord.UserName)
		logger.Debug(userRecord.Type)
		logger.Debug(userRecord.UserId)
		logger.Infof("username %s", userRecord.UserName)
		logger.Infof("type %s", userRecord.Type)
		logger.Infof("userId %s", userRecord.UserId)
		userRecords = append(userRecords, userRecord)
	}
	return userRecords, nil
}



func (t *ShoppingCart) extractServiceRecord(row shim.Row) UserRecord {
	return UserRecord{
		UserId:  row.Columns[0].GetString_(),
		UserName: row.Columns[1].GetString_(),
		Type: row.Columns[2].GetString_(),
		logger.Infof("username %s", UserName)
		logger.Infof("type %s", Type)
		logger.Infof("userId %s", UserId)
	}
}


func main() {
	
	logger.SetLevel(shim.LogDebug)
	err := shim.Start(new(ShoppingCart))
	if err != nil {
		fmt.Printf("Error starting Energy trading chaincode: %s", err)
	}
}
