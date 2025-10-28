package patterns

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Name string
}

func fetch(_ context.Context, user User) (string, error) {
	time.Sleep(10 * time.Millisecond)

	return user.Name, nil
}

func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0)

	for _, u := range users {
		name, err := fetch(ctx, u)
		if err != nil {
		}

		names[name] = names[name] + 1
	}

	return names, nil
}

// 1. Распараллелить запросы
// 2. Прерывать запросы при возникновении ошибки хотя бы в одном
func main() {
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background()

	start := time.Now()
	res, err := process(ctx, names)
	if err != nil {
		fmt.Println("an error occurred:", err.Error())
	}
	fmt.Println("time:", time.Since(start))
	fmt.Println(res)
}
