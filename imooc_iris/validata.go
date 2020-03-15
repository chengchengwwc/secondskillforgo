package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"secondskillforgo/imooc_iris/common"
	"secondskillforgo/imooc_iris/encrypt"
	"strconv"
	"sync"
)

//设置集群地址，最好的内网IP
var hostArray = []string{"127.0.0.1", "127.0.0.1"}
var localHost = "127.0.0.1"
var port = "8081"
var hashConsistent = common.NewConsistent()

type AccessControl struct {
	//用来存放用户信息
	sourcesArray map[int]interface{}
	*sync.RWMutex
}

var accessControl = &AccessControl{sourcesArray: make(map[int]interface{})}

//读取记录
func (m *AccessControl) GetNewRecord(uid int) interface{} {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	data := m.sourcesArray[uid]
	return data
}

//设置记录
func (m *AccessControl) SetNewRecord(uid int) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	m.sourcesArray[uid] = "hhhh"
}

func (m *AccessControl) GetDistributeRight(req *http.Request) bool {
	uid, err := req.Cookie("uid")
	if err != nil {
		return false
	}
	//采用一致性hash算法，根据用户ID，判断获取具体机器
	hostRequest, err := hashConsistent.Get(uid.Value)
	if err != nil {
		return false
	}
	//判断是否是本机
	if hostRequest == localHost {
		//执行本机的数据读取和校验
		return m.GetDataFromMap(uid.Value)
	} else {
		//不是本机，则充当代理，访问数据返回结果
		return GetDataFromOtherMap(hostRequest, req)

	}

}

//获取其他节点处理结果
func GetDataFromOtherMap(host string, req *http.Request) bool {
	uidPre, err := req.Cookie("uid")
	if err != nil {
		return false
	}
	//获取sign
	uidSign, err := req.Cookie("sign")
	if err != nil {
		return false
	}
	//模拟接口访问
	client := &http.Client{}
	rq, err := http.NewRequest("GET", "http://"+host+":"+port+"/access", nil)
	if err != nil {
		return false
	}
	// 手动指定cookies
	cookieUid := &http.Cookie{Name: "uid", Value: uidPre.Value, Path: "/"}
	cookieSign := &http.Cookie{Name: "sign", Value: uidSign.Value, Path: "/"}
	rq.AddCookie(cookieUid)
	rq.AddCookie(cookieSign)
	response, err := client.Do(rq)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}
	//判断状态
	if response.StatusCode == 200 {
		if string(body) == "true" {
			return true
		} else {
			return false
		}
	}
	return false

}

func (m *AccessControl) GetDataFromMap(uid string) bool {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return false
	}
	data := m.GetNewRecord(uidInt)
	if data != nil {
		return true
	}
	return false

}

//统一验证拦截器
func Auth(w http.ResponseWriter, r *http.Request) error {
	err := CheckUserInfo(r)
	if err != nil {
		return err
	}

	return nil
}

func Check(w http.ResponseWriter, r *http.Request) {

}

//身份校验函数
func CheckUserInfo(r *http.Request) error {
	//获取uid cookie
	uidCookie, err := r.Cookie("uid")
	if err != nil {
		return errors.New("no uidcookie")
	}
	signCookie, err := r.Cookie("sign")
	if err != nil {
		return errors.New("no sign")
	}
	//对信息进行解密
	signByte, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		return nil
	}
	if checkInfo(uidCookie.Value, string(signByte)) {
		return nil
	}
	return err

}

func checkInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}

func main() {
	//负载均衡器的设置
	//采用一致性哈希算法
	for _, v := range hostArray {
		hashConsistent.Add(v)
	}

	//1 过滤器
	filter := common.NewFilter()
	//2 注册拦截器
	filter.RegisterFilterUri("/check", Auth)
	//3 启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	http.ListenAndServe("8003", nil)

}
