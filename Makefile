
.PHONY: force

all: junos_exporter

force:
	@true

junos_exporter:
	go build .
	gzip -9 < junos_exporter.8 > junos_exporter.8.gz

install:
	mkdir -p $(DESTDIR)/etc/default/
	mkdir -p $(DESTDIR)/usr/lib/systemd/system/
	mkdir -p $(DESTDIR)/usr/bin/
	mkdir -p $(DESTDIR)/usr/share/man/man8/
	echo 'ARGS=""' > $(DESTDIR)/etc/default/junos_exporter
	install -m 644 junos_exporter.service $(DESTDIR)/usr/lib/systemd/system/
	install -m 755 junos_exporter $(DESTDIR)/usr/bin/
	install -m 644 junos_exporter.8.gz $(DESTDIR)/usr/share/man/man8/

clean:
	rm -f junos_exporter junos_exporter.8.gz
