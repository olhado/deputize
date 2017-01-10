package config

import (
  "encoding/json"
  "io/ioutil"
)

var Config DeputizeConfig

func init() {
  Config = NewConfig("config.json")
}

// DeputizeConfig is our config struct
type DeputizeConfig struct {
  BaseDN string
  GrayLogEnabled bool
  GrayLogAddress string
  LDAPServer string
  LDAPPort int
  MailAttribute string
  MemberAttribute string
  ModUserDN string
  OnCallGroup string
  OnCallGroupDN string
  OnCallSchedules []string
  RootCAFile string
  SlackChan string
  SlackEnabled bool
  TokenPath string
  UserAttribute string
  VaultSecretPath string
  VaultServer string
}

func NewConfig(fname string) DeputizeConfig {
  data,err := ioutil.ReadFile(fname)
  if err != nil{
    panic(err)
  }
  config := DeputizeConfig{}
  err = json.Unmarshal(data,&config)
  if err != nil {
    panic(err)
  }
  return config
}