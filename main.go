package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v2"
)

type Config struct {
    QQ       int64  `yaml:"qq"`
    Pwd      string `yaml:"pwd"`
    Platform int    `yaml:"platform"`
}

func addToYaml(config Config) {
    //读取./data/ysQQ/qq.yaml
    data, err := ioutil.ReadFile("./data/ysQQ/qq.yaml")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 创建一个 Config 类型的变量
    var configList []Config
    // 将文件内容转换为 Config 结构体
    yaml.Unmarshal(data, &configList)
    //追加到configList
    configList = append(configList, config)
    // 将 Config 结构体转换为 yaml 格式的字符串
    data, err = yaml.Marshal(configList)
    if err != nil {
        fmt.Println(err)
        return
    }
    // 将字符串写入到文件中
    ioutil.WriteFile("./data/ysQQ/qq.yaml", data, 0666)
}

func writeQQData(QQNumber int64, Pwd string, Platform int) {
    if Platform == 0 {
        Platform = 5
    }
    // 创建一个 Config 类型的变量
    config := Config{
        QQ:       QQNumber,
        Pwd:      Pwd,
        Platform: Platform,
    }
    // 将 Config 结构体转换为 yaml 格式的字符串
    data, err := yaml.Marshal(config)
    if err != nil {
        fmt.Println(err)
        return
    }
    // 将字符串写入到文件中
    ioutil.WriteFile("./config/config/qq.yaml", data, 0666)
}

func readConfig() Config {
    // 读取 yaml 文件的内容
    data, err := ioutil.ReadFile("./config/config/qq.yaml")
    if err != nil {
        fmt.Println(err)
        return Config{}
    }
    // 创建一个 Config 类型的变量
    var config Config
    // 将文件内容转换为 Config 结构体
    yaml.Unmarshal(data, &config)
    // 输出里面的 qq 字段
    return config
}

func readQQNum() int64 {
    //检查是否存在./config/config/qq.yaml，没有就return "1"
    _, err := os.Stat("./config/config/qq.yaml")
    if err != nil {
        return 0
    }
    // 读取 yaml 文件的内容
    data, err := ioutil.ReadFile("./config/config/qq.yaml")
    // 创建一个 Config 类型的变量
    var config Config
    // 将文件内容转换为 Config 结构体
    yaml.Unmarshal(data, &config)
    // 输出里面的 qq 字段
    return config.QQ
}

func readInt() int64 {
    var num int64
    fmt.Scanln(&num)
    return num
}

func readString() string {
    var str string
    fmt.Scanln(&str)
    return str
}

func changeQQAndPwd() {
    c := readConfig()
    if c.QQ != 0 {
        fmt.Print("请输入QQ号(直接回车将使用 ", c.QQ, " 登录)：")
        var i int64
        fmt.Scanln(&i)
        //检测用户是否直接回车
        if i == 0 {
            i = c.QQ
        }
        fmt.Print("请输入QQ密码(直接回车将使用 ", c.Pwd, " 登录)：")
        p := readString()
        //检测用户是否直接回车
        if len(p) == 0 {
            p = c.Pwd
        }
        fmt.Print("请选择登录方式(1:安卓手机、 2:aPad 、 3:安卓手表、 4:MacOS 、 5:iPad(默认))：")
        var platform int
        fmt.Scanln(&platform)
        writeQQData(i, p, platform)
        fmt.Println("修改成功！")
    } else {
        fmt.Print("请输入QQ号：")
        i := readInt()
        fmt.Print("请输入QQ密码：")
        p := readString()
        fmt.Print("请选择登录方式(1:安卓手机、 2:aPad 、 3:安卓手表、 4:MacOS 、 5:iPad(默认))：")
        var platform int
        fmt.Scanln(&platform)
        writeQQData(i, p, platform)
        fmt.Println("修改成功！")
    }
}

func switchToPwdLogin() {
    c := readConfig()
    fmt.Print("请输入QQ密码：")
    newPwd := readString()
    writeQQData(c.QQ, newPwd, c.Platform)
    fmt.Println("修改成功！")
}

func changeAccount() {
    //读取./data/ysQQ/qq.yaml
    data, err := ioutil.ReadFile("./data/ysQQ/qq.yaml")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 创建一个 Config 类型的变量
    var config []Config
    // 将文件内容转换为 Config 结构体
    yaml.Unmarshal(data, &config)
    //输出所有信息
    for i, v := range config {
        fmt.Printf("%d. 账号: %d\t\t密码: %s\t\t登录方式: %d\n", i+1, v.QQ, v.Pwd, v.Platform)
    }
    //选择要切换的账号
    fmt.Print("请输入要切换的账号(1-", len(config), ")：")
    var num int
    fmt.Scanln(&num)
    num--
    //切换账号
    writeQQData(config[num].QQ, config[num].Pwd, config[num].Platform)
    fmt.Println("切换成功！")
}

func initialization() {
    //检测./data是否存在名为ysQQ的文件夹，没有就创建,并且创建qq.yaml
    _, err := os.Stat("./data/ysQQ")
    if err != nil {
        os.MkdirAll("./data/ysQQ", 0666)
    }
    _, err = os.Stat("./data/ysQQ/qq.yaml")
    if err != nil {
        os.Create("./data/ysQQ/qq.yaml")
    } else {
        return
    }
    nowQQInfo := readConfig()
    //如果QQ号不存在，就不记录，如果存在，则创建结构体，记录QQ号和密码，并写入到./data/ysQQ/
    var config []Config
    if nowQQInfo.QQ != 0 {
        config = []Config{
            {QQ: nowQQInfo.QQ, Pwd: nowQQInfo.Pwd, Platform: nowQQInfo.Platform},
        }
    }

    //写入到./data/ysQQ/qq.yaml
    data, err := yaml.Marshal(config)
    if err != nil {
        fmt.Println(err)
        return
    }
    // 将字符串写入到文件中
    ioutil.WriteFile("./data/ysQQ/qq.yaml", data, 0666)
}

func menu() {
    fmt.Println("1.修改QQ号、密码、登录方式")
    fmt.Println("2.切换为密码登录")
    // fmt.Println("3.切换QQ账号")
    fmt.Println("0.退出")
    fmt.Print("请选择(0-2)：")
}

func main() {

    // initialization()
    //检测当前文件夹是否存在config文件夹，没有就输出"请将本程序复制到云崽根目录下使用"
    _, err := os.Stat("./config")
    if err != nil {
        fmt.Println("请将本程序复制到云崽根目录下使用")
        return
    }

    menu()
    for {
        //读取用户输入
        var input int
        fmt.Scanln(&input)
        loop := true
        switch input {
        case 1:
            changeQQAndPwd()
            loop = false
        case 2:
            switchToPwdLogin()
            loop = false
        // case 3:
        //     changeAccount()
        //     loop = false
        case 0:
            return
        default:
            fmt.Println("输入错误，请重新输入")
        }

        if loop == false {
            break
        }
    }
    fmt.Println("按回车退出...")
    fmt.Scanln()
}
