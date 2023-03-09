package indexer

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
)

type Instance struct {
	Env            string
	DB             *pg.DB
	Daemon         *xelisDaemon.RPC
	ScanHeightPath string
}

func (i *Instance) GetScanHeight() (uint64, error) {
	_, err := os.Stat(i.ScanHeightPath)
	if os.IsNotExist(err) {
		return 0, nil
	}

	buf, err := os.ReadFile(i.ScanHeightPath)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint64(buf), nil
}

func (i *Instance) SaveScanHeight(scanHeight uint64) error {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, scanHeight)
	return os.WriteFile(i.ScanHeightPath, buf, 0644)
}

func LoadInstance(envPath string) (*Instance, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	daemon, err := xelisDaemon.NewRPC(os.Getenv("DAEMON_URL"))
	if err != nil {
		return nil, err
	}

	addr := os.Getenv("PG_ADDR")
	password := os.Getenv("PG_PASSWORD")
	user := os.Getenv("PG_USER")
	database := os.Getenv("PG_DB")
	ssl := os.Getenv("PG_SSL")

	dbOptions := pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: database,
	}

	if ssl == "true" {
		dbOptions.TLSConfig = &tls.Config{
			ClientAuth:         tls.RequestClientCert,
			InsecureSkipVerify: true,
		}
	}

	db := pg.Connect(&dbOptions)

	ctx := context.Background()
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	env := os.Getenv("ENV")
	scanHeightPath := os.Getenv("SCAN_HEIGHT_PATH")

	instance := Instance{
		Env:            env,
		Daemon:         daemon,
		DB:             db,
		ScanHeightPath: scanHeightPath,
	}

	return &instance, nil
}

func (instance *Instance) Close() error {
	if instance.DB != nil {
		err := instance.DB.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
