package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"flag"
	"log"
	"os"

	"strings"
)

type Response struct {
	Successful      bool
	Text	        string
	Translated_text string
}

const CODIC_TOKEN_PATH   = "/tmp/token_codic"
const CODIC_CASING_PATH  = "/tmp/casing_codic"
const DEFALUT_TOKEN_STR  = "default"
const DEFALUT_CASING_STR = "camel"

func main() {
	casing := setUpOptionsAndGetCasing()
	args   :=  flag.Args()

	if len(args) > 0{
		getCodic(args[0], casing)
	}else{
		flag.Usage()
		os.Exit(0)
	}
}

func setUpOptionsAndGetCasing() string{
	// token, query
	token  := flag.String("token", DEFALUT_TOKEN_STR, "initial setting token.")
	casing := flag.String("casing",DEFALUT_CASING_STR,"[camel, pascal, lower_underscore, upper_underscore, hyphen]")
	flag.Parse()
	c := *casing

	if *token != DEFALUT_TOKEN_STR{
		setAccessTokenToTmp(*token)
	}
	if c != DEFALUT_CASING_STR {
		setCasingToTmp(c)
	}else {
		c = getCasingFromTmp()
	}
	return c
}

// MARK: AccessToken

func setAccessTokenToTmp(token string){
	content := []byte(token)
	ioutil.WriteFile(CODIC_TOKEN_PATH, content, os.ModePerm)
}

func getAccessTokenFromTmp() string{
	// ファイルの読み込み
	contents,err := ioutil.ReadFile(CODIC_TOKEN_PATH) // ReadFileの戻り値は []byte
	if err != nil {
		println("たぶんtoken setしてないので-token=XXXみたいにCodicでログインしてやってください")
		os.Exit(0)
	}
	return string(contents)
}

// MARK: Casing

func setCasingToTmp(token string){
	content := []byte(token)
	ioutil.WriteFile(CODIC_CASING_PATH, content, os.ModePerm)
}

func getCasingFromTmp() string{
	// ファイルの読み込み
	contents,err := ioutil.ReadFile(CODIC_CASING_PATH) // ReadFileの戻り値は []byte
	if err != nil {
		return DEFALUT_CASING_STR
	}
	return string(contents)
}

// MARK: Codic

func getCodic(query string, casing string){
	// build url
	url := "https://api.codic.jp/v1/engine/translate.json?text="
	url += urlencode(query)
	// CLI引数の関係で_をwhite spaceに変換
	url += "&casing="+urlencode(strings.Replace(casing,"_"," ",-1))
	//println(url)
	//os.Exit(0)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	token := getAccessTokenFromTmp()

	req.Header.Add("Authorization", "Bearer "+token)
	r, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// tokenなし
	if r.StatusCode == http.StatusUnauthorized {
		log.Fatal("AccessTokenがないかExpireしてますよ => ここから取得 https://codic.jp/my/api_status")
		os.Exit(0)
	}

	defer r.Body.Close()

	var datas []Response
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &datas)

	if datas[0].Successful{
		fmt.Println("[",datas[0].Text,"]=>",datas[0].Translated_text)
	}else{
		fmt.Println("これはだめや")
		print()
	}
}

// 参考
// http://qiita.com/hnaohiro/items/80e9112c78f2975f9699
func urlencode(s string) (result string){
	for _, c := range(s) {
		if c <= 0x7f { // single byte
			result += fmt.Sprintf("%%%X", c)
		} else if c > 0x1fffff {// quaternary byte
			result += fmt.Sprintf("%%%X%%%X%%%X%%%X",
				0xf0 + ((c & 0x1c0000) >> 18),
				0x80 + ((c & 0x3f000) >> 12),
				0x80 + ((c & 0xfc0) >> 6),
				0x80 + (c & 0x3f),
			)
		} else if c > 0x7ff { // triple byte
			result += fmt.Sprintf("%%%X%%%X%%%X",
				0xe0 + ((c & 0xf000) >> 12),
				0x80 + ((c & 0xfc0) >> 6),
				0x80 + (c & 0x3f),
			)
		} else { // double byte
			result += fmt.Sprintf("%%%X%%%X",
				0xc0 + ((c & 0x7c0) >> 6),
				0x80 + (c & 0x3f),
			)
		}
	}

	return result
}
