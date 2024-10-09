package tests

import (
	Controllers "StoryPlatforn_GIN/internal/app/controller"
	"StoryPlatforn_GIN/internal/app/repository"
	"StoryPlatforn_GIN/internal/app/service"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

var (
	dbURI, dbName string
)

func init() {
	dbURI = os.Getenv("DATABASE_URI")
	dbName = os.Getenv("DB_NAME")
}

type APITestSuite struct {
	suite.Suite

	db           *pgx.Conn
	controllers  *Controllers.Controller
	services     *service.Service
	repositories *repository.Repository
}

func TestApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(APITestSuite))
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func (s *APITestSuite) SetupSuite() {
	if dbURI == "" {
		s.FailNow("DATABASE_URI is not set")
	}
	if dbName == "" {
		s.FailNow("DB_NAME is not set")
	}

	conn, err := pgx.Connect(context.Background(), dbURI)
	if err != nil {
		s.FailNow("failed to connect to database: %v", err)
	} else {
		s.db = conn
	}

	s.initDeps()

	err = s.populateDB()
	if err != nil {
		s.FailNow("failed to populate database: %v", err)
	}
}

func (s *APITestSuite) TearDownSuite() {
	err := s.db.Close(context.Background())
	if err != nil {
		s.FailNow("failed to close conn to database: %v", err)
	}
}

func (s *APITestSuite) initDeps() {
	s.repositories = repository.New(s.db)
	s.services = service.New(s.repositories)
	s.controllers = Controllers.New(s.services)
}

func (s *APITestSuite) populateDB() error {
	for _, user := range users {
		_, err := s.db.Exec(context.Background(), "INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, $4)", user.ID, user.Username, user.Password, user.CreatedAt)
		if err != nil {
			return err
		}
	}

	for _, session := range sessions {
		_, err := s.db.Exec(context.Background(), "INSERT INTO sessions (user_id, session, created_at) VALUES ($1, $2, $3)", session.UserID, session.SessionID, session.CreatedAt)
		if err != nil {
			return err
		}
	}

	for _, story := range stories {
		_, err := s.db.Exec(context.Background(), "INSERT INTO stories (id, user_id, author, title, text, rating, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", story.ID, story.UserID, story.Author, story.Title, story.Text, story.Rating, story.CreatedAt)
		if err != nil {
			return err
		}
	}

	for _, rate := range rates {
		_, err := s.db.Exec(context.Background(), "INSERT INTO story_ratings (user_id, story_id) VALUES ($1, $2)", rate.UserID, rate.StoryID)
		if err != nil {
			return err
		}
	}

	return nil
}
