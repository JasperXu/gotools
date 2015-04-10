package gotools

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type Rule struct {
	Func      RuleFunc
	Operator  string
	ParamNums int
}

type RuleFunc func(value interface{}, params ...interface{}) bool

var rules map[string][]Rule = make(map[string][]Rule, 0)

func init() {
	//rules of string
	item := Rule{}
	item.Operator = "Required"
	item.ParamNums = 0
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		if len(strValue) == 0 {
			return false
		}
		return true
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "IsEmail"
	item.ParamNums = 0
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		m, _ := regexp.MatchString("^[\\w-]+(\\.[\\w-]+)*@[\\w-]+(\\.[\\w-]+)+$", strValue)
		return m
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "IsPhone"
	item.ParamNums = 0
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		m, _ := regexp.MatchString("^1[0-9]{10}$", strValue)
		return m
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "IsIP"
	item.ParamNums = 0
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		m, _ := regexp.MatchString("^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])$", strValue)
		return m
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "IsIDCard" //身份证
	item.ParamNums = 0
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		m, _ := regexp.MatchString("^\\d{14}\\d{3}?\\w$", strValue)
		return m
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "IsMoney"
	item.ParamNums = 0
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		m, _ := regexp.MatchString("^[+-]{0,1}\\d{1,8}\\.\\d{2}$", strValue)
		return m
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "IsTime"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		param := params[0].(string)
		_, err := time.Parse(param, strValue)
		if err != nil {
			return false
		}
		return true
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "len>="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		param := params[0].(int)
		if len(strValue) >= param {
			return true
		}
		return false
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "len<="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		param := params[0].(int)
		if len(strValue) <= param {
			return true
		}
		return false
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "len>"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		param := params[0].(int)
		if len(strValue) > param {
			return true
		}
		return false
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "len<"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		param := params[0].(int)
		if len(strValue) < param {
			return true
		}
		return false
	}
	AddRule("string", item)

	item = Rule{}
	item.Operator = "lenRange"
	item.ParamNums = 2
	item.Func = func(value interface{}, params ...interface{}) bool {
		strValue := value.(string)
		paramA := params[0].(int)
		paramB := params[1].(int)
		if len(strValue) >= paramA && len(strValue) <= paramB {
			return true
		}
		return false
	}
	AddRule("string", item)

	//rules of int64
	item = Rule{}
	item.Operator = ">"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(int64)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue > param {
			return true
		}
		return false
	}
	AddRule("int64", item)

	item = Rule{}
	item.Operator = ">="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(int64)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue >= param {
			return true
		}
		return false
	}
	AddRule("int64", item)

	item = Rule{}
	item.Operator = "<"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(int64)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue < param {
			return true
		}
		return false
	}
	AddRule("int64", item)

	item = Rule{}
	item.Operator = "<="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(int64)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue <= param {
			return true
		}
		return false
	}
	AddRule("int64", item)

	item = Rule{}
	item.Operator = "range"
	item.ParamNums = 2
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(int64)
		paramA, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		paramB, converted := ConvToInt64(params[1])
		if !converted {
			return false
		}
		if trueValue >= paramA && trueValue <= paramB {
			return true
		}
		return false
	}
	AddRule("int64", item)

	//rules of int
	item = Rule{}
	item.Operator = ">"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue, _ := ConvToInt64(value)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue <= param {
			return true
		}
		return false
	}
	AddRule("int", item)

	item = Rule{}
	item.Operator = ">="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue, _ := ConvToInt64(value)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue >= param {
			return true
		}
		return false
	}
	AddRule("int", item)

	item = Rule{}
	item.Operator = "<"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue, _ := ConvToInt64(value)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue < param {
			return true
		}
		return false
	}
	AddRule("int", item)

	item = Rule{}
	item.Operator = "<="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue, _ := ConvToInt64(value)
		param, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		if trueValue <= param {
			return true
		}
		return false
	}
	AddRule("int", item)

	item = Rule{}
	item.Operator = "range"
	item.ParamNums = 2
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue, _ := ConvToInt64(value)
		paramA, converted := ConvToInt64(params[0])
		if !converted {
			return false
		}
		paramB, converted := ConvToInt64(params[1])
		if !converted {
			return false
		}
		if trueValue >= paramA && trueValue <= paramB {
			return true
		}
		return false
	}
	AddRule("int", item)

	//rules of time.Time
	item = Rule{}
	item.Operator = "IsMinTime"
	item.ParamNums = 0
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(time.Time)
		minTime, _ := time.Parse("2006-01-02 15:04:05 -0700", "2000-01-01 00:00:00 +0000")
		duration := trueValue.Sub(minTime)
		if duration == 0 {
			return true
		}
		return false
	}
	AddRule("time", item)

	item = Rule{}
	item.Operator = "<="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(time.Time)
		param := params[0].(time.Time)
		duration := trueValue.Sub(param)
		if duration <= 0 {
			return true
		}
		return false
	}
	AddRule("time", item)

	item = Rule{}
	item.Operator = "<"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(time.Time)
		param := params[0].(time.Time)
		duration := trueValue.Sub(param)
		if duration < 0 {
			return true
		}
		return false
	}
	AddRule("time", item)

	item = Rule{}
	item.Operator = ">="
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(time.Time)
		param := params[0].(time.Time)
		duration := trueValue.Sub(param)
		if duration >= 0 {
			return true
		}
		return false
	}
	AddRule("time", item)

	item = Rule{}
	item.Operator = ">"
	item.ParamNums = 1
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(time.Time)
		param := params[0].(time.Time)
		duration := trueValue.Sub(param)
		if duration > 0 {
			return true
		}
		return false
	}
	AddRule("time", item)

	item = Rule{}
	item.Operator = "range"
	item.ParamNums = 2
	item.Func = func(value interface{}, params ...interface{}) bool {
		trueValue := value.(time.Time)
		paramA := params[0].(time.Time)
		paramB := params[1].(time.Time)
		durationA := trueValue.Sub(paramA)
		durationB := trueValue.Sub(paramB)
		if durationA >= 0 && durationB <= 0 {
			return true
		}
		return false
	}
	AddRule("time", item)
}

func ConvToInt64(i interface{}) (int64, bool) {
	iValue := reflect.ValueOf(i)

	defer func() {
		recover()
	}()

	return iValue.Int(), true
}

//添加验证规则
func AddRule(typeName string, item Rule) {
	if item.Operator != "" && item.Func != nil {
		typeName = strings.ToLower(typeName)
		items := rules[typeName]
		items = append(items, item)
		rules[typeName] = items
	}
}

//删除验证规则
func RemoveRule(typeName, operator string) bool {
	typeName = strings.ToLower(typeName)
	items, ok := rules[typeName]
	if !ok {
		fmt.Println("TypeName:", typeName, "的Rule不存在！")
		return false
	}
	index := -1
	for i, v := range items {
		if v.Operator == operator {
			index = i
			break
		}
	}
	if index < 0 {
		fmt.Println("TypeName:", typeName, "Operator:", operator, "的Rule不存在！")
		return false
	}
	prevSlice := items[0:index]
	nextSlice := items[index+1 : len(rules)]
	items = append(prevSlice, nextSlice...)
	rules[typeName] = items
	return true
}

//验证
func Check(input interface{}, params ...interface{}) bool {
	typeName := reflect.TypeOf(input).Name()
	typeName = strings.ToLower(typeName)
	if len(params) == 0 {
		return false
	}
	return checkRules(input, typeName, params...)
}

//检测参数是否符合定义的规则
func checkRules(input interface{}, typeName string, params ...interface{}) bool {
	i := 0
	for i < len(params) {
		items, ok := rules[typeName]
		if !ok || len(items) == 0 {
			return false
		}
		index, item := indexOfRule(items, params[i].(string))
		if index < 0 {
			return false
		}
		i = i + 1

		if checked := item.Func(input, params[i:i+item.ParamNums]...); !checked {
			return false
		}
		i = i + item.ParamNums
	}
	return true
}

//查找指定类型名称和指定操作符的规则
func indexOfRule(items []Rule, operator string) (index int, item *Rule) {
	for i, v := range items {
		if v.Operator == operator {
			return i, &v
		}
	}
	return -1, nil
}
