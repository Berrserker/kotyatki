package router

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/auth0/go-jwt-middleware"

	"github.com/dgrijalva/jwt-go"

	"os"

	"fmt"

	"github.com/gorilla/mux"

	u "../utils"

	"strconv"

	"github.com/buger/jsonparser"

	"github.com/gorilla/handlers"
)

type Response struct {
	Message string `json:"message"`
}

func Route(config *[]byte) {

	//initiate var to read config, use custom jsonparser
	var parsed_port int64
	var parsed_db, parsed_views, parsed_static /*parsed_files,*/, parsed_img, parsed_secret string
	var parsed_error error

	//parse config

	parsed_db, parsed_error = jsonparser.GetString(*config, "db_name")
	parsed_port, parsed_error = jsonparser.GetInt(*config, "port")
	parsed_views, parsed_error = jsonparser.GetString(*config, "path_to_views")
	parsed_static, parsed_error = jsonparser.GetString(*config, "path_to_static")
	//parsed_files, parsed_error = jsonparser.GetString(*config, "path_to_files")
	parsed_img, parsed_error = jsonparser.GetString(*config, "path_to_img")
	parsed_secret, parsed_error = jsonparser.GetString(*config, "secret")
	parse_file, _ := jsonparser.GetBoolean(*config, "parse_file")
	csv_path, _ := jsonparser.GetString(*config, "csv_path")
	delimeter, _ := jsonparser.GetString(*config, "delimeter")
	os.Setenv("secret", parsed_secret)
	if parsed_error != nil {
		fmt.Print(parsed_error)
		panic(parsed_error)
	}

	//initiate utils
	u.Init(parsed_db)
	//fmt.Print(u.SetPwd("Password") + "\n")
	//deinit utils
	//defer u.Kill()
	//u.PPrint(parse_file)
	if parse_file == true {
		go u.Csv(csv_path, delimeter)
	}
	//set port to http server
	port := strconv.Itoa(int(parsed_port))

	//set cors up (dev only)
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	//set jwt up
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(parsed_secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	//u.PPrint(port)
	//u.PPrint(parsed_files)
	//u.PPrint(parsed_views)
	//create router
	router := mux.NewRouter()

	//set static routes
	//u.ListDirByReadDir(parsed_views)
	//fs := http.FileServer(http.Dir(parsed_views))
	//http.Handle("/", http.StripPrefix("/", fs))

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir(parsed_static)))

	//set api handlers

	//router.HandleFunc("/status", u.Status).Methods("GET")
	router.Handle("/status", jwtMiddleware.Handler(http.HandlerFunc(u.Status))).Methods("GET")
	router.Handle("/get_voc", jwtMiddleware.Handler(http.HandlerFunc(u.ShowListvoc))).Methods("GET")
	router.Handle("/get_voc", jwtMiddleware.Handler(http.HandlerFunc(u.Showvoc))).Methods("POST")
	router.Handle("/set_voc", jwtMiddleware.Handler(http.HandlerFunc(u.Setvoc))).Methods("POST")
	//router.Handle("/showcase", u.Showcase).Methods("GET")
	//router.Handle("/showcase{id}", u.ShowcaseId).Methods("GET")
	//router.Handle("/set_case", jwtMiddleware.Handler(http.HandlerFunc(u.SetNewCase))).Methods("POST")
	//router.Handle("/showcase{id}", jwtMiddleware.Handler(http.HandlerFunc(u.ShowcaseWrite))).Methods("POST")
	//router.Handle("/showcase{name}", jwtMiddleware.Handler(u.Showcase)).Methods("GET")
	router.Handle("/profile", jwtMiddleware.Handler(http.HandlerFunc(u.Profile))).Methods("GET")
	//router.Handle("/profile", jwtMiddleware.Handler(u.ProfileSave)).Methods("POST")

	router.HandleFunc("/auth", u.Auth).Methods("POST")
	//router.HandleFunc("/showcase/{slug}/feedback"u.Showcase).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(parsed_static))))
	router.PathPrefix("/img/").Handler(http.FileServer(http.Dir(parsed_img)))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(parsed_views)))
	//port = os.Getenv("PORT")
	fmt.Println("Service start at port: ", port)
	err := http.ListenAndServe(":"+port, corsWrapper.Handler(handlers.LoggingHandler(os.Stdout, router)))
	if err != nil {
		fmt.Print(err)
	}
}
