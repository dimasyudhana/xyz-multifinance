package migration

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/dimasyudhana/xyz-multifinance/app/config"
)

func Migrate(cfg config.AppConfig, arg string) {

	var migratePath string

	if len(cfg.AppPath) > 0 {
		migratePath = fmt.Sprintf("%s/migration/", cfg.AppPath)
	} else {
		migratePath = "migration/"
	}

	rawCMD := "echo \"y\" | migrate -database \"mysql://%s:%s@tcp(%s:%d)/%s\" -path %s %s"
	command := fmt.Sprintf(rawCMD,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Hostname,
		cfg.DB.Port,
		cfg.DB.Name,
		migratePath,
		arg)

	fmt.Println(command)

	cmd := exec.Command("bash", "-c", command)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, _ := cmd.CombinedOutput()
	log.Printf("%s", out)
}
