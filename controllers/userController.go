package controllers

import (
	"github.com/Christomesh/go-jwt-project/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword()

func VerifyPassword()

func Login()

func Signup()

func GetAllUsers()

func GetUser()
