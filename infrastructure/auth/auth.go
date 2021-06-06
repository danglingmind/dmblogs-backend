package auth

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type AuthInterface interface {
	CreateAuth(uuid.UUID, *TokenDetails) error
	FetchAuth(string) (uuid.UUID, error)
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
	UserId    uuid.UUID
}

//Save token metadata to Redis
func (tk *ClientData) CreateAuth(userid uuid.UUID, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	_, err := tk.client.Do("SET", td.TokenUuid, userid)
	if err != nil {
		return err
	}
	_, err = tk.client.Do("EXPIRE", td.TokenUuid, int(at.Sub(now).Seconds()))
	if err != nil {
		return err
	}

	_, err = tk.client.Do("SET", td.RefreshUuid, userid)
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
func (tk *ClientData) FetchAuth(tokenUuid string) (userUUID uuid.UUID, err error) {
	userid, err := redis.Bytes(tk.client.Do("GET", tokenUuid))
	if err != nil {
		return
	}
	userUUID, err = uuid.FromBytes(userid)
	return
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
