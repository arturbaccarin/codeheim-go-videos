package examplewithstruct

import (
	"log/slog"
	"os"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) LogValue() slog.Value {
	return slog.Int64Value(u.ID)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr,
		&slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))

	slog.SetDefault(logger)

	u := User{
		ID:       102312,
		Name:     "Darth Vader",
		Email:    "darth_vapder@example.com",
		Password: "galactic_empire",
	}

	slog.Info("User Info", "user", u)
}
