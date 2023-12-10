package handle

import (
	"RedRockHomework6/dao"
	"RedRockHomework6/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	var info dao.Userinfo
	if err := ctx.ShouldBind(&info); err != nil {
		ctx.JSON(400, gin.H{
			"error": "请完整输入",
		})
		return
	}
	var setIdentify dao.Identify
	setIdentify.Question = ctx.PostForm("question")

	//判断用户是否存在
	n := dao.Db.Model(&dao.Userinfo{}).Where("username = ?", info.Username).First(&dao.Userinfo{}).RowsAffected
	if n != 0 {
		ctx.JSON(400, gin.H{
			"error": "该用户已存在，请更换用户名重试",
		})
		return
	}
	if setIdentify.Question != "" {
		setIdentify.Answer = ctx.PostForm("answer")
		setIdentify.Username = info.Username
		setIdentify.Password = info.Password
		dao.Db.Model(&dao.Identify{}).Create(&setIdentify)
	}
	if err := dao.Db.Model(&dao.Userinfo{}).Create(&info).Error; err != nil {
		ctx.JSON(400, gin.H{
			"error": "创建失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
	return
}
func Login(ctx *gin.Context) {
	var input dao.Userinfo
	var output dao.Userinfo
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(400, gin.H{
			"error": "请完整输入",
		})
		return
	}

	//判断用户是否存在
	n := dao.Db.Model(&dao.Userinfo{}).Where("username = ?", input.Username).Find(&dao.Userinfo{}).RowsAffected
	if n == 0 {
		ctx.JSON(400, gin.H{
			"message": "该用户不存在，请先注册",
		})
		return
	}
	dao.Db.Model(&dao.Userinfo{}).Where("username = ?", input.Username).First(&output)
	if output.Password != input.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "密码不正确，请重试或找回密码",
		})
		return
	}
	token, err := middleware.GetToken(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "颁发token失败",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
	return
}
func CheckQuestion(ctx *gin.Context) {
	var info dao.Identify
	username := ctx.PostForm("username")
	n := dao.Db.Model(&dao.Identify{}).Where("username = ?", username).Find(&info).RowsAffected
	if n == 0 {
		ctx.JSON(400, gin.H{
			"message": "该用户未设置找回密码的问题",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"question": info.Question,
	})
	return
}
func FindPassword(ctx *gin.Context) {
	var answer dao.Identify
	username := ctx.PostForm("username")
	value := ctx.PostForm("answer")
	if username == "" {
		ctx.JSON(400, gin.H{
			"error": "请输入用户名",
		})
		return
	}
	if value == "" {
		ctx.JSON(400, gin.H{
			"error": "请回答问题",
		})
		return
	}
	dao.Db.Model(&dao.Identify{}).Where("username = ?", username).First(&answer)
	if answer.Answer != value {
		ctx.JSON(400, gin.H{
			"message": "回答错误",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":  "回答正确",
		"password": answer.Password,
	})
	return
}
func GetMessageList(ctx *gin.Context) {
	username, ok := ctx.Get("username")
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "登录状态异常",
		})
		return
	}
	var sentMessage []dao.Message
	var receiveMessage []dao.Message
	dao.Db.Model(&dao.Message{}).Where("src_user = ?", username).Find(&sentMessage)
	dao.Db.Model(&dao.Message{}).Where("dest_user = ?", username).Find(&receiveMessage)

	ctx.JSON(200, gin.H{
		"sent message ":    sentMessage,
		"receive message ": receiveMessage,
	})
	return
}
func SendMessage(ctx *gin.Context) {
	var sent dao.Message
	if err := ctx.ShouldBind(&sent); err != nil {
		ctx.JSON(400, gin.H{
			"error": "请完整输入",
		})
		return
	}
	name, ok := ctx.Get("username")
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "登录状态异常",
		})
		return
	}
	sent.SrcUser = name.(string)
	if err := dao.Db.Model(&dao.Message{}).Create(&sent).Error; err != nil {
		ctx.JSON(500, gin.H{
			"message": "发送失败",
			"error":   err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "发送成功",
	})
	return
}
