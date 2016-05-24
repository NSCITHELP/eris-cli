package errno

import (
	"errors"
	"fmt"
	"runtime"
)

// ------------------ base error framework ------------------------
type ErisError struct {
	Code int
	ErrMsg error
	FixMsg string
}

// Take a string error defined in this file & concates with thrown error
// TODO make BaseErrorEE
func BaseError(errMsg string, thrownError error) error {
	return fmt.Errorf(errMsg, thrownError)
}

// takes two strings & returns an error
func BaseErrorES(errMsg, thing string) error {
	return fmt.Errorf(errMsg, thing)
}

// takes an erro and two strings & returns an error
func BaseErrorESE(errMsg, thing string, err error) error {
	return errors.New(fmt.Sprintf(errMsg, thing, err))
}

func (e *ErisError) Error() string {
	return fmt.Sprintf("error code %d/nerror: %v/Try fixing it with: %s/n", e.Code, e.ErrMsg, e.FixMsg)
}
// ----------------------------------------------------------------

// ------------------ somewhat general errors ---------------------
var (
	ErrNonInteractiveExec    = errors.New("Non-interactive exec sessions must provide arguments to execute")
	ErrRenaming              = errors.New("Cannot rename to same name")
	ErrCannotFindService     = errors.New("I cannot find that service. Please check the service name you sent me. To find known services use [eris services ls --known]")
	ErrCantFindChain         = fmt.Errorf("I don't known of that chain.\nPlease retry with a known chain.\nTo find known chains use: eris chains ls --known")
	ErrCantFindData          = errors.New("I cannot find that data container. Please check the data container name you sent me with [eris data ls]")
	ErrCantFindAction        = errors.New("The marmots could not find the action definition file. Please check your actions with [eris actions ls]")
	ErrContainerNameNotFound = errors.New("container name not found")

	ErrNoChainName       = errors.New("Cannot start a chain without a name") // TODO generalize this error message
	ErrInvalidPkgJSON    = "Sorry, the marmots could not figure that package.json out.\nPlease check your package.json is properly formatted:"
	ErrNoChainSpecified  = "Marmot disapproval face.\nYou tried to start a service which has a `$chain` variable but didn't give us a chain.\nPlease rerun the command either after [eris chains checkout CHAINNAME] *or* with a --chain flag.\n"
	ErrServiceNotRunning = "The requested service is not running, start it with [eris services start <serviceName>]"

	ErrNeedChainCheckedOut = errors.New("A checked out chain is needed to continue. Please check out the appropriate chain or rerun with a chain flag")
	ErrContainerExists     = errors.New("container exists")
	ErrImageNotExist       = errors.New("Image does not exist. Something went wrong. Exiting")
	ErrNoFile              = errors.New("Cannot find that file. Sorry.")
	ErrNoPermGiven         = errors.New("No permission given. Exiting.")
	ErrNeedGitAndGo        = errors.New("either git or go is not installed. both are required for non-binary update")
	ErrMarmotWTF           = errors.New("The marmots could not figure out how eris was installed. Exiting.")
	ErrNotLetMePull        = errors.New("Cannot start a container based on an image you will not let me pull")
	ErrMergeParameters     = errors.New("parameters are not pointers to struct")
	
	ErrCreatingDataCont = "error creating data container:%v\n"
	ErrPermissionNotGiven = "Permission to %s denied. Exiting."
	ErrPathIsNotDirectory = "path (%s) is not a directory; please provide a path to a directory"
	ErrPathDoesNotExist = "path (%s) does not exist; please rerun command with a proper path"
	ErrBadConfigOptions = "Config options should be <key>=<value> pairs. Got %s"
	ErrUnknownCatCmd = "unknown cat subcommand: %s"

	ErrRunningCommand = "error running command (%s): %v"
	ErrMakingDirectory = "The marmots could neither find, nor had access to make the directory: (%s)\nerror: %v\n"
	ErrBadTemplate = "%stemplate error: %v\n"
	ErrWritingFile = "Cannot add default %s: %v\n"

	ErrMigratingDirs = "error migrating directories: %v\n"
	ErrInitErisRoot ="error initializing the Eris root directory: %v\n"
	ErrInitDefaults = "error instantiating default files: %v\n"
	ErrDropDefaults = "error dropping default files:%v\ntoadserver may be down: re-run [eris init] with [--source=rawgit]"

	// -------- files --------------
	BadGatewayURL = "Invalid gateway URL provided %v\n"
	ErrEnsureRunningIPFS = "Failed to ensure IPFS is running: %v"
	ErrNoFileToExport = errors.New("error: no file to export")
	WarnAllOrNothing = errors.New("Either remove a file by hash or all of them.")

	ErrStartingService = "error starting service: %v\n"
	ErrNoServiceGiven = errors.New("no service given")

	ErrWritingDefinitionFile = "error writing definition file: %v\n"

	// -------- chains -------------
	ErrReadingGenesisFile = "error reading genesis file: %v\n"
	ErrStartingChain = "error starting chain: %v\n"
	ErrReadingFromGenesisFile = "error reading %s genesis file: %v\n"
	ErrExecChain = "error %s: %v\n"

	// -------- util ---------------
	ErrListingContainers = "error listing containers: %v\n"
	ErrRemovingContainer = "error removing container: %v\n"

	ErrBadReport = "The marmots had an error trying to print a nice report: %v\n"

	ParseIPFShost = "parse the URL"
	SplitHP       = "split the host and port"
	ErrConnectDockerTLS = "Failed to connect to Docker Backend via TLS.\nerror:%v\n"
	ErrStartingDockerMachine = "There was an error starting the newly created docker-machine.\nerror:%v\n"
	ErrDockerWindows = "Could not add ssh.exe to PATH.\nerror:%v\n"
	ErrConnectDockerMachine = "Could not evaluate the env vars for the %s docker-machine.\nerror:%v\n"
	ErrParseIPFS = "The marmots could not %s for the DockerHost to populate the IPFS Host.\nPlease check that your docker-machine VM is running with [docker-machine ls]\nerror: %v\n"

)
// ----------------------------------------------------------------

