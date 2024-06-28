package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	APP_NAME    = "gw"
	CONFIG_FILE = "config.toml"
)

type Config struct {
	Path   string
	Config GwtConfig
}

type GwtConfig map[string]RepoConfig

type RepoConfig struct {
	Path         string   `toml:"path"`
	WorkTreePath string   `toml:"worktree_path"`
	Script       []string `toml:"script"`
}

func (r RepoConfig) Name() string {
	return filepath.Base(r.Path)
}

func (r RepoConfig) WorkDir() string {
	if r.WorkTreePath == "" {
		return filepath.Join(r.Path, ".worktree")
	}
	return filepath.Join(r.WorkTreePath, r.Name())
}

func (r RepoConfig) Cmd() []string {
	return r.Script
}

func New(path string) *Config {
	config := &Config{
		Path:   path,
		Config: make(map[string]RepoConfig),
	}

	if err := config.load(); err != nil {
		log.Fatalf("failed to load config: %v", err)
		return nil
	}

	return config
}

func (c *Config) FindRepoByPath(cwd string) *RepoConfig {
	for _, value := range c.Config {
		if strings.HasPrefix(cwd, value.Path) {
			return &value
		}
	}
	return nil
}

func (c *Config) FindRepoByName(name string) *RepoConfig {
	for _, value := range c.Config {
		if name == value.Name() {
			return &value
		}
	}
	return nil
}

func (c *Config) AddRepo(basePath, worktreePath, script string) error {
	if basePath == "" {
		return fmt.Errorf("path is required")
	}

	//check if path is a git repo
	if _, err := os.Stat(filepath.Join(basePath, ".git")); os.IsNotExist(err) {
		return fmt.Errorf("not a git repo: %s", basePath)
	}

	config := RepoConfig{
		Path:         basePath,
		WorkTreePath: worktreePath,
		Script:       strings.Split(script, ";"),
	}

	name := config.Name()

	if repo := c.FindRepoByName(name); repo != nil {
		return fmt.Errorf("name %s already exists", name)
	}

	c.Config[name] = config
	return c.save()
}

func (c *Config) RemoveRepo(name string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}

	if _, ok := c.Config[name]; !ok {
		return fmt.Errorf("name %s does not exist", name)
	}
	delete(c.Config, name)
	return c.save()
}

func (c *Config) init() error {
	// Ensure the directory exists
	dir := filepath.Dir(c.Path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Ensure the file exists
	file, err := os.OpenFile(c.Path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return nil
}

func (c *Config) load() error {
	if err := c.init(); err != nil {
		return fmt.Errorf("failed to init config: %w", err)
	}

	_, err := toml.DecodeFile(c.Path, &c.Config)
	if err != nil {
		return fmt.Errorf("failed to decode config: %w", err)
	}

	return nil
}

func (c *Config) save() error {
	file, err := os.OpenFile(c.Path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	err = toml.NewEncoder(file).Encode(c.Config)
	if err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}
	return nil
}

func GetConfigPath() string {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		log.Fatal("XDG_CONFIG_HOME not set")
	}
	return filepath.Join(configDir, APP_NAME, CONFIG_FILE)
}

func GetCWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}
	return cwd
}
