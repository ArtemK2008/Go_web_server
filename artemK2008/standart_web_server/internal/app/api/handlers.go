package api

import (
	"encoding/json"
	"fmt"
	"github.com/artemK2008/standart_web_server/internal/app/middleware"
	"github.com/artemK2008/standart_web_server/internal/app/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	isError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get All Articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "we have some troubles, try later",
			isError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article by ID /api/v1/article/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} params", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Cant parse id to int",
			isError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindById(id)
	if err != nil {
		api.logger.Info("Troubles while access DB with with id, err", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Cant access DB, try again",
			isError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("cant fing article with id")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that id does not exist in DB",
			isError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}
func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Article by ID /api/v1/article/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} params", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Cant parse id to int",
			isError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.Article().FindById(id)
	if err != nil {
		api.logger.Info("Troubles while access DB with with id, err", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Cant access DB, try again",
			isError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("cant find article with id")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that id does not exist in DB",
			isError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Error while deleting user ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Could not delete,try again",
			isError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with ID %d deleted", id),
		isError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post articles POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid JSON recieved")
		msg := Message{
			StatusCode: 400,
			Message:    "Privvided JSON invalid",
			isError:    true,
		}
		json.NewEncoder(writer).Encode(msg)
		writer.WriteHeader(400)
		return

	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("troubles while creating new article")
		msg := Message{
			StatusCode: 501,
			Message:    "troubles accesing DB, try again",
			isError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}
func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post user Register POST /api/v1/user/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid JSON recieved")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON invalid",
			isError:    true,
		}
		json.NewEncoder(writer).Encode(msg)
		writer.WriteHeader(400)
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while access DB with with login, err", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Cant access DB, try again",
			isError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("user already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User already exists",
			isError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	u, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("troubles while creating new user")
		msg := Message{
			StatusCode: 500,
			Message:    "troubles accesing DB, try again",
			isError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(u)
}

func (api *API) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post uto auth POST /api/v1/user/auth")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid JSON recieved")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON invalid",
			isError:    true,
		}
		json.NewEncoder(writer).Encode(msg)
		writer.WriteHeader(400)
		return
	}
	userInDb, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while access DB  with login, err", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Cant access DB, try again",
			isError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("You Are not registered")
		msg := Message{
			StatusCode: 400,
			Message:    "No Such User",
			isError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if userInDb.Password != user.Password {
		api.logger.Info("invalid credentials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			isError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // доп действия(в форме мапы) для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		api.logger.Info("Can not claim jwt-token", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have troubles, try again",
			isError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		isError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
