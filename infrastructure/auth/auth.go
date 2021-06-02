package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type AuthInterface interface {
	CreateAuth(uint64, *TokenDetails) error
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type ClientData struct {
	client redis.Conn
}

var _ AuthInterface = &ClientData{}

func NewAuth(client redis.Conn) *ClientData {
	return &ClientData{client: client}
}

type AccessDetails struct {
	TokenUuid string
	UserId    uint64
}

//Save token metadata to Redis
func (tk *ClientData) CreateAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	_, err := tk.client.Do("SET", td.TokenUuid, strconv.Itoa(int(userid)))
	if err != nil {
		return err
	}
	_, err = tk.client.Do("EXPIRE", td.TokenUuid, int(at.Sub(now).Seconds()))
	if err != nil {
		return err
	}

	_, err = tk.client.Do("SET", td.RefreshUuid, strconv.Itoa(int(userid)))
	if err != nil {
		return err
	}
	_, err = tk.client.Do("EXPIRE", td.RefreshUuid, int(rt.Sub(now).Seconds()))
	if err != nil {
		return err
	}
	return nil
}

//Check the metadata saved
func (tk *ClientData) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := redis.Uint64(tk.client.Do("GET", tokenUuid))
	if err != nil {
		return 0, err
	}
	return userid, nil
}

//Once a user row in the token table
func (tk *ClientData) DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)
	//delete access token
	_, err := tk.client.Do("DEL", authD.TokenUuid)
	if err != nil {
		return err
	}
	//delete refresh token
	_, err = tk.client.Do("DEL", refreshUuid)
	if err != nil {
		return err
	}

	return nil
}

func (tk *ClientData) DeleteRefresh(refreshUuid string) error {
	//delete refresh token

	deleted, err := tk.client.Do("DEL", refreshUuid)
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
