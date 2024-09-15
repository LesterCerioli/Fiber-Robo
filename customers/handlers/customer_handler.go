package hendlers

import(
	"github.com/gofiber/fiber/v2"
    "project-root/customers/repository"
    "project-root/customers/models"
)

func CreateCustomer(c *fiber.Ctx) error {
    customer := new(models.Customer)

    if err := c.BodyParser(customer); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := repository.CreateCustomer(customer); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create customer",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(customer)

}