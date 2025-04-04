package core

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func BooleanTo[T interface{}](p bool, trueValue T, falseValue T) T {
	if p {
		return trueValue
	} else {
		return falseValue
	}
}
func BooleanFun(p bool, trueFun func(), falseFunc ...func()) {
	if p {
		trueFun()
	} else {
		for _, f := range falseFunc {
			f()
		}
	}
}

func StringToInt32(strings []string) []int32 {
	int32Slice := make([]int32, len(strings))
	for i, str := range strings {
		num, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			fmt.Printf("转换失败，无效的字符串: %s\n", str)
			return nil
		}
		int32Slice[i] = int32(num)
	}

	return int32Slice
}
func Int32ToString(numbers []int32) []string {
	result := make([]string, len(numbers))
	for i, num := range numbers {
		result[i] = strconv.Itoa(int(num))
	}
	return result
}
func StringExist(slice []string, target string) bool {
	found := false
	for _, s := range slice {
		if s == target {
			found = true
			break
		}
	}
	return found
}

func ToInterface[T any](slc []T) []interface{} {
	result := make([]interface{}, len(slc))
	for i, v := range slc {
		result[i] = v
	}
	return result
}

func SnakeToLowerCamelCase(columnName string) string {
	// 去掉下划线和横杠，并将后面的字母改为大写
	var modifiedName string
	for i := 0; i < len(columnName); i++ {
		if columnName[i] == '_' {
			if i+1 < len(columnName) {
				modifiedName += strings.ToUpper(string(columnName[i+1]))
				i++
			}
		} else {
			modifiedName += string(columnName[i])
		}
	}

	return modifiedName
}
func LowerCamelCaseToSnake(columnName string) string {

	// 使用正则表达式将大写字母前面插入下划线，并将所有字母转换为小写
	re := regexp.MustCompile(`(.)([A-Z][a-z]+)`)
	converted := re.ReplaceAllString(columnName, "${1}_${2}")

	// 将所有字母转换为小写
	converted = strings.ToLower(converted)

	return converted

}
func SnackToPath(columnName string) string {
	return strings.ReplaceAll(columnName, "_", "/")
}

func SnackLastName(columnName string) string {
	split := strings.Split(columnName, "_")
	return split[len(split)-1]
}

func SnakeToUpperCamelCase(columnName string) string {
	lowUpStr := SnakeToLowerCamelCase(columnName)
	upUpStr := strings.Title(lowUpStr)
	return upUpStr
}
func GoTypeConversion(columnType string) string {
	m := map[string]string{
		"varchar": "string",
		"int":     "int",
	}
	goLangType := m[columnType]
	if goLangType == "" {
		goLangType = "string"
	}
	return goLangType
}
func TsTypeConversion(columnType string) string {
	m := map[string]string{
		"varchar": "string",
		"int":     "number",
	}
	goLangType := m[columnType]
	if goLangType == "" {
		goLangType = "string"
	}
	return goLangType
}

func CopyFrom[T any](key any) T {
	t := new(T)
	_ = copier.Copy(t, key)
	return *t
}

func CopyListFrom[T any, F any](target []F) []T {
	result := make([]T, 0, len(target))
	for _, item := range target {
		var newItem T
		// 使用 copier 进行复制
		if err := copier.Copy(&newItem, item); err != nil {
			fmt.Println("复制失败:", err)
			continue
		}
		result = append(result, newItem)
	}
	return result
}

// HasField 检查结构体是否包含指定字段
func HasField(obj interface{}, fieldName string) bool {
	v := reflect.ValueOf(obj)
	// 确保传入的是结构体
	if v.Kind() != reflect.Struct {
		return false
	}

	// 获取字段
	field := v.FieldByName(fieldName)
	return field.IsValid()
}

func SetField(obj interface{}, fieldName string, value interface{}) error {
	// 获取对象的反射值
	v := reflect.ValueOf(obj).Elem()
	// 获取字段
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("无效的字段名: %s", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("字段 %s 不能被设置", fieldName)
	}

	// 设置值
	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return fmt.Errorf("无法将 %v 设置为 %s, 类型不匹配", value, fieldName)
	}

	field.Set(val)
	return nil
}

func Instance[T any]() (result T) {
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	return *thisResult
}

