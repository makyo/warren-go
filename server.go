package main

import (
	"io/ioutil"
	"os"

	"github.com/go-martini/martini"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"

	"github.com/makyo/warren-go/handlers"
)

type Mongo struct {
	Host string `yaml:"host"`
	DB   string `yaml:"db"`
}

type Config struct {
	AuthKey       string `yaml:"auth-key"`
	EncryptionKey string `yaml:"encryption-key"`
	Mongo         Mongo
}

var (
	store sessions.Store
	db    *mgo.Database
)

func init() {
	var config Config
	file := os.Args[1]
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	store = sessions.NewCookieStore([]byte(config.AuthKey), []byte(config.EncryptionKey))

	dbSession, err := mgo.Dial(config.Mongo.Host)
	if err != nil {
		panic(err)
	}

	db = dbSession.DB(config.Mongo.DB)
}

func main() {
	m := martini.Classic()

	h := handlers.New(store, db)

	m.Get("/", h.Front)

	m.Get("/login", h.DisplayLogin)
	m.Post("/login", h.Login)
	m.Get("/logout", h.Logout)
	m.Get("/register", h.DisplayRegister)
	m.Post("/register", h.Register)
	m.Get("/confirm/:confirmation", h.Confirm)
	m.Get("/~:username", h.DisplayUser)
	m.Post("/~:username/follow", h.FollowUser)
	m.Post("/~:username/unfollow", h.UnfollowUser)

	m.Get("/(?P<post>\\d+)", h.DisplayPost)
	m.Get("/(?P<post>\\d+)/delete", h.DisplayDeletePost)
	m.Get("/(?P<post>\\d+)/delete", h.DeletePost)
	m.Post("/(?P<post>\\d+)/share", h.SharePost)
	m.Get("/post", h.DisplayCreatePost)
	m.Post("/post", h.CreatePost)

	m.Use(h.SessionMiddleware)
	m.Use(h.AuthenticationMiddleware)

	m.Run()
}
