/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main


import (
        "bytes"
        "encoding/json"
        "fmt"
        "strconv"
        "time"

        "github.com/hyperledger/fabric/core/chaincode/shim"
        sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the bag
type Bag struct {
        TagId                   string          `json:"tagId"`
        MfgDate                 string          	`json:"mfgDate"`
        Mfg    					string   		`json:"mfg"`
        Location            	string          `json:"location"`
        DistributorId           string          `json:"distributorId"`
        RetailerId              string          `json:"retailerId"`
        Owner                   string          `json:"owner"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

        // Retrieve the requested Smart Contract function and arguments
        function, args := APIstub.GetFunctionAndParameters()
        // Route to the appropriate handler function to interact with the ledger appropriatel+++y
        if function == "createBag" {
                return s.createBag(APIstub, args)
        } else if function == "assignDistributor" {
                return s.assignDistributor(APIstub, args)
        } else if function == "assignRetailer" {
                return s.assignRetailer(APIstub, args)
        }else if function == "assignOwner" {
                return s.assignOwner(APIstub, args)
        }else if function == "getBagHistory" {
                return s.getBagHistory(APIstub, args)
        }

        return shim.Error("Invalid Smart Contract function name.")
}





// This function is initiate by LV
func (s *SmartContract) createBag(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {


        BG := Bag{}

        err  := json.Unmarshal([]byte(args[0]),&BG)
		if err != nil {
                return shim.Error("Not able to parse args into BG")
        }
        BGBytes, err := json.Marshal(BG)
		APIstub.PutState(BG.TagId,BGBytes)
        fmt.Println("Bag Created -> ", BG)



        return shim.Success(nil)
}

// This function is initiate by Seller
func (s *SmartContract) assignDistributor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        bagID := struct {
                TagId  string `json:"tagId"`
				DistributorId string `json:"distributorId"`
        }{}
				
        err  := json.Unmarshal([]byte(args[0]),&bagID)
        if err != nil {
                return shim.Error("Not able to parse args into bagID")
        }

        BagAsBytes, _ := APIstub.GetState(bagID.TagId)

        bg := Bag{}

        err = json.Unmarshal(BagAsBytes, &bg)

        if err != nil {
                return shim.Error("Issue with bg json unmarshaling")
        }

        //BG := bag{tagId: bg.tagId, mfgDate: bg.mfgDate, mfg: bg.mfg, location: bg.location, distributorId: bagID.distributorId, retailerId: bg.retailerId, owner: bg.owner}
		bg.DistributorId=bagID.DistributorId
        BGBytes, err := json.Marshal(bg)

        if err != nil {
                return shim.Error("Issue with bag json marshaling")
        }

		APIstub.PutState(bg.TagId,BGBytes)
        fmt.Println("Distributor ID Assigned -> ", bg)


        return shim.Success(nil)
}

func (s *SmartContract) assignRetailer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        bagID := struct {
                TagId  string `json:"tagId"`
				RetailerId string `json:"retailerId"`
        }{}
        err  := json.Unmarshal([]byte(args[0]),&bagID)
        if err != nil {
                return shim.Error("Not able to parse args into bagID")
        }

        BGAsBytes, _ := APIstub.GetState(bagID.TagId)

        bg := Bag{}

        err = json.Unmarshal(BGAsBytes, &bg)

        if err != nil {
                return shim.Error("Issue with bg json unmarshaling")
        }


        //LC := LetterOfCredit{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Accepted"}
		bg.RetailerId = bagID.RetailerId
        BGBytes, err := json.Marshal(bg)

        if err != nil {
                return shim.Error("Issue with BG json marshaling")
        }

		APIstub.PutState(bg.TagId,BGBytes)
        fmt.Println("Retailer ID Assigned -> ", bg)

        return shim.Success(nil)
}

func (s *SmartContract) assignOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        bagID := struct {
                TagId  string `json:"tagId"`
				Owner string `json:"owner"`
        }{}
        err  := json.Unmarshal([]byte(args[0]),&bagID)
        if err != nil {
                return shim.Error("Not able to parse args into bagID")
        }

        BGAsBytes, _ := APIstub.GetState(bagID.TagId)

        bg := Bag{}

        err = json.Unmarshal(BGAsBytes, &bg)

        if err != nil {
                return shim.Error("Issue with bg json unmarshaling")
        }


        //LC := LetterOfCredit{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Accepted"}
		bg.Owner = bagID.Owner
        BGBytes, err := json.Marshal(bg)

        if err != nil {
                return shim.Error("Issue with BG json marshaling")
        }

		APIstub.PutState(bg.TagId,BGBytes)
        fmt.Println("Owner Assigned -> ", bg)

        return shim.Success(nil)
}

func (s *SmartContract) getBagHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        bgId := args[0];



        resultsIterator, err := APIstub.GetHistoryForKey(bgId)
        if err != nil {
                return shim.Error("Error retrieving bag history.")
        }
        defer resultsIterator.Close()

        // buffer is a JSON array containing historic values for the marble
        var buffer bytes.Buffer
        buffer.WriteString("[")

        bArrayMemberAlreadyWritten := false
        for resultsIterator.HasNext() {
                response, err := resultsIterator.Next()
                if err != nil {
                        return shim.Error("Error retrieving bag history.")
                }
                // Add a comma before array members, suppress it for the first array member
                if bArrayMemberAlreadyWritten == true {
                        buffer.WriteString(",")
                }
                buffer.WriteString("{\"TxId\":")
                buffer.WriteString("\"")
                buffer.WriteString(response.TxId)
                buffer.WriteString("\"")

                buffer.WriteString(", \"Value\":")
                // if it was a delete operation on given key, then we need to set the
                //corresponding value null. Else, we will write the response.Value
                //as-is (as the Value itself a JSON marble)
                if response.IsDelete {
                        buffer.WriteString("null")
                } else {
                        buffer.WriteString(string(response.Value))
                }

                buffer.WriteString(", \"Timestamp\":")
                buffer.WriteString("\"")
                buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
                buffer.WriteString("\"")

                buffer.WriteString(", \"IsDelete\":")
                buffer.WriteString("\"")
                buffer.WriteString(strconv.FormatBool(response.IsDelete))
                buffer.WriteString("\"")

                buffer.WriteString("}")
                bArrayMemberAlreadyWritten = true
        }
        buffer.WriteString("]")

        fmt.Printf("- getBagHistory returning:\n%s\n", buffer.String())



        return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

        // Create a new Smart Contract
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }
}

