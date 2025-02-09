# libsisimai.org/sisimai/Developers.mk
#  ____                 _                                       _    
# |  _ \  _____   _____| | ___  _ __   ___ _ __ ___   _ __ ___ | | __
# | | | |/ _ \ \ / / _ \ |/ _ \| '_ \ / _ \ '__/ __| | '_ ` _ \| |/ /
# | |_| |  __/\ V /  __/ | (_) | |_) |  __/ |  \__ \_| | | | | |   < 
# |____/ \___| \_/ \___|_|\___/| .__/ \___|_|  |___(_)_| |_| |_|_|\_\
#                              |_|                                   
# -------------------------------------------------------------------------------------------------
SHELL := /bin/sh
HERE  := $(shell pwd)
NAME  := sisimai
MKDIR := mkdir -p
LS    := ls -1
CP    := cp

LIBSISIMAI := libsisimai.org
SISIMAIDIR := address arf fact lda lhost mail message reason rfc1123 rfc1894 rfc2045 rfc3464 \
			  rfc3834 rfc5322 rfc5965 rfc791 rhost sis smtp string
COVERAGETO := coverage.txt

# -------------------------------------------------------------------------------------------------
.PHONY: clean

list-test-files:
	@ find $(SISIMAIDIR) -type f -name '*_test.go'

count-test-cases:
	@ go test -v `find $(SISIMAIDIR) -type f -name '*_test.go' | xargs dirname | sort | uniq | sed -e 's|^|./|g'` | \
		grep 'The number of ' | awk '{ cx += $$7 } END { print cx }'

loc:
	@ find libsisimai.go $(SISIMAIDIR) -type f -name '*.go' -not -name '*_test.go' | \
		xargs grep -vE '(^$$|^//|/[*]|[*]/|^ |^--)' | grep -vE "\t+//" | wc -l

clean:

