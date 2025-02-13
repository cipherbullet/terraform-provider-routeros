# routeros_radius (Resource)


## Example Usage
```terraform
resource "routeros_radius" "user_manager" {
  address = "127.0.0.1"
  service = ["ppp", "login"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) IPv4 or IPv6 address of RADIUS server.

### Optional

- `accounting_backup` (Boolean) An option whether the configuration is for the backup RADIUS server.
- `accounting_port` (Number) RADIUS server port used for accounting.
- `authentication_port` (Number) RADIUS server port used for authentication.
- `called_id` (String) RADIUS calling station identifier.
- `certificate` (String) Certificate to use for communication with RADIUS Server with RadSec enabled.
- `comment` (String)
- `disabled` (Boolean)
- `domain` (String) Microsoft Windows domain of client passed to RADIUS servers that require domain validation.
- `protocol` (String) An option specifies the protocol to use when communicating with the RADIUS Server.
- `realm` (String) Explicitly stated realm (user domain), so the users do not have to provide proper ISP domain name in the user name.
- `secret` (String, Sensitive) The shared secret to access the RADIUS server.
- `service` (Set of String) A set of router services that will use the RADIUS server. Possible values: `hotspot`, `login`, `ppp`, `wireless`, `dhcp`, `ipsec`, `dot1x`.
- `src_address` (String) Source IPv4/IPv6 address of the packets sent to the RADIUS server.
- `timeout` (String) A timeout, after which the request should be resent.

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/radius get [print show-ids]]
terraform import routeros_radius.user_manager *1
```
