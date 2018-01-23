package configmap

import "github.com/Unknwon/goconfig"

//parse configmap to a map[string]map[string]string struct
//necessary: the necessary section names
//optional: the optional section names
//example: see config.ini and test
func ParseConfig(filename string, necessary []string, optional []string)  (map[string]map[string]string, error){
	c, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		return map[string]map[string]string{}, err
	}

	configMap := make(map[string]map[string]string)

	//necessary := []string{"sectionA"}
	for _, sectionName := range necessary {
		tempMap, err := c.GetSection(sectionName)
		if err != nil {
			return map[string]map[string]string{}, err
		} else {
			configMap[sectionName] = tempMap
		}
	}

	//optional := []string{"sectionB"}
	for _, sectionName := range optional {
		tempMap, err := c.GetSection(sectionName)
		if err != nil {
			continue
		} else {
			configMap[sectionName] = tempMap
		}
	}

	return configMap, nil
}
