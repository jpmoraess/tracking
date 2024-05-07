package main

import (
	"context"
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/tracking/configs"
	"github.com/jpmoraess/tracking/internal/app/usecase"
	"github.com/jpmoraess/tracking/internal/infra/handlers"
	"github.com/jpmoraess/tracking/internal/infra/persistence"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "The listen address of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.DBURI))
	if err != nil {
		panic(err)
	}

	var (
		// initialize repositories
		customerRepository            = persistence.NewCustomerMongo(client)
		customerSystemRepository      = persistence.NewCustomerSystemMongo(client)
		shippingCompanyRepository     = persistence.NewShippingCompanyMongo(client)
		customerIntegrationRepository = persistence.NewCustomerIntegrationMongo(client)

		// initialize usecases
		createCustomer            = usecase.NewCreateCustomer(customerRepository)
		createCustomerSystem      = usecase.NewCreateCustomerSystem(customerSystemRepository)
		createShippingCompany     = usecase.NewCreateShippingCompany(shippingCompanyRepository)
		createCustomerIntegration = usecase.NewCreateCustomerIntegration(customerRepository, customerSystemRepository, shippingCompanyRepository, customerIntegrationRepository)

		// initialize handlers
		customerHandler            = handlers.NewCustomerHandler(*createCustomer)
		customerSystemHandler      = handlers.NewCustomerSystemHandler(*createCustomerSystem)
		shippingCompanyHandler     = handlers.NewShippingCompanyHandler(*createShippingCompany)
		customerIntegrationHandler = handlers.NewCustomerIntegrationHandler(*createCustomerIntegration)

		// initialize fiber
		app   = fiber.New()
		apiV1 = app.Group("/api/v1")
	)

	apiV1.Post("/customer", customerHandler.HandlePostCustomer)

	apiV1.Post("/customer-system", customerSystemHandler.HandlePostCustomerSystem)

	apiV1.Post("/shipping-company", shippingCompanyHandler.HandlePostShippingCompany)

	apiV1.Post("/customer-integration", customerIntegrationHandler.HandlePostCustomerIntegration)

	app.Listen(*listenAddr)
}
