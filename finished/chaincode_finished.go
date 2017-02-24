package main

import (
	
	"fmt"
	"errors"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)


type SimpleChaincode struct {
}

func main(){
	
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("sudheer")
	}	
}


func(t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("expected 1 argument in INIT method")
	}
	err := stub.PutState("hello world", []byte(args[0]))
	if err != nil {
		 return nil, err
	}
return nil, nil
}


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	
	if function == "init"{
		return t.Init(stub, "init", args)
	}
	if function == "write"{
		return t.write(stub, args)
	}
	return nil, nil
}


func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	
	if function == "read"{
		t.read(stub, args)
	}
	return nil, nil
}


func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface,args []string) ([]byte, error){
	
	var key,jsonRESP string
	var err error
	
	if len(args) !=1 {
		
	} 
	
	key = args[0]
	valAsBytes, err := stub.GetState(key)
	if err != nil {
		jsonRESP = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonRESP)
	}
	return valAsBytes, nil
}

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
	var key, value string
	var err error
	
	if len(args)!=2{
		return nil, errors.New("2 arg expected in write method")
	}
	
	key = args[0]
	value = args[1]
    err = stub.PutState(key, []byte(value))
    if err != nil {
		return nil, err
	}
	return nil, nil
}
