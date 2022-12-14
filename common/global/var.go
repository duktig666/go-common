// description:
// @author renshiwei
// Date: 2022/10/5 17:06

package global

const VERSION = "v0.0.1"

var (
	Config ConfigYaml
)

type ConfigYaml struct {
	Server struct {
		Name    string
		Version string
		Http    struct {
			Host string
			Port int
		}
	}

	Cli struct {
		Name string
	}

	Log struct {
		Level struct {
			Server string
		}
	}
}
