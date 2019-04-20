package common

func Close() error {
	if DefaultDB != nil {
		DefaultDB.Close()
	}

	return nil
}
