package cmds

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
)

var (
	//
	// DQL related flags.
	//
	fsDQLName  = "dql"
	fsDQL      = pflag.NewFlagSet(fsDQLName, pflag.ContinueOnError)
	fsDQLUsage = func() {
		fmt.Printf("usage: datakit dql [options]\n\n")
		fmt.Printf("DQL used to query data from DataFlux. If no option specified, query interactively. Other available options:\n\n")
		fmt.Println(fsDQL.FlagUsagesWrapped(0))
	}

	flagDQLJSON        = fsDQL.BoolP("json", "J", false, "output in json format")
	flagDQLAutoJSON    = fsDQL.Bool("auto-json", false, "pretty output string if field/tag value is JSON")
	flagDQLVerbose     = fsDQL.BoolP("verbose", "V", false, "verbosity mode")
	flagDQLString      = fsDQL.StringP("run", "R", "", "run single DQL")
	flagDQLToken       = fsDQL.StringP("token", "T", "", "run query for specific token(workspace)")
	flagDQLCSV         = fsDQL.String("csv", "", "Specify the directory")
	flagDQLForce       = fsDQL.BoolP("force", "F", false, "overwrite csv if file exists")
	flagDQLDataKitHost = fsDQL.StringP("host", "H", "", "specify datakit host to query")
	flagDQLLogPath     = fsDQL.String("log", commonLogFlag(), "command line log path")

	//
	// running mode. (not used).
	//
	fsRunName          = "run"
	fsRun              = pflag.NewFlagSet(fsRunName, pflag.ContinueOnError)
	FlagRunInContainer = fsRun.BoolP("container", "c", false, "running in container mode")
	// flagRunLogPath     = fsRun.String("log", commonLogFlag(), "command line log path").
	fsRunUsage = func() {
		fmt.Printf("usage: datakit run [options]\n\n")
		fmt.Printf("Run used to select different datakit running mode.\n\n")
		fmt.Println(fsRun.FlagUsagesWrapped(0))
	}

	//
	// pipeline related flags.
	//
	fsPLName          = "pipeline"
	debugPipelineName = ""
	fsPL              = pflag.NewFlagSet(fsPLName, pflag.ContinueOnError)
	flagPLLogPath     = fsPL.String("log", commonLogFlag(), "command line log path")
	flagPLTxtData     = fsPL.StringP("txt", "T", "", "text string for the pipeline or grok(json or raw text)")
	flagPLTxtFile     = fsPL.StringP("file", "F", "", "text file path for the pipeline or grok(json or raw text)")
	flagPLTable       = fsPL.Bool("tab", false, "output result in table format")
	flagPLDate        = fsPL.Bool("date", false, "append date display(according to local timezone) on timestamp")
	// flagPLGrokQ       = fsPL.BoolP("grokq", "G", false, "query groks interactively").
	fsPLUsage = func() {
		fmt.Printf("usage: datakit pipeline [pipeline-script-name.p] [options]\n\n")
		fmt.Printf("Pipeline used to debug exists pipeline script.\n\n")
		fmt.Println(fsPL.FlagUsagesWrapped(0))
	}

	//
	// version related flags.
	//
	fsVersionName                    = "version"
	fsVersion                        = pflag.NewFlagSet(fsVersionName, pflag.ContinueOnError)
	flagVersionLogPath               = fsVersion.String("log", commonLogFlag(), "command line log path")
	flagVersionDisableUpgradeInfo    = fsVersion.Bool("upgrade-info-off", false, "do not show upgrade info")
	flagVersionUpgradeTestingVersion = fsVersion.BoolP("testing", "T", false, "show testing version upgrade info")
	fsVersionUsage                   = func() {
		fmt.Printf("usage: datakit version [options]\n\n")
		fmt.Printf("Version used to handle version related functions.\n\n")
		fmt.Println(fsVersion.FlagUsagesWrapped(0))
	}

	//
	// service management related flags.
	//
	fsServiceName        = "service"
	fsService            = pflag.NewFlagSet(fsServiceName, pflag.ContinueOnError)
	flagServiceLogPath   = fsService.String("log", commonLogFlag(), "command line log path")
	flagServiceRestart   = fsService.BoolP("restart", "R", false, "restart datakit service")
	flagServiceStop      = fsService.BoolP("stop", "T", false, "stop datakit service")
	flagServiceStart     = fsService.BoolP("start", "S", false, "start datakit service")
	flagServiceUninstall = fsService.BoolP("uninstall", "U", false, "uninstall datakit service")
	flagServiceReinstall = fsService.BoolP("reinstall", "I", false, "reinstall datakit service")
	fsServiceUsage       = func() {
		fmt.Printf("usage: datakit service [options]\n\n")
		fmt.Printf("Service used to manage datakit service\n\n")
		fmt.Println(fsService.FlagUsagesWrapped(0))
	}

	//
	// monitor related flags.
	//
	fsMonitorName              = "monitor"
	fsMonitor                  = pflag.NewFlagSet(fsMonitorName, pflag.ContinueOnError)
	flagMonitorTo              = fsMonitor.String("to", "localhost:9529", "specify the DataKit(IP:Port) to show its statistics")
	flagMonitorMaxTableWidth   = fsMonitor.IntP("max-table-width", "W", 16, "set max table cell width")
	flagMonitorOnlyInputs      = fsMonitor.StringSliceP("input", "I", nil, "show only specified inputs stats, seprated by ',', i.e., -I cpu,mem")
	flagMonitorLogPath         = fsMonitor.String("log", commonLogFlag(), "command line log path")
	flagMonitorRefreshInterval = fsMonitor.DurationP("refresh", "R", 5*time.Second, "refresh interval")
	flagMonitorVerbose         = fsMonitor.BoolP("verbose", "V", false, "show all statistics info, default not show goroutine and inputs config info")
	fsMonitorUsage             = func() {
		fmt.Printf("usage: datakit monitor [options]\n\n")
		fmt.Printf("Monitor used to show datakit running statistics\n\n")
		fmt.Println(fsMonitor.FlagUsagesWrapped(0))
	}

	//
	// install related flags.
	//
	fsInstallName       = "install"
	fsInstall           = pflag.NewFlagSet(fsInstallName, pflag.ContinueOnError)
	flagInstallLogPath  = fsInstall.String("log", commonLogFlag(), "command line log path")
	flagInstallTelegraf = fsInstall.Bool("telegraf", false, "install Telegraf")
	flagInstallScheck   = fsInstall.Bool("scheck", false, "install SCheck")
	flagInstallIPDB     = fsInstall.String("ipdb", "", "install IP database(currently only iploc available)")
	fsInstallUsage      = func() {
		fmt.Printf("usage: datakit install [options]\n\n")
		fmt.Printf("Install used to install DataKit related packages and plugins\n\n")
		fmt.Println(fsInstall.FlagUsagesWrapped(0))
	}

	//
	// debug related flags.
	//

	fsDebugName            = "debug"
	fsDebug                = pflag.NewFlagSet(fsDebugName, pflag.ContinueOnError)
	flagDebugLogPath       = fsDebug.String("log", commonLogFlag(), "command line log path")
	flagDebugCloudInfo     = fsDebug.String("show-cloud-info", "", "show current host's cloud info(aliyun/tencent/aws)")
	flagDebugIPInfo        = fsDebug.String("ipinfo", "", "show IP geo info")
	flagDebugWorkspaceInfo = fsDebug.Bool("workspace-info", false, "show workspace info")
	flagDebugCheckConfig   = fsDebug.Bool("check-config", false, "check inputs configure and main configure")
	flagDebugCmdLog        = fsDebug.String("cmd-log", "/dev/null", "command line log path")
	flagDebugDumpSamples   = fsDebug.String("dump-samples", "", "dump all inputs samples")
	flagDebugLoadLog       = fsDebug.Bool("upload-log", false, "upload log")
	fsDebugUsage           = func() {
		fmt.Printf("usage: datakit debug [options]\n\n")
		fmt.Printf("Various tools for debugging\n\n")
		fmt.Println(fsDebug.FlagUsagesWrapped(0))
	}
)

