package subcmd

import "flag"
import "fmt"
import "os"

type subCommand interface {
	Name() string
	DefineFlags(*flag.FlagSet)
	Run() int
	Usage()
}

type subCommandParser struct {
	cmd subCommand
	fs  *flag.FlagSet
}

// Parse ...
func Parse(usage func(), commands ...subCommand) int {
	scp := make(map[string]*subCommandParser, len(commands))
	for _, cmd := range commands {
		name := cmd.Name()
		scp[name] = &subCommandParser{cmd, flag.NewFlagSet(name, flag.ExitOnError)}
		cmd.DefineFlags(scp[name].fs)
	}

	flag.Usage = func() {
		usage()
		for _, sc := range scp {
			sc.cmd.Usage()
			fmt.Fprintf(os.Stderr, "\n") // nolint: gas
		}
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return 64
	}

	cmdName := flag.Arg(0)
	if sc, ok := scp[cmdName]; ok {
		err := sc.fs.Parse(flag.Args()[1:])
		if err != nil {
			panic(err)
		}
		return sc.cmd.Run()
	}

	fmt.Fprintf(os.Stderr, "invalid command %s", cmdName) // nolint: gas
	flag.Usage()
	return 64
}
