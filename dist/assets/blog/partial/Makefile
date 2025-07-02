IMAGE = my-image 
WORKDIR = /post
QMD = content.qmd
OUT = content.html
LOCALDIR = .

.PHONY: build render clean

build:
	docker build -t $(IMAGE) .

render:
	docker run --rm -v "$(LOCALDIR):/workspace" $(IMAGE) \
		bash -c "quarto render /workspace/$(QMD)"

clean:
	rm -f $(OUT)

