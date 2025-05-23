package utilities

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type JSONMap map[string]any

func ParseGraphQLQueryCopy(bodyString string) (JSONMap, error) {
	allFieldMap, entitlementIdMap, err := ParseSchema()
	if err != nil {
		return nil, fmt.Errorf("Error parsing schema: " + err.Error())
	}

	policiesList, err := GetApolloPoliciesRequiredHeders()
	if err != nil {
		return nil, fmt.Errorf("Error Getting Policies schema: " + err.Error())
	}
	if len(policiesList) == 0 {
		return nil, fmt.Errorf("no valid policies provided")
	}
	var data JSONMap
	err = json.Unmarshal([]byte(bodyString), &data)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON body: " + err.Error())
	}
	policyMap := make(map[string]map[string]map[string]any)

	for _, policy := range policiesList {
		parts := splitPoliciesAndRemoveSpace(policy["key"].(string), ".")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid policy format")
		}
		typename := parts[0]
		field := parts[1]

		if typename == "Query" {
			if dataField, ok := data["data"].(map[string]any); ok {
				delete(dataField, field)
			}
			continue
		}

		engineResponse, error := getEngineResponseBasedOnPolicy(policy["key"].(string))
		if error != nil {
			return nil, fmt.Errorf("Error getting engine response: " + error.Error())

		}

		if _, ok := policyMap[typename]; !ok {
			policyMap[typename] = make(map[string]map[string]any)
		}
		node, ok := policy["node"].(map[string]any)
		var entitlementVal string
		if !ok {
			entitlementVal = ResolveRefIdNameFallback(typename)
		} else {
			entitlementVal = node["entitlementIdentifier"].(string)
		}
		policyMap[typename][field] = make(map[string]any)
		policyMap[typename][field]["engineresponse"] = engineResponse
		policyMap[typename][field]["entitlementIdentifier"] = entitlementVal
	}

	if len(policyMap) != 0 {
		data = traverseAndRedactCopy(data["data"].(map[string]any), allFieldMap, policyMap, entitlementIdMap, "", "")
	}
	return data, nil
}

func traverseAndRedactCopy(jsonMap map[string]interface{}, fieldMap map[string]string, policyMap map[string]map[string]map[string]any, entitlementIdMap map[string]string, typename string, refid string) map[string]interface{} {
	for key, value := range jsonMap {
		if typename != "" {
			normalizedType := normalizeTypeName(typename)
			var refIdField string
			if normalizeTypeName(typename) == "Account" {
				if _, ok := policyMap[normalizedType][key]; ok {
					refIdField = policyMap[normalizedType][key]["entitlementIdentifier"].(string)
				} else {
					refIdField = ResolveRefIdNameFallback(normalizedType)
				}
				if tempRefid := findReferenceID(jsonMap, refIdField); tempRefid != "" {
					refid = tempRefid
				}
				//if refIdVal, ok := jsonMap["accountReferenceId"].(string); ok {
				//refid = refIdVal
				//}
			}
			if normalizeTypeName(typename) == "Card" {
				if refIdVal, ok := jsonMap["cardReferenceId"].(string); ok {
					refid = refIdVal
				}
			}
			if fieldsMap, exists := policyMap[normalizedType]; exists {
				if engineResponse, fieldExists := fieldsMap[key]["engineresponse"].(map[string]string); fieldExists {
					fmt.Println(reflect.TypeOf(value))
					switch accounValue := value.(type) {
					case []interface{}:
						var filtered []interface{}
						for _, obj := range accounValue {
							if m, ok := obj.(map[string]interface{}); ok {
								refIdField = policyMap[normalizedType][key]["entitlementIdentifier"].(string)
								refid = findReferenceID(m, refIdField)
								//refid, _ := m["accountReferenceId"].(string)
								if processEngineResonse(engineResponse, refid) {
									filtered = append(filtered, m)
								}
							}
						}
						jsonMap[key] = filtered
					case any:
						if !processEngineResonse(engineResponse, refid) {
							delete(jsonMap, key)
							continue
						}
					}
				}
			}
		}
		switch v := jsonMap[key].(type) {
		case map[string]interface{}:
			jsonMap[key] = traverseAndRedactCopy(v, fieldMap, policyMap, entitlementIdMap, fieldMap[key], refid)
		case []interface{}:
			for i, item := range v {
				if obj, ok := item.(map[string]interface{}); ok {
					v[i] = traverseAndRedactCopy(obj, fieldMap, policyMap, entitlementIdMap, fieldMap[key], refid)
				}
			}
			jsonMap[key] = v
		}
	}
	return jsonMap
}