// GetFieldValueSlice 从一个结构体切片中提取指定字段的值，返回一个该字段值的切片。
// T 是一个泛型参数，表示返回的切片中元素的类型。
// slice 参数是输入的结构体切片，fieldName 是可变参数，用于指定要提取的字段名。
// 如果 fieldName 没有提供，则默认提取 "UID" 字段。
func GetFieldValueSlice[T any](slice interface{}, fieldName ...string) []T {
	// 设置默认字段名为 "UID"，如果提供了 fieldName 参数，则使用第一个作为字段名。
	_fieldName := "UID"
	if len(fieldName) > 0 {
		_fieldName = fieldName[0]
	}

	// 获取 slice 的反射值。
	v := reflect.ValueOf(slice)
	// 检查反射值的类型是否为切片，如果不是，则记录错误并返回空切片。
	if v.Kind() != reflect.Slice {
		zap.L().Error("GetFieldValueSlice: not a slice")
		return []T{}
	}

	// 初始化一个用于存储提取字段值的切片。
	var values []T
	// 遍历切片中的每个元素。
	for i := 0; i < v.Len(); i++ {
		// 获取当前元素的反射值。
		item := v.Index(i)
		// 尝试根据字段名获取当前元素的指定字段的反射值。
		field := item.FieldByName(_fieldName)
		// 如果字段有效（即存在该字段），则将字段的值转换为 T 类型，并添加到 values 切片中。
		if field.IsValid() {
			values = append(values, field.Interface().(T))
		} else {
			// 如果字段无效（即不存在该字段），则返回空切片。
			return []T{}
		}
	}
	// 返回存储了所有元素的指定字段值的切片。
	return values
}

func GetSliceLast[T any](slice []T, compareValue any, fieldName ...string) T {
	var _fieldName string = "UID"

	if len(fieldName) > 0 {
		_fieldName = fieldName[0]
	}
	var zeroValue T // 用于存储默认值
	// 确保切片不为空
	if len(slice) == 0 {
		return zeroValue
	}
	// 遍历切片，从后向前查找
	for i := len(slice) - 1; i >= 0; i-- {
		v := reflect.ValueOf(slice[i])
		field := v.FieldByName(_fieldName)

		// 确保字段有效
		if field.IsValid() && field.Interface() == compareValue {
			return slice[i]
		}
	}
	return zeroValue
}

// GetSliceLastPointer 通过反射在切片中找到最后一个与指定比较值匹配的元素指针。
// 这个函数适用于任何类型的切片，只要切片的元素类型包含一个可以与比较值匹配的字段。
// 参数:
//
//	slice []T: 要搜索的切片。
//	compareValue any: 用于比较的值。
//	fieldName ...string: 可变参数，用于指定要比较的字段名，默认为 "UID"。
//
// 返回值:
//
//	*T: 指向匹配元素的指针，如果没有找到匹配的元素，则返回指向零值的指针。
func GetSliceLastPointer[T any](slice []T, compareValue any, fieldName ...string) *T {
	// 默认比较字段名设为 "UID"。
	var _fieldName string = "UID"

	// 如果提供了字段名参数，使用第一个字段名作为比较字段。
	if len(fieldName) > 0 {
		_fieldName = fieldName[0]
	}
	// zeroValue 用于存储 T 类型的零值。
	var zeroValue T
	// 如果切片为空，返回指向零值的指针。
	if len(slice) == 0 {
		return &zeroValue
	}
	// 从切片的最后一个元素开始向前遍历。
	for i := len(slice) - 1; i >= 0; i-- {
		// 获取当前元素的反射值。
		v := reflect.ValueOf(slice[i])
		// 尝试通过字段名获取字段值。
		field := v.FieldByName(_fieldName)
		// 检查字段是否有效，并且字段值与比较值相等。
		if field.IsValid() && field.Interface() == compareValue {
			// 如果找到匹配的元素，返回指向该元素的指针。
			return &slice[i]
		}
	}
	// 如果没有找到匹配的元素，返回指向零值的指针。
	return &zeroValue
}

// GetSlicePointers 通过反射在切片中找到最后一个与指定比较值匹配的元素指针Slice。
func GetSlicePointers[T any](slice []T, compareValue any, fieldName ...string) []*T {
	// 默认比较字段名设为 "UID"。
	var _fieldName string = "UID"

	// 如果提供了字段名参数，使用第一个字段名作为比较字段。
	if len(fieldName) > 0 {
		_fieldName = fieldName[0]
	}
	// 初始化一个用于存储匹配元素指针的切片。
	var result []*T

	// 如果切片为空，直接返回空切片。
	if len(slice) == 0 {
		return result
	}

	// 遍历切片中的每个元素。
	for i := 0; i < len(slice); i++ {
		// 获取当前元素的反射值。
		v := reflect.ValueOf(slice[i])
		// 尝试通过字段名获取字段值。
		field := v.FieldByName(_fieldName)
		// 检查字段是否有效，并且字段值与比较值相等。
		if field.IsValid() && field.Interface() == compareValue {
			// 如果找到匹配的元素，将其指针添加到结果切片中。
			result = append(result, &slice[i])
		}
	}
	// 返回存储了所有匹配元素指针的切片。
	return result
}

