package guacbypasser

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

func W32_RegistryChecker(key32 registry.Key, path string, key string) (err error, code int) {
	var wkey32 registry.Key
	wkey32, err = registry.OpenKey(
		key32, path, registry.QUERY_VALUE|registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = wkey32.GetStringValue(key)
	if err != nil {
		log.Fatal(err)
	}

	return err, registry.QUERY_VALUE
}
