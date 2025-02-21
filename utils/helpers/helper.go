package helpers

import (
	"uni_app/database"

	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

func GetIDFromContxt(c echo.Context) (ID database.PID, err error) {
	ID = database.Parse(c.Param("id"))
	if !ID.IsValid() {
		return database.NilPID, database.ErrInvalidPID
	}

	return ID, nil
}

const (
	// Secret ...
	Secret = "123ChainStore@AdminSecretKey"
	// JWT Claims -----------------------------------------------------------------------
	// ClaimsClientKey ...
	ClaimsClientKey = "key"
	// ClaimsUsername ...
	ClaimsUsername = "uname"
	// ClaimsUserID ...
	ClaimsUserID = "uid"
	// ClaimsRole ...
	ClaimsRole = "rol"
	// ClaimsExpireTime ...
	ClaimsExpireTime = "exp"
	// ClaimsStoreID
	ClaimsStoreID = "sid"
	// Context Headers ------------------------------------------------------------------
	// HeadersUser ...
	HeadersUser = "x-user"
	// HeadersUserID ...
	HeadersUserID = "x-user-id"
	// HeadersStore ...
	HeadersStore = "x-store"
	// HeadersBranch ...
	HeadersBranch = "x-branch"
	// HeadersStock ...
	HeadersStock = "x-stock"
	// HeadersStoreID ...
	HeadersStoreID = "x-store-id"
	// HeadersBranchID ...
	HeadersBranchID = "x-branch-id"
	// HeadersStockID ...
	HeadersStockID = "x-stock-id"
	// HeadersRemoteStockID ...
	HeadersRemoteStockID = "x-remote-stock-id"
	// HeadersStockIDs ...
	HeadersStockIDs = "x-stock-ids"
	// HeadersStoreNo ...
	HeadersBoothID = "x-booth-id"
	// HeadersRemoteStockIDs ...
	HeadersRemoteStockIDs = "x-remote-stock-ids"
	// HeadersBranches ...
	HeadersBranches = "x-branches"
	// HeadersAdapterID ...
	HeadersAdapterID = "x-adapter-id"
	// HeadersAdapterAlias ...
	HeadersAdapterAlias = "x-adapter-alias"
	// Adapter
	HeadersAdapter       = "x-adapter"
	HeadersAdapterStore  = "x-adapter-store"
	HeadersAdapterBranch = "x-adapter-branch"
	HeadersAdapterStock  = "x-adapter-stock"
	// HeadersClient ...
	HeadersClient = "x-client"
	// HeadersClient ...
	HeadersClientID = "x-client-id"
	// HeadersCustomer ...
	HeadersCustomer = "x-customer"
	// HeadersCustomer ...
	HeadersCustomerID = "x-customer-id"
	// HeadersAuthenticated ...
	HeadersAuthenticated = "x-authenticated"
	// HeadersTokenKey ...
	HeadersTokenKey = "x-token-key"
	// HeadersPosType ...
	HeadersPosType = "x-pos-type"
	// HaedersProxyDestination ...
	HaedersProxyDestination = "x-proxy-destination"
	// HeadersTunnelKey ...
	// HeadersTunnelKey = "x-tunnel-key"
	HeadersTunnelKey = "x-client-id"
	// HeadersPlatform pwa, wordpress, next, pl, cr, ...
	HeadersPlatform = "x-platform"
)

var (
	alphanumericChars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// letterChars       = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	verificationChars = []byte("123456789")
	digitChars        = []byte("0123456789")
)

// GetAbsPath ...
func GetAbsPath(filename string) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return exPath + "/" + filename
}

func Replace(text string, old []string, new []string) (string, error) {
	if len(old) != len(new) {
		return "", errors.New("the length of old and new must be same")
	}

	var s string

	for i, o := range old {
		s = strings.Replace(text, o, new[i], -1)
	}

	return s, nil
}

