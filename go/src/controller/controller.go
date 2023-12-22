package controller

import (
	"src/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var newUser model.User

	c.ShouldBindJSON(&newUser)

	// ユーザーがいるかどうか探す
	userID, err := model.GetUserID(&newUser)

	if err != nil{
		c.JSON(401, gin.H{"error":"ユーザーが見つかりません"})
		return
	}

	tokenString, err := model.GenerateToken(userID)
	if err != nil{
		c.JSON(401, gin.H{"error":"トークンの作成に失敗しました"})
	}

	c.JSON(200, gin.H{"token":tokenString})
}

func logout(c *gin.Context) {
	userID := getUserIDforHeader(c)
    model.DeleteToken(userID)
    c.JSON(200, gin.H{"message":"ok"})

}

func signUp(c *gin.Context) {
	var newUser model.User

	// リクエストボディからユーザーが入力した情報を取得
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "無効なリクエスト"})
		return
	}

	// 同じ名前のユーザーがいないか
	_, result := model.GetUserID(&newUser)

	if result == nil{
		c.JSON(200, gin.H{"error":"ユーザーネームはすでに使用されています"})
		return
	}

	// パスワードをハッシュ化する処理

	// modelパッケージのCreateUserを呼び出す
	result = model.CreateUser(&newUser)

	if result != nil{
		c.JSON(500, gin.H{"error": "ユーザーの作成に失敗しました"})
		return
	}
	userID, result := model.GetUserID(&newUser)

	if result != nil{
		c.JSON(500, gin.H{"error": "ユーザーの作成に失敗しました"})
		return
	}

	model.InitHighScore(userID)

	model.InitHighScore(newUser.UserID)

	c.JSON(200, gin.H{"message": "ユーザーが正常に登録されました"})
}

func getRanking(c *gin.Context){

	var ranking model.Ranking

	scores, err := model.GetRanking()

	if err != nil{
		c.JSON(400, gin.H{"error":"ランキングの取得に失敗しました"})
		return
	}

	if scores == nil{
		c.JSON(400, gin.H{"error":"scoresがnilです"})
	}

	for _, userScore := range scores{
		fmt.Println(userScore.UserID)
		fmt.Println(userScore.Score)
		fmt.Println(userScore.ScoreID)

		username, result := model.GetUserName(userScore.UserID)
		fmt.Println(username)

		if result != nil{
			c.JSON(400, gin.H{"error":"ランキングの取得に失敗しましたあ"})
			return
		}

		score := model.UserScore{
			UserName: username,
			Score:    userScore.Score,
		}

		ranking.Ranking = append(ranking.Ranking, score)
	}

	c.JSON(200, ranking)
}

// scoreを返す、ハイスコアが更新されたらtrue
func getResult(c *gin.Context){

	// 投げられてくる形を定義
	type jsonScore struct{
		Score int `json:"score"`
	}

	var userScore jsonScore

	if err := c.ShouldBindJSON(&userScore); err != nil {
		c.JSON(400, gin.H{"error": "無効なリクエスト"})
		return
	}

	userID := getUserIDforHeader(c)

	dbScore,err := model.GetHighScore(userID)

	if err != nil{
		c.JSON(400, gin.H{"error":"ハイスコアの取得に失敗しました"})
		return
	}
	
	// ハイスコアが更新されたか
	if dbScore < userScore.Score{
		newScore := model.Scores{
			Score: userScore.Score,
			UserID: userID,
		}
		
		model.UpdateHighScore(&newScore)

		c.JSON(200, gin.H{"score":newScore.Score, "isUpdateHighScore":true})
		return
	}

	c.JSON(200, gin.H{"score":userScore.Score, "isUpdateHighScore":false})	
}

func getScore(c *gin.Context){

	type jsonCity struct {
		Cities []string `json:"cities"`
	}

	var cities jsonCity

	if err := c.ShouldBindJSON(&cities); err != nil {
		c.JSON(400, gin.H{"error": "無効なリクエスト"})
		return
	}
	
	result, err := model.GetScore(cities.Cities)

	if err != nil{
		c.JSON(400, gin.H{"error": "スコアの取得に失敗しました"})
	}

	c.JSON(200, gin.H{"score":result})
}

func getUserIDforHeader(c *gin.Context) int {

	claims := c.MustGet("claims").(jwt.MapClaims)
	userID := int(claims["user_id"].(float64)) // ユーザーIDの取得

	return userID
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}