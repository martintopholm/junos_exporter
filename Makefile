
.PHONY: force

all: junos_exporter

force:
	@true

junos_exporter:
	go build .
	gzip -9 < junos_exporter.8 > junos_exporter.8.gz

install:
	mkdir -p $(DESTDIR)/usr/sbin/
	mkdir -p $(DESTDIR)/usr/share/man/man8
	install -m 755 junos_exporter $(DESTDIR)/usr/sbin/
	install -m 755 junos_exporter.8.gz $(DESTDIR)/usr/share/man/man8
