package service

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"9-session-login/config"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo"
)

type SessionService struct {
	store *pgstore.PGStore
}

func NewSessionService() *SessionService {
	store, err := pgstore.NewPGStore(config.DB_URI, []byte(config.AUTH_KEY), []byte(config.ENCRYPTION_KEY))
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}
	return &SessionService{store: store}
}

func (s *SessionService) SetSession(c echo.Context, sessionID string, username string) error {
	session, err := s.store.Get(c.Request(), sessionID)
	if err != nil {
		return err
	}

	session.Values["username"] = username

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (s *SessionService) GetSession(c echo.Context, sessionID string) (string, error) {
	session, err := s.store.Get(c.Request(), sessionID)
	if err != nil {
		return "", err
	}

	if len(session.Values) == 0 {
		return "", fmt.Errorf("empty session")
	}

	data := session.Values["username"]
	return fmt.Sprintf("%v", data), nil
}