func normalizeTypeName(name string) string {
	re := regexp.MustCompile(`\[|\]`) // Remove brackets
	return re.ReplaceAllString(name, "")
}

func findReferenceID(data map[string]interface{}, refIdName string) string {
	if refIdName == "" {
		return ""
	}
	for key, value := range data {
		if key == refIdName { //check if refName is not empty
			if str, ok := value.(string); ok {
				return str
			}
		}
		if nestedMap, ok := value.(map[string]interface{}); ok {
			if ref := findReferenceID(nestedMap, refIdName); ref != "" {
				return ref
			}
		}
		if array, ok := value.([]interface{}); ok {
			for _, item := range array {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if ref := findReferenceID(itemMap, refIdName); ref != "" {
						return ref
					}
				}
			}
		}
	}
	return ""
}

func splitPoliciesAndRemoveSpace(policies string, delimeter string) []string {
	parts := strings.Split(policies, delimeter)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func getEngineResponseBasedOnPolicy(policy string) (map[string]string, error) {
	permissions := map[string]map[string]string{
		"Account.balance": {
			"acc123": "ALLOW",
			"acc456": "DENY",
		},
		"Card.cardNumber": {
			"card123": "ALLOW",
			"card456": "DENY",
		},
		"AvailableCreditAmount.availableSpendingCreditAmount": {
			"acc123":  "ALLOW",
			"acc456":  "DENY",
			"card123": "DENY",
			"card456": "ALLOW",
		},
		"Account.availableCreditAmount": {
			"acc123": "ALLOW",
			"acc456": "DENY",
		},
		"Card.cardReferenceId": {
			"card123": "ALLOW",
			"card456": "DENY",
		},
		"Customer.accounts": {
			"acc123": "ALLOW",
			"acc456": "DENY",
		},
	}

	if val, ok := permissions[policy]; ok {
		return val, nil
	}
	//return nil, errors.New("no engine response found for policy: " + policy)
	return nil, nil
}

func processEngineResonse(engineResonse map[string]string, refids string) bool {
	valueInEngineResponse := engineResonse[refids]
	if valueInEngineResponse == "ALLOW" {
		return true
	} else if valueInEngineResponse == "DENY" {
		return false
	} else {
		return false
	}
}

func GetApolloPoliciesRequiredHeders() ([]map[string]interface{}, error) {
	jsonData := `
	{
	"version": 1,
	"stage": "SupergraphRequest",
	"control": "continue",
	"id": "ee6c1a03-8ddb-41f2-b30a-0583b0d07a4d",
	"headers": {
		"accept": ["/"],
		"accept-encoding": ["gzip, deflate, br, zstd"],
		"accept-language": ["en-US,en;q=0.9"],
		"connection": ["keep-alive"],
		"content-length": ["358"],
		"content-type": ["application/json"],
		"customeridtoken": ["CustIDToken-DENY"],
		"host": ["localhost:4000"],
		"origin": ["https://studio.apollographql.com"],
		"sec-ch-ua": ["\"Google Chrome\";v=\"135\", \"Not-A.Brand\";v=\"8\", \"Chromium\";v=\"135\""],
		"sec-ch-ua-mobile": ["?0"],
		"sec-ch-ua-platform": ["\"macOS\""],
		"sec-fetch-dest": ["empty"],
		"sec-fetch-mode": ["cors"],
		"sec-fetch-site": ["cross-site"],
		"user-agent": ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36"]
	},
	"body": {
		"query": "query AccountQuery($customerReferenceId: String!) {\\n customers {\\n address\\n age\\n name\\n referenceId\\n }\\n accounts(customerReferenceId: $customerReferenceId) {\\n countryCode\\n referenceId\\n status\\n }\\n}",
		"operationName": "AccountQuery",
		"variables": {
		"accountReferenceId": "zzzzzzzzzz1",
		"customerReferenceId": "aaaaaaaaaa"
		}
	},
	"context": {
		"entries": {
		"apollo_authorization::policies::required": {
			"{ \"key\": \"Account.balance\"}": null
		},
		"operation_kind": "query",
		"operation_name": "AccountQuery"
		}
	}
	}
	`
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	keys := []map[string]interface{}{}
	if contextMap, ok := result["context"].(map[string]interface{}); ok {
		if entriesMap, ok := contextMap["entries"].(map[string]interface{}); ok {
			if requiredPolicies, ok := entriesMap["apollo_authorization::policies::required"].(map[string]interface{}); ok {
				for k := range requiredPolicies {
					var obj map[string]interface{}
					if err := json.Unmarshal([]byte(k), &obj); err == nil {
						keys = append(keys, obj)
					}
				}
			}
		}
	}
	return keys, nil
}
