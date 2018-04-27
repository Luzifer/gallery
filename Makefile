default:

lint:
	docker run --rm -ti -v $(CURDIR):$(CURDIR) -w $(CURDIR) luzifer/eslint frontend/*.js

pack:
	go-bindata -modtime 1 frontend/...

publish:
	curl -sSLo golang.sh https://raw.githubusercontent.com/Luzifer/github-publish/master/golang.sh
	bash golang.sh
