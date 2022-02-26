# SSOh-No 

This tool is designed to enumerate users, password spray and perform brute force attacks against Azure AD via a seamless SSO autologon endpoint. 

Generally, this endpoint provides extremely verbose errors which can be leveraged to enumerate users and validate their passwords via brute force/spraying attacks. 

This tool is a weaponised version of a PoC demonstrated in the [arstechnica article research article](https://arstechnica.com/information-technology/2021/09/new-azure-active-directory-password-brute-forcing-flaw-has-no-fix/]) which discusses the techniques utilised to exploit the endpoint.

This endpoint is known to Microsoft however, in typical fashion it has been branded a feature, not a bug.

This endpoint does enforce "smart locking" which can be bypassed by rotating IP. 

### Why Is This Unique?
The SSO Autologon endpoint does not contain logging of any sort bar potentially updating the users "Last Logon" time. 

The following have been tested and contain no logs:
- AzureAD
- Sentinel
- Defender for Identity (Formerly Advanced Thread Protection)
- Defender for Cloud Apps

## Usage
```
$ ./SSOh-No -h
usage: SSOh-No [-h|--help] [-e|--email "<value>"] [-p|--password "<value>"]
               [-U|--userlist "<value>"]

               Enumerate and abuse a sub-par Azure SSO endpoint.

Arguments:

  -h  --help      Print help information
  -e  --email     Email address to query. Example: user@domain.com
  -p  --password  Password to spray. Example: Password123!
  -U  --userlist  Specify userlist to enumerate
```


## Upcoming Features 

- Proxy Implementation to bypass smart lock
- Password brute force from password lists (single user- No plans for password list brute force against a userlist)
- Output files for successful user enum/password spraying
