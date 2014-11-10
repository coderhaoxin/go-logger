package logger

import "strconv"
import "time"
import "log"
import "os"

var l *log.Logger

type config struct {
	Name     string
	Prefix   string
	Logdir   string
	Duration int64
}

var c config

func init() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	c = config{
		Name:     "log",        // log file name
		Prefix:   "",           // log prefix
		Logdir:   "",           // log file dir
		Duration: 24 * 60 * 60, // log file duration
	}

	l = log.New(os.Stdout, c.Prefix, log.LstdFlags)
}

func Config(name string, prefix string, logdir string, duration int64) {
	c.Name = name
	c.Prefix = prefix
	c.Logdir = logdir
	c.Duration = duration

	var err error
	var file *os.File

	if c.Name != "" && c.Logdir != "" && c.Duration != 0 {
		file, err = fresh(nil)

		if err != nil {
			panic(err)
		}

		ticker := time.NewTicker(time.Second * time.Duration(c.Duration))

		go func() {
			for _ = range ticker.C {
				file, err = fresh(file)

				if err != nil {
					panic(err)
				}
			}
		}()
	}
}

func fresh(oldFile *os.File) (*os.File, error) {
	oldFile.Close()

	f, err := os.OpenFile(c.Logdir+"/"+c.Name+"-"+strconv.FormatInt(time.Now().Unix(), 10)+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	l = log.New(f, c.Prefix, log.Lshortfile)
	return f, nil
}

func Fatal(v ...interface{}) {
	l.Fatal(v)
}

func Fatalf(format string, v ...interface{}) {
	l.Fatalf(format, v)
}

func Fatalln(v ...interface{}) {
	l.Fatalln(v)
}

func Flags() int {
	return l.Flags()
}

func Output(calldepth int, s string) error {
	return l.Output(calldepth, s)
}

func Panic(v ...interface{}) {
	l.Panic(v)
}

func Panicf(format string, v ...interface{}) {
	l.Panicf(format, v)
}

func Panicln(v ...interface{}) {
	l.Panicln(v)
}

func Prefix() string {
	return l.Prefix()
}

func Print(v ...interface{}) {
	l.Print(v)
}

func Printf(format string, v ...interface{}) {
	l.Printf(format, v)
}

func Println(v ...interface{}) {
	l.Println(v)
}

func SetFlags(flag int) {
	l.SetFlags(flag)
}

func SetPrefix(prefix string) {
	l.SetPrefix(prefix)
}
