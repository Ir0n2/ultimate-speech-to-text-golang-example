# Example code how to do Speech-To-Text in Go

We use a simple HTTP client with `net/http` and the [AssemblyAI API](https://www.assemblyai.com).

- Step 1 Upload a file: [upload.go](upload.go)
- Step 2 Transcribe: [transcribe.go](transcribe.go)
- Step 3 Poll the result: [poll.go](poll.go)

These steps have been simplified into 3 simple functions. Just add in your api key and you should be good to go!

ultimate Transcriber.go is like an all in one script, no need for 3 individual go files. Code should be pretty self explanatory.
