package codic_tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"flag"
	"log"
	"strings"
)

const CODIC_TOKEN_PATH   = "/tmp/token_codic"
const CODIC_CASING_PATH  = "/tmp/casing_codic"
const DEFALUT_TOKEN_STR  = "XXXXX"
const DEFALUT_CASING_STR = "camel"

type Codic struct {
	token  string
	casing string
	query  string
}


func (this *Codic) Run() {
	this.argumentConfiguration()
	result := this.requestCodicApi()
	print(fmt.Sprintf("[%v] => %v\n", this.query, result))
}


func (this *Codic) argumentConfiguration(){
	// token, query
	token  := flag.String("token", DEFALUT_TOKEN_STR, "Need to set AccessToken.")
	casing := flag.String("casing",DEFALUT_CASING_STR,"Optional set casing. [camel, pascal, lower_underscore, upper_underscore, hyphen]")
	flag.Parse()
	if(len(flag.Args()) > 0){
		this.query = flag.Args()[0]
	}else{
		log.Fatal("Please set argument word.")
	}

	// token
	if *token != DEFALUT_TOKEN_STR{
		setToTmp(*token, CODIC_TOKEN_PATH)
	}
	this.token = getFromTmp(CODIC_TOKEN_PATH)
	if this.token == ""{
		log.Fatal("!Unset Token!")
	}

	// casing
	if *casing != DEFALUT_CASING_STR || getFromTmp(CODIC_CASING_PATH) == ""{
		setToTmp(*casing, CODIC_CASING_PATH)
	}
	this.casing = getFromTmp(CODIC_CASING_PATH)
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
	}

	defer r.Body.Close()

	var datas []Response
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &datas)

	if datas[0].Successful{
		return datas[0].Translated_text
	}else{
		println(string(body))
		log.Fatal("Request Failed")
	}
	return ""
}

