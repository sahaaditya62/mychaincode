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

package interfaces

/*
 * Project Name : THD
 * Created by Santanu on 04/08/18
 * EMAIL: santamit@in.ibm.com
 */

/*
	Functions - this interface describes the variadic functions for individual operations

	Member Functions:

		SetOperationNames - is the primary variadic functions with arguments representing

		@parameters:

		arg[0] - name of the function
		arg[1 ... n] - the parameters to the function

		GetOperationNames - returns the Name of the function along with the parameters to it
		WorkerMethodsAndParameters - function selector and calls the specific function from the implementation
*/
type Functions interface {
	SetOperationNames(name ...string)
	GetOperationNames() []string
	WorkerMethodsAndParameters()
	BuildAssetJson([]map[string]interface{})
	BuildASNJson(poNumber string)
	ValidateASN(targetData []string, compareDataBase []string) (bool, map[string]string)
}