func ErrContainerExit(id string, code int) error {
	return fmt.Errorf("Container %s exited with status %d", id, code)

}

func ErrRunningArguments(args []string, err error) error {
	return fmt.Errorf("error running args: %v\n%v\n", args, err)
}

func ErrRemovingDataCont(err1, err2 error) error {
	return fmt.Errorf("error removing data container after executing (%v): %v\n", err1, err2)
}

func ErrChainMissing(ch1, ch2 string) error {
	return fmt.Errorf("chain %s depends on chain %s but %s is not running", ch1, ch2, ch2)
}

// only a warning
func ErrSettingUpChain(err error) string {
	return fmt.Sprintf("error setting up chain: %v\nCleaning up...", err)
}

func ErrCleaningUpChain(contName string, err1, err2 error) error {
	return fmt.Errorf("Tragic! Our marmots encountered an error during setupChain for %s.\nThey also failed to cleanup after themselves (remove containers) due to another error.\nFirst error: %v\nCleanup error: %v\n", contName, err1, err2)

}

func ErrWrongLength(thing string, length int) error {
	return fmt.Errorf("%s length !=%d", thing, length)
}

// --- from cmd/eris ----
func ErrBadCommandLength(typ, numStr string) error {
	return fmt.Errorf("**Note** you sent our marmots the wrong number of %s.\nPlease send the marmots %s", typ, numStr)
}


func ErrNoDirectories(path1, path2 string) error {
	return fmt.Errorf("neither deprecated (%s) or new (%s) exists. please run `init` prior to `update`\n", path1, path2)
}

func ErrCheckKeysAndCerts(thing, file string, err error) error {
	return fmt.Errorf("The marmots could not find a file that was required to connect to Docker. %s\n%s\nFile needed: %s\nerror:", thing, file, err)
}

func MustInstallDockerError() error {
	install := `The marmots cannot connect to Docker. Do you have Docker installed?
If not, please visit here: https://docs.docker.com/installation/`

	switch runtime.GOOS {
	case "linux":
		run := `Do you have Docker running? If not, please type [sudo service docker start].
Also check that your user is in the "docker" group. If not, you can add it
using the [sudo usermod -a -G docker $USER] command or rerun as [sudo eris]`

		return fmt.Errorf("%slinux/\n\n%s", install, run)
	case "darwin":
		return fmt.Errorf("%smac/", install)
	case "windows":
		return fmt.Errorf("%swindows/", install)
	}
	return fmt.Errorf(install)
}

var ErrConnectDockerDaemon = "There was an error connecting to your Docker daemon.\nCome back after you have resolved the issue and the marmots will be happy to service your blockchain management needs: %v"

func ErrBadWhaleVersions(thing, verMin, verDetected string) error {
	return fmt.Errorf("Eris requires %s version >= %s\nThe marmots have detected docker version: %s\nCome back after you have upgraded and the marmots will be happy to service your blockchain management needs", thing, verMin, verDetected)

}

func ErrRmDataContainer(err, err2 error) error {
	return fmt.Errorf("Tragic! Error removing data container after executing (%v): %v", err, err2)
}

func ErrLoadViperConfig(typ string, err error) error {
	return fmt.Errorf("Check your known %ss with [eris %ss ls --known]\nThere may also be an error with the formatting of the .toml file:\n%v", typ, typ, err)
}

func ErrLoadingDefFile(typ string) error {
	return fmt.Errorf("error loading %s definition file:", typ)
}
