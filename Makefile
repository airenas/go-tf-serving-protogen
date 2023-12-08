####################################################
tf?=2.10.0
tfs?=$(tf)
buf?=1.28.1

####################################################
generate: _tfcode/.serving.$(tfs) _tfcode/.tf.$(tf)
	docker run --volume "$(CURDIR):/workspace" --workdir /workspace bufbuild/buf:$(buf) generate
	go mod tidy
.PHONY: generate

####################################################
_tfcode/.serving.$(tfs):
	rm -rf _tfcode/serving
	git clone --branch $(tfs) --depth 1 https://github.com/tensorflow/serving.git _tfcode/serving
	touch $@

_tfcode/.tf.$(tf):
	rm -rf _tfcode/tensorflow
	git clone --branch v$(tf) --depth 1 https://github.com/tensorflow/tensorflow.git _tfcode/tensorflow
	touch $@
####################################################
clean:
	rm -rf _tfcode	
.PHONY: clean
####################################################
