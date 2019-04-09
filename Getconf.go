package Getconf

/******************
import config.ini
by:sea
Can't Support nested
The local file's format like this:
*******************
[MYSQL]
DBNAME = comments
DBHOST = 10.10.0.211
DBPORT = 30521
DBUSER = root
DBPASS = 123456

[LISTEN]
HOST = 127.0.0.1
PORT = 8080
********************
will conver to:
********************
map[string] string {
    "MYSQL.DBNAME" : "comments",
    ....
    "LISTEN.HOST" : "127.0.0.1"
}
*******************/

import(
    "bufio"
    "errors"
    "os"
    "strings"
)

// key value
type Config struct{
    conf map[string] string
}

//Get one config by name
func (cf *Config) Get(key string) string {
    return cf.conf[key]
}

// init file
func NewFileConf(filePath string) (*Config, error) {

    cf := &Config {
        conf: make(map[string] string, 20),
    }

    f, err := NewFileReader(filePath)
    if err != nil {
        return nil, errors.New("Error:can not read file \"" + filePath + "\"")
    }
    defer f.Close()

    //配置文件[]中的内容
    tag := ""

    //读取文件句柄
    buf := bufio.NewReader(f)

    //替换掉一行中所有空格
    replacer := strings.NewReplacer(" ", "")

    for {
        //按行读取
        lstr, err := buf.ReadString('\n')
        if err != nil && err != errors.New("EOF") {
            break
        }
        lstr = strings.TrimSpace(lstr)
        if lstr == "" {
            continue
        }
        
        if idx := strings.Index(lstr, "["); idx != -1 {
            //如果是标签
            if lstr[len(lstr) - 1:] != "]" {
                return nil, errors.New("Error:field to parse this symbol style:\"" + lstr + "\"")
            }
            tag = lstr[1:len(lstr) - 1]
        } else {
            lstr = replacer.Replace(lstr)
            
            //注释忽略
            if lstr[0:1] == ";" {
                continue
            }

            //按=辟开单行配置
            spl := strings.Split(lstr, "=")
            
            //错误配置
            if len(spl) < 2 {
                return nil, errors.New("error:" + lstr)
            }
            
            //标签+配置项名称 = value值
            cf.conf[tag + "." + spl[0]] = spl[1]
        }
    }

    return cf, nil
}

// 打开一个文件句柄
func NewFileReader(filePath string) (*os.File, error) {
    if !PathExists(filePath) {
        return nil, errors.New("Error:File not exists:" + filePath)
    }

    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }

    return f, nil
}

// 检查文件或文件夹是否存在
func PathExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }

    if os.IsNotExist(err) {
        return false
    }

    return false
}
