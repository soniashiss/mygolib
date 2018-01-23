package configmap

import "testing"

func TestParseConfig(t *testing.T) {

	wantMap := map[string]map[string]string{
		"sectionA":
			{"aa": "bb",},
		"sectionB":
			{"cc": "dd",},
	}

	funcMap, err := ParseConfig("config.ini", []string{"sectionA"}, []string{"sectionB"})
	if err != nil {
		t.Errorf("ParseConfig get err: %v", err.Error())
	} else {
		for k, v := range funcMap {
			for kk, vv := range v {
				ok1 := false
				ok2 := false
				_, ok1 = wantMap[k]
				if ok1 {
					_, ok2 = wantMap[k][kk]
				}
				if !ok1 || !ok2 || vv != wantMap[k][kk] {
					t.Errorf("Map not Match. Get %v, Want %v", vv, wantMap[k][kk])
				}
			}
		}
		//fmt.Println(funcMap)
	}
}