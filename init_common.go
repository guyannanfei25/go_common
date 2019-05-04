package common

func Close() error {
	if DefaultDB != nil {
		DefaultDB.Close()
	}

	if DefaultRedis != nil {
		DefaultRedis.Close()
	}

	return nil
}
