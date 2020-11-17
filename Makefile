
.PHONY: force

all: junos_exporter

force:
	@true

junos_exporter:
	go build .

install:
	mkdir -p $(DESTDIR)/usr/sbin/
	install -m 755 junos_exporter $(DESTDIR)/usr/sbin/
