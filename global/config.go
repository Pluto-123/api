package g

import (
	"api_project/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
)

type Config struct {
	Server struct {
		Port          string
		DbAutoMigrate bool // 是否自动迁移数据库表结构
		DbLogMode     string
	}
	JWT struct {
		Secret string
		Expire int64 // hour
		Issuer string
	}
	Mysql struct {
		Host     string // 服务器地址
		Port     string // 端口
		Config   string // 高级配置
		Dbname   string // 数据库名
		Username string // 数据库用户名
		Password string // 数据库密码
	}
}

var Conf *Config

func GetConfig() *Config {
	if Conf == nil {
		log.Panic("配置文件未初始化")
		return nil
	}
	return Conf
}

// ReadConfig 从指定路径读取配置文件
func ReadConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	v.AutomaticEnv()                                   // 允许使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // SERVER_APPMODE => SERVER.APPMODE

	if err := v.ReadInConfig(); err != nil {
		panic("配置文件读取失败: " + err.Error())
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic("配置文件反序列化失败: " + err.Error())
	}

	log.Println("配置文件内容加载成功: ", path)
	return Conf
}

// DbDSN 数据库连接字符串
func (*Config) DbDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		Conf.Mysql.Username, Conf.Mysql.Password, Conf.Mysql.Host, Conf.Mysql.Port, Conf.Mysql.Dbname, Conf.Mysql.Config,
	)
}

// InitDatabase 根据配置文件初始化数据库
func InitDatabase(conf *Config) *gorm.DB {
	dsn := conf.DbDSN()

	var db *gorm.DB
	var err error

	var level logger.LogLevel
	switch conf.Server.DbLogMode {
	case "silent":
		level = logger.Silent
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		fallthrough
	default:
		level = logger.Error
	}

	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		SkipDefaultTransaction:                   true, // 禁用默认事务（提高运行速度）
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true, // 单数表名
		//},
	}
	db, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	log.Println("数据库连接成功", dsn)

	if conf.Server.DbAutoMigrate {
		if err := model.MakeMigrate(db); err != nil {
			log.Fatal("数据库迁移失败", err)
		}
		log.Println("数据库自动迁移成功")
	}
	return db
}
