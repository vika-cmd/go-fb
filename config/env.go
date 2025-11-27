package config

import (
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"os"
	"strconv"
)

// счит из env ф-ла и помещ в перем окр, откуда мы можем исп. их
func Init(){
	//fmt.Println(os.Getwd())
	if err :=godotenv.Load(); err !=nil {
		log.Println("Err-env#Init: not load")
		return
	}
	fmt.Println(".env file loadet!")
	//godotenv.Load(".env")
}

type DatabaseConfig struct{
	Url string
}

func NewDatebaseConfig() *DatabaseConfig{
	return  &DatabaseConfig{
		//Url: os.Getenv("DATABASE_URL"),
		Url: getString("DATABASE_URL", "Any default"),
	}
}

//Default value
func getString(key, defaultValue string) string{
	value :=os.Getenv(key)
	if value ==""{
		return defaultValue
	}
	return value
}

//Bool
type Flag struct{
	Tumbler bool
}

func NewFlag() *Flag{
	return &Flag{
		Tumbler: getBool("FLAG", false),
	}
}
func getBool(key string, defaultValue bool) bool{
	value :=os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err !=nil{
		return defaultValue
	}
	return b
}

//Int CURRENT_VALUE
type CurrentValue struct{
	CurrVal int
}
func NewCurrentVal() *CurrentValue{
	return &CurrentValue{
		CurrVal: getInt("CURRENT_VALUE", 0),
	}
}

func getInt(key string, defaultValue int) int{
	value :=os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err !=nil{
		return defaultValue
	}
	return i
}

