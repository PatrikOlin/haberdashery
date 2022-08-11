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
	IsOrphan 	bool		`json:"is_orphan" db:"is_orphan"`
}

func Ptr[T any](v T) *T {
	return &v
}

func GetAllGarments(c *fiber.Ctx) error {
	includeOrphans := c.Query("includeOrphans") != ""
	var err error
	var garments []Garment

	garments, err = getGarments(c, includeOrphans)

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
		"data":    garments,
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
		IsOrphan:    true,
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

func UpdateGarment(c *fiber.Ctx) error {
	g := new(Garment)

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Could not read ID",
			"error":   err,
			"data":    nil,
		})
	}

	g.ID = id
	g.IsOrphan = false

	if err := c.BodyParser(g); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Could not parse body",
			"error":   err,
			"data":    nil,
		})
	}

	stmt := `UPDATE garments SET 
			color=:color, purchase_price=:purchase_price,
			purchased_at=:purchased_at, last_worn=:last_worn,
			last_washed=:last_washed, times_worn=:times_worn,
			times_washed=:times_washed, brand=:brand,
			is_orphan=:is_orphan
			WHERE id = :id`

	rows, err := db.DBClient.NamedQuery(stmt, g)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Could not update db",
			"error":   err,
			"data":    nil,
		})
	}

	if rows.Next() {
		rows.Scan(&g)
	}
	
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data": g,
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

func getGarments(c *fiber.Ctx, includeOrphans bool) ([]Garment, error) {
	var garments []Garment
	var stmt string
	if includeOrphans == true {
		stmt = `SELECT * FROM garments g ORDER BY g.last_worn DESC`
	} else {
		stmt = `SELECT * FROM garments g WHERE g.is_orphan = false ORDER BY g.last_worn DESC`
	}

	err := db.DBClient.Select(&garments, stmt)

	if err != nil {
		return garments, err
	}

	return garments, nil
}

func persistGarment(g Garment) (Garment, error) {
	var garment Garment
	stmt := `INSERT INTO garments (
			id, color, purchase_price,
			purchased_at, last_worn, last_washed,
			times_worn, times_washed, brand,
			image_url, added_at, is_orphan)
			VALUES (
			:id, :color, :purchase_price,
			:purchased_at, :last_worn, :last_washed,
			:times_worn, :times_washed, :brand,
			:image_url, :added_at, :is_orphan)
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
