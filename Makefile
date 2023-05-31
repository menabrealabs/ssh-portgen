ssh-portgen: 
	go build -o ./build/bin/ssh-portgen

clean:
	rm -rf ./build/bin

clean-dist:
	rm -rf ./build

install : ssh-portgen
	# Installing binary to /usr/local/bin/ssh-portgen
	cp ./build/bin/ssh-portgen /usr/local/bin/

uninstall:
	# Uninstalling from /usr/local/bin
	rm /usr/local/bin/ssh-portgen