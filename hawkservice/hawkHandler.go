package hawkservice


import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"gopkg.in/yaml.v2"
	"../utils"
	"net/http"
	//"../config"
	"encoding/json"
	//"os"
	//"path/filepath"
)



type cfg struct {
	
    Service struct{ 
	Arnrole string `yaml: "arnrole"`
    Image struct{
    Repository string `yaml: "repository"`
    } 
    
    }`yaml: "service"`
	HelmVersion 	 string `yaml:"helmVersion"`
	Error string
   }


   type Body struct {
	ResponseCode int
	Message      string
}


/*type config struct{
	A string `yaml: "a"`
	//B string `yaml: "b"`
}*/

func (c *cfg) getConfig(path string) *cfg {
	//var c config
	
	yamlFile, err := ioutil.ReadFile(path)
if err != nil {
	log.Printf("yamlFile.Get err   #%v ", err)

}

//fmt.Println(yamlFile)
err = yaml.Unmarshal(yamlFile, &c)
if err != nil {
	log.Fatalf("Unmarshal: %v", err)
}else{
	return c	
}

return c
}


//func to get index value for env

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
 }


func listDir( path string) []string{

	fmt.Println(path)
dirList := make([]string, 100)

files, err := ioutil.ReadDir("C:\\HAWK\\Repo\\ccgf-qastaging-hawk-config")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
       // fmt.Println(file.Name())
		dirList=append(dirList,file.Name())
    }


//fmt.Println(dirList)

	

	/***
	files, err := ioutil.Read(path)
	if err != nil {
		fmt.Println("cannot read folder")
		log.Fatal(err)
	}
	for _, f := range files {
		dirList=append(dirList,f.Name())
	}

	***/

/*	for _,fn:=range list {

		dirList=append(dirList,fn)
	}*/

	return dirList

}

func compareEnv(w http.ResponseWriter, r *http.Request){



//QaStagingPath:="C:\\Projects\\Go\\hawk\\Automation\\ccgf-qastaging-hawk-config-master\\"
//PerfPath:="C:\\Projects\\Go\\hawk\\Automation\\Perf\\"

conf := utils.ReadConfig()

repoPath:=conf.RepoPath
fmt.Println(repoPath)

env1:=strings.TrimSpace(r.URL.Query().Get("env1"))
env2:=strings.TrimSpace(r.URL.Query().Get("env2"))

env1Path:=strings.TrimSpace(repoPath+"\\"+env1)
//env1Path:="C:\\HAWK\\Repo\\ccgf-qastaging-hawk-config\\"
env2Path:=strings.TrimSpace(repoPath+"\\"+env2)
//env2Path:="C:\\HAWK\\Repo\\ccgf-perf-hawk-config\\"


//get the actual env name from properties file
env1=conf.EnvList[indexOf(env1,conf.EnvRepo)]
env2=conf.EnvList[indexOf(env2,conf.EnvRepo)]

fmt.Printf("%s,%s",env1,env2)
fmt.Println(env1Path)

CC := r.URL.Query().Get("Email") 
//CC:="usudhakar@informatica.com"
var qaConfig cfg
var perfConfig cfg
var colorService string
//var colorHelm string

//dirListQA:=listDir(QaStagingPath)
dirListPerf:=listDir(env2Path)

fmt.Println(dirListPerf)

htmlData:="<html> <table style='backgound:#fff;border-collapse: collapse;' border = '1' cellpadding = '6'> <tr style='background:#000;color:#fff'><th>Service Name</th> <th>"+env1+" Service Version</th>  <th>"+env2+" Service Version </th> <th> "+env1+" Helm Version</th>  <th>"+env2+" Helm Version </th></tr>  "

//fmt.Println(dirListPerf)
for _,i := range dirListPerf{

	if i!="" && i!="README.md" && i!=".git" {
	
	qaConfig.getConfig(env1Path+"\\"+i+"\\configuration.yaml")
	perfConfig.getConfig(env2Path+"\\"+i+"\\configuration.yaml")
	
	//skip the service if repo empty
	if qaConfig.Service.Image.Repository=="" || perfConfig.Service.Image.Repository=="" {
		continue
	}

	repoQA:=strings.Split(qaConfig.Service.Image.Repository,":")
	repoPerf:=strings.Split(perfConfig.Service.Image.Repository,":")
	
	fmt.Println(i,repoPerf)
	if repoQA[1]!=repoPerf[1]{
		colorService="#ff8080"
	}else{
		colorService="green"
	}
	//fmt.Println(repoQA,repoPerf)
	helmVersionQA:=""
	helmVersionPerf:=""
	

	if qaConfig.HelmVersion !="" {
		helmVersionQA=qaConfig.HelmVersion
		helmVersionPerf=perfConfig.HelmVersion
		//fmt.Println(helmVersionQA,helmVersionPerf)
	}

	htmlData=htmlData+"<tr> <td style='background:"+colorService+"'> <b>"+i+"</b></td><td>"+repoQA[1]+"</td><td>"+repoPerf[1]+"</td><td>"+helmVersionQA+"</td><td>"+helmVersionPerf+"</td></tr>"


}

}
//fmt.Println(htmlData)
subject:="Service Version Comparison "+env1+" vs "+env2

utils.SendMail(htmlData, subject, CC)

utils.RespondWithJSON("Email Sent Successfully", w, r)



}



func test(w http.ResponseWriter, r *http.Request) {

	body := Body{ResponseCode: 200, Message: "OK"}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBody)

}