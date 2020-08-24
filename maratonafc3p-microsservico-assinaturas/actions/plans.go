package actions

import (
	"fmt"
	"net/http"
	"subscription_service/models"

	"github.com/gobuffalo/pop/v5"

	"github.com/gobuffalo/buffalo"
)

// PlansIndex default implementation.
func PlansIndex(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)

	plans := models.Plans{}
	err := tx.All(&plans)

	if err != nil {
		fmt.Print("ERROR!\n")
		fmt.Printf("%v\n", err)
	}

	c.Set("plans", plans)
	return c.Render(http.StatusOK, r.HTML("plans/index.html"))
}
