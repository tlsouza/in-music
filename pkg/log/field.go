package log

import "go.uber.org/zap"

func Any(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func String(key string, val string) zap.Field {
	return zap.String(key, val)
}

func Bool(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}

func Error(err error) zap.Field {
	return zap.Error(err)
}
