## INSTALL

from: https://www.vultr.com/docs/using-strongswan-for-ipsec-vpn-on-centos-7

## config

download, but always change
```
yum install http://dl.fedoraproject.org/pub/epel/7/x86_64/e/epel-release-7-10.noarch.rpm -y
yum install http://dl.fedoraproject.org/pub/epel/7/x86_64/Packages/e/epel-release-7-11.noarch.rpm -y
visit http://dl.fedoraproject.org/pub/epel to fix
```

ipsec.conf
```
config setup
    uniqueids=never
    charondebug="cfg 2, dmn 2, ike 2, net 0"

conn %default
    left=%defaultroute
    leftsubnet=0.0.0.0/0
    leftcert=vpnHostCert.pem
    right=%any
    rightsourceip=172.16.1.100/16
    ike=aes256-sha256-modp2048,aes256-sha256-modp1024,aes256-sha1-modp1024!

conn CiscoIPSec
    keyexchange=ikev1
    fragmentation=yes
    rightauth=pubkey
    rightauth2=xauth
    leftsendcert=always
    rekey=no
    auto=add

conn XauthPsk
    keyexchange=ikev1
    leftauth=psk
    rightauth=psk
    rightauth2=xauth
    auto=add

#conn XauthPsk-Android
#    keyexchange=ikev1
#    leftauth=psk
#    rightauth=psk
#    rightauth2=xauth
#    aggressive=yes
#    dpdaction=clear
#    ikelifetime=500h
#    lifetime=200s
#    auto=add

conn IpsecIKEv2
    keyexchange=ikev2
    leftauth=pubkey
    rightauth=pubkey
    leftsendcert=always
    auto=add

conn IpsecIKEv2-EAP
    keyexchange=ikev2
    rekey=no
    leftauth=pubkey
    leftsendcert=always
    rightauth=eap-mschapv2
    eap_identity=%any
    auto=add

```

strongswan.conf
```
charon {
    load_modular = yes
    #i_dont_care_about_security_and_use_aggressive_mode_psk = yes
    duplicheck.enable = no
    compress = yes
    plugins {
            include strongswan.d/charon/*.conf
    }
    dns1 = 8.8.8.8
    dns2 = 8.8.4.4
    nbns1 = 8.8.8.8
    nbns2 = 8.8.4.4
}

include strongswan.d/*.conf
```

ipsec.secrets
```
: RSA vpnHostKey.pem
: PSK "helloyaoguai"
client %any : EAP "xxxx"
client %any : XAUTH "xxxx"
litao %any : EAP "xxxx"
litao %any : XAUTH "xxxx"
```

