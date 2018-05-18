package main

import("fmt"
		 "os"
        "io/ioutil"
         "gopkg.in/mgo.v2"
        "encoding/json"
        "log"
        )
type JsonObject struct{
	SystemInfo SystemInfo
}
type SystemInfo struct {
	GrorVersion string `bson:"grorVersion"`
	Name string `bson:"name"`
}

type AuthDatas struct{
	AuthDatas []AuthData `bson:"authData"`

}
type AuthData struct {  
	Password string `bson:"password"`
    Key string `bson:"key"`
    Username string `bson:"username"`
    Auth string `bson:"auth"`
    Email string  `bson:"email"`
}

type Hosts struct{

	Hosts []Host `bson:"hosts"`
}

type Host struct {

	Protocol string `bson:"protocol"`
	ApiVersion string `bson:"apiVersion"`
	HostType string `bson:"hostType"`
	DockerVersion string `bson:"dockerVersion"`
	Alias string `bson:"alias"`
	CertPathForDockerDaemon string `bson:"certPathForDockerDaemon"`
	IP string `bson:"ip"`
	DockerPort string `bson:"dockerPort"`
}

type Components struct {

	Components []Component `bson:"components"`
}

type Component struct{
	
	Instances []Instance  `bson:"instances"`
	Name string `bson:"name"`
}

type Instance struct {

	EnvMap EnvMap
	PortMapping string `bson:"portMapping"`
	AuthId string `bson:"authId"`
	HostId string `bson:"hostId"`
	VolumeMapping VolumeMapping
	VolumesFrom []string `bson:"volumesFrom"`
	CommandToBeExecuted string `bson:"commandToBeExecuted"`
	Links string `bson:"links"`
	ImageName string `bson:"imageName"`
	Tag string `bson:"tag"`
	HostsMapping string `bson:"hostsMapping"`
	Name string `bson:"name"`

}

type EnvMap struct {
	
	CASSANDRA_BROADCAST_ADDRESS string `bson:"CASSANDRA_BROADCAST_ADDRESS"`
	CASSANDRA_SEEDS string `bson:"CASSANDRA_SEEDS"`

}


type VolumeMapping struct{

	 CassData string `bson:"/home/ubuntu/cass-data"`
	 CassConfig string `bson:"/home/ubuntu/cass-config"`

}

func main(){

	jsonFile,err := os.Open("sample_gror.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened sample_gror.json")
	defer jsonFile.Close()

	byteValue,_ := ioutil.ReadAll(jsonFile)

	
	var components Components
	var hosts Hosts
	var authDatas AuthDatas
	var jsonObject JsonObject

	json.Unmarshal(byteValue, &hosts)
	json.Unmarshal(byteValue, &authDatas)
	json.Unmarshal(byteValue,&components)
	json.Unmarshal(byteValue,&jsonObject)
	

	

	session,err := mgo.Dial("mongodb://127.0.0.1:27017/")
	 if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("demoDB").C("samples")
	 	
	err = c.Insert(&SystemInfo{jsonObject.SystemInfo.GrorVersion,jsonObject.SystemInfo.Name})

	if err !=nil {
		log.Fatal(err)
	}


	for i:=0;i<len(authDatas.AuthDatas);i++{

		err = c.Insert(&AuthData{authDatas.AuthDatas[i].Password,authDatas.AuthDatas[i].Key,authDatas.AuthDatas[i].Username,authDatas.AuthDatas[i].Auth,authDatas.AuthDatas[i].Email})

		if err!=nil{
			log.Fatal(err)
		}
	}

	for i:=0;i<len(hosts.Hosts);i++{

	
		err = c.Insert(&Host{hosts.Hosts[i].Protocol,hosts.Hosts[i].ApiVersion,hosts.Hosts[i].HostType,hosts.Hosts[i].DockerVersion,hosts.Hosts[i].Alias,hosts.Hosts[i].CertPathForDockerDaemon,hosts.Hosts[i].IP,hosts.Hosts[i].DockerPort})

		if err != nil {
                log.Fatal(err)
        }

       
	}


	
	
	for i:=0;i<len(components.Components);i++{

		err = c.Insert(&Component{components.Components[i].Instances,components.Components[i].Name})

		if err!=nil{
		log.Fatal(err)
		}
	
	}
}

	





