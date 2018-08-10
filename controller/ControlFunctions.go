/*
    * * Copyright IBM Corp.  All Rights Reserved.
    * * Licensed under the Apache License, Version 2.0 (the "License");
    * * you may not use this file except in compliance with the License.
    * * You may obtain a copy of the License at
	* *	 http://www.apache.org/licenses/LICENSE-2.0
    * * Unless required by applicable law or agreed to in writing, software
    * * distributed under the License is distributed on an "AS IS" BASIS,
    * * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    * * See the License for the specific language governing permissions and
    * * limitations under the License.
*/

package controller

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/orlabdev/THD/global"
	"strconv"
)

/*
 Project Name : THD
 Created by Santanu on 06/08/18
 EMAIL: santamit@in.ibm.com
*/
var skuArray []string
var skuArrayFromASN []string
var ASNStatus map[string]string

type ControlFunctions struct {
	*DataController
}

func (*ControlFunctions) CreatePo() {}

func (*ControlFunctions) CompareWithUoM() {
	data := (*ControlFunctions).JSONBuilder(&ControlFunctions{}, "/Users/mgxc2/Documents/myGoProjects/src/github.com/orlabdev/THD/data/static/THD_PO_June.txt")
	_, strLines := (*ControlFunctions).FileReader(&ControlFunctions{}, "/Users/mgxc2/Documents/myGoProjects/src/github.com/orlabdev/THD/data/static/THD_UoM.txt")

	fmt.Print("\nEnter the Record you want to view (Line Number) : ")

	var i int

	fmt.Scanf("%d", &i)

	number := i

	var counter int
	var m = 0
	for _ = range data[number]["poItems"].([]interface{}) {
		fmt.Print()
		counter++
	}
	skuArray = make([]string, counter)

	for _, value := range data[number]["poItems"].([]interface{}) {
		//fmt.Println("\n",  key, "  ", value)
		for k, v := range value.(map[string]interface{}) {
			//fmt.Print("\n", k, "   ", v)

			if k == "sku" {
				skuArray[m] = v.(string)
				m++
			}

		}
		//counter++
	}
	fmt.Printf("\n Total number of PoItems : [ %d ] under PoNumber : [ %s ] ", len(skuArray), data[number]["poNumber"].(string))

	(*ControlFunctions).FindItemsInData(&ControlFunctions{}, skuArray, strLines)

	AssetJsonPoItemsBlock := (*ControlFunctions).BuildAssetJson(&ControlFunctions{}, data[number]["poItems"].([]interface{}))
	AssetJson := global.AssetJson{data[number]["poNumber"].(string), data[number]["partnerId"].(string),
		data[number]["poCreationDate"].(string), data[number]["ifc"].(string), AssetJsonPoItemsBlock}

	c, _ := json.MarshalIndent(AssetJson, "", "  ")
	_ = c

	var px global.ValidatedAssetModel
	t := json.Unmarshal(c, &px)
	_ = t
	//fmt.Print("\n", px)
	finalAssetJson := string(c)
	fmt.Print("\n", finalAssetJson)

	(*ControlFunctions).BuildASNJson(&ControlFunctions{}, data[number]["poNumber"].(string))
	//(*ControlFunctions).BuildASNJson(&ControlFunctions{})
	//found, value := (*GenericOps).FindItemsInData(&GenericOps{}, skuArray, strLines)
	//fmt.Print("\n", identifiedLines)

}

