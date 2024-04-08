package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	sqlc "github.com/iKayrat/test/internal/app/db/sqlc"

	_ "github.com/lib/pq"
)

type Control struct {
	Ctx      context.Context
	Store    sqlc.Store
	Response []Response
}

type Response struct {
	Shelf           string `json:"shelf"`
	Name            string `json:"name"`
	AdditionalShelf string `json:"additional_shelf"`
	Id              int    `json:"id"`
	OrderNum        int    `json:"order_num"`
	Quantity        int    `json:"quantity"`
}

func New(ctx context.Context, store sqlc.Store) ControllerI {

	return &Control{
		Ctx:   ctx,
		Store: store,
	}
}

func (c *Control) GetPages(args []string) ([]Response, error) {
	splitted := validate(args)
	if splitted == nil {
		return nil, fmt.Errorf("nil slice: %v", splitted)
	}

	// m := make(map[string]Response, 0)``
	s := make([]Response, 0)

	for _, a := range splitted {

		get, err := c.Store.GetOrderByGrouping(c.Ctx, a)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}

		for _, g := range get {
			orderNum, _ := strconv.Atoi(g.OrderNumber)

			str, err := inToString(g.AdditionalShelves)
			if err != nil {
				log.Println(err)
			}

			// m[g.PrimaryShelf] = Response{}
			// s = append(s, m[g.PrimaryShelf])
			s = append(s, Response{
				Shelf:           g.PrimaryShelf,
				Name:            g.Title,
				Id:              int(g.ProductID.Int32),
				OrderNum:        orderNum,
				Quantity:        int(g.Quantity),
				AdditionalShelf: str,
			})
		}
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].OrderNum < s[j].OrderNum
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i].Shelf < s[j].Shelf
	})

	c.Response = s

	return s, nil
}

func (c *Control) BeautyPrint(args []string) {
	fmt.Println("=+=+=+=")

	fmt.Printf("Страница сборки заказов %s\n\n", strings.Join(args, ","))

	fmt.Printf("===Стеллаж %s\n", c.Response[0].Shelf)
	current := c.Response[0].Shelf

	for _, items := range c.Response {

		if items.Shelf != current {
			fmt.Printf("===Стеллаж %s\n", items.Shelf)
		}

		current = items.Shelf

		fmt.Printf("%s (id=%d)\nзаказ %d,%d шт\n", items.Name, items.Id, items.OrderNum, items.Quantity)
		if items.AdditionalShelf != "" {
			fmt.Printf("доп стеллаж: %s\n", items.AdditionalShelf)
		}

		fmt.Println()
	}
}

func validate(args []string) []string {

	splitted := strings.Split(args[0], ",")

	for _, a := range splitted {
		num, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			log.Println("не валидная страница сборки заказов: ", num)
			return nil
		}
	}
	return splitted
}

func inToString(in interface{}) (string, error) {

	if str, ok := in.(string); ok {
		return str, nil
	}
	return "", fmt.Errorf("value is not a string: %v", in)
}