func commonLogFlag() string {
	if runtime.GOOS == datakit.OSWindows {
		return "nul" // under windows, nul is /dev/null
	}
	return "/dev/null"
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "DataKit is a collect client.\n")
	fmt.Fprintf(os.Stderr, "\nUsage:\n\n")

	fmt.Fprintf(os.Stderr, "\tdatakit <command> [arguments]\n\n")

	fmt.Fprintf(os.Stderr, "The commands are:\n\n")

	fmt.Fprintf(os.Stderr, "\tdql        query DQL for various usage\n")
	fmt.Fprintf(os.Stderr, "\trun        select DataKit running mode(defaul running as service)\n")
	fmt.Fprintf(os.Stderr, "\tpipeline   debug pipeline\n")
	fmt.Fprintf(os.Stderr, "\tservice    manage datakit service\n")
	fmt.Fprintf(os.Stderr, "\tmonitor    show datakit running statistics\n")
	fmt.Fprintf(os.Stderr, "\tinstall    install DataKit related packages and plugins\n")
	fmt.Fprintf(os.Stderr, "\tdebug      methods of all debug datakits\n")

	// TODO: add more commands...

	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Use 'datakit help <command>' for more information about a command.\n\n")
}

func runHelpFlags() {
	switch len(os.Args) {
	case 2: // only 'datakit help'
		printHelp()
	case 3: // need help for various commands
		switch os.Args[2] {
		case fsPLName:
			fsPLUsage()

		case fsDQLName:
			fsDQLUsage()

		case fsRunName:
			fsRunUsage()

		case fsVersionName:
			fsVersionUsage()

		case fsServiceName:
			fsServiceUsage()

		case fsMonitorName:
			fsMonitorUsage()

		case fsInstallName:
			fsInstallUsage()

		case fsDebugName:
			fsDebugUsage()

		default: // add more
			fmt.Fprintf(os.Stderr, "flag provided but not defined: %s", os.Args[2])
			printHelp()
			os.Exit(-1)
		}
	}
}

