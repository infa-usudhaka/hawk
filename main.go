package main


import (
	"net/http"
	"./hawkservice"
)

/*

type config struct {
	
    Service struct{ 
	Arnrole string `yaml: "arnrole"`
    Image struct{
    Repository string `yaml: "repository"`
    } 
    
    }`yaml: "service"`
	HelmVersion 	 string `yaml:"helmVersion"`
	Error string
   }



func (c *config) getConfig(path string) *config {
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

func listDir( path string) []string{
	dirList := make([]string, 50)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		dirList=append(dirList,f.Name())
	}

	return dirList

}
*/
func main(){


	hawkservice.Init()
	http.ListenAndServe(":4040", nil)

/*
QaStagingPath:="C:\\Projects\\Go\\hawk\\Automation\\ccgf-qastaging-hawk-config-master\\"
PerfPath:="C:\\Projects\\Go\\hawk\\Automation\\Perf\\"

var qaConfig config
var perfConfig config
var colorService string
//var colorHelm string

//dirListQA:=listDir(QaStagingPath)
dirListPerf:=listDir(PerfPath)

fmt.Println(dirListPerf)

htmlData:="<html> <table style='backgound:#fff;border-collapse: collapse;' border = '1' cellpadding = '6'> <tr><th>Service Name</th> <th> QA Service Version</th>  <th>Perf Service Version </th> <th> QA Helm Version</th>  <th>Perf Helm Version </th></tr>  "

//fmt.Println(dirListPerf)
for _,i := range dirListPerf{

	if i!="" && i!="README.md" {
	
	qaConfig.getConfig(QaStagingPath+i+"\\configuration.yaml")
	perfConfig.getConfig(PerfPath+i+"\\configuration.yaml")
	
	//skip the service if repo empty
	if qaConfig.Service.Image.Repository=="" || perfConfig.Service.Image.Repository=="" {
		continue
	}

	repoQA:=strings.Split(qaConfig.Service.Image.Repository,":")
	repoPerf:=strings.Split(perfConfig.Service.Image.Repository,":")
	
	fmt.Println(i,repoPerf)
	if repoQA[1]!=repoPerf[1]{
		colorService="red"
	}else{
		colorService="green"
	}
	//fmt.Println(repoQA,repoPerf)
	helmVersionQA:=""
	helmVersionPerf:=""
	

	if qaConfig.HelmVersion !="" {
		helmVersionQA=qaConfig.HelmVersion
		helmVersionPerf=qaConfig.HelmVersion
		//fmt.Println(helmVersionQA,helmVersionPerf)
	}

	htmlData=htmlData+"<tr> <td style='background:"+colorService+"'> "+i+"</td><td>"+repoQA[1]+"</td><td>"+repoPerf[1]+"</td><td>"+helmVersionQA+"</td><td>"+helmVersionPerf+"</td></tr>"


}

}
//fmt.Println(htmlData)
subject:="Service Version Comparison"
cc:="asdash@informatica.com"
utils.SendMail(htmlData, subject, cc)





//helmVersion:=c.HelmVersion
//fmt.Println(c)


//fmt.Println(repo)
*/

}
