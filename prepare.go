package main

import (
	"NetWatchDashboard/models"
	_ "NetWatchDashboard/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

var NwdData = "./data"
var userDefaultLogin = "admin"
var userDefaultPassword = "admin12345"

func prepareRunApplication() error {

	//beego.BConfig.WebConfig.Session.SessionProvider = "file"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"

	if err := prepareDataFolder(); err != nil {
		return err
	}
	if err := prepareConfig(); err != nil {
		return err
	}

	return nil
}

func prepareConfig() error {
	orm.Debug = true
	err := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		return err
	}
	err = orm.RegisterDataBase("config", "sqlite3", fmt.Sprintf("%s/config.sqlite", NwdData))
	if err != nil {
		return err
	}
	orm.RegisterModel(new(models.User))
	err = orm.RunSyncdb("config", true, true)
	if err != nil {
		return err
	}

	o := orm.NewOrmUsingDB("config")
	count, err := o.QueryTable("user").Filter("Role", 1).Filter("Enabled", true).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		password := []byte(userDefaultPassword)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u := models.User{
			Login:    userDefaultLogin,
			Password: string(hashedPassword),
			Role:     1,
			Enabled:  true,
		}
		_, err = o.Insert(&u)
		if err != nil {
			return err
		}

		log.Printf("Создан пользователь %s, пароль: %s", userDefaultLogin, userDefaultPassword)
	}
	return nil
}

func prepareDataFolder() error {
	// Получаем путь к директории из переменной окружения NWD_DATA
	dataPath := os.Getenv("NWD_DATA")
	if dataPath != "" {
		// Если переменная окружения не установлена, используем глобальное значение
		NwdData = dataPath
	}

	// Проверяем существует ли директория
	_, err := os.Stat(NwdData)
	if err != nil {
		if os.IsNotExist(err) {
			// Если директория не существует, создаем её
			err := os.MkdirAll(NwdData, 0755) // 0755 - права доступа к создаваемой директории
			if err != nil {
				log.Printf("Ошибка при создании директории %s: %s", NwdData, err.Error())
				return err
			}
		} else {
			log.Printf("Ошибка при проверке директории %s: %s", NwdData, err.Error())
			return err
		}
	}

	// Проверяем права на запись
	if _, err := os.Stat(NwdData); os.IsPermission(err) {
		fmt.Printf("У вас нет прав на запись в эту директорию %s", NwdData)
		return err
	}
	return nil
}
