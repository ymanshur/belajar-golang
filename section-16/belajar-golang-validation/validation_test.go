package belajar_golang_validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidateVar(t *testing.T) {
	validate := validator.New()

	var user string = "Yusuf"

	err := validate.Var(user, "required")
	if err != nil {
		t.Log(err.Error())
	}
}

func TestValidateVarWithValue(t *testing.T) {
	validate := validator.New()

	var (
		password        string = "pass"
		confirmPassword string = "passs"
	)

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		t.Log(err.Error())
	}
}

func TestValidateMultiTag(t *testing.T) {
	validate := validator.New()

	var user string = "Yusuf123"

	err := validate.Var(user, "required,alpha")
	if err != nil {
		t.Log(err.Error())
	}
}

func TestValidateTagParam(t *testing.T) {
	validate := validator.New()

	var user string = "Yusuf123Manshur"

	err := validate.Var(user, "required,alpha,min=3,max=10")
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestValidateStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"` // reflection tag
		Password string `validate:"required,min=5"` // reflection tag
	}

	validate := validator.New()

	request := LoginRequest{
		Username: "yusuf@example.com",
		Password: "password",
	}

	err := validate.Struct(request)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestValidateError(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"` // reflection tag
		Password string `validate:"required,min=5"` // reflection tag
	}

	validate := validator.New()

	request := LoginRequest{
		Username: "yusuf",
		Password: "pass",
	}

	err := validate.Struct(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			t.Logf("error %s on tag %s", fieldErr.Field(), fieldErr.Tag())
		}
	}
}

func TestValidateCrossField(t *testing.T) {
	type LoginRequest struct {
		Username        string `validate:"required,email"`                  // reflection tag
		Password        string `validate:"required,min=5"`                  // reflection tag
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"` // reflection tag
	}

	validate := validator.New()

	request := LoginRequest{
		Username:        "yusuf@example.com",
		Password:        "secretpassword",
		ConfirmPassword: "secretpasswordd",
	}

	err := validate.Struct(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			t.Logf("error %s on tag %s", fieldErr.Field(), fieldErr.Tag())
		}
	}
}

func TestValidateNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      int     `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()

	user := User{
		Id:      0,
		Name:    "",
		Address: Address{},
	}

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			t.Logf("error %s on tag %s", fieldErr.Field(), fieldErr.Tag())
		}
	}
}

func TestValidateCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        int       `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()

	user := User{
		Id:   0,
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

func TestValidateBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        int       `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()

	user := User{
		Id:   0,
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{""},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

func TestValidateMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        int               `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	}

	validate := validator.New()

	user := User{
		Id:   1,
		Name: "Yusuf Manshur",
		Addresses: []Address{
			{
				City:    "Sleman",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{"Coding"},
		Schools: map[string]School{
			"SD": {
				Name: "",
			},
			"SMP": {
				Name: "SMPIT As Syifa",
			},
			"": {
				Name: "",
			},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

func TestValidateBasicMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        int               `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=8000"`
	}

	validate := validator.New()

	user := User{
		Id:   1,
		Name: "Yusuf Manshur",
		Addresses: []Address{
			{
				City:    "Sleman",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{"Coding"},
		Schools: map[string]School{
			"SMP": {
				Name: "SMPIT As Syifa",
			},
		},
		Wallet: map[string]int{
			"BCA":       1000000,
			"ShopeePay": 1000,
		},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

func TestValidateAliasTag(t *testing.T) {
	type User struct {
		Id      int      `validate:"required"`
		Name    string   `validate:"varchar"`
		Hobbies []string `validate:"required,dive,varchar"`
	}

	validate := validator.New()
	validate.RegisterAlias("varchar", "required,min=5")

	user := User{
		Id:      1,
		Name:    "Yusuf Manshur",
		Hobbies: []string{"abc"},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

func ValidUsername(level validator.FieldLevel) bool {
	value, ok := level.Field().Interface().(string)
	if ok {
		if value == strings.ToLower(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestValidateCustomTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", ValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"` // reflection tag
		Password string `validate:"required,min=5"`    // reflection tag
	}

	request := LoginRequest{
		Username: "YUSUF",
		Password: "password",
	}

	err := validate.Struct(request)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func ValidPin(level validator.FieldLevel) bool {
	length, err := strconv.Atoi(level.Param())
	if err != nil {
		panic(err)
	}

	value := level.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestValidateCustomParam(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", ValidPin)

	type LoginRequest struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := LoginRequest{
		Phone: "082344084475",
		Pin:   "123123",
	}

	err := validate.Struct(request)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

func TestValidateOrRule(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", ValidPin)

	type LoginRequest struct {
		Username string `validate:"required,email|number"`
		Pin      string `validate:"required,pin=6"`
	}

	request := LoginRequest{
		//Username: "082344084475",
		Username: "yusuf@example.com",
		Pin:      "123123",
	}

	err := validate.Struct(request)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

func EqualsIgnoreCase(level validator.FieldLevel) bool {
	value, _, _, b2 := level.GetStructFieldOK2()
	if !b2 {
		panic("field not ok")
	}

	firstValue := strings.ToLower(level.Field().String())
	secondValue := strings.ToLower(value.String())

	return firstValue == secondValue
}

func TestValidateCustomCrossField(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("qeigfield", EqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,qeigfield=Email|qeigfield=Phone"`
		Email    string `validate:"required"`
		Phone    string `validate:"required"`
	}

	user := User{
		Username: "yusuf@example.com",
		//Username: "082344084475",
		Email: "yusuf@example.com",
		Phone: "082344084475",
	}

	err := validate.Struct(user)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required"`
	Password string `validate:"required"`
}

func ValidRegisterRequest(level validator.StructLevel) {
	request := level.Current().Interface().(RegisterRequest)

	if request.Username == request.Phone || request.Username == request.Email {

	} else {
		level.ReportError(request.Username, "Username", "RegisterRequest", "username", "")
	}
}

func TestValidateStructLevel(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(ValidRegisterRequest, new(RegisterRequest))

	request := RegisterRequest{
		Username: "yusuf",
		//Username: "082344084475",
		Email:    "yusuf@example.com",
		Phone:    "082344084475",
		Password: "1234",
	}

	err := validate.Struct(request)
	if err != nil {
		t.Logf("\n%s", err.Error())
	}
}