// GetSlice 通过反射在切片中找到最后一个与指定比较值匹配的元素指针Slice。
func GetSlice[T any](slice []T, compareValue any, fieldName ...string) []T {
	// 默认比较字段名设为 "UID"。
	var _fieldName string = "UID"

	// 如果提供了字段名参数，使用第一个字段名作为比较字段。
	if len(fieldName) > 0 {
		_fieldName = fieldName[0]
	}
	// 初始化一个用于存储匹配元素指针的切片。
	var result []T

	// 如果切片为空，直接返回空切片。
	if len(slice) == 0 {
		return result
	}

	// 遍历切片中的每个元素。
	for i := 0; i < len(slice); i++ {
		// 获取当前元素的反射值。
		v := reflect.ValueOf(slice[i])
		// 尝试通过字段名获取字段值。
		field := v.FieldByName(_fieldName)
		// 检查字段是否有效，并且字段值与比较值相等。
		if field.IsValid() && field.Interface() == compareValue {
			// 如果找到匹配的元素，将其指针添加到结果切片中。
			result = append(result, slice[i])
		}
	}
	// 返回存储了所有匹配元素指针的切片。
	return result
}

// GetSlicePointersInSlice 通过反射在切片中找到最后一个与指定比较值匹配的元素指针Slice。
func GetSlicePointersInSlice[T any](slice []T, compareValues []any, fieldName ...string) []*T {
	// 默认比较字段名设为 "UID"。
	var _fieldName string = "UID"

	// 如果提供了字段名参数，使用第一个字段名作为比较字段。
	if len(fieldName) > 0 {
		_fieldName = fieldName[0]
	}
	// 初始化一个用于存储匹配元素指针的切片。
	var result []*T

	// 如果切片为空，直接返回空切片。
	if len(slice) == 0 {
		return result
	}

	// 遍历切片中的每个元素。
	for i := 0; i < len(slice); i++ {
		// 获取当前元素的反射值。
		v := reflect.ValueOf(slice[i])
		// 尝试通过字段名获取字段值。
		field := v.FieldByName(_fieldName)
		// 检查字段是否有效，并且字段值与比较值相等。
		if field.IsValid() && slices.Contains(compareValues, field.Interface()) {
			// 如果找到匹配的元素，将其指针添加到结果切片中。
			result = append(result, &slice[i])
		}
	}
	// 返回存储了所有匹配元素指针的切片。
	return result
}

// SliceToPointers 通过反射在切片中找到所有与指定比较值匹配的元素指针Slice。
func SliceToPointers[T any](slice []T) []*T {
	var result []*T
	for i := 0; i < len(slice); i++ {
		result = append(result, &slice[i])
	}
	return result
}

// SHA1Encrypt 对输入的字符串进行 SHA-1 加密
func SHA1Encrypt(data string) string {
	// 创建一个新的 SHA-1 哈希对象
	hash := sha1.New()
	// 向哈希对象中写入要加密的数据
	hash.Write([]byte(data))
	// 计算哈希值
	hashedData := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hexString := hex.EncodeToString(hashedData)
	return hexString
}

// HashPassword 使用 bcrypt 对密码进行加密
func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

// ComparePasswords 比对明文密码和加密后的密码
func ComparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func GetNowDateTimeNoSymbolStr() string {
	format := time.Now().In(getLocation()).Format(time.DateTime)
	format = strings.ReplaceAll(format, "-", "")
	format = strings.ReplaceAll(format, " ", "")
	format = strings.ReplaceAll(format, ":", "")
	return format
}

func GetNowDateTimeStr() string {
	return time.Now().In(getLocation()).Format(time.DateTime)
}
func GetNowTimeOnlyStr() string {
	return time.Now().In(getLocation()).Format(time.TimeOnly)
}
func GetNowDateOnlyStr() string {
	return time.Now().In(getLocation()).Format(time.DateOnly)
}
func GetNowTimeUnixMilli() int64 {
	return time.Now().In(getLocation()).UnixMilli()
}
func GetNowTimeUnix() int64 {
	return time.Now().In(getLocation()).Unix()
}

func getLocation() *time.Location {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return location
}
