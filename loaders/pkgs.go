package loaders

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/eris-ltd/eris-cli/config"
	"github.com/eris-ltd/eris-cli/definitions"
	. "github.com/eris-ltd/eris-cli/errors"

	"github.com/spf13/viper"

	log "github.com/Sirupsen/logrus"
)

func LoadPackage(path, chainName string) (*definitions.Package, error) {
	var name string
	var dir bool

	var err error
	f, err := os.Stat(path)
	if err != nil {
		return nil, &ErisError{404, err, ""}
	}
	if f.IsDir() {
		name = filepath.Base(path)
		dir = true
	} else {
		name = filepath.Base(filepath.Dir(path))
		dir = false
	}

	var pkgConf *viper.Viper
	var pkg *definitions.Package
	if dir {
		pkgConf, err = loadPackage(path)
	} else {
		pkgConf, err = loadPackage(filepath.Dir(path))
	}

	if err != nil {
		log.Info("The marmots could not read that package.json. Will use defaults.")
		pkg = DefaultPackage(name, chainName)
	} else {
		// marshal chain and always reset the operational requirements
		// this will make sure to sync with docker so that if changes
		// have occured in the interim they are caught.
		pkg, err = marshalPackage(pkgConf)
		if err != nil {
			return nil, &ErisError{404, err, ""}
		}
	}

	if err := checkName(pkg, chainName); err != nil {
		return nil, &ErisError{404, err, ""}
	}

	return pkg, nil
}

// read the config file into viper
func loadPackage(path string) (*viper.Viper, error) {
	return config.LoadViperConfig(path, "package", "pkg")
}

// set's the defaults
func DefaultPackage(name, chainName string) *definitions.Package {
	pkg := definitions.BlankPackage()
	pkg.Name = name
	pkg.ChainName = chainName
	pkg.PackageID = "" // TODO hash it
	return pkg
}

func marshalPackage(pkgConf *viper.Viper) (*definitions.Package, error) {
	pkgDef := definitions.BlankPackageDefinition()
	if err := pkgConf.Unmarshal(pkgDef); err != nil {
		return nil, fmt.Errorf("%s\n%v", ErrInvalidPkgJSON, err)
	}
	pkg := pkgDef.Package
	if pkgDef.Name != "" {
		pkg.Name = pkgDef.Name
	}

	return pkg, nil
}

func checkName(pkg *definitions.Package, name string) error {
	if strings.Contains(pkg.Name, " ") {
		newName := strings.Replace(pkg.Name, " ", "_", -1)
		log.WithFields(log.Fields{
			"old": pkg.Name,
			"new": newName,
		}).Debug("Correcting package name.")
		pkg.Name = newName
	}

	return nil
}
