package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/spf13/viper"

	"github.com/abdfnx/gosh"
	"github.com/julienschmidt/httprouter"
)

type Result struct {
	Message string `json:"message"`
}

const (
	baseLinux   = "systemctl"
	baseWindows = "shutdown -f"
)

var (
	shutdownFunc func(command string) error
	options      = []string{"suspend", "poweroff", "hibernate", "reboot"}
)

func init() {
	switch runtime.GOOS {
	case "windows":
		shutdownFunc = shutdownWindows
	case "linux":
		fmt.Println("###############\nPlease be sure that this script has sudo priviledges in order to run commands from this script\n################")
		shutdownFunc = shutdownLinux
	default:
		panic("Your os is not yet supported")
	}
}

func StartServer() {
	port := viper.GetString("port")
	router := httprouter.New()
	router.POST("/:command", handleCommand)
	fmt.Println("Port is " + port)
	http.ListenAndServe(":"+port, router)
}

func handleCommand(w http.ResponseWriter, r *http.Request, prms httprouter.Params) {
	command := prms.ByName("command")
	fmt.Println("Command is " + command)
	enc := json.NewEncoder(w)
	if !stringInSlice(command, options) {
		handleError(w, errors.New("Option not available"), enc, http.StatusMethodNotAllowed)
		return
	}
	err := shutdownFunc(command)
	if err != nil {
		handleError(w, err, enc, http.StatusInternalServerError)
		return
	}
}

func handleError(w http.ResponseWriter, err error, enc *json.Encoder, code int) {
	w.WriteHeader(code)
	res := Result{Message: err.Error()}

	enc.Encode(&res)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		fmt.Println("Current command is " + b)
		if b == a {
			return true
		}
	}
	return false
}

func shutdownLinux(command string) error {
	var commandEx *exec.Cmd
	switch command {
	case "suspend":
		commandEx = exec.Command(baseLinux, "suspend")
	case "poweroff":
		commandEx = exec.Command(baseLinux, "poweroff")
	case "hibernate":
		commandEx = exec.Command(baseLinux, "hibernate")
	case "reboot": //Really?
		commandEx = exec.Command(baseLinux, "reboot")
	}
	return commandEx.Run()
}

func shutdownWindows(command string) error {

	switch command {
	case "suspend":
		gosh.Run("rundll32 powrprof.dll,SetSuspendState 0,1,0")
	case "poweroff":
		gosh.Run(baseWindows + " -s")
	case "hibernate":
		gosh.Run(baseWindows + " -h")
	case "reboot": //Really?
		gosh.Run(baseWindows + " -r")
	}
	return nil
}
