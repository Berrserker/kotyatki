package utils

import (
	//st "github.com/fatih/structs"

	"encoding/json"

	"net/http"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	//"github.com/buger/jsonparser"

	"fmt"

	m "../models"

	"io/ioutil"

	//"github.com/recoilme/pudge"
	//"sync"
	//"log"

	"os"

	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var PPrint = fmt.Println

var Client *mongo.Client
var Users *mongo.Collection
var Animals *mongo.Collection
var Vocs *mongo.Collection

func Init(base string) {
	var c_error error
	//Client, c_error = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	Client, c_error = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.drvrt.mongodb.net/test"))
	c_error = Client.Connect(context.TODO())
	if c_error != nil {
		PPrint(c_error)
	}
	Users = Client.Database(base).Collection("users")
	Animals = Client.Database(base).Collection("animals")
	Vocs = Client.Database(base).Collection("vocs")
}

func Message(status int32, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Message(0, "Server Online")
	Respond(w, response)
}

/*
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &m.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := Create(&account) //Создать аккаунт
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	Respond(w, resp)
}
*/

var Auth = func(w http.ResponseWriter, r *http.Request) {

	account := &m.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		Respond(w, Message(1, "Invalid request"))
		return
	}

	resp := Login(account.Email, account.Password)
	//PPrint(resp)
	Respond(w, resp)
}

/*
func Validate(account *m.Account) (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email должен быть уникальным
	temp := &m.Account{}

	//проверка на наличие ошибок и дубликатов электронных писем
	//err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func Create(account *m.Account) map[string]interface{} {

	if resp, ok := Validate(&account); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Создать новый токен JWT для новой зарегистрированной учётной записи
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("secret")))
	account.Token = tokenString

	account.Password = "" //удалить пароль

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}
*/

func Login(email, password string) map[string]interface{} {

	account := &m.Account{}
	//PPrint(email)
	filter := bson.M{
		"login": email,
	}
	count, _ := Users.CountDocuments(context.TODO(), bson.M{})
	//PPrint(count)
	cur, _ := Users.Find(
		context.TODO(),
		filter,
	)
	//PPrint(c_err)
	resp := Message(1, "Invalid login credentials. Please try again")
	if (count > 0) && (count < 2) {
		for cur.Next(context.TODO()) {
			var elem m.User
			err := cur.Decode(&elem)
			if err != nil {
				PPrint(err)
			}
			err = bcrypt.CompareHashAndPassword([]byte(elem.Password), []byte(password))
			if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
				//PPrint("Пароль не совпадает!!")
				break
			}

			//Создать токен JWT
			// Устанавливаем набор параметров для токена
			atClaims := jwt.MapClaims{}
			atClaims["admin"] = true
			atClaims["login"] = elem.Login
			atClaims["exp"] = time.Now().Add(time.Minute * 24).Unix()

			// Подписываем токен нашим секретным ключем
			at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			tokenString, err := at.SignedString([]byte(os.Getenv("secret")))

			account.Token = tokenString // Сохраним токен в ответе

			account.Password = "" //удалить пароль

			resp = Message(0, "Logged In")
			resp["account"] = account
			break
		}
	}
	return resp
}

/*
func GetUser(u uint) *m.Account {

	acc := &m.Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //Пользователь не найден!
		return nil
	}

	acc.Password = ""
	return acc
}
*/
func Kill() {
	defer Client.Disconnect(context.TODO())
}

func ListDirByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}

func SetPwd(pass string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func Showvoc(w http.ResponseWriter, r *http.Request) {
	rvoc := &m.Rvoc{}
	err := json.NewDecoder(r.Body).Decode(rvoc) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		Respond(w, Message(1, "Invalid query"))
		return
	}
	voc(w, rvoc)
}

func ShowListvoc(w http.ResponseWriter, r *http.Request) {
	rvoc := &m.Rvoc{}
	rvoc.Request = []string{"list"}
	voc(w, rvoc)
}

func voc(w http.ResponseWriter, rvoc *m.Rvoc) {

	flag := true
	var voc []m.Voc
	resp := Message(0, "Query is ok")
	for _, value := range rvoc.Request {

		switch value {
		case "all":
			cur, cur_err := Vocs.Find(context.TODO(), bson.M{})
			if cur_err != nil {
				PPrint(cur_err)
				resp = Message(2, "Server error")
				Respond(w, resp)
				flag = false
				break
			}
			for cur.Next(context.TODO()) {
				var elem m.Voc
				_ = cur.Decode(&elem)
				voc = append(voc, elem)
			}
			flag = false
		case "list":
			cur, cur_err := Vocs.Find(context.TODO(), bson.M{"name": "list"})
			count, _ := Vocs.CountDocuments(context.TODO(), bson.M{"name": "list"})
			if cur_err != nil {
				PPrint(cur_err)
				resp = Message(2, "Server error")
				Respond(w, resp)
				flag = false
				break
			}

			if (count > 0) && (count < 2) {
				for cur.Next(context.TODO()) {
					var elem m.Voc
					_ = cur.Decode(&elem)
					voc = append(voc, elem)
					break
				}
			}
			flag = false
		default:
			cur, cur_err := Vocs.Find(context.TODO(), bson.M{"name": value})
			count, _ := Vocs.CountDocuments(context.TODO(), bson.M{"name": value})
			if cur_err != nil {
				PPrint(cur_err)
				resp = Message(2, "Server error")
				Respond(w, resp)
				flag = false
				break
			}
			if (count > 0) && (count < 2) {
				for cur.Next(context.TODO()) {
					var elem m.Voc
					_ = cur.Decode(&elem)
					voc = append(voc, elem)
				}
			}
		}
		if !flag {
			break
		}
	}
	resp["vocs"] = voc
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	PPrint(resp)
	Respond(w, resp)
}

func Setvoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Message(0, "Server Online")
	Respond(w, response)
}