func ReplaceMap(text string, m map[string]interface{}) string {
	if len(m) == 0 {
		return text
	}

	for k, v := range m {
		text = strings.Replace(text, k, fmt.Sprintf("%v", v), -1)
	}

	return text
}

// GetUintFromString ...
func GetUintFromString(s string) (uint, error) {
	idP, err := strconv.Atoi(s)
	uID := uint(idP)

	return uID, err
}

// GetIntFromString ...
func GetIntFromString(s string) (int, error) {
	idP, err := strconv.Atoi(s)

	return idP, err
}

// GetStringsDefault ...
func GetStringsDefault(s []string, def []string) []string {
	if len(s) > 0 {
		return s
	}

	return def
}

// GetStringDefault ...
func GetStringDefault(s string, def string) string {
	if strings.TrimSpace(s) == "" {
		return def
	}

	return s
}

func GetEnvDefault(envKey string, def string) string {
	var (
		exists bool
		val    string
	)

	if val, exists = os.LookupEnv(envKey); !exists {
		return def
	}

	return val
}

// GetIntDefault ...
func GetIntDefault(s string, def int) int {
	v, err := strconv.ParseInt(s, 0, 0)

	if err != nil {
		return def
	}

	return int(v)
}

// GetBoolsDefault ...
func GetBoolsDefault(params []string, def []bool) (values []bool) {
	if len(params) > 0 {
		for _, param := range params {
			if value, err := strconv.ParseBool(param); err == nil {
				values = append(values, value)
			}
		}

		return values
	}

	return def
}

// GetBoolDefault ...
func GetBoolDefault(s string, def bool) bool {
	v, err := strconv.ParseBool(s)

	if err != nil {
		return def
	}

	return v
}

// GetBool ...
func GetBool(s string) (bool, error) {
	if s == "" {
		return false, nil
	}
	return strconv.ParseBool(s)
	// if s == "true" {
	// 	return true, nil
	// } else if s == "false" || s == "" {
	// 	return false, nil
	// } else {
	// 	return false, errors.New("invalid value")
	// }
}

// GenerateVerificationCode ...
func GenerateVerificationCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = verificationChars[int(b[i])%len(verificationChars)]
	}
	return string(b)
}

// GenerateCode ...
func GenerateCode(max int) (string, error) {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = digitChars[int(b[i])%len(digitChars)]
	}

	return string(b), nil
}

// RandStringCode ...
func RandStringCode(max int) (string, error) {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}
	for i := range b {
		b[i] = alphanumericChars[int(b[i])%len(alphanumericChars)]
	}
	return string(b), nil
}

// GenerateUUID generate uuid with custom separator
func GenerateUUID(sep string) string {
	return strings.ReplaceAll(uuid.New().String(), "-", sep)
}

// GenerateToken ...
/*func GenerateToken(username string, expTime time.Time, role int) (string, error) {

    token := jwt.New(jwt.SigningMethodHS256)

    // Setting Claims
    claims := token.Claims.(jwt.MapClaims)
    claims[ClaimsRole] = role
    claims[ClaimsUsername] = username
    claims[ClaimsExpireTime] = expTime

    // Generate encoded token and save it into db.
    tokenKey, err := token.SignedString([]byte(Secret))

    return tokenKey, err
}*/

// // GenerateToken ...
// // func GenerateToken(auth map[string]string, role int, username string) (tokenKey string, expTime time.Time, err error) {
// func GenerateToken(auth map[string]string, userID, storeID database.PID) (tokenKey string, expTime time.Time, err error) {
// 	var (
// 		alg, secret string
// 		expire      int64
// 		token       *jwt.Token
// 	)

// 	alg = auth["alg"]
// 	secret = auth["secret"]
// 	expire, _ = strconv.ParseInt(auth["expire"], 10, 64)

// 	switch alg {
// 	case "HS256":
// 		token = jwt.New(jwt.SigningMethodHS256)
// 	default:
// 		token = jwt.New(jwt.SigningMethodHS256)
// 	}

