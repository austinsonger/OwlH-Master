package models 

import (
    "owlhmaster/changeControl"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/changecontrol/ \
// }
func GetChangeControl()(data map[string]map[string]string, err error) {
    data, err = changecontrol.GetChangeControl()
    return data, err
}