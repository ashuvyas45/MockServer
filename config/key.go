package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func mustHaveString(key string) string {
	mustHave(key)
	return viper.GetString(key)
}

func mustHaveBool(key string) bool {
	mustHave(key)
	return viper.GetBool(key)
}

func mustGetTimeDurationInS(key string) time.Duration {
	return time.Second * time.Duration(mustGetInt(key))
}

func mustGetTimeDurationInMs(key string) time.Duration {
	return time.Millisecond * time.Duration(mustGetInt(key))
}

func mustGetInt(key string) int {
	mustHave(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid integer value", key))
	}
	return v
}

func mustHave(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}