// 	expTime = time.Now().Local().Add(time.Hour * time.Duration(expire))

// 	// Setting Claims
// 	claims := token.Claims.(jwt.MapClaims)
// 	// claims[ClaimsRole] = role
// 	// claims[ClaimsUsername] = username
// 	claims[ClaimsStoreID] = storeID
// 	claims[ClaimsUserID] = userID
// 	claims[ClaimsExpireTime] = expTime

// 	// Generate encoded token and save it into db.
// 	tokenKey, err = token.SignedString([]byte(secret))

// 	return
// }

// GetStringFromUInt ...
func GetStringFromUInt(i uint) string {
	return strconv.FormatInt(int64(i), 10)
}

// IsRequestValid ...
func IsRequestValid(i interface{}) (bool, error) {

	validate := validator.New()

	err := validate.Struct(i)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateDirIfNotExist ...
func CreateDirIfNotExist(assetsPath, dir string) (os.FileInfo, string, error) {
	// rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// d := rootDir + dir

	d := path.Join(assetsPath, dir)
	// fmt.Println(d)
	fInfo, err := os.Stat(d)

	if err == nil {
		return fInfo, d, err
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(d, 0755)
		if err != nil {
			return fInfo, "", err
		}
	}
	return fInfo, d, err
}

// Cast ...
func Cast(src interface{}, dst interface{}) {
	dstValue := reflect.ValueOf(dst)
	dstElem := dstValue.Elem()
	countDstElems := dstElem.NumField()
	dstType := dstElem.Type()

	for i := 0; i < countDstElems; i++ {
		dstField := dstElem.Field(i)
		dstFieldName := dstType.Field(i).Name
		// dstFieldType := dstField.Type()

		srcValue := reflect.ValueOf(src)
		srcElem := srcValue.Elem()
		countSrcElems := srcElem.NumField()
		srcType := srcElem.Type()

		if i == 0 {
			valID := reflect.Indirect(srcValue).Field(i)
			dstField.Set(valID)
			continue
		}

		for j := 1; j < countSrcElems; j++ {

			// srcField := srcElem.Field(j)
			srcFieldName := srcType.Field(j).Name
			// srcFieldType := srcField.Type()

			if dstFieldName == srcFieldName { // && dstFieldType == srcFieldType {
				// FOUND
				val := reflect.Indirect(srcValue).Field(j)
				// f := reflect.Indirect(r).FieldByName("Mobile")
				dstField.Set(val)
				break
			}
		}
	}

	// s := reflect.ValueOf(src).Elem()

	// fmt.Println(s.NumField())
	// typeOfR := s.Type()

	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("%d: %s %s = %v\n", i, typeOfR.Field(i).Name, f.Type(), f.Interface())
	// }
}

// CastToInt ...
func CastToInt(iface interface{}) (int, error) {
	// iaface = indirect(iface)

	if iface == nil {
		return 0, nil
	}

	var err error
	switch s := iface.(type) {
	case int:
		return s, err
	case int32:
		return int(s), err
	case int16:
		return int(s), err
	case int8:
		return int(s), err
	case uint:
		return int(s), err
	case uint64:
		return int(s), err
	case uint32:
		return int(s), err
	case uint16:
		return int(s), err
	case uint8:
		return int(s), err
	case float64:
		return int(s), err
	case float32:
		return int(s), err
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", iface, iface)
	}
}

// TimeTrack measuring executing function
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	color.Set(color.FgCyan, color.Bold)
	fmt.Printf("[%s elapsed]\t%s\n", elapsed, name)
	color.Unset()
	// return fmt.Sprintf("[%s]\t%s", elapsed, name)
}

func GetBaseName(src string) string {
	ext := path.Ext(src)
	return strings.TrimSuffix(path.Base(src), ext)
}

func GetName(src string) string {
	return path.Base(src)
}

