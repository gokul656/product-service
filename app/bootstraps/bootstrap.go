package bootstraps

import "github.com/gokul656/product-service/app/configs"

func Setup() {
	configs.MigrationIntializer()
}
