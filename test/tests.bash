
trap "make consul_kill" INT TERM EXIT

export OCTO_DEBUG=1

T_06runbinary() {
  result="$(bin/octo)"
}

T_10configset() {
  result="$(bin/octo config set -c darron --key testing --value='This value is set for the darron container.')"
  configset="$(consul-cli kv-read octohost/darron/TESTING | grep 'darron container')"
  [[ "$?" -eq "0" ]]
}

T_12configget() {
  result="$(bin/octo config get -c darron --key testing | grep 'darron container')"
  [[ "$?" -eq "0" ]]
}
