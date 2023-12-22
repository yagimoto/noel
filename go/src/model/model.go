package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type score struct {
	ScoreID int `json:"score_id"`
	Score   int `json:"score"`
	UserID  int `json:"user_id"`
}

type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type cities struct {
	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`
	Score    int    `json:"score"`
}

var db *gorm.DB

var Score []score
var City []cities
var rank []score

// var user []User

func Connect() *gorm.DB {
	user := "root"
	pass := "pass"
	progocol := "tcp(hackdenoel-mysql:3306)"
	dbname := "noel-db"
	connect := user + ":" + pass + "@" + progocol + "/" + dbname
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetHighScore(db *gorm.DB, userID int) (int, error) {
	//ユーザーIDからハイスコアを取得
	err := db.Where("user_id = ?", userID).Find(&Score)
	return Score[0].Score, err.Error
}

func UpdateHighScore(db *gorm.DB, newHighScore int, userid int) error {
	// DBに格納されたハイスコアを引数で受け取った値に変更
	err := db.Model(&score{}).Where("user_id = ?", userid).Updates(map[string]interface{}{"score": newHighScore}).Error
	return err
}

func GetRanking(db *gorm.DB) ([]score, error) {
	// スコアテーブルからすべてのユーザーのハイスコアを取得しソート
	err := db.Order("score").Find(&rank).Error
	return rank, err
}

func CreateUser(db *gorm.DB, newUser *User) error {
	// 引数で受け取ったユーザーの情報をDBに格納
	err := db.Create(&newUser).Error
	return err
}

func GetScore(db *gorm.DB, cities []string) (int, []error) {
	// 都市の名前を受け取り、DBからスコアを取得し、合計値を返す
	var err []error
	var result int
	for i := 0; i < len(cities); i++ {
		e := db.Where("city_name = ?", cities[i]).Find(&City)
		err = append(err, e.Error)
		result += City[0].Score
	}
	return result, err
}

func GenerateToken(userID int) (string, error){
    expirationTime := time.Now().Add(time.Hour * 1).Unix()
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     expirationTime, // トークンの有効期限（1時間）
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // secretKey を使ってトークンを署名
    signedToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}


func main() {

}
