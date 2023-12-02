package find_answer

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var key = []byte("u2oh6Vu^HWe4_AES")

// 登录接口（post）
const LOGIN_URL = "http://passport2.chaoxing.com/fanyalogin"

// 获取课程（get）
const GET_COURSE_URl = "https://mooc2-ans.chaoxing.com/mooc2-ans/visit/courses/list"

// 获取课程作业（get）
const GET_WORK_URL = "https://mooc1.chaoxing.com/mooc2/work/list"

// 提交作业（post）
const COMMIT_WORK = "https://mooc1.chaoxing.com/work/addStudentWorkNewWeb"

// 确认提交（get）
const IS_COMMIT = "https://mooc1.chaoxing.com/work/validate"

// 获取个人信息(get)
const SELF_INFO = "http://passport2.chaoxing.com/mooc/accountManage"

// 获取二维码图片（get）
const GET_LOGIN_QR = "https://passport2.chaoxing.com/createqr"

// 判断是否二维码登录成功(post)
const IS_QR_LOGIN = "https://passport2.chaoxing.com/getauthstatus"

// 登录页面首页（get）
const HOME_LOGIN = "https://passport2.chaoxing.com/login"

// 重做作业（get）
const REDO_WORK = "https://mooc1.chaoxing.com/work/phone/redo"

// 答案题目类型
var answerType = []map[string]string{
	{"type": "单选题", "fun": "multipleChoice", "key": "0"},
	{"type": "多选题", "fun": "multipleChoices", "key": "1"},
	{"type": "判断题", "fun": "judgeChoice", "key": "3"},
	{"type": "填空题", "fun": "comprehensive", "key": "2"},
	{"type": "简答题", "fun": "shortAnswer", "key": "4"},
	{"type": "论述题", "fun": "essayQuestion", "key": "6"},
	{"type": "编程题", "fun": "programme", "key": "9"},
	{"type": "其他", "fun": "other", "key": "8"},
}

// 问题题目类型
var questionType = []map[string]string{
	{"type": "单选题", "fun": "multipleChoice"},
	{"type": "多选题", "fun": "multipleChoices"},
	{"type": "判断题", "fun": "judgeChoice"},
	{"type": "填空题", "fun": "comprehensive"},
	{"type": "简答题", "fun": "shortAnswer"},
	{"type": "论述题", "fun": "essayQuestion"},
	{"type": "编程题", "fun": "programme"},
	{"type": "其他", "fun": "other"},
}

func encryptAES(plainText string) string {
	block, _ := aes.NewCipher(key)
	encrypted := make([]byte, len(plainText))
	block.Encrypt(encrypted, []byte(plainText))
	return base64.StdEncoding.EncodeToString(encrypted)
}

func decryptAES(cipherText string) string {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(cipherText))
	block.Decrypt(decrypted, []byte(cipherText))
	return string(decrypted)
}

// END: ed8c6549bwf9

func getWorkScore(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "无"
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "无"
	}

	score := doc.Find(".score").Text()
	if score == "" {
		score = doc.Find(".p").Children().Text()
		if score == "" {
			score = "无"
		}
	}

	return score
}

func getIsRedo(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "no"
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "no"
	}

	redo := doc.Find(".a").Text()
	if strings.Contains(redo, "重做") {
		return "yes"
	}

	return "no"
}
