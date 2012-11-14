
dwmstatus.go
============


Outputs cpu usage, mem usage, weather and date.

    dwmstatus | while read -r; do xsetroot -name "$REPLY"; done &

### Sample Output

	 1 | 15 | 2.3 Fog | Wed Nov 14 13:37

Dependencies
------------

* wu (https://github.com/sramsay/wu)

### Note on wu

To let dwmstatus not query wunderground.com every two seconds, a cronjob can be set up, e.g.

	SHELL=/bin/zsh

	*/10 * * * * < =(wu -conditions) | awk '/Temperature/{gsub("[(]","",$4); t=$4}; \
	/Conditions/{c=($3 " " $4)}END{print t,c}' > /tmp/wucond.txt

	0    0 * * * < =(wu -astro) | awk '/Sunrise/{r=$3}; \
	/Sunset/{s=$3}END{print r,s}' > /tmp/wuastro.txt

See https://github.com/sramsay/wu for further information , e.g. the API key.

Todo
----

* Make use of x-go-binding.
* Import wu and don't use cronjobs.


### Maybe

* Implement more status outputs.
* Implement some option switches.

Credits
-------

* cornu (https://github.com/Cornu/dwmstatus)
