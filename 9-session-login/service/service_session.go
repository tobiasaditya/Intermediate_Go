package service

import (
	"fmt"
	"log"
	"os"

	"9-session-login/config"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
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
		return err
	}
	return nil
}

func (s *SessionService) GetSession(c echo.Context, sessionID string) (*sessions.Session, error) {
	session, err := s.store.Get(c.Request(), sessionID)
	if err != nil {
		return nil, err
	}

	if len(session.Values) == 0 {
		return nil, fmt.Errorf("empty session")
	}

	return session, nil
}

func (s *SessionService) DeleteSession(c echo.Context, sessionID string) error {
	session, err := s.GetSession(c, sessionID)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1 //forced to be expired
	session.Save(c.Request(), c.Response())

	return nil
}