func (*ControlFunctions) CompareWithUoM_Working_Copy(stub shim.ChaincodeStubInterface, dataPO []byte, dataUOM []byte, objIdKey string) (bool, string) {
	//TODO: Copy code from the CompareWithUoM function
	//TODO: remove data[number] simply use data

	//poStructure := make([]interface{}, 0)
	var assetJson []global.AssetJson

	inStreamPOJSON := make(map[string]interface{})
	inStreamUOMJSON := make(map[string]interface{})
	err := json.Unmarshal(dataPO, &inStreamPOJSON)
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Print("\n", inStreamPOJSON)

	skuArrayFromPO := make([]string, 0)
	skuArrayFromUoM := make([]string, 0)
	//fmt.Print("\n", poItems)
	for _, v := range inStreamPOJSON["poItems"].([]interface{}) {
		//fmt.Print("\n", a, "    ", v)
		AssetJsonPoItemsBlock := (*ControlFunctions).BuildAssetJson(&ControlFunctions{}, inStreamPOJSON["poItems"].([]interface{}))
		AssetJson := global.AssetJson{inStreamPOJSON["poNumber"].(string), inStreamPOJSON["partnerId"].(string),
			inStreamPOJSON["poCreationDate"].(string), inStreamPOJSON["ifc"].(string), AssetJsonPoItemsBlock}

		assetJson = append(assetJson, AssetJson)

		for picker, value := range v.(map[string]interface{}) {
			//fmt.Print("\n", picker, "   ", value)
			if picker == "sku" {
				skuArrayFromPO = append(skuArrayFromPO, value.(string))
			}
		}

	}

	fmt.Print("\n", skuArrayFromPO)

	errUoM := json.Unmarshal(dataUOM, &inStreamUOMJSON)
	if errUoM != nil {
		fmt.Print(errUoM.Error())
	}

	for _, v := range inStreamUOMJSON["uomData"].([]interface{}) {
		for picker, value := range v.(map[string]interface{}) {
			if picker == "sku" {
				skuArrayFromUoM = append(skuArrayFromUoM, value.(string))
			}
		}
	}

	fmt.Print("\n", skuArrayFromUoM)

	//--------------- Compare Arrays ---------------//
	func() {
		for index, targetElement := range skuArrayFromPO {
			for _, referenceElement := range skuArrayFromUoM {
				if targetElement == referenceElement {

					global.FinalAssetJson = assetJson[index]

					break
				}
			}
		}
	}()

	c, _ := json.MarshalIndent(global.FinalAssetJson, "", "  ")
	_ = c

	putStateError := stub.PutState(objIdKey, c)
	if putStateError != nil {
		return false, "Revise data for error ..!! [ Status : Not Saved ]"
	}

	fmt.Print("\n", string(c))

	return true, "[PO] Message saved to ledger"
}

func (*ControlFunctions) BuildASNJson(poNumber string) {

	arrayOfPoNumber := make([]string, 1)
	arrayOfPoNumber[0] = poNumber

	//fmt.Print("\n", arrayOfPoNumber)
	//data := (*ControlFunctions).JSONBuilder(&ControlFunctions{}, "/Users/mgxc2/Documents/myGoProjects/src/github.com/orlabdev/THD/data/static/BOSCH_ASN_POLine_Json.txt")
	_, strLines := (*ControlFunctions).FileReader(&ControlFunctions{}, "/Users/mgxc2/Documents/myGoProjects/src/github.com/orlabdev/THD/data/static/BOSCH_ASN_POLine_Json.txt")
	matchedLines := (*ControlFunctions).FindItemsInData(&ControlFunctions{}, arrayOfPoNumber, strLines)
	//fmt.Print("\n", matchedLines)

	var px map[string]interface{}
	err := json.Unmarshal([]byte(matchedLines[0]), &px)
	if err != nil {
		fmt.Print(err)
	}

	var marker = 0

	var index = 0
	rx := px["asnItems"].([]interface{})
	//fmt.Printf("\n%#v", len(rx))
	elements := make([]map[string]interface{}, len(rx))
	cx := make([]global.AssetJsonASNItemsBlock, len(rx))
	skuArrayFromASN = make([]string, len(rx))

	ASNStatus = make(map[string]string, len(skuArrayFromASN))

	for _, v := range rx {
		//fmt.Print("\n", k, "    ", v)
		elements[index] = v.(map[string]interface{})

		skuArrayFromASN[marker] = elements[index]["sku"].(string)
		marker++

		cx[index] = global.AssetJsonASNItemsBlock{elements[index]["ucc128"].(string), elements[index]["sku"].(string),
			elements[index]["vendorItemId"].(string), elements[index]["qty"].(float64), elements[index]["rdc"].(string)}
		_ = cx[index]
		index++
	}

	fmt.Print("\nNo of ASN SKU # ", len(skuArrayFromASN), "\n [ ", skuArrayFromASN, " ] \nNo of PO Item SKU # [ ", len(skuArray), " ]\n", skuArray)

	fmt.Print("\n", cx)

	isFound, status := (*ControlFunctions).ValidateASN(&ControlFunctions{}, skuArrayFromASN, skuArray)
	fmt.Print("\n", isFound, "        ", status)

}

func (*ControlFunctions) ValidateASN(targetData []string, reference []string) (bool, map[string]string) {

	c := make(chan map[string]string)
	go func() {
		var isFound = false
		var index = 0
		tmp := make([]string, 0)
		for _, v := range targetData {
			//fmt.Println(sort.SearchStrings(reference, v))
			for _, r := range reference {
				if v == r {
					if !isFound {

						fmt.Printf("\n [ %s ] Found", v)
						ASNStatus[""+strconv.Itoa(index)+""] = v
						index++
						isFound = false
						break
					}
				} else {
					if !isFound {
						tmp = append(tmp, v)
						isFound = false
					}
				}
			}

		}
		c <- ASNStatus

		fmt.Print("\n", tmp, " -------- Not Found")
		tmp = tmp[:0]
	}()

	ASNStatus = <-c
	fmt.Print("\n", ASNStatus)
	return false, nil
}

/*/

 */
