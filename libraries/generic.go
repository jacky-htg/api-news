package libraries

import (
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"
)

var file io.Writer

func init() {
	var err error
	file, err = os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)

	rand.Seed(time.Now().UnixNano())
}

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func Pagination(limit int, page int) int {
	var offset int

	offset = page*limit - limit

	return offset
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
