package errno

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrorNonInteractiveExec    = errors.New("Non-interactive exec sessions must provide arguments to execute")
	ErrorRenaming              = errors.New("Cannot rename to same name")
	ErrorCannotFindService     = errors.New("I cannot find that service. Please check the service name you sent me. To find known services use [eris services ls --known]")
	ErrorCantFindChain         = fmt.Errorf("I don't known of that chain.\nPlease retry with a known chain.\nTo find known chains use: eris chains ls --known")
	ErrorCantFindData          = errors.New("I cannot find that data container. Please check the data container name you sent me with [eris data ls]")
	ErrorCantFindAction        = errors.New("The marmots could not find the action definition file. Please check your actions with [eris actions ls]")
	ErrorContainerNameNotFound = errors.New("container name not found")

	ErrorNoChainName       = errors.New("Cannot start a chain without a name") // TODO generalize this error message
	ErrorInvalidPkgJSON    = "Sorry, the marmots could not figure that package.json out.\nPlease check your package.json is properly formatted:"
	ErrorNoChainSpecified  = "Marmot disapproval face.\nYou tried to start a service which has a `$chain` variable but didn't give us a chain.\nPlease rerun the command either after [eris chains checkout CHAINNAME] *or* with a --chain flag.\n"
	ErrorServiceNotRunning = "The requested service is not running, start it with [eris services start <serviceName>]"

	ErrorNeedChainCheckedOut = errors.New("A checked out chain is needed to continue. Please check out the appropriate chain or rerun with a chain flag")
	ErrorContainerExists     = errors.New("container exists")
	ErrorImageNotExist       = errors.New("Image does not exist. Something went wrong. Exiting")
	ErrorNoFile              = errors.New("Cannot find that file. Sorry.")
	ErrorNoPermGiven         = errors.New("No permission given. Exiting.")
	ErrorNeedGitAndGo        = errors.New("either git or go is not installed. both are required for non-binary update")
	ErrorMarmotWTF           = errors.New("The marmots could not figure out how eris was installed. Exiting.")
	ErrorNotLetMePull        = errors.New("Cannot start a container based on an image you will not let me pull")
	ErrorMergeParameters     = errors.New("parameters are not pointers to struct")

	tb = "This error was thrown by: "
)

func ErrorContainerExit(id string, code int) error {
	return fmt.Errorf("Container %s exited with status %d", id, code)

}
var (
	ErrorCreatingDataCont = "error creating data container:%v\n"
	ErrorPermissionNotGiven = "Permission to %s denied. Exiting."
	ErrorPathIsNotDirectory = "path (%s) is not a directory; please provide a path to a directory"
	ErrorPathDoesNotExist = "path (%s) does not exist; please rerun command with a proper path"
	ErrorBadConfigOptions = "Config options should be <key>=<value> pairs. Got %s"
	ErrorUnknownCatCmd = "unknown cat subcommand: %s"
)
	
var (
	ErrorRunningCommand = "error running command (%s): %v"
	ErrorMakingDirectory = "The marmots could neither find, nor had access to make the directory: (%s)\nerror: %v\n"
	ErrorBadTemplate = "%stemplate error: %v\n"
	ErrorWritingFile = "Cannot add default %s: %v\n"
)

func ErrorRunningArguments(args []string, err error) error {
	return fmt.Errorf("error running args: %v\n%v\n", args, err)
}

func ErrorRemovingDataCont(err1, err2 error) error {
	return fmt.Errorf("error removing data container after executing (%v): %v\n", err1, err2)
}

var (
ErrorMigratingDirs = "error migrating directories: %v\n"
ErrorInitErisRoot ="error initializing the Eris root directory: %v\n"
ErrorInitDefaults = "error instantiating default files: %v\n"
ErrorDropDefaults = "error dropping default files:%v\ntoadserver may be down: re-run [eris init] with [--source=rawgit]"
)

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

// -------- files --------------
var (
BadGatewayURL = "Invalid gateway URL provided %v\n"
ErrorEnsureRunningIPFS = "Failed to ensure IPFS is running: %v"
ErrorNoFileToExport = errors.New("error: no file to export")
WarnAllOrNothing = errors.New("Either remove a file by hash or all of them.")
)

// -------- chains -------------
var (
ErrorReadingGenesisFile = "error reading genesis file: %v\n"
ErrorStartingChain = "error starting chain: %v\n"
ErrorWriteChainFile = "error writing chain definition file: %v\n"
ErrorReadingFromGenesisFile = "error reading %s genesis file: %v\n"
ErrorExecChain = "error %s: %v\n"
)

func ErrorChainMissing(ch1, ch2 string) error {
	return fmt.Errorf("chain %s depends on chain %s but %s is not running", ch1, ch2, ch2)
}