func doParseAndRunFlags() {
	pflag.Usage = printHelp
	pflag.ErrHelp = errors.New("")

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			runHelpFlags()
			os.Exit(0)
		}

		switch os.Args[1] {
		case fsDQLName:
			setCmdRootLog(*flagDQLLogPath)
			if err := fsDQL.Parse(os.Args[2:]); err != nil {
				errorf("Parse: %s\n", err)
				fsDQLUsage()
				os.Exit(-1)
			}

			tryLoadMainCfg()

			if err := runDQLFlags(); err != nil {
				errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsPLName:
			setCmdRootLog(*flagPLLogPath)

			debugPipelineName = os.Args[2]

			// NOTE: args[2] must be the pipeline source name
			if err := fsPL.Parse(os.Args[3:]); err != nil {
				errorf("Parse: %s\n", err)
				fsPLUsage()
				os.Exit(-1)
			}

			tryLoadMainCfg()

			if err := runPLFlags(); err != nil {
				errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsVersionName:
			setCmdRootLog(*flagVersionLogPath)
			if err := fsVersion.Parse(os.Args[2:]); err != nil {
				errorf("Parse: %s\n", err)
				fsVersionUsage()
				os.Exit(-1)
			}

			tryLoadMainCfg()

			if err := runVersionFlags(); err != nil {
				errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsServiceName:
			setCmdRootLog(*flagServiceLogPath)
			if err := fsService.Parse(os.Args[2:]); err != nil {
				errorf("Parse: %s\n", err)
				fsServiceUsage()
				os.Exit(-1)
			}

			if err := runServiceFlags(); err != nil {
				errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsMonitorName:
			setCmdRootLog(*flagMonitorLogPath)
			if err := fsMonitor.Parse(os.Args[2:]); err != nil {
				errorf("Parse: %s\n", err)
				fsMonitorUsage()
				os.Exit(-1)
			}

			if err := runMonitorFlags(); err != nil {
				errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsInstallName:
			// TODO
			setCmdRootLog(*flagInstallLogPath)
			if err := fsInstall.Parse(os.Args[2:]); err != nil {
				errorf("Parse: %s\n", err)
				fsInstallUsage()
				os.Exit(-1)
			}

			if err := installPlugins(); err != nil {
				errorf("%s\n", err)
				os.Exit(-1)
			}
			os.Exit(0)

		case fsDebugName:
			setCmdRootLog(*flagDebugLogPath)
			if err := fsDebug.Parse(os.Args[2:]); err != nil {
				errorf("Parse: %s\n", err)
				fsDebugUsage()
				os.Exit(-1)
			}

			err := runDebugFlags()
			if err != nil {
				errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		default:
			errorf("unknown command `%s'\n", os.Args[1])
			printHelp()
		}
	}
}

func ParseFlags() {
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "-") {
			parseOldStyleFlags()
		} else {
			doParseAndRunFlags()
		}
	}
}

func showDeprecatedInfo() {
	infof("\nFlag %s deprecated, please use datakit help to use recommend flags.\n\n", os.Args[1])
}

func RunCmds() {
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "-") {
			showDeprecatedInfo()
			runOldStyleCmds()
		}
	}
}

func init() { //nolint:gochecknoinits
	initOldStyleFlags()
}
