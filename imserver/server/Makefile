start_mosca:
	node mosca.js --internalPort 1883 --internalUsername testuser --internalPassword testpassword \
		--externalPort 8883 --externalCert ./server.crt --externalCertKey ./server.key \
		--agentHost 127.0.0.1 --agentPort 50050 &
	node mosca.js --internalPort 1884 --internalUsername testuser --internalPassword testpassword \
		--externalPort 8884 --externalCert ./server.crt --externalCertKey ./server.key \
		--agentHost 127.0.0.1 --agentPort 50050 &

stop_mosca:
	ps aux  |  grep -i '[n]ode mosca.js'  |  awk '{print $$2}'  |  xargs kill -9