// only a warning
func ErrorSettingUpChain(err error) string {
	return fmt.Sprintf("error setting up chain: %v\nCleaning up...", err)
}

func ErrorCleaningUpChain(contName string, err1, err2 error) error {
	return fmt.Errorf("Tragic! Our marmots encountered an error during setupChain for %s.\nThey also failed to cleanup after themselves (remove containers) due to another error.\nFirst error: %v\nCleanup error: %v\n", contName, err1, err2)

}

// -------- util ---------------
var (
ErrorListingContainers = "error listing containers: %v\n"
ErrorRemovingContainer = "error removing container: %v\n"
)

func ErrorWrongLength(thing string, length int) error {
	return fmt.Errorf("%s length !=%d", thing, length)
}

// --- from cmd/eris ----
func ErrorBadCommandLength(typ, numStr string) error {
	return fmt.Errorf("**Note** you sent our marmots the wrong number of %s.\nPlease send the marmots %s", typ, numStr)
}

// ---------------------

var ErrorBadReport = "The marmots had an error trying to print a nice report: %v\n"

func ErrorNoDirectories(path1, path2 string) error {
	return fmt.Errorf("neither deprecated (%s) or new (%s) exists. please run `init` prior to `update`\n", path1, path2)
}

var (
	ParseIPFShost = "parse the URL"
	SplitHP       = "split the host and port"
	ErrorConnectDockerTLS = "Failed to connect to Docker Backend via TLS.\nerror:%v\n"
	ErrorStartingDockerMachine = "There was an error starting the newly created docker-machine.\nerror:%v\n"
	ErrorDockerWindows = "Could not add ssh.exe to PATH.\nerror:%v\n"
	ErrorConnectDockerMachine = "Could not evaluate the env vars for the %s docker-machine.\nerror:%v\n"
	ErrorParseIPFS = "The marmots could not %s for the DockerHost to populate the IPFS Host.\nPlease check that your docker-machine VM is running with [docker-machine ls]\nerror: %v\n"
)

func ErrorCheckKeysAndCerts(thing, file string, err error) error {
	return fmt.Errorf("The marmots could not find a file that was required to connect to Docker. %s\n%s\nFile needed: %s\nerror:", thing, file, err)
}

func MustInstallDockerError() error {
	errBase := "The marmots cannot connect to Docker.\nDo you have docker installed?\nIf not please visit here:\t"
	dInst := "https://docs.docker.com/installation/"

	switch runtime.GOOS {
	case "linux":
		return fmt.Errorf("%s%s\nDo you have docker installed and running?\nIf not please [sudo service docker start] on Ubuntu.\nAlso check that your user is in the docker group (or rerun with sudo).\nTo fix this please run [sudo usermod -a -G docker $USER] on Ubuntu with your user substituted.", errBase, dInst)
	case "darwin":
		return fmt.Errorf("%s%s\n", errBase, (dInst + "mac/"))
	case "windows":
		return fmt.Errorf("%s%s\n", errBase, (dInst + "windows/"))
	}

	return fmt.Errorf("%s%s\n", errBase, dInst)
}

var ErrorConnectDockerDaemon = "There was an error connecting to your Docker daemon.\nCome back after you have resolved the issue and the marmots will be happy to service your blockchain management needs: %v"

func ErrorBadWhaleVersions(thing, verMin, verDetected string) error {
	return fmt.Errorf("Eris requires %s version >= %s\nThe marmots have detected docker version: %s\nCome back after you have upgraded and the marmots will be happy to service your blockchain management needs", thing, verMin, verDetected)

}

// -------- services ------------
type ServiceError struct {
	Command  string
	ErrMsg   error
	ThrownBy string
}

func ErrorRmDataContainer(err, err2 error) error {
	return fmt.Errorf("Tragic! Error removing data container after executing (%v): %v", err, err2)

}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("error %s-ing service: %v\n%s%s", e.Command, e.ErrMsg, tb, e.ThrownBy)
}

// --------- loaders ---------------
type InvalidLoadingError struct {
	TypeFile string
	ErrMsg   error
	ThrownBy string
}

func ErrorLoadViperConfig(typ string, err error) error {
	return fmt.Errorf("Check your known %ss with [eris %ss ls --known]\nThere may also be an error with the formatting of the .toml file:\n%v", typ, typ, err)
}

func ErrorLoadingDefFile(typ string) string {
	return fmt.Sprintf("error loading %s definition file:", typ)
}
func (e *InvalidLoadingError) Error() string {
	return fmt.Sprintf("%s\n%s\n%s%s", ErrorLoadingDefFile(e.TypeFile), e.ErrMsg, tb, e.ThrownBy)

}
