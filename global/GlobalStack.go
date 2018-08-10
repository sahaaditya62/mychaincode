/*
 * Copyright IBM Corp. 2018 All Rights Reserved.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 		 http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package global

/**
 * Project Name : THD
 * Created by Santanu on 04/08/18
 * EMAIL: santamit@in.ibm.com
**/

var FunctionPointers []string
var Line []string
var JSONData []map[string]interface{}
var ASNJSON []AssetJson

/*
	@Description {GlobalStack} - This is the identifier for the Global Stack and the reference for the Global Variables in the project
*/
type GlobalStack struct{}

type ASNJson struct {
	PONumber        string                   `json:"-"`
	ASNShipmentID   string                   `json:"-"`
	ASNNumber       string                   `json:"asnNumber"`
	ASNCreationDate string                   `json:"asnCreationDate"`
	Ifc             string                   `json:"ifc, omitempty"`
	AsnStatus       string                   `json:"asn_status, omitempty"`
	ASNItems        []AssetJsonASNItemsBlock `json:"asnItems"`
}

type AssetJsonASNItemsBlock struct {
	UCC128       string  `json:"ucc128"`
	Sku          string  `json:"sku"`
	VendorItemId string  `json:"vendorItemId"`
	QTY          float64 `json:"qty"`
	Rdc          string  `json:"rdc"`
}
type AssetJsonPoItemsBlock struct {
	//Key interface{} `json:"poItems"` //poItems : {format: ["sku":"1000054082", "vendorItemId":"GCM12SD","qty":2,"rdc":"5084"]}
	//Value interface{} `json:"value"`
	//PoItems []interface{} //(["sku":"", "vendorItemId":"", "qty", 0, "rdc":""])

	Sku          string  `json:"sku"`
	VendorItemId string  `json:"vendorItemId"`
	Qty          float64 `json:"qty"`
	Rdc          string  `json:"rdc"`
}

type AssetJson struct {
	PoNumber       string                  `json:"poNumber"`
	PartnerId      string                  `json:"partnerId"`
	PoCreationDate string                  `json:"poCreationDate"`
	Ifc            string                  `json:"ifc"`
	PoItems        []AssetJsonPoItemsBlock `json:"poItems"`
}

var FinalAssetJson AssetJson

type ValidatedAssetModel interface{}

/*
"ucc128": "",
	"sku": "",
	"vendorItemId": "",
	"qty": 0,
	"rdc": ""
*/
