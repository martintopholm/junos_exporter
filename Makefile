
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
	echo 'ARGS=""' > $(DESTDIR)/etc/defaults/junos_exporter
	install -m 755 junos_exporter $(DESTDIR)/usr/sbin/
	install -m 644 junos_exporter.service $(DESTDIR)/usr/lib/systemd/system
	install -m 644 junos_exporter.8.gz $(DESTDIR)/usr/share/man/man8
