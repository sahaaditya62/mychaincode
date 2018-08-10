/*
Copyright IBM Corp. 2018 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/orlabdev/THD/global"
	"github.com/orlabdev/THD/interfaces"
	"os"
	"strings"
)

/**
 * Project Name : THD
 * Created by Santanu on 04/08/18
 * EMAIL: santamit@in.ibm.com
 */

type DataController struct {
	DataHandler interfaces.DataHandler
	*global.GlobalStack
	//*model.DataModel
}

type ReturnItem struct {
	Status   bool   `json:"status"`
	DataLine string `json:"dataline"`
}

var counter int

func (dataController *DataController) FileReader(file string) (error, []string) {
	fileHandle, err := os.Open(file)

	if err != nil {
		return err, nil
	}
	defer fileHandle.Close()
	var marker int

	reader := bufio.NewReader(fileHandle)
	//go readrRoutine(reader)
	channelLineCount := make(chan int)
	go func() {
		for {
			_, err = reader.ReadString('\n')
			counter++
			if err != nil {
				break
			}
		}
		channelLineCount <- counter
		counter = 0
	}()

	counter := <-channelLineCount
	fmt.Print("\n# Of Lines : ", counter)

	global.Line = make([]string, counter)

	channelLineItems := make(chan []string)

	/*
		Description: Each line of the provided data files are complex JSON structure. Thus, to extract the right JSON out of the lines
		we need a non-Blocking code to read the data file, read the lines. As the data files are huge thus non-blocking
		codes are recommended otherwise the file pointers unexpectedly throwing EOF and terminating the operation
	*/
	go func() {
		fs, err := os.Open(file)
		if err != nil {
			return
		}
		defer fs.Close()
		r := bufio.NewReader(fs)
		for marker = 0; marker < counter; marker++ {
			//reader.Reset(reader)
			global.Line[marker], err = r.ReadString('\n')

			if err != nil {
				//fmt.Print("Error : ", err)
				break
			}
		}

		channelLineItems <- global.Line
	}()

	lineItems := <-channelLineItems

	//fmt.Print(global.Line [9])*/
	counter = 0
	return err, lineItems
}

func (dataController *DataController) FileWriter(file string) (error, bool) {
	panic("implement me")
}

func (dataController *DataController) JSONBuilder(file string) []map[string]interface{} {

	_, strLines := dataController.FileReader(file)
	global.JSONData = make([]map[string]interface{}, len(strLines))

	channelJSONData := make(chan []map[string]interface{})

	/*
		Description: goroutine to Unmarshal the JSON line-by-line. With goroutine this piece of code remains non-blocking but, concurrent
		tags: goroutine, non-Blocking, JSONBuilder
	*/
	go func() {
		var marker int

		for marker = 0; marker < len(strLines); marker++ {
			var data map[string]interface{}
			byteArray := []byte(strLines[marker])
			if err := json.Unmarshal(byteArray, &data); err != nil {
				break
			}
			global.JSONData[marker] = data
			//fmt.Println("\n marker : ", marker, "========= >>> ", global.JSONData[marker])
		}

		channelJSONData <- global.JSONData
		//fmt.Print("\n\n", global.JSONData[19]["poItems"])
	}()

	jsonItems := <-channelJSONData
	return jsonItems
}

func (dataController *DataController) FindItemsInData(items []string, data []string) []string {
	skuArray := make([]string, len(items))
	skuArray = items

	var filteredLines []string
	filteredLines = make([]string, 0)

	finderChannel := make(chan []string)

	go func() {
		var targetItem int
		var lineMarker int
		var n = 0
		for targetItem = 0; targetItem < len(items); targetItem++ {
			for lineMarker = 0; lineMarker < len(data); lineMarker++ {
				if strings.Contains(data[lineMarker], skuArray[targetItem]) {
					lineItem := ReturnItem{Status: true, DataLine: data[lineMarker]}
					filteredLines = append(filteredLines, lineItem.DataLine)
					//fmt.Print("\n", filteredLines[n])
					n++
				}
			}
		}
		if len(filteredLines) > 0 {
			finderChannel <- filteredLines
		} else {
			finderChannel <- nil
		}

	}()

	matchedRegion := <-finderChannel

	filteredLines = filteredLines[:0]
	return matchedRegion
}

func (dataController *DataController) BuildAssetJson(poItems []interface{}) []global.AssetJsonPoItemsBlock {

	//for _, v := range data {
	var n = 0
	b := make([]global.AssetJsonPoItemsBlock, len(poItems))

	for _, v := range poItems {
		r := v.(map[string]interface{})
		b[n] = global.AssetJsonPoItemsBlock{r["sku"].(string), r["vendorItemId"].(string), r["qty"].(float64), r["rdc"].(string)}
		n++

	}

	return b
}

func (*ControlFunctions) JSONBuilderWithBodyLine(body string, typeOfElement string) (map[string]interface{}, []map[string]interface{}) {

	if typeOfElement == "interface{}" {
		data := make(map[string]interface{}, len(body))
		if err := json.Unmarshal([]byte(body), &data); err != nil {
			fmt.Print(err)
		} else {
			return data, nil
		}
	}

	return nil, nil
}

/*
   "asnCreationDate": "",
   "asnNumber": "",
   "asnShipmentID": "",
   "asnItems": [
       {
           "ucc128": "",
           "sku": "",
           "vendorItemId": "",
           "qty": 0,
           "rdc": ""
       }
   ],
   "asn_status": "",
*/
/*
poNumber string `json:"poNumber"`
partnerId string `json:"partnerId"`
poCreationDate string `json:"poCreationDate"`
ifc string `json:"ifc"`
poItems types.Array `json:"poItems"`
*/
