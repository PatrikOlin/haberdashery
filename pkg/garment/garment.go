package garment

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PatrikOlin/haberdashery/pkg/db"
	"github.com/gofiber/fiber/v2"
)

type Garment struct {
	ID          int64      `json:"id" db:"id"`
	Color       *string    `json:"color" db:"color"`
	Price       *int       `json:"price" db:"purchase_price"`
	PurchasedAt *time.Time `json:"purchased_at" db:"purchased_at"`
	LastWorn    *time.Time `json:"last_worn" db:"last_worn"`
	LastWashed  *time.Time `json:"last_washed" db:"last_washed"`
	TimesWorn   *int       `json:"times_worn" db:"times_worn"`
	TimesWashed *int       `json:"times_washed" db:"times_washed"`
	Brand       *string    `json:"brand" db:"brand"`
	ImageURL    *string    `json:"image_url" db:"image_url"`
	AddedAt     *time.Time `json:"added_at" db:"added_at"`
}

func Ptr[T any](v T) *T {
	return &v
}

func GetAllGarments(c *fiber.Ctx) error {
	var shirts []Garment
	stmt := `SELECT * FROM garments g ORDER BY g.last_worn DESC`

	err := db.DBClient.Select(&shirts, stmt)

	fmt.Println(err)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to fetch garments",
			"error":   err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    shirts,
	})
}

func CreateGarment(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Could not read ID",
			"error":   err,
			"data":    nil,
		})
	}

	g := Garment{
		ID:          id,
		Color:       Ptr(""),
		Price:       Ptr(0),
		PurchasedAt: Ptr(time.Now()),
		LastWorn:    Ptr(time.Now()),
		LastWashed:  Ptr(time.Now()),
		TimesWorn:   Ptr(0),
		TimesWashed: Ptr(0),
		Brand:       Ptr(""),
		ImageURL:    Ptr(""),
		AddedAt:     Ptr(time.Now()),
	}

	garment, err := persistGarment(g)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Could not create garment",
			"error":   err,
			"data":    nil,
		})
	}


	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    garment,
	})
}

func IncrementGarment(c *fiber.Ctx) error {
	freshlyWashed := c.Query("washed")
	id := fmt.Sprintf("%s", c.Params("id"))
	var stmt string
	var results []Garment
	t := time.Now().Local().Format("2006-01-02 15:04:05")

	if freshlyWashed == "" {
		stmt = getWornStmt(id)
	} else {
		stmt = getWashedStmt(id)
	}

	err := db.DBClient.Select(&results, stmt, t, id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"error":   err,
			"data":    nil,
		})
	}

	if results == nil {
		return CreateGarment(c)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    results,
	})
}

func persistGarment(g Garment) (Garment, error) {
	var garment Garment
	stmt := `INSERT INTO garments (
			id, color, purchase_price,
			purchased_at, last_worn, last_washed,
			times_worn, times_washed, brand,
			image_url, added_at)
			VALUES (
			:id, :color, :purchase_price,
			:purchased_at, :last_worn, :last_washed,
			:times_worn, :times_washed, :brand,
			:image_url, :added_at)
			RETURNING *`

	rows, err := db.DBClient.NamedQuery(stmt, g)

	if err != nil {
		return garment, err
	}

	if rows.Next() {
		rows.Scan(&garment)
	}

	return garment, nil
}

func getWornStmt(id string) string {
	return `UPDATE garments
			SET times_worn = times_worn + 1, last_worn = $1
			WHERE id = $2
			RETURNING *`
}

func getWashedStmt(id string) string {
	return `UPDATE garments
			SET times_washed = times_washed + 1, last_washed = $1
			WHERE id = $2
			RETURNING *`
}
