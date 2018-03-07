$ORIGIN demons-souls.com.
@	3600 IN	SOA ns1.demons-souls.com. ns2.demons-souls.com. (
				2017042745 ; serial
				7200       ; refresh (2 hours)
				3600       ; retry (1 hour)
				1209600    ; expire (2 weeks)
				3600       ; minimum (1 hour)
				)

	3600 IN NS ns1.demons-souls.com.
	3600 IN NS ns2.demons-souls.com.

c   IN A     192.168.15.10
ns1 IN A     192.168.15.10
ns2 IN A     192.168.15.10
