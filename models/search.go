package models 

import (
    "owlhmaster/search"
)

// curl -X PUT \
//   https://52.47.197.22:50002/v1/search/getRulesetsBySearch \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "search": "v",
//     "rulesetName": "v"
//  }
func GetRulesetsBySearch(anode map[string]string)(data interface{}, err error) {
    data, err = search.GetRulesetsBySearch(anode)
    return data, err
}