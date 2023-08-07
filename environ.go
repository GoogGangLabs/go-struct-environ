package environ

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Load(path string, environStruct interface{}) (error) {
	structPtr := reflect.ValueOf(environStruct)
	if structPtr.Kind() != reflect.Ptr { return errors.New("argument must be pointer") }
	structElem := structPtr.Elem()
	if structElem.Kind() != reflect.Struct { return errors.New("argument's type must be struct") }

	file, err := os.Open(path)
	if err != nil { return err }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineBypassRegexp, _ := regexp.Compile(`^[\s]*$|^$|^#`)
	lineRuleRegexp, _ := regexp.Compile(`^[^=\s]+=[^=\s]+$`)
	var isAllowLine bool

	for scanner.Scan() {
		line := scanner.Text()

		isAllowLine = lineBypassRegexp.MatchString(line)
		if isAllowLine { continue }

		isAllowLine = lineRuleRegexp.MatchString(line)
		if !isAllowLine { return errors.New("invalid env line") }

		entry := strings.Split(line, "=")

		key, value := entry[0], entry[1]
		field := structElem.FieldByName(key)
		if !field.Comparable() { continue }

		os.Setenv(key, value)
		fieldValue := os.Getenv(key)

		if field.CanInt() {
			intValue, err := strconv.ParseInt(fieldValue, 10, 64)
			if err != nil { return fmt.Errorf("[%s] must be an integer", key) }
			field.SetInt(intValue)
		} else {
			if fieldValue == "" { return fmt.Errorf("[%s] must not be blank", key) }
			field.SetString(fieldValue)
		}
	}

	for i := 0; i < structElem.NumField(); i++ {
		key := structElem.Type().Field(i).Name
		field := structElem.Field(i)
		
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() == 0 { return fmt.Errorf("[%s] must not be empty", key) }
		case reflect.String:
			if field.Len() == 0 { return fmt.Errorf("[%s] must not be empty", key) }
		default:
			return fmt.Errorf("[%s] must be string or int", key)
		}
	}

	return nil
}
