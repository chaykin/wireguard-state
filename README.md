# WireGuard-State
System tray icon for WireGuard VPN

Simple utility for control WireGuard VPN (check/enable/disable) by system tray icon. Tested on Ubuntu 21.04 only.

## Installation
1. Since You don't want to run this as root, You could delegate some work to Network Manager (It supports WireGuard since 1.26.2).
So, You can create and configure a WireGuard network connection by `nm-connection-editor` 
2. Download last [release archive](https://github.com/chaykin/wireguard-state/releases) and unpack it
3. Check config file `wgs.yaml`. It contains bash-script commands to control VPN. Adjust it, if needed: 
provide actual Network Manager connection id (by default, `wg0`) for `wgUpCommand`, `wgStateCommand`, `wgDownCommand` and 
WireGuard interface name (usually, `wg0`) for `rxCommand` and `txCommand`. Or, You can write own completely different command,
if You use different Linux distributive
4. Run it.
