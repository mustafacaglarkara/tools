package EnvIslemleri

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type EnvIslemleri struct {
}

func (e *EnvIslemleri) GetEnvStr(key string) string {
	// .env dosyası yüklenir. hata var ise hata döner.
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Env dosyası okunurken hata meydana geldi. \n Hata: %s", err)
	}
	var tmpDeger = os.Getenv(key)
	if tmpDeger == "" {
		fmt.Errorf("Belirtilen Key bulunamadı. %s", key)
		return ""
	}
	return tmpDeger
}
