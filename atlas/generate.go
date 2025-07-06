package main

import (
	"log"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	DevName  string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка при загрузке .env: ", err)
	}

	cfg := DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		DevName:  os.Getenv("DB_DEV_NAME"),
	}

	tmpl, err := template.ParseFiles("atlas/atlas.hcl.templ")
	if err != nil {
		log.Fatal("Ошибка парсинга шаблона: ", err)
	}

	out, err := os.Create("atlas.hcl")
	if err != nil {
		log.Fatal("Ошибка создания файла atlas.hcl: ", err)
	}
	defer out.Close()

	err = tmpl.Execute(out, cfg)
	if err != nil {
		log.Fatal("Ошибка генерации файла: ", err)
	}

	log.Println("✅ atlas.hcl успешно сгенерирован!")

	// Если хочешь сразу запускать миграции (убери комментарии)
	// cmd := exec.Command("atlas", "migrate", "apply", "--env", "gorm")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatal("Ошибка выполнения atlas: ", err)
	// }
}