// // SearchSlice ...
// func SearchSlice(key string, value interface{}, slice []interface{}) (index int, found bool) {
// 	// found = false

// 	for index, item := range slice {
// 		if item[key] == value {
// 			found = true

// 		}

// 	}
// }

// AnyItemExists ...
func AnyItemExists(searchArrayType interface{}, checkArrayType interface{}) bool {
	if searchArrayType == nil || checkArrayType == nil {
		return false
	}

	var (
		searchArr = reflect.ValueOf(searchArrayType)
	)

	if searchArr.Kind() != reflect.Array && searchArr.Kind() != reflect.Slice {
		// panic("Invalid data-type")
		return false
	}

	for i := 0; i < searchArr.Len(); i++ {
		if ItemExists(checkArrayType, searchArr.Index(i).Interface()) {
			return true
		}
	}

	return false
}

// ReverseSlice reverece order slice
func ReverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}

// PurgeArray purge array
func PurgeArray(arrayType interface{}) interface{} {
	if arrayType == nil {
		return nil //, errors.New("arrayType is nil")
	}

	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		return nil //, errors.New("arrayType must be array or slice")
	}

	check := make(map[interface{}]int)

	for i := 0; i < arr.Len(); i++ {
		check[arr.Index(i).Interface()] = 1
	}

	typ := reflect.TypeOf(arrayType).Elem()
	temp := make([]interface{}, len(check))
	for key := range check {
		temp = append(temp, key)
	}

	t := reflect.MakeSlice(reflect.SliceOf(typ), 0, len(check))
	reflect.Copy(t, reflect.ValueOf(temp))

	return t.Interface() //, nil
}

// ItemExists ...
func ItemExists(arrayType interface{}, item interface{}) bool {
	if arrayType == nil {
		return false
	}

	arr := reflect.ValueOf(arrayType)

	// fmt.Println("array:", arr.Kind().String())

	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		return false
		// panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

// RemoveIndex remove index from slice or array
func RemoveIndex(arrayType interface{}, i int) (interface{}, error) {
	if arrayType == nil {
		return nil, errors.New("array is nil")
	}

	arr := reflect.ValueOf(arrayType)
	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		return nil, errors.New("arrayType must be array or slice")
	}

	if i >= arr.Len() {
		return nil, errors.New("index is upper than arrayType length")
	}

	// a[i] = a[len(a)-1]
	// a[len(a)-1] = nil
	// a = a[:len(a)-1]

	// elem := arr.Index(i)
	// lastElem := arr.Index(arr.Len() - 1)
	// elem.Set(lastElem)
	// lastElem.Set(reflect.ValueOf(nil))

	newArr := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(arrayType).Elem()), 0, arr.Len()-1)

	for idx := 0; idx < arr.Len(); idx++ {
		if idx == i {
			continue
		}

		newArr = reflect.Append(newArr, arr.Index(idx).Elem())
	}

	return newArr, nil
}

// TrimString ...
func TrimString(s string) (value string) {
	value = strings.TrimSpace(s)
	value = strings.ToValidUTF8(value, "")
	return value
}

// MergeMaps overwriting duplicate keys, you should handle that if there is a need
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func JSONTo(src interface{}, dst interface{}) (err error) {
	var js []byte

	js, err = json.Marshal(src)
	if err != nil {
		return
	}

	err = json.Unmarshal(js, dst)
	if err != nil {
		return
	}

	return
}

// Deprecated: use JSONTo instead
func MyAssertion(src interface{}, dst interface{}) (err error) {
	return JSONTo(src, dst)
}

func HashPassword(purePass string) (hashedPass string) {
	bytePass := []byte(purePass)
	bytePass, _ = bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	hashedPass = string(bytePass)
	return
}

func ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func ContextUserID(ctx echo.Context) (userID database.PID) {
	if userID, err := database.ParsePID(GetFromContext(ctx, HeadersUserID)); err != nil {
		return database.NilPID
	} else {
		return userID
	}
}

