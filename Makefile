ssh-portgen: 
	go build -o ./build/bin/ssh-portgen

clean:
	rm -rf ./build/bin

install : ssh-portgen
	cp ./build/bin/ssh-portgen /usr/local/bin/

uninstall:
	rm /usr/local/bin/ssh-portgen
