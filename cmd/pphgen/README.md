# pphgen

pphgen is a simple tool to generate secure passphrases

Usage:
```
pphgen [-n num_words]  [-list wordlist] [-entropy]
  -entropy
    	print estimated entropy.
  -list string
    	select wordlist to use. Choose from 'eff' (default), 'eff_short' and 'de'.
  -n int
    	number of words to use (default 5).
```