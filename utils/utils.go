package utils

import (
	"encoding/csv"

	"strconv"

	//st "github.com/fatih/structs"

	"strings"

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

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"

	"reflect"
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

func Profile(w http.ResponseWriter, r *http.Request) {
	login, token_error := ExtractTokenMetadata(r)
	if token_error != nil {
		Respond(w, Message(1, "Invalid token"))
	}

	resp := Message(1, "Invalid login credentials. Please try again")
	account, _, invalid := GetUserByLogin(login)
	if invalid != nil {
		Respond(w, resp)
	}
	resp = Message(0, "Profile found")
	resp["account"] = account
	Respond(w, resp)
}

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
	count, _ := Users.CountDocuments(context.TODO(), filter)
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
			atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			// Подписываем токен нашим секретным ключем
			at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			tokenString, err := at.SignedString([]byte(os.Getenv("secret")))
			account.Email = elem.Email
			account.Address = elem.Address
			account.ApiToken = elem.Token
			account.Phone = elem.Phone
			account.Role = elem.Role
			account.Manager = elem.Manager
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
			cur, cur_err := Vocs.Find(context.TODO(), bson.M{})
			count, _ := Vocs.CountDocuments(context.TODO(), bson.M{})
			if cur_err != nil {
				PPrint(cur_err)
				resp = Message(2, "Server error")
				Respond(w, resp)
				flag = false
				break
			}
			if count > 0 {
				var result m.Voc
				result.Name = "list"
				result.Login = "base"
				for cur.Next(context.TODO()) {
					var elem m.Voc
					_ = cur.Decode(&elem)
					result.Voc = append(result.Voc, elem.Name)
				}
				voc = append(voc, result)
			}
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
	//PPrint(resp)
	Respond(w, resp)
}

func Setvoc(w http.ResponseWriter, r *http.Request) {

	wrvoc := &m.Wrvoc{}
	err := json.NewDecoder(r.Body).Decode(wrvoc) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		Respond(w, Message(1, "Invalid input"))
		return
	}
	//PPrint(wrvoc)
	for _, value := range wrvoc.Vocs {
		//PPrint(value)
		if value.Name == "Вид" || value.Name == "Пол" {
			continue
		}
		filter := bson.M{
			"name": bson.M{
				"$eq": value.Name, // check if bool field has value of 'false'
			},
		}
		update := bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "voc", Value: value.Voc},
			primitive.E{Key: "login", Value: value.Login},
			primitive.E{Key: "name", Value: value.Name},
		}}}
		opts := options.Update().SetUpsert(true)
		result, U_err := Vocs.UpdateOne(
			context.Background(),
			filter,
			update,
			opts,
		)
		if U_err != nil {
			panic(U_err)
		} else {
			PPrint("UpdateOne() result:", result)
			PPrint("UpdateOne() result TYPE:", reflect.TypeOf(result))
			PPrint("UpdateOne() result MatchedCount:", result.MatchedCount)
			PPrint("UpdateOne() result ModifiedCount:", result.ModifiedCount)
			PPrint("UpdateOne() result UpsertedCount:", result.UpsertedCount)
			PPrint("UpdateOne() result UpsertedID:", result.UpsertedID)
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Message(0, "Updated")
	Respond(w, response)
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (string, error) {
	tokenString := ExtractToken(r)
	type MyCustomClaims struct {
		Login string `json:"login"`
		jwt.StandardClaims
	}

	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret")), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.Login, nil
	} else {
		return "error", nil
	}
}

func GetUserByLogin(login string) (*m.Account, string, error) {

	account := &m.Account{}
	filter := bson.M{
		"login": login,
	}
	count, _ := Users.CountDocuments(context.TODO(), filter)
	//PPrint(count)
	cur, _ := Users.Find(
		context.TODO(),
		filter,
	)
	//PPrint(c_err)
	if (count > 0) && (count < 2) {
		var pwd string
		for cur.Next(context.TODO()) {
			var elem m.User
			err := cur.Decode(&elem)
			if err != nil {
				PPrint(err)
			}
			account.Email = elem.Email
			account.Address = elem.Address
			account.ApiToken = elem.Token
			account.Phone = elem.Phone
			account.Role = elem.Role
			account.Manager = elem.Manager
			pwd = elem.Password
			account.Password = "" //удалить пароль
			break
		}
		return account, pwd, nil
	} else {
		err := fmt.Errorf("Invalid Login")
		return nil, "", err
	}
}

func Csv(path string, del string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	run := rune(del[0])
	reader.Comma = run

	for {
		record, e := reader.Read()
		if e != nil {
			break
		}
		//PPrint(record)
		///*
		data := m.CSV{
			N:             record[0],
			Card_N:        record[1],
			Type:          record[2],
			Year:          record[3],
			Mass:          record[4],
			Name:          record[5],
			Sex:           record[6],
			Sub_type:      record[7],
			Color:         record[8],
			Fur:           record[9],
			Ear:           record[10],
			Tale:          record[11],
			Size:          record[12],
			Uni:           record[13],
			Volier_n:      record[14],
			Id_m:          record[15],
			Sterilized:    record[16],
			Doctor:        record[17],
			Social:        record[18],
			Act_in:        record[19],
			Date_in:       record[20],
			Adm_place:     record[21],
			Act_get:       record[22],
			Get_Place:     record[23],
			UL_name:       record[24],
			UL_adr:        record[25],
			UL_phone:      record[26],
			Name_own:      record[27],
			Addres_own:    record[28],
			Phone_own:     record[29],
			FL_name:       record[30],
			Pasport:       record[31],
			P_place:       record[32],
			P_date:        record[33],
			FL_adress:     record[34],
			FL_Phone:      record[35],
			Date_in_event: record[36],
			Act_in_event:  record[37],
			Date_out:      record[38],
			Out:           record[39],
			Act_out:       record[40],
			Priute_adress: record[41],
			Expl_org:      record[42],
			Manager:       record[43],
			Resp_own:      record[44],
			N_pp_apar:     record[45],
			Date_apar:     record[46],
			Drug_apar:     record[47],
			Dose_apar:     record[48],
			N_pp_vac:      record[49],
			Date_vac:      record[50],
			Drug_vac:      record[51],
			Serie_vac:     record[52],
			Date_status:   record[53],
			Status:        record[54],
		}
		animal := &m.Animal{}
		animal.Medicine.Status = data.Status
		//PPrint(TimeParser(data.Date_status))
		animal.Medicine.Status_date = TimeParser(data.Date_status)
		animal.Medicine.Doctor = data.Doctor
		switch data.Sterilized {
		case "стерилизована ранее":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "ранее стерилизована":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "ранее стерилизован":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "кастрирован ранее":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "поступила стерилизованная":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "ранее кастрирован":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "ранее стерилизо-вана":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "ранее стер.и чипир.":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "стерилизация по возрасту":
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "Зона карантина":
			animal.Medicine.Sterilized.Status = false
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "без стерилизации по возрасту":
			animal.Medicine.Sterilized.Status = false
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "без стер.ввиду возраста":
			animal.Medicine.Sterilized.Status = false
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "без стерилизации ввиду возраста":
			animal.Medicine.Sterilized.Status = false
			animal.Medicine.Sterilized.Comment = data.Sterilized
		case "не рекомендовано по возрасту":
			animal.Medicine.Sterilized.Status = false
			animal.Medicine.Sterilized.Comment = data.Sterilized
		default:
			animal.Medicine.Sterilized.Status = true
			animal.Medicine.Sterilized.Date = TimeParser(data.Sterilized)
		}
		switch data.Social {
		case "да":
			animal.Medicine.Socialize = true
		case "нет":
			animal.Medicine.Socialize = false
		}

		for i, value := range strings.Split(data.N_pp_vac, "|") {
			PPrint(data.Serie_vac)
			vaccine := m.Vaccine{}
			dat := strings.Split(data.Date_vac, "|")
			ser := strings.Split(data.Serie_vac, "|")
			drug := strings.Split(data.Drug_vac, "|")
			if len(ser) > i+1 {
				vaccine.Serie = ser[i-1]
			} else {
				vaccine.Serie = ser[0]
			}
			if len(dat) > i+1 {
				vaccine.Date = TimeParser(dat[i-1])
			} else {
				vaccine.Date = TimeParser(dat[0])
			}
			if len(drug) > i+1 {
				vaccine.Drug = drug[i-1]
			} else {
				vaccine.Drug = drug[0]
			}
			vaccine.N_pp = value
			animal.Medicine.Vaccine = append(animal.Medicine.Vaccine, vaccine)
		}

		for i, value := range strings.Split(data.N_pp_apar, "|") {
			aparazite := m.Aparazite{}
			dat := strings.Split(data.Date_apar, "|")
			dos := strings.Split(data.Dose_apar, "|")
			drug := strings.Split(data.Drug_apar, "|")
			if len(dos) > i+1 {
				aparazite.Dose = dos[i+1]
			} else {
				aparazite.Dose = dos[0]
			}
			if len(dat) > i+1 {
				aparazite.Date = TimeParser(dat[i-1])
			} else {
				aparazite.Date = TimeParser(dat[0])
			}
			if len(drug) > i+1 {
				aparazite.Drug = drug[i-1]
			} else {
				aparazite.Drug = drug[0]
			}
			aparazite.N_pp = value
			animal.Medicine.Aparazite = append(animal.Medicine.Aparazite, aparazite)
		}

		animal.Name = data.Name
		switch data.Type {
		case "кошка":
			animal.Type = "Кошка"
		case "Кошка":
			animal.Type = "Кошка"
		case "собака":
			animal.Type = "Собака"
		case "Собака":
			animal.Type = "Собака"
		}
		yr, _ := strconv.Atoi(data.Year)
		mas, _ := strconv.ParseFloat(data.Mass, 64)
		animal.Login = "@mail"
		animal.Socialize = animal.Medicine.Socialize
		animal.Ready = true
		animal.Params.Type = animal.Type
		animal.Params.Year = int64(yr)
		animal.Params.Mass = mas
		animal.Params.Sex = data.Year
		animal.Params.Unique = data.Uni
		animal.Params.Size = data.Size
		animal.Params.Ears = data.Ear
		animal.Params.Fur = data.Fur
		animal.Params.Color = data.Color
		animal.Params.Tale = data.Color
		animal.Params.Name = animal.Name
		switch data.Sub_type {
		case "метис":
			animal.Params.Sub_type = data.Sub_type
		case "Метис":
			animal.Params.Sub_type = "метис"
		default:
			animal.Params.Sub_type = data.Sub_type
		}
		if data.Date_in_event != "" {
			event := m.Event{}
			event.Date = TimeParser(data.Date_in_event)
			event.Comment = ""
			event.Act_n = data.Act_in_event
			event.Status = "IN"
			animal.Events = append(animal.Events, event)
		}

		if data.Date_out != "" {
			event := m.Event{}
			event.Date = TimeParser(data.Date_out)
			event.Comment = data.Out
			event.Act_n = data.Act_out
			event.Status = "OUT"
			animal.Events = append(animal.Events, event)
		}
		animal.Priute.Aviary = data.Volier_n
		animal.Priute.Documents.CardN = data.Card_N
		animal.Priute.Documents.Id_metka = data.Id_m
		animal.Priute.Documents.Act_in = data.Act_in
		animal.Priute.Documents.Date_in = TimeParser(data.Date_in)
		animal.Priute.Documents.Adm_place = data.Adm_place
		animal.Priute.Documents.Act_get = data.Act_get
		animal.Priute.Documents.Place_get = data.Get_Place
		animal.Priute.Documents.Act_out = data.Act_out

		animal.Priute.Responsible.Adress = data.Priute_adress
		animal.Priute.Responsible.Expl_org = data.Expl_org
		animal.Priute.Responsible.Manager = data.Manager
		animal.Priute.Responsible.Resp_own = data.Resp_own

		animal.Owner.FL.Name = data.FL_name
		animal.Owner.FL.Id_num = data.Pasport
		PPrint(data.P_date)
		animal.Owner.FL.Id_date = TimeParser(data.P_date)
		animal.Owner.FL.Id_plase = data.P_place
		animal.Owner.FL.Reg_place = data.FL_adress
		animal.Owner.FL.Phone = data.FL_Phone

		animal.Owner.UL.Name = data.UL_name
		animal.Owner.UL.UL_place = data.UL_adr
		animal.Owner.UL.UL_phone = data.UL_phone
		animal.Owner.UL.Name_owner = data.Name_own
		animal.Owner.UL.Place = data.Addres_own
		animal.Owner.UL.Phone = data.Phone_own

		PPrint(animal)
		result, insertErr := Animals.InsertOne(context.TODO(), animal)
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
			os.Exit(1) // safely exit script on error
		} else {
			fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
			fmt.Println("InsertOne() API result:", result)

			// get the inserted ID string
			newID := result.InsertedID
			fmt.Println("InsertOne() newID:", newID)
			fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
		}
		//*/
	}
}

func TimeParser(s string) time.Time {
	layoutISO := "2006-01-02"
	t, _ := time.Parse(layoutISO, swap_date(s))
	PPrint(t)
	return t
}

func swap_date(in string) string {

	//PPrint := fmt.Println
	var ans, y, m, d string
	if len(in) == 10 {
		d = in[:2]
		m = in[3:5]
		y = in[6:10]
		//PPrint(m)
		ans = y + "-" + m + "-" + d
		//PPrint(ans)
	} else {
		ans = ""
	}
	return ans
}
