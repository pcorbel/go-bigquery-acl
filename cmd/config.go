package main

import (
        "io/ioutil"

        "gopkg.in/yaml.v2"
)

// Config object to be loaded with configuration from YAML file
type Config struct {
        Project  string `yaml:"project,omitempty"`
        Datasets []struct {
                Name  string `yaml:"name,omitempty"`
                Owner struct {
                        GroupByEmail []string `yaml:"group_by_email,omitempty"`
                        UserByEmail  []string `yaml:"user_by_email,omitempty"`
                        SpecialGroup []string `yaml:"special_group,omitempty"`
                } `yaml:"owner,omitempty"`
                Writer struct {
                        GroupByEmail []string `yaml:"group_by_email,omitempty"`
                        UserByEmail  []string `yaml:"user_by_email,omitempty"`
                        SpecialGroup []string `yaml:"special_group,omitempty"`
                } `yaml:"writer,omitempty"`
                Reader struct {
                        GroupByEmail []string `yaml:"group_by_email,omitempty"`
                        UserByEmail  []string `yaml:"user_by_email,omitempty"`
                        SpecialGroup []string `yaml:"special_group,omitempty"`
                } `yaml:"reader,omitempty"`
                View []struct {
                        DatasetID string `yaml:"dataset_id,omitempty"`
                        ViewID    string `yaml:"view_id,omitempty"`
                } `yaml:"view,omitempty"`
        } `yaml:"datasets,omitempty"`
}

// LoadFromFile return Config object according to a YAML file
func (conf *Config) LoadFromFile(file string) error {
        yamlFile, err := ioutil.ReadFile(file)
        if err == nil {
                err = yaml.Unmarshal(yamlFile, conf)
        }
        return err
}