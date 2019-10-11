package search

import(
	// "regexp"
    // "os"
	// "os/exec"
	// "encoding/json"
	"owlhmaster/database"
	"time"
    "owlhmaster/ruleset"
    // "errors"
	// "database/sql"
    "strings"
	// "strconv"
	"regexp"
	// "io/ioutil"
    "github.com/astaxie/beego/logs"
)


//Struct for search over a ruleset files
type Rule struct {
	Sid 	  	string		`json:"sid"`
	Msg       	string		`json:"msg"`
	Rulesets  	[]Ruleset	
}
type Ruleset struct {
	File      	string		`json:"file"`
	Status    	string		`json:"status"`
	Name    	string		`json:"name"`
	Uuid    	string		`json:"uuid"`
	Node    	[]string	`json:"node"`
}
var RulesIndex []Rule = nil

func GetRulesetsBySearch(anode map[string]string)(data []Rule, err error) {
	logs.Debug(anode["search"])
	var count int
	var matchingRules []Rule = nil
	var isSID = regexp.MustCompile(`(\d+)`)
	sid := isSID.FindStringSubmatch(anode["search"])
	
	for w := range RulesIndex {
		if sid != nil{
			if strings.Contains(RulesIndex[w].Sid, anode["search"]){
				currentRulesets := RulesIndex[w].Rulesets
				if anode["rulesetName"] == ""{
					// for z := range currentRulesets {
						matchingRules = append(matchingRules, RulesIndex[w])
					// }
				}else{					
					for z := range currentRulesets {
						if currentRulesets[z].Name == anode["rulesetName"]{
							// if RulesIndex[w].Rulesets.Uuid != currentRulesets[z].Uuid{
								matchingRules = append(matchingRules, RulesIndex[w])
							// }
						}
					}
				}
			}
		}else {
			if strings.Contains(strings.ToLower(RulesIndex[w].Msg), strings.ToLower(anode["search"])){
				count++
				currentRulesets := RulesIndex[w].Rulesets
				if anode["rulesetName"] == ""{
					// for z := range currentRulesets {
					// 	Rsets := Ruleset{}
					// 	Rsets.Name = currentRulesets[z].Name
					// 	Rsets.Status = currentRulesets[z].Status
					// 	Rsets.File = currentRulesets[z].File
					// 	Rsets.Uuid = currentRulesets[z].Uuid
						
					// 	var NewRule Rule
					// 	NewRule.Sid = RulesIndex[w].Sid
					// 	NewRule.Msg = RulesIndex[w].Msg
					// 	NewRule.Rulesets = append(NewRule.Rulesets, Rsets) 

						// matchingRules = append(matchingRules, NewRule)
						matchingRules = append(matchingRules, RulesIndex[w])
					// }
				}else{
					for z := range currentRulesets {
						if currentRulesets[z].Name == anode["rulesetName"]{
							// Rsets := Ruleset{}
							// Rsets.Name = currentRulesets[z].Name
							// Rsets.Status = currentRulesets[z].Status
							// Rsets.File = currentRulesets[z].File
							// Rsets.Uuid = currentRulesets[z].Uuid

							// var NewRule Rule
							// NewRule.Sid = RulesIndex[w].Sid
							// NewRule.Msg = RulesIndex[w].Msg
							// NewRule.Rulesets = append(NewRule.Rulesets, Rsets) 
	
							// matchingRules = append(matchingRules, NewRule)
							matchingRules = append(matchingRules, RulesIndex[w])
						}
					}
				}
			}	
		}
	}
	logs.Warn(count)
	return matchingRules, err
}

func Init()(){
	for {

		RulesIndex = nil
		exists := false
		allRulesets,err := ndb.GetAllRuleFiles()
		if err != nil {logs.Error("Search/Init error: %s", err.Error())}
		for x,_ := range allRulesets {	
			rset := Ruleset{}
			currentRules, _ := ruleset.ReadRuleset(allRulesets[x]["path"])
			rset.File = allRulesets[x]["path"]
			rset.Name = allRulesets[x]["name"]
			for y := range currentRules {				
				rset.Status = currentRules[y]["enabled"]
				rset.Uuid = x
				rule := Rule{}
				rule.Rulesets = append(rule.Rulesets, rset)
				rule.Sid = currentRules[y]["sid"]
				rule.Msg = currentRules[y]["msg"]
				exists = false

				for w := range RulesIndex {
					if RulesIndex[w].Sid == rule.Sid{	
						RulesIndex[w].Rulesets = append(RulesIndex[w].Rulesets, rset)
						exists=true
						break
					}
				}
				if !exists {
					RulesIndex = append(RulesIndex, rule)
				}
			}
		}

		// for w := range RulesIndex {

		// }

		logs.Info("Ruleset list has been updated.")
		time.Sleep(5 * time.Minute)
	}
}




// nodes,err := ndb.GetAllNodes()
// if err != nil {logs.Error("Error getting all ruleset values: "+err.Error())}
// for f := range nodes{
// 	if allRulesets[x]["sourceUUID"] == f{
// 		nodeName,err := ndb.ObtainNodeName(nodes[f])
// 		if err != nil {logs.Error("Error getting node name from their ruleset: "+err.Error())}
// 		logs.Warn(nodes[f]+"  ->  "+nodeName)
		
		
// 		// // isNodeYet := false
// 		// for r := range rset.Node{
// 		// 	if rset.Node[r] == nodeName{
// 		// 		// isNodeYet = true
// 		// 		break;
// 		// 	}else{
// 		// 		rset.Node = append(rset.Node, nodeName)
// 		// 	}
// 		// }
// 		// // if !isNodeYet{
// 		// // }
// 	}
// }