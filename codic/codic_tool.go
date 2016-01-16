package codic_tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"flag"
	"log"
	"os"
	"strings"
)

const CODIC_TOKEN_PATH   = "/tmp/token_codic"
const CODIC_CASING_PATH  = "/tmp/casing_codic"
const DEFALUT_TOKEN_STR  = "default"
const DEFALUT_CASING_STR = "camel"

type Codic struct {
	token  string
	casing string
	query  string
}


func (this *Codic) Run() {
	this.argumentConfiguration()
	result := this.requestCodicApi()
	fmt.Println("[％s] => %s", this.query, result)
}


func (this *Codic) argumentConfiguration(){
	// token, query
	token  := flag.String("token", DEFALUT_TOKEN_STR, "initial setting token.")
	casing := flag.String("casing",DEFALUT_CASING_STR,"[camel, pascal, lower_underscore, upper_underscore, hyphen]")
	if(len(flag.Args()) > 0){
		this.query = flag.Args()[0]
	}
	flag.Parse()

	// token
	if *token != DEFALUT_TOKEN_STR{
		setToTmp(*token, CODIC_TOKEN_PATH)
	}
	this.token = getFromTmp(CODIC_TOKEN_PATH, "!Unset Token!")

	// casing
	if *casing != DEFALUT_CASING_STR{
		setToTmp(*casing, CODIC_CASING_PATH)
	}
	this.casing = getFromTmp(CODIC_CASING_PATH, "")
}

func (this *Codic) requestCodicApi()(result string){
	// build url
	requestUrl := "https://api.codic.jp/v1/engine/translate.json?text="
	requestUrl += url.QueryEscape(this.query)
	// CLI引数の関係で_をwhite spaceに変換
	requestUrl += "&casing="+url.QueryEscape(strings.Replace(this.casing,"_"," ",-1))

	client := &http.Client{}

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+this.token)
	r, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if r.StatusCode == http.StatusUnauthorized {
		log.Fatal("Maybe AccessToken is expired. Please reset.")
		os.Exit(0)
	}

	defer r.Body.Close()

	var datas []Response
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &datas)

	if datas[0].Successful{
		return datas[0].Translated_text
	}else{
		log.Fatal("Request Failed")
		os.Exit(0)
	}
	return ""
}
