package account

import (
	"path/filepath"
)

const rootDir = "/media/mahdi/Cloud/Happle"
const serviceDir = "services"

func GetServicesPath(file string) string {
	return filepath.Join(rootDir, serviceDir, file)
}
