build:
	GOOS=linux go build -o app
	docker build -t gcr.io/text-to-speech/text-to-speech .
	rm -f app

push:
	gcloud docker -- push gcr.io/text-to-speech/text-to-speech