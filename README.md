
dwmstatus.go
============

![Screenshot](http://github.com/rwilhelm/dwmstatus.go/raw/master/screenshot.png)

Outputs cpu usage, mem usage, weather and date. For now it can be used by e.g.

    dwmstatus | while read -r; do xsetroot -name "$REPLY"; done &

Dependencies
------------

* wu (https://github.com/sramsay/wu)

### Note on wu

To let dwmstatus not query wunderground.com every two seconds, a cronjob can be set up, e.g.

    */10 * * * * $HOME/bin/wu > /tmp/wu.txt && cat /tmp/wu.txt | awk \
	'/Temperature/{gsub("[(]","",$4); t=$4}; /Conditions/{c=($3 " " $4)}END{print t,c}' > /tmp/weather.txt

See https://github.com/sramsay/wu for further information , e.g. the API key.

Todo
----

* Make use of x-go-binding.


### Maybe

* Implement some more status outputs (battery, GMail).
* Make single status outputs optional.

Credits
-------

* cornu (https://github.com/Cornu/dwmstatus)
