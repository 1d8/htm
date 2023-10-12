This a compilation of notes for security camera hacking for the Foscam, specifically the **INSERT MODEL HERE** model.

When we first attempt to visit the webpage of the camera, we are asked to login:

![](https://i.imgur.com/A2eD6XZ.png)

When we send our username & password to the website, it first combines the username & password with the **:** delimeter and then it base64 encodes the username & password combination and adds it to an **Authorization** header. This is known as a **Basic Authorization Header**. Here's an example request along with an example response:

![](https://i.imgur.com/azX3kKs.png)

![](https://i.imgur.com/Ef384oO.png)

And if we get the password correct, here's an example response:

![](https://i.imgur.com/KchFud4.png)

Knowing this, I built a Python script that can automate the brute forcing process while taking a username & password list. 

Example usage of this script is:

`python3 bf.py --target <TARGET IP ADDRESS> --username userlist.txt --password passlist.txt`:

![](https://i.imgur.com/jFtfUES.png)

And when it finds the correct **username:password** combination, it'll stop running!
