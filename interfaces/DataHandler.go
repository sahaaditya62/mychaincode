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

package interfaces

/*
 * Project Name : THD
 * Created by Santanu on 04/08/18
 * EMAIL: santamit@in.ibm.com
 */

/*
	Description: This interface mostly deals with various data operations and file handling
*/
type DataHandler interface {
	FileReader(file string) (error, []string)
	FileWriter(file string) (error, bool)
	JSONBuilder(file string) []map[string]interface{}
	JSONBuilderWithBodyLine(body string, typeOfElement string) map[string]interface{}
	FindItemsInData(items []string, data []string) []string
}
