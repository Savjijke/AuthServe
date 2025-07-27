package database

import (
	"SRC/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=% s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)

	}
	err = DB.Ping()

	if err != nil {
		log.Fatal("База данных не отвечает:", err)
	}

	fmt.Println("База данных успешно подключилась")
}

func CreateUser(user models.User) error {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`
	err := DB.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		log.Fatal("Ошибка при создании пользователя:", err)
		return err
	}

	return err

}
func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `SELECT id, username, email, password_hash FROM users WHERE email = $1`

	err := DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	return user, err
}
