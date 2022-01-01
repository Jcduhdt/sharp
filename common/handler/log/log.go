package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"reflect"
	"sharp/common/consts"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func Init() {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		LevelKey:    "level",                     // 日志级别对应的key名
		MessageKey:  "msg",                       // 日志内容对应的key名，此参数必须不为空
		EncodeLevel: zapcore.CapitalLevelEncoder, //大写不带颜色
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(time.RFC3339))
		},
		CallerKey:    "caller",
		LineEnding:   zapcore.DefaultLineEnding,
		EncodeCaller: zapcore.ShortCallerEncoder, // 配置短路径显示
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter("./log/sharp.log")
	warnWriter := getWriter("./log/sharp.log.wf")

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	// 传入 zap.AddCaller() 才会显示打日志点的文件名和行数，使用sugar便于使用printf
	// 添加zap.AddCallerSkip(1)会打印调用着所在行
	Logger = zap.New(core, zap.AddCaller()).Sugar()
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(3*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func BuildLogByMap(ctx context.Context,params map[string]interface{}) string {
	var (
		buffer         bytes.Buffer
		equalSymbol    = []byte(consts.LogEqualSymbol)
		paramDelimiter = []byte(consts.LogParamDelimiter)
	)

	params["traceid"]=ctx.Value("traceid")
	for k, v := range params {
		var val string
		kind := reflect.ValueOf(v).Kind()
		if kind == reflect.Map || kind == reflect.Struct || kind == reflect.Ptr {
			b, _ := json.Marshal(v)
			val = string(b)
		} else {
			val = fmt.Sprint(v)
		}

		buffer.Write([]byte(fmt.Sprint(k)))
		buffer.Write(equalSymbol)
		buffer.Write([]byte(val))
		buffer.Write(paramDelimiter)
	}
	return buffer.String()
}