func ContextCustomerID(ctx echo.Context) (customerID database.PID) {
	if customerID, err := database.ParsePID(GetFromContext(ctx, HeadersCustomerID)); err != nil {
		return database.NilPID
	} else {
		return customerID
	}
}

func ContextStoreID(ctx echo.Context) (storeID database.PID) {
	if storeID, err := database.ParsePID(GetFromContext(ctx, HeadersStoreID)); err != nil {
		return database.NilPID
	} else {
		return storeID
	}
}

func ContextBranchID(ctx echo.Context) (branchID database.PID) {
	if branchID, err := database.ParsePID(GetFromContext(ctx, HeadersBranchID)); err != nil {
		return database.NilPID
	} else {
		return branchID
	}
}

func ContextStockID(ctx echo.Context) (stockID database.PID) {
	if stockID, err := database.ParsePID(GetFromContext(ctx, HeadersStockID)); err != nil {
		return database.NilPID
	} else {
		return stockID
	}
}

func ContextRemoteStockID(ctx echo.Context) (remoteStockID string) {
	if remoteStockID := cast.ToString(GetFromContext(ctx, HeadersRemoteStockID)); remoteStockID == "" {
		return ""
	} else {
		return remoteStockID
	}
}

func ContextPlatform(ctx echo.Context) (platform string) {
	if platform := cast.ToString(GetFromContext(ctx, HeadersPlatform)); platform == "" {
		return ""
	} else {
		return platform
	}
}

func ContextAdapterAlias(ctx echo.Context) (alias string) {
	if alias := cast.ToString(GetFromContext(ctx, HeadersAdapterAlias)); alias == "" {
		return ""
	} else {
		return alias
	}
}

func GetFromContext(ctx echo.Context, key string) interface{} {
	if v := ctx.Get(key); v == nil && ctx.Request() != nil {
		return ctx.Request().Header.Get(key)
	} else {
		return v
	}
}

func SetToContext(ctx echo.Context, key string, value interface{}) {
	ctx.Set(key, value)
}

func SetHeaderToRequest(ctx echo.Context, key string, value interface{}) {
	if ctx.Request() == nil {
		return
	}

	switch v := value.(type) {
	case string:
		ctx.Request().Header.Set(key, v)
	case database.PID:
		ctx.Request().Header.Set(key, v.String())
	default:
		ctx.Request().Header.Set(key, cast.ToString(v))
	}
}

func ContextToHeader(ctx echo.Context) (headers map[string]string) {
	return map[string]string{
		HeadersStoreID:       ContextStoreID(ctx).String(),
		HeadersBranchID:      ContextBranchID(ctx).String(),
		HeadersStockID:       ContextStockID(ctx).String(),
		HeadersRemoteStockID: ContextRemoteStockID(ctx),
		HeadersAdapterAlias:  ContextAdapterAlias(ctx),
		HeadersPlatform:      ContextPlatform(ctx),
	}
}

func GetIntersection(slice1, slice2 []interface{}) []interface{} {
	elementMap := make(map[interface{}]bool)
	intersection := []interface{}{}

	for _, item := range slice1 {
		if _, ok := elementMap[item]; !ok {
			elementMap[item] = true
		}
	}

	for _, item := range slice2 {
		if _, ok := elementMap[item]; ok {
			intersection = append(intersection, item)
			delete(elementMap, item)
		}
	}

	return intersection
}

func SetTransaction(ctx echo.Context, tx *gorm.DB) {
	ctx.Set("db", tx)
}

func GetTransaction(ctx echo.Context) (tx *gorm.DB, err error) {
	txInterface := ctx.Get("db")
	if txInterface == nil {
		return nil, errors.New("transaction not found in context")
	}
	tx, ok := txInterface.(*gorm.DB)
	if !ok {
		return nil, errors.New("transaction in context is not of type *gorm.DB")
	}

	return tx, nil
}
