pphgen is a simple tool to generate secure *p*ass*ph*rases

install (assumes go 1.13 or newer in module mode):

go get github.com/mxrth/pphgen/cmd/pphgen

Usage of pphgen:
  -entropy
    	print estimated entropy
  -list string
    	select wordlist to use. Choose from 'eff', 'eff_short' and 'de' (default "eff")
  -n int
    	number of words to use (default 5)
