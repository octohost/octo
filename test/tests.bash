
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

T_14configdelrm() {
  bin/octo config del -c darron --key="KEY1"
  bin/octo config rm -c darron --key="KEY2"
}

T_16configshow() {
  result="$(bin/octo config show -c darron)"
  [[ "$result" == '/darron/TESTING:"This value is set for the darron container."' ]]
}

T_18configexport() {
  result="$(bin/octo config export -c darron)"
  [[ "$result" == 'octo config set -c="darron" --key="TESTING" --value="This value is set for the darron container."' ]]
}

T_20containerGetSpace() {
  result="$(bin/octo config get -c 'there should not')"
  [[ "$result" == 'A container cannot contain a space.' ]]
}

T_22getSpace() {
  result="$(bin/octo config get -c 'darron' --key 'This key should fail.' )"
  [[ "$result" == 'A key cannot contain a space.' ]]
}

T_24containerSetSpace() {
  result="$(bin/octo config set -c 'there should not')"
  [[ "$result" == 'A container cannot contain a space.' ]]
}

T_26setSpace() {
  result="$(bin/octo config set -c 'darron' --key 'This key should fail.' )"
  [[ "$result" == 'A key cannot contain a space.' ]]
}

T_28containerDelSpace() {
  result="$(bin/octo config del -c 'there should not')"
  [[ "$result" == 'A container cannot contain a space.' ]]
}

T_30delSpace() {
  result="$(bin/octo config del -c 'darron' --key 'This key should fail.' )"
  [[ "$result" == 'A key cannot contain a space.' ]]
}

T_32containerShow() {
  result="$(bin/octo config show -c 'there should not')"
  [[ "$result" == 'A container cannot contain a space.' ]]
}

T_34containerShow() {
  result="$(bin/octo config export -c 'there should not')"
  [[ "$result" == 'A container cannot contain a space.' ]]
}
