package configuration

import (
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
)

// IViper 🐍 is the interface for Viper methods.
type IViper interface {
	OnConfigChange(run func(in fsnotify.Event))
	WatchConfig()
	SetConfigFile(in string)
	SetEnvPrefix(in string)
	GetEnvPrefix() string
	AllowEmptyEnv(allowEmptyEnv bool)
	ConfigFileUsed() string
	AddConfigPath(in string)
	AddRemoteProvider(provider, endpoint, path string) error
	AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
	SetTypeByDefaultValue(enable bool)
	Get(key string) any
	Sub(key string) *viper.Viper
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetUint(key string) uint
	GetUint16(key string) uint16
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]any
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetSizeInBytes(key string) uint
	UnmarshalKey(key string, rawVal any, opts ...viper.DecoderConfigOption) error
	Unmarshal(rawVal any, opts ...viper.DecoderConfigOption) error
	UnmarshalExact(rawVal any, opts ...viper.DecoderConfigOption) error
	BindPFlags(flags *pflag.FlagSet) error
	BindPFlag(key string, flag *pflag.Flag) error
	BindFlagValues(flags viper.FlagValueSet) (err error)
	BindFlagValue(key string, flag viper.FlagValue) error
	BindEnv(input ...string) error
	MustBindEnv(input ...string)
	IsSet(key string) bool
	AutomaticEnv()
	SetEnvKeyReplacer(r *strings.Replacer)
	RegisterAlias(alias, key string)
	InConfig(key string) bool
	SetDefault(key string, value any)
	Set(key string, value any)
	ReadInConfig() error
	MergeInConfig() error
	ReadConfig(in io.Reader) error
	MergeConfig(in io.Reader) error
	MergeConfigMap(cfg map[string]any) error
	WriteConfig() error
	SafeWriteConfig() error
	WriteConfigAs(filename string) error
	SafeWriteConfigAs(filename string) error
	ReadRemoteConfig() error
	WatchRemoteConfig() error
	WatchRemoteConfigOnChannel() error
	AllKeys() []string
	AllSettings() map[string]any
	SetFs(fs afero.Fs)
	SetConfigName(in string)
	SetConfigType(in string)
	SetConfigPermissions(perm os.FileMode)
	Debug()
	DebugTo(w io.Writer)
}
