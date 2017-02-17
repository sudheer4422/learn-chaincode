package main

import(

"fmt"
"errors"
"github.com/hyperledger/fabric/core/chaincode/shim"
)


type SimpleChaincode struct {


}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


func (t *SimpleChaincode) Init (stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

err := createTableone(stub)
err1 := createTabletwo(stub)
if err != nil{
return nil, fmt.Errorf("Error creating table one during init. %s", err)
}
if err1 != nil{
return nil, fmt.Errorf("Error creating table one during init. %s", err)
}
return nil, errors.New("Unsupported operation")
}

func createTabletwo(stub shim.ChaincodeStubInterface) error {

var columnDefsTableTwo []*shim.ColumnDefinition

rollnumber := shim.ColumnDefinition{Name: "colOneTableTwo",Type: shim.ColumnDefinition_STRING, Key: true}
Name := shim.ColumnDefinition{Name: "coltwoTableTwo",Type: shim.ColumnDefinition_STRING, Key: true}
Gender := shim.ColumnDefinition{Name: "colthreeTableTwo",Type: shim.ColumnDefinition_STRING, Key: true}

columnDefsTableTwo = append (columnDefsTableTwo, &rollnumber)
columnDefsTableTwo = append (columnDefsTableTwo, &Name)
columnDefsTableTwo = append (columnDefsTableTwo, &Gender)
return stub.CreateTable("tableTwo", columnDefsTableTwo)

}

func createTableone(stub shim.ChaincodeStubInterface) error {

var columnDefsTableOne []*shim.ColumnDefinition

rollnumber := shim.ColumnDefinition{Name: "colOneTableOne",Type: shim.ColumnDefinition_STRING, Key: true}
Name := shim.ColumnDefinition{Name: "colTwoTableOne",Type: shim.ColumnDefinition_STRING, Key: true}
Gender := shim.ColumnDefinition{Name: "colThreeTableOne",Type: shim.ColumnDefinition_STRING, Key: true}

columnDefsTableOne = append (columnDefsTableOne, &rollnumber)
columnDefsTableOne = append (columnDefsTableOne, &Name)
columnDefsTableOne = append (columnDefsTableOne, &Gender)
return stub.CreateTable("tableOne", columnDefsTableOne)

}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
                case "getRowTableOne":
		if len(args) < 1 {
			return nil, errors.New("getRowTableOne failed. Must include 1 key value")
		}

		col1Val := args[0]
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		columns = append(columns, col1)
                row, err := stub.GetRow("tableOne", columns)
		if err != nil {
			return nil, fmt.Errorf("getRowTableOne operation failed. %s", err)
		}

		rowString := fmt.Sprintf("%s", row)
		return []byte(rowString), nil
	}
return nil, errors.New("Unsupported operation")
	}



func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

switch function {

       case "insertRowTableOne":

       if len(args) < 3 {
           return nil, errors.New("insertTableOne failed. Must include 3 column values")
       }
       
       colval1 := args[0]
       colval2 := args[1]
       colval3 := args[2]


       var columns []*shim.Column
       col1 :=  shim.Column {Value : &shim.Column_String_{String_: colval1}}
       col2 :=  shim.Column {Value : &shim.Column_String_{String_: colval2}}
       col3 :=  shim.Column {Value : &shim.Column_String_{String_: colval3}}
       columns = append(columns, &col1)
       columns = append(columns, &col2)
       columns = append(columns, &col3)

       row := shim.Row{Columns: columns}
       ok, err := stub.InsertRow("tableOne", row)
       if err != nil {
	return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
	}
	if !ok {
	return nil, errors.New("insertTableOne operation failed. Row with given key already exists")
	}
	return nil, errors.New("Unsupported operation")
}
return nil, errors.New("Unsupported operation")
}
