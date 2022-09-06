#!/usr/bin/python
from datetime import date

lc = [13,0,0,0,0]
d = (date.today() - date(2012, 12, 21)).days # days since epoch

while d >= 144000: # baktun
	lc[0] += 1
	d -= 144000
while d >= 7200: # katun
	lc[1] += 1
	d -= 7200
while d >= 360: # tun
	lc[2] += 1
	d -= 360
while d >= 20: # uinal
	lc[3] += 1
	d -= 20
lc[4] = d # kin

print(*lc, sep=".